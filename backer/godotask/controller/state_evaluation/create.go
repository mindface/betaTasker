package state_evaluation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// CreateEvaluation: POST /api/state-evaluations
func (ctrl *StateEvaluationController) CreateEvaluation(c *gin.Context) {
	var req model.EvaluationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	evaluation, err := ctrl.Service.EvaluateState(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   evaluation,
	})
}