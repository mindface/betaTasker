package teaching_free_control

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteTeachingFreeControl: DELETE /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) DeleteTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteTeachingFreeControl(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete teaching free control",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"message": "Teaching free control deleted",
		"teaching_free_control_id": id,
	})
}
