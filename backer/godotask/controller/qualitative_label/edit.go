package qualitative_label

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditQualitativeLabel: PUT /api/qualitative_label/:id
func (ctl *QualitativeLabelController) EditQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	var qualitativeLabel model.QualitativeLabel
	if err := c.ShouldBindJSON(&qualitativeLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateQualitativeLabel(id, &qualitativeLabel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit qualitative label"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Qualitative label edited", "qualitative_label": qualitativeLabel})
}
