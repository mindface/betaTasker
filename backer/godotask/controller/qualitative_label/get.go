package qualitative_label

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetQualitativeLabel: GET /api/qualitative_label/:id
func (ctl *QualitativeLabelController) GetQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	qualitativeLabel, err := ctl.Service.GetQualitativeLabelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Qualitative label not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qualitative_label": qualitativeLabel})
}
