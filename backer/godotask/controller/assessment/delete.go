package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// DeleteAssessment: DELETE /api/assessment/:id
func (ctl *AssessmentController) DeleteAssessment(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteAssessment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assessment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Assessment deleted"})
}
