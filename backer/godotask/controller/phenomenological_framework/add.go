package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddPhenomenologicalFramework: POST /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) AddPhenomenologicalFramework(c *gin.Context) {
	var phenomenologicalFramework model.PhenomenologicalFramework
	if err := c.ShouldBindJSON(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreatePhenomenologicalFramework(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add phenomenological framework"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phenomenological framework added", "phenomenological_framework": phenomenologicalFramework})
}
