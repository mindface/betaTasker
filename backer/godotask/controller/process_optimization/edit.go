package process_optimization

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditProcessOptimization: PUT /api/process_optimization/:id
func (ctl *ProcessOptimizationController) EditProcessOptimization(c *gin.Context) {
	id := c.Param("id")
	var processOptimization model.ProcessOptimization
	if err := c.ShouldBindJSON(&processOptimization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateProcessOptimization(id, &processOptimization); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit process optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Process optimization edited", "process_optimization": processOptimization})
}
