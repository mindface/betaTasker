package rag

import "sort"

// OntologyConcept is a compact in-memory representation of a concept loaded
// from seed data, RDF/SPARQL results, or a domain ontology definition.
type OntologyConcept struct {
	ID         string                 `json:"id"`
	Label      string                 `json:"label"`
	Definition string                 `json:"definition"`
	ParentIDs  []string               `json:"parent_ids"`
	RelatedIDs []string               `json:"related_ids"`
	Metadata   map[string]interface{} `json:"metadata"`
	Domain     string                 `json:"domain"`
}

// OntologyDocument is the search target after document ingestion.
type OntologyDocument struct {
	ID           string                 `json:"id"`
	Text         string                 `json:"text"`
	OntologyTags []string              `json:"ontology_tags"`
	Metadata     map[string]interface{} `json:"metadata"`
}

// OntologyRegistry stores concepts and graph-derived features used for ranking.
type OntologyRegistry struct {
	concepts   map[string]OntologyConcept
	centrality map[string]float64
}

func NewOntologyRegistry() *OntologyRegistry {
	return &OntologyRegistry{
		concepts:   map[string]OntologyConcept{},
		centrality: map[string]float64{},
	}
}

func (r *OntologyRegistry) RegisterMany(concepts []OntologyConcept) {
	for _, concept := range concepts {
		r.concepts[concept.ID] = concept
	}
}

func (r *OntologyRegistry) Concept(id string) (OntologyConcept, bool) {
	concept, ok := r.concepts[id]
	return concept, ok
}

func (r *OntologyRegistry) BuildFeatures() {
	degrees := map[string]int{}
	for id := range r.concepts {
		degrees[id] = 0
	}

	for id, concept := range r.concepts {
		for _, linkedID := range append(concept.ParentIDs, concept.RelatedIDs...) {
			degrees[id]++
			if _, ok := degrees[linkedID]; ok {
				degrees[linkedID]++
			}
		}
	}

	maxDegree := 0
	for _, degree := range degrees {
		if degree > maxDegree {
			maxDegree = degree
		}
	}

	for id, degree := range degrees {
		if maxDegree == 0 {
			r.centrality[id] = 0
			continue
		}
		r.centrality[id] = float64(degree) / float64(maxDegree)
	}
}

func (r *OntologyRegistry) Centrality(id string) float64 {
	return r.centrality[id]
}

func (r *OntologyRegistry) ExpandConceptIDs(ids []string) map[string]float64 {
	expanded := map[string]float64{}
	for _, id := range ids {
		expanded[id] = maxFloat(expanded[id], 1.0)

		concept, ok := r.concepts[id]
		if !ok {
			continue
		}
		for _, parentID := range concept.ParentIDs {
			expanded[parentID] = maxFloat(expanded[parentID], 0.65)
		}
		for _, relatedID := range concept.RelatedIDs {
			expanded[relatedID] = maxFloat(expanded[relatedID], 0.45)
		}
	}
	return expanded
}

// BuildScoreInput converts ontology and document values to the normalized
// features consumed by SearchScoringPolicy.
func (r *OntologyRegistry) BuildScoreInput(queryConceptIDs []string, textSimilarity float64, doc OntologyDocument) SearchScoreInput {
	docTags := setFromStrings(doc.OntologyTags)
	queryExpansion := r.ExpandConceptIDs(queryConceptIDs)

	var matchedConceptIDs []string
	var matchedWeight float64
	var possibleWeight float64
	for conceptID, weight := range queryExpansion {
		possibleWeight += weight
		if docTags[conceptID] {
			matchedConceptIDs = append(matchedConceptIDs, conceptID)
			matchedWeight += weight
		}
	}

	ontologyMatch := 0.0
	if possibleWeight > 0 {
		ontologyMatch = matchedWeight / possibleWeight
	}

	return SearchScoreInput{
		DocumentID:      doc.ID,
		TextSimilarity:  textSimilarity,
		OntologyMatch:   ontologyMatch,
		RiskPriority:    r.riskPriority(doc, matchedConceptIDs),
		GraphCentrality: r.graphCentrality(doc.OntologyTags, matchedConceptIDs),
	}
}

func (r *OntologyRegistry) riskPriority(doc OntologyDocument, matchedConceptIDs []string) float64 {
	risk := numberFromMetadata(doc.Metadata, "risk_level")
	for _, conceptID := range matchedConceptIDs {
		concept, ok := r.concepts[conceptID]
		if !ok {
			continue
		}
		risk = maxFloat(risk, numberFromMetadata(concept.Metadata, "risk_level"))
		risk = maxFloat(risk, numberFromMetadata(concept.Metadata, "urgency"))
		risk = maxFloat(risk, numberFromMetadata(concept.Metadata, "criticality"))
	}
	return clamp01(risk / 5.0)
}

func (r *OntologyRegistry) graphCentrality(docTags []string, matchedConceptIDs []string) float64 {
	ids := matchedConceptIDs
	if len(ids) == 0 {
		ids = docTags
	}

	if len(ids) == 0 {
		return 0
	}

	var total float64
	for _, id := range ids {
		total += r.Centrality(id)
	}
	return clamp01(total / float64(len(ids)))
}

func setFromStrings(values []string) map[string]bool {
	result := map[string]bool{}
	for _, value := range values {
		result[value] = true
	}
	return result
}

func numberFromMetadata(metadata map[string]interface{}, key string) float64 {
	if metadata == nil {
		return 0
	}

	switch value := metadata[key].(type) {
	case int:
		return float64(value)
	case int64:
		return float64(value)
	case float64:
		return value
	case float32:
		return float64(value)
	default:
		return 0
	}
}

func maxFloat(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func RankOntologyDocuments(registry *OntologyRegistry, policy SearchScoringPolicy, queryConceptIDs []string, textSimilarities map[string]float64, docs []OntologyDocument) ([]SearchScoreResult, error) {
	inputs := make([]SearchScoreInput, 0, len(docs))
	for _, doc := range docs {
		inputs = append(inputs, registry.BuildScoreInput(queryConceptIDs, textSimilarities[doc.ID], doc))
	}

	results, err := policy.Rank(inputs)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].DocumentID < results[j].DocumentID
		}
		return results[i].Score > results[j].Score
	})

	return results, nil
}
