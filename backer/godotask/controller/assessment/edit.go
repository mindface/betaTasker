package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// EditAssessment: PUT /api/assessment/:id
func (ctl *AssessmentController) EditAssessment(c *gin.Context) {
	id := c.Param("id")
	var assessment model.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateAssessment(id, &assessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit assessment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Assessment edited", "assessment": assessment})
}
