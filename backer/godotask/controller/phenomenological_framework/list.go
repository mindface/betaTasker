package phenomenological_framework

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListPhenomenologicalFrameworks: GET /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) ListPhenomenologicalFrameworks(c *gin.Context) {
	phenomenologicalFrameworks, err := ctl.Service.ListPhenomenologicalFrameworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list PhenomenologicalFrameworks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"PhenomenologicalFrameworks": phenomenologicalFrameworks})
}