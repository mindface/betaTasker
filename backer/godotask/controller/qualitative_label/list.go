package qualitative_label

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListQualitativeLabels: GET /api/qualitative_label
func (ctl *QualitativeLabelController) ListQualitativeLabels(c *gin.Context) {
	qualitativeLabels, err := ctl.Service.ListQualitativeLabels()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
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
		"message": "Qualitative labels retrieved",
		"qualitative_labels": qualitativeLabels,
	})
}