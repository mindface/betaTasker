package task

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
)

// AddTaskWithErrorCode: エラーコードを使用したタスク追加の実装例
func (ctl *TaskController) AddTaskWithErrorCode(c *gin.Context) {
	var task model.Task
	
	// リクエストボディのバインドとバリデーション
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

	// 必須フィールドのチェック
	if task.Title == "" {
		appErr := errors.NewAppError(
			errors.VAL_MISSING_FIELD,
			errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
			"タイトルは必須項目です",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// タスクの作成
	if err := ctl.Service.CreateTask(&task); err != nil {
		// エラーの種類に応じて適切なエラーコードを設定
		var appErr *errors.AppError

		// 重複エラーの例
		if isDuplicateError(err) {
			appErr = errors.NewAppError(
				errors.VAL_DUPLICATE_ENTRY,
				errors.GetErrorMessage(errors.VAL_DUPLICATE_ENTRY),
				"同じタイトルのタスクが既に存在します",
			)
		} else if isDBConnectionError(err) {
			appErr = errors.NewAppError(
				errors.DB_CONNECTION_FAILED,
				errors.GetErrorMessage(errors.DB_CONNECTION_FAILED),
				"",
			)
		} else {
			appErr = errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				err.Error(),
			)
		}
		
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// 成功レスポンス
	c.JSON(200, gin.H{
		"success": true,
		"message": "タスクが正常に作成されました",
		"data": gin.H{
			"task": task,
		},
	})
}

// エラー判定のヘルパー関数（実際の実装では適切なエラー判定を行う）
func isDuplicateError(err error) bool {
	// データベースエラーメッセージなどから重複エラーを判定
	// 例: strings.Contains(err.Error(), "duplicate key")
	return false
}

func isDBConnectionError(err error) bool {
	// データベース接続エラーを判定
	// 例: strings.Contains(err.Error(), "connection refused")
	return false
}