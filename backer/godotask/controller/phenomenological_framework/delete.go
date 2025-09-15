package phenomenological_framework

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeletePhenomenologicalFramework: DELETE /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) DeletePhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeletePhenomenologicalFramework(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete phenomenological framework"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phenomenological framework deleted"})
}
