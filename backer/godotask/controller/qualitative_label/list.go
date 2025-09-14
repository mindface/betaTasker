package qualitative_label

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListQualitativeLabels: GET /api/qualitative_label
func (ctl *QualitativeLabelController) ListQualitativeLabels(c *gin.Context) {
	qualitativeLabels, err := ctl.Service.ListQualitativeLabels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list qualitative labels"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qualitative_labels": qualitativeLabels})
}