package state_evaluation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetEvaluationHistory: GET /api/state-evaluations/user/:user_id
func (ctrl *StateEvaluationController) GetEvaluationHistory(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	limitStr := c.Query("limit")
	limit := 0
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	evaluations, err := ctrl.Service.GetEvaluationHistory(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   evaluations,
		"count":  len(evaluations),
	})
}