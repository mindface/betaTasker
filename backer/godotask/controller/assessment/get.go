package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

// GetAssessment: GET /api/assessment/:id
func (ctl *AssessmentController) GetAssessment(c *gin.Context) {
	id := c.Param("id")
	assessment, err := ctl.Service.GetAssessmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assessment": assessment})
}
