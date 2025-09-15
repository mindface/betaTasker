package phenomenological_framework

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// PhenomenologicalFramework: GET /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) GetPhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	phenomenologicalFramework, err := ctl.Service.GetPhenomenologicalFrameworkByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phenomenological framework not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"phenomenological_framework": phenomenologicalFramework})
}
