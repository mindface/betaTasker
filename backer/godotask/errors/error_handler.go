package errors

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	// AppErrorの場合
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	// 想定外エラー（panic防止）
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    SYS_INTERNAL_ERROR,
		"message": "internal server error",
	})
}

// ErrorHandlerMiddleware はアプリケーション全体のエラーハンドリングを行うミドルウェア
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err interface{}) {
		// パニックをキャッチしてAppErrorに変換
		appErr := NewAppError(
			SYS_INTERNAL_ERROR,
			GetErrorMessage(SYS_INTERNAL_ERROR),
			"システムでパニックが発生しました",
		)

		// ログに記録
		log.Printf("Panic recovered: %v", err)

		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
	})
}
