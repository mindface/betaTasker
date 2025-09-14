package qualitative_label

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddQualitativeLabel: POST /api/qualitative_label
func (ctl *QualitativeLabelController) AddQualitativeLabel(c *gin.Context) {
	var QualitativeLabel model.QualitativeLabel
	if err := c.ShouldBindJSON(&QualitativeLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateQualitativeLabel(&QualitativeLabel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add qualitative label"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Qualitative label added", "qualitative_label": QualitativeLabel})
}
