package teaching_free_control

import (
	"net/http"

	"github.com/godotask/infrastructure/db/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// AddTeachingFreeControl: POST /api/teaching_free_control
func (ctl *TeachingFreeControlController) AddTeachingFreeControl(c *gin.Context) {
	var teachingFreeControl model.TeachingFreeControl
	if err := c.ShouldBindJSON(&teachingFreeControl); err != nil {
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

	if err := ctl.Service.CreateTeachingFreeControl(&teachingFreeControl); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to add teaching free control",
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
		"message": "Teaching free control added",
		"teaching_free_control": teachingFreeControl,
	})
}
