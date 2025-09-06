package state_evaluation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvaluation: GET /api/state-evaluations/:id
func (ctrl *StateEvaluationController) GetEvaluation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Evaluation ID is required"})
		return
	}

	evaluation, err := ctrl.Service.GetEvaluationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   evaluation,
	})
}