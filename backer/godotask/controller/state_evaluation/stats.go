package state_evaluation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLevelProgression: GET /api/state-evaluations/user/:user_id/progression
func (ctrl *StateEvaluationController) GetLevelProgression(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	progression, err := ctrl.Service.GetLevelProgression(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   progression,
	})
}

// GetEvaluationStats: GET /api/state-evaluations/stats
func (ctrl *StateEvaluationController) GetEvaluationStats(c *gin.Context) {
	userID := c.Query("user_id")
	
	// Get recent evaluations for statistics
	limit := 10
	if userID != "" {
		evaluations, err := ctrl.Service.GetEvaluationHistory(userID, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Calculate basic statistics
		if len(evaluations) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": gin.H{
					"total_evaluations": 0,
					"average_score":     0.0,
					"improvement_trend": "No data available",
				},
			})
			return
		}

		totalScore := 0.0
		completedCount := 0
		for _, eval := range evaluations {
			if eval.Status == "completed" {
				totalScore += eval.EvaluationScore
				completedCount++
			}
		}

		averageScore := 0.0
		if completedCount > 0 {
			averageScore = totalScore / float64(completedCount)
		}

		// Simple trend analysis
		trend := "Stable"
		if len(evaluations) >= 2 && evaluations[0].Status == "completed" && evaluations[1].Status == "completed" {
			if evaluations[0].EvaluationScore > evaluations[1].EvaluationScore {
				trend = "Improving"
			} else if evaluations[0].EvaluationScore < evaluations[1].EvaluationScore {
				trend = "Declining"
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"total_evaluations": len(evaluations),
				"completed_evaluations": completedCount,
				"average_score":     averageScore,
				"improvement_trend": trend,
				"latest_score":      evaluations[0].EvaluationScore,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required for statistics"})
	}
}