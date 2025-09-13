package process_optimization

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddProcessOptimization: POST /api/process_optimization
func (ctl *ProcessOptimizationController) AddProcessOptimization(c *gin.Context) {
	var processOptimization model.ProcessOptimization
	if err := c.ShouldBindJSON(&processOptimization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateProcessOptimization(&processOptimization); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add process optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Process optimization added", "process_optimization": processOptimization})
}
