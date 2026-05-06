package rag

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
