package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/controller/user" 
	"github.com/godotask/errors"
)

// ListTasks: GET /api/task
func (ctl *TaskController) ListTasks(c *gin.Context) {
   // コンテキストから userID を取得
    userID, ok := user.GetUserIDFromContext(c)
    if !ok {
			appErr := errors.NewAppError(
				errors.AUTH_UNAUTHORIZED,
				"Unauthorized",
				"User ID not found in context",
			)
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			return
    }

    // userID でフィルタしてタスク取得
    tasks, err := ctl.Service.ListTasksByUser(userID)
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
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
