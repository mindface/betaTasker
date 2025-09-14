package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditPhenomenologicalFramework: PUT /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) EditPhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	var phenomenologicalFramework model.PhenomenologicalFramework
	if err := c.ShouldBindJSON(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdatePhenomenologicalFramework(id, &phenomenologicalFramework); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit phenomenological framework"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phenomenological framework edited", "phenomenological_framework": phenomenologicalFramework})
}
