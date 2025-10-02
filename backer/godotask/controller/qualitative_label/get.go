package qualitative_label

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetQualitativeLabel: GET /api/qualitative_label/:id
func (ctl *QualitativeLabelController) GetQualitativeLabel(c *gin.Context) {
	id := c.Param("id")
	qualitativeLabel, err := ctl.Service.GetQualitativeLabelByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"Memory not found",
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
		"message": "Qualitative label retrieved",
		"qualitative_label": qualitativeLabel,
	})
}
