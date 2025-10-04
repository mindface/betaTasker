package task

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
)

// AddTask: POST /api/task
func (ctl *TaskController) AddTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
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
	
	// 必須フィールドのバリデーション
	if task.Title == "" {
		appErr := errors.NewAppError(
			errors.VAL_MISSING_FIELD,
			errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
			"Missing required field: title",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	if err := ctl.Service.CreateTask(&task); err != nil {
		var appErr *errors.AppError
		
		// エラー内容に応じた適切なエラーコードを設定
		errMsg := err.Error()
		if strings.Contains(errMsg, "duplicate") || strings.Contains(errMsg, "UNIQUE constraint") {
			appErr = errors.NewAppError(
				errors.VAL_DUPLICATE_ENTRY,
				errors.GetErrorMessage(errors.VAL_DUPLICATE_ENTRY),
				"Note: Target task with the same title already exists",
			)
		} else if strings.Contains(errMsg, "connection refused") || strings.Contains(errMsg, "no such host") {
			appErr = errors.NewAppError(
				errors.DB_CONNECTION_FAILED,
				errors.GetErrorMessage(errors.DB_CONNECTION_FAILED),
				"",
			)
		} else {
			appErr = errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				"",
			)
		}

		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Task added",
		"task": task,
	})
}
