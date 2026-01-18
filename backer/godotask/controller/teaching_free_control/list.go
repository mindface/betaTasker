package teaching_free_control

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/controller/user" 
	"github.com/godotask/errors"
)

// ListTeachingFreeControls: GET /api/teaching_free_control
func (ctl *TeachingFreeControlController) ListTeachingFreeControls(c *gin.Context) {
  userID, _ := user.GetUserIDFromContext(c)
	teachingFreeControls, err := ctl.Service.ListTeachingFreeControls(userID)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list teaching freeControls",
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
		"message": "Teaching free controls retrieved",
		"teaching_free_controls": teachingFreeControls,
	})
}