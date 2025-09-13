package process_optimization

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetProcessOptimization: GET /api/process_optimization/:id
func (ctl *ProcessOptimizationController) GetProcessOptimization(c *gin.Context) {
	id := c.Param("id")
	processOptimization, err := ctl.Service.GetProcessOptimizationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Process optimization not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"process_optimization": processOptimization})
}
