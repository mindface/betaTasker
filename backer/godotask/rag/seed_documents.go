package rag

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const defaultEmbeddingDimension = 16

type SeedDocumentOptions struct {
	Sources            []string `json:"sources"`
	EmbeddingDimension int      `json:"embedding_dimension"`
}

type SeedOntologyRankRequest struct {
	Policy             *SearchScoringPolicy `json:"policy"`
	QueryConceptIDs    []string             `json:"query_concept_ids" binding:"required"`
	QueryText          string               `json:"query_text"`
	Sources            []string             `json:"sources"`
	EmbeddingDimension int                  `json:"embedding_dimension"`
	Limit              int                  `json:"limit"`
}

type SeedEmbeddingResult struct {
	StoredCount int                `json:"stored_count"`
	Documents   []OntologyDocument `json:"documents"`
}

func BuildSeedOntologyDocuments(db *gorm.DB, sources []string) ([]OntologyDocument, error) {
	sourceSet := seedSourceSet(sources)
	documents := []OntologyDocument{}

	if sourceSet["knowledge_patterns"] {
		var patterns []model.KnowledgePattern
		if err := db.Order("created_at DESC, id DESC").Find(&patterns).Error; err != nil {
			return nil, err
		}
		for _, pattern := range patterns {
			documents = append(documents, knowledgePatternDocument(pattern))
		}
	}

	if sourceSet["phenomenological_frameworks"] {
		var frameworks []model.PhenomenologicalFramework
		if err := db.Order("created_at DESC, id DESC").Find(&frameworks).Error; err != nil {
			return nil, err
		}
		for _, framework := range frameworks {
			documents = append(documents, phenomenologicalFrameworkDocument(framework))
		}
	}

	if sourceSet["memories"] {
		var memories []model.Memory
		if err := db.Order("created_at DESC, id DESC").Find(&memories).Error; err != nil {
			return nil, err
		}
		for _, memory := range memories {
			documents = append(documents, memoryDocument(memory))
		}
	}

	return documents, nil
}

func StoreSeedDocumentEmbeddings(db *gorm.DB, docs []OntologyDocument, dimension int) (int, error) {
	dimension = normalizeEmbeddingDimension(dimension)
	rows := make([]model.RagDocumentEmbedding, 0, len(docs))

	for _, doc := range docs {
		sourceTable, sourceID := splitDocumentID(doc.ID)
		rows = append(rows, model.RagDocumentEmbedding{
			DocumentID:  doc.ID,
			SourceTable: sourceTable,
			SourceID:    sourceID,
			Text:        doc.Text,
			OntologyTags: model.JSON{
				"values": doc.OntologyTags,
			},
			Metadata:  model.JSON(doc.Metadata),
			Embedding: model.JSON{"values": DeterministicEmbedding(doc.Text, dimension)},
		})
	}

	if len(rows) == 0 {
		return 0, nil
	}

	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "document_id"}},
		UpdateAll: true,
	}).Create(&rows).Error
	if err != nil {
		return 0, err
	}

	return len(rows), nil
}

func RankSeedOntologyDocuments(db *gorm.DB, req SeedOntologyRankRequest) ([]SearchScoreResult, error) {
	docs, err := BuildSeedOntologyDocuments(db, req.Sources)
	if err != nil {
		return nil, err
	}

	dimension := normalizeEmbeddingDimension(req.EmbeddingDimension)
	if _, err := StoreSeedDocumentEmbeddings(db, docs, dimension); err != nil {
		return nil, err
	}

	registry := NewOntologyRegistry()
	registry.RegisterMany(ConceptsFromDocuments(docs))
	registry.BuildFeatures()

	textSimilarities := map[string]float64{}
	queryEmbedding := DeterministicEmbedding(req.QueryText, dimension)
	for _, doc := range docs {
		textSimilarities[doc.ID] = CosineSimilarity(queryEmbedding, DeterministicEmbedding(doc.Text, dimension))
	}

	policy := DefaultSearchScoringPolicy()
	if req.Policy != nil {
		policy = *req.Policy
	}

	results, err := RankOntologyDocuments(registry, policy, req.QueryConceptIDs, textSimilarities, docs)
	if err != nil {
		return nil, err
	}

	if req.Limit > 0 && req.Limit < len(results) {
		return results[:req.Limit], nil
	}
	return results, nil
}

func ConceptsFromDocuments(docs []OntologyDocument) []OntologyConcept {
	seen := map[string]OntologyConcept{}
	for _, doc := range docs {
		for _, tag := range doc.OntologyTags {
			concept := seen[tag]
			concept.ID = tag
			concept.Label = tag
			concept.Definition = "seed derived ontology tag"
			concept.Domain = fmt.Sprint(doc.Metadata["domain"])
			if concept.Metadata == nil {
				concept.Metadata = map[string]interface{}{}
			}
			concept.Metadata["risk_level"] = maxFloat(numberFromMetadata(concept.Metadata, "risk_level"), numberFromMetadata(doc.Metadata, "risk_level"))
			seen[tag] = concept
		}
	}

	concepts := make([]OntologyConcept, 0, len(seen))
	for _, concept := range seen {
		concepts = append(concepts, concept)
	}
	sort.Slice(concepts, func(i, j int) bool {
		return concepts[i].ID < concepts[j].ID
	})
	return concepts
}

func DeterministicEmbedding(text string, dimension int) []float64 {
	dimension = normalizeEmbeddingDimension(dimension)
	vector := make([]float64, dimension)
	for _, token := range strings.Fields(strings.ToLower(text)) {
		hash := sha256.Sum256([]byte(token))
		index := int(binary.BigEndian.Uint64(hash[:8]) % uint64(dimension))
		sign := 1.0
		if hash[8]%2 == 0 {
			sign = -1.0
		}
		vector[index] += sign
	}
	normalizeVector(vector)
	return vector
}

func CosineSimilarity(a []float64, b []float64) float64 {
	if len(a) == 0 || len(a) != len(b) {
		return 0
	}

	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return clamp01((dot/math.Sqrt(normA*normB) + 1) / 2)
}

func knowledgePatternDocument(pattern model.KnowledgePattern) OntologyDocument {
	return OntologyDocument{
		ID:   "knowledge_patterns:" + pattern.ID,
		Text: strings.Join([]string{pattern.TacitKnowledge, pattern.ExplicitForm}, "\n"),
		OntologyTags: compactStrings([]string{
			"source:knowledge_patterns",
			"domain:" + pattern.Domain,
			"type:" + pattern.Type,
			"level:" + pattern.AbstractLevel,
			fmt.Sprintf("task:%d", pattern.TaskID),
		}),
		Metadata: map[string]interface{}{
			"source_table":   "knowledge_patterns",
			"source_id":      pattern.ID,
			"domain":         pattern.Domain,
			"risk_level":     riskLevelFromQuality(pattern.Accuracy, pattern.Coverage, pattern.Consistency),
			"accuracy":       pattern.Accuracy,
			"coverage":       pattern.Coverage,
			"consistency":    pattern.Consistency,
			"abstract_level": pattern.AbstractLevel,
		},
	}
}

func phenomenologicalFrameworkDocument(framework model.PhenomenologicalFramework) OntologyDocument {
	return OntologyDocument{
		ID: "phenomenological_frameworks:" + framework.ID,
		Text: strings.Join([]string{
			framework.Name,
			framework.Description,
			framework.Goal,
			framework.Scope,
			framework.GoalFunction,
		}, "\n"),
		OntologyTags: compactStrings([]string{
			"source:phenomenological_frameworks",
			"domain:" + framework.Domain,
			"level:" + framework.AbstractLevel,
			fmt.Sprintf("task:%d", framework.TaskID),
		}),
		Metadata: map[string]interface{}{
			"source_table":   "phenomenological_frameworks",
			"source_id":      framework.ID,
			"domain":         framework.Domain,
			"risk_level":     riskLevelFromRange(framework.LimitMin, framework.LimitMax),
			"limit_min":      framework.LimitMin,
			"limit_max":      framework.LimitMax,
			"abstract_level": framework.AbstractLevel,
		},
	}
}

func memoryDocument(memory model.Memory) OntologyDocument {
	return OntologyDocument{
		ID: "memories:" + fmt.Sprint(memory.ID),
		Text: strings.Join([]string{
			memory.Title,
			memory.Notes,
			memory.Factor,
			memory.Process,
			memory.EvaluationAxis,
			memory.InformationAmount,
		}, "\n"),
		OntologyTags: compactStrings(append([]string{
			"source:memories",
			"source_type:" + memory.SourceType,
			"status:" + memory.ReadStatus,
		}, strings.Split(memory.Tags, ",")...)),
		Metadata: map[string]interface{}{
			"source_table": "memories",
			"source_id":    fmt.Sprint(memory.ID),
			"risk_level":   1,
		},
	}
}

func seedSourceSet(sources []string) map[string]bool {
	if len(sources) == 0 {
		return map[string]bool{
			"knowledge_patterns":           true,
			"phenomenological_frameworks": true,
			"memories":                    true,
		}
	}

	result := map[string]bool{}
	for _, source := range sources {
		result[source] = true
	}
	return result
}

func normalizeEmbeddingDimension(dimension int) int {
	if dimension <= 0 {
		return defaultEmbeddingDimension
	}
	return dimension
}

func normalizeVector(vector []float64) {
	var norm float64
	for _, value := range vector {
		norm += value * value
	}
	if norm == 0 {
		return
	}
	norm = math.Sqrt(norm)
	for i := range vector {
		vector[i] = vector[i] / norm
	}
}

func splitDocumentID(documentID string) (string, string) {
	parts := strings.SplitN(documentID, ":", 2)
	if len(parts) != 2 {
		return "unknown", documentID
	}
	return parts[0], parts[1]
}

func compactStrings(values []string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}

func riskLevelFromQuality(values ...float64) float64 {
	if len(values) == 0 {
		return 1
	}
	var total float64
	for _, value := range values {
		total += value
	}
	average := total / float64(len(values))
	if average >= 0.9 {
		return 2
	}
	if average >= 0.75 {
		return 3
	}
	return 4
}

func riskLevelFromRange(minValue float64, maxValue float64) float64 {
	if maxValue == 0 {
		return 1
	}
	ratio := math.Abs(maxValue - minValue) / math.Abs(maxValue)
	if ratio <= 0.1 {
		return 4
	}
	if ratio <= 0.3 {
		return 3
	}
	return 2
}
