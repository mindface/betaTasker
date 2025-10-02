package qualitative_label

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteQualitativeLabel: DELETE /api/process_optimization/:id
func (ctl *QualitativeLabelController) DeleteQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteQualitativeLabel(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"Qualitative label not found",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Qualitative label deleted",
		"qualitative_label_id": id,
	})
}
