package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext" 
	"github.com/godotask/errors"
)

// ListTasks: GET /api/task
func (ctl *TaskController) ListTasks(c *gin.Context) {
	userID, _ := authcontext.UserID(c)
	// userID でフィルタしてタスク取得
	tasks, err := ctl.Service.ListTasks(userID)
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
		"message": "Tasks retrieved",
		"tasks": tasks,
	})
}