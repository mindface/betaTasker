package rag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	ragengine "github.com/godotask/rag"
)

type OntologySearchRequest struct {
	Policy           *ragengine.SearchScoringPolicy `json:"policy"`
	QueryConceptIDs  []string                       `json:"query_concept_ids" binding:"required"`
	Concepts         []ragengine.OntologyConcept    `json:"concepts" binding:"required"`
	Documents        []ragengine.OntologyDocument   `json:"documents" binding:"required"`
	TextSimilarities map[string]float64             `json:"text_similarities"`
}

// ListOntologySearchPolicies: GET /api/rag/ontology/policies
func (ctl *OntologySearchController) ListOntologySearchPolicies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ontology search scoring policies retrieved",
		"policies": gin.H{
			"default": ragengine.DefaultSearchScoringPolicy(),
			"medical": ragengine.MedicalSearchScoringPolicy(),
		},
		"score_formula": "score = alpha * text_similarity + beta * ontology_match + gamma * risk_priority + delta * graph_centrality",
	})
}

// RankOntologySearch: POST /api/rag/ontology/rank
func (ctl *OntologySearchController) RankOntologySearch(c *gin.Context) {
	var req OntologySearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid ontology search request",
			"detail":  err.Error(),
		})
		return
	}

	policy := ragengine.DefaultSearchScoringPolicy()
	if req.Policy != nil {
		policy = *req.Policy
	}

	registry := ragengine.NewOntologyRegistry()
	registry.RegisterMany(req.Concepts)
	registry.BuildFeatures()

	results, err := ragengine.RankOntologyDocuments(
		registry,
		policy,
		req.QueryConceptIDs,
		req.TextSimilarities,
		req.Documents,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed to rank ontology search results",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ontology search results ranked",
		"policy":  policy,
		"results": results,
	})
}

// StoreSeedOntologyEmbeddings: POST /api/rag/ontology/seed/embeddings
func (ctl *OntologySearchController) StoreSeedOntologyEmbeddings(c *gin.Context) {
	var req ragengine.SeedDocumentOptions
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid seed embedding request",
			"detail":  err.Error(),
		})
		return
	}

	documents, err := ragengine.BuildSeedOntologyDocuments(model.DB, req.Sources)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to build seed ontology documents",
			"detail":  err.Error(),
		})
		return
	}

	storedCount, err := ragengine.StoreSeedDocumentEmbeddings(model.DB, documents, req.EmbeddingDimension)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to store seed ontology embeddings",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "seed ontology embeddings stored",
		"result": ragengine.SeedEmbeddingResult{
			StoredCount: storedCount,
			Documents:   documents,
		},
	})
}

// RankSeedOntologySearch: POST /api/rag/ontology/seed/rank
func (ctl *OntologySearchController) RankSeedOntologySearch(c *gin.Context) {
	var req ragengine.SeedOntologyRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid seed ontology rank request",
			"detail":  err.Error(),
		})
		return
	}

	policy := ragengine.DefaultSearchScoringPolicy()
	if req.Policy != nil {
		policy = *req.Policy
	}

	results, err := ragengine.RankSeedOntologyDocuments(model.DB, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to rank seed ontology documents",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "seed ontology search results ranked",
		"policy":  policy,
		"results": results,
	})
}
