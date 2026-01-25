package teaching_free_control

import (
	"net/http"

	"github.com/godotask/infrastructure/db/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// EditTeachingFreeControl: PUT /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) EditTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
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
	if err := ctl.Service.UpdateTeachingFreeControl(id, &teachingFreeControl); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to edit teaching free control",
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
		"message": "Teaching free control edited",
		"teaching_free_control": teachingFreeControl,
	})
}
