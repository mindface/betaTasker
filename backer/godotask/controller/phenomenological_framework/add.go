package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddProcessOptimization: POST /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) AddProcessOptimization(c *gin.Context) {
	var phenomenologicalFramework model.PhenomenologicalFrameworkController
	if err := c.ShouldBindJSON(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreatePhenomenologicalFramework(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add process optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Process optimization added", "phenomenological_framework": phenomenologicalFramework})
}
