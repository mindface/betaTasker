package teaching_free_control

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetTeachingFreeControl: GET /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) GetTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
	teachingFreeControl, err := ctl.Service.GetTeachingFreeControlByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Teaching free control not found",
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
		"message": "Teaching free control retrieved",
		"teaching_free_control": teachingFreeControl,
	})
}
