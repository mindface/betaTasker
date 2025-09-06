package tool_matching

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// FindOptimalTools: POST /api/tool-matching
func (ctrl *ToolMatchingController) FindOptimalTools(c *gin.Context) {
	var req model.ToolMatchingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.Service.FindOptimalTools(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   result,
	})
}