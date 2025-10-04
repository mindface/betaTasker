package qualitative_label

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// EditQualitativeLabel: PUT /api/qualitative_label/:id
func (ctl *QualitativeLabelController) EditQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	var qualitativeLabel model.QualitativeLabel
	if err := c.ShouldBindJSON(&qualitativeLabel); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.UpdateQualitativeLabel(id, &qualitativeLabel); err != nil {
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
		"message": "Qualitative label edited",
		"qualitative_label": qualitativeLabel,
	})
}
