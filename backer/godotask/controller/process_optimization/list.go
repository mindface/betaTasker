package process_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListProcessOptimizations: GET /api/process_optimization
func (ctl *ProcessOptimizationController) ListProcessOptimizations(c *gin.Context) {
	ProcessOptimizations, err := ctl.Service.ListProcessOptimizations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list ProcessOptimizations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ProcessOptimizations": ProcessOptimizations})
}