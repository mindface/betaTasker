package tool_matching

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMatchingHistory: GET /api/tool-matching/state-evaluation/:state_evaluation_id
func (ctrl *ToolMatchingController) GetMatchingHistory(c *gin.Context) {
	stateEvaluationID := c.Param("state_evaluation_id")
	if stateEvaluationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State evaluation ID is required"})
		return
	}

	results, err := ctrl.Service.GetMatchingHistory(stateEvaluationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   results,
		"count":  len(results),
	})
}

// GetAvailableTools: GET /api/tool-matching/available-tools
func (ctrl *ToolMatchingController) GetAvailableTools(c *gin.Context) {
	// This could be extended to get available tools from the service
	// For now, return a simple response indicating the endpoint is available
	
	domain := c.Query("domain")
	level := c.Query("level")

	response := gin.H{
		"status": "success",
		"message": "Available tools endpoint - implementation can be extended",
		"filters": gin.H{
			"domain": domain,
			"level":  level,
		},
	}

	// In a full implementation, you would:
	// 1. Get available robots from database
	// 2. Get available optimization models
	// 3. Filter by domain and level if provided
	// 4. Return structured tool information

	c.JSON(http.StatusOK, response)
}