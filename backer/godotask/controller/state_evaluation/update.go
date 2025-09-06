package state_evaluation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateEvaluationResults: PUT /api/state-evaluations/:id/results
func (ctrl *StateEvaluationController) UpdateEvaluationResults(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Evaluation ID is required"})
		return
	}

	var req struct {
		Results          map[string]interface{} `json:"results" binding:"required"`
		LearnedKnowledge string                 `json:"learned_knowledge"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Service.UpdateEvaluationResults(id, req.Results, req.LearnedKnowledge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Evaluation results updated successfully",
	})
}