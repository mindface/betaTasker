package qualitative_label

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteQualitativeLabel: DELETE /api/process_optimization/:id
func (ctl *QualitativeLabelController) DeleteQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteQualitativeLabel(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete qualitative label"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Qualitative label deleted"})
}
