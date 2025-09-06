package tool_matching

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// GetRecommendations: POST /api/tool-matching/recommendations
func (ctrl *ToolMatchingController) GetRecommendations(c *gin.Context) {
	var req struct {
		StateEvaluationID string                 `json:"state_evaluation_id" binding:"required"`
		Requirements      map[string]interface{} `json:"requirements"`
		UserPreferences   map[string]interface{} `json:"user_preferences"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to tool matching request
	toolReq := &model.ToolMatchingRequest{
		StateEvaluationID: req.StateEvaluationID,
		Requirements:      req.Requirements,
		Preferences:       req.UserPreferences,
	}

	result, err := ctrl.Service.FindOptimalTools(toolReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return only the recommendations part for this endpoint
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"matching_score":       result.MatchingScore,
			"recommendations":      result.Recommendations,
			"expected_performance": result.ExpectedPerformance,
			"optimal_parameters":   result.Parameters,
		},
	})
}