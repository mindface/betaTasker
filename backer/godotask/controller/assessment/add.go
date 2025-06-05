package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"github.com/godotask/service"
)

// AddAssessment: POST /api/assessment
func (ctl *AssessmentController) AddAssessment(c *gin.Context) {
	var assessment model.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateAssessment(&assessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add assessment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Assessment added", "assessment": assessment})
}
