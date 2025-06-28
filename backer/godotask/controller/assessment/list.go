package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ListAssessments: GET /api/assessment
func (ctl *AssessmentController) ListAssessments(c *gin.Context) {
	assessments, err := ctl.Service.ListAssessments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list assessments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assessments": assessments})
}
