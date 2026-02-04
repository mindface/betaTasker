package errors

import (
	"log"
	"github.com/gin-gonic/gin"
)

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

// CORSMiddleware はCORSエラーを適切に処理するミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 許可されたオリジンのチェック（本番環境では厳密に設定）
		allowedOrigins := []string{
			"http://localhost:3000",
			"http://localhost:8080",
		}

		isAllowed := false
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				isAllowed = true
				break
			}
		}

		if !isAllowed && origin != "" {
			appErr := NewAppError(
				AUTH_UNAUTHORIZED,
				GetErrorMessage(AUTH_UNAUTHORIZED),
				"許可されていないオリジンからのアクセスです",
			)
			
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			c.Abort()
			return
		}
	
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// RequestValidationMiddleware は基本的なリクエスト検証を行うミドルウェア
func RequestValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content-Typeの検証（POST、PUT、PATCHメソッドの場合）
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if contentType != "" && contentType != "application/json" {
				appErr := NewAppError(
					VAL_INVALID_FORMAT,
					GetErrorMessage(VAL_INVALID_FORMAT),
					"Content-Typeはapplication/jsonである必要があります",
				)
				
				c.JSON(appErr.HTTPStatus, gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
					"detail":  appErr.Detail,
				})
				c.Abort()
				return
			}
		}
		
		// リクエストサイズの検証（1MBを超える場合はエラー）
		if c.Request.ContentLength > 1024*1024 {
			appErr := NewAppError(
				VAL_CONSTRAINT_FAILED,
				GetErrorMessage(VAL_CONSTRAINT_FAILED),
				"リクエストサイズが1MBを超えています",
			)
			
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			c.Abort()
			return
		}
		
		c.Next()
	}
}

// RateLimitMiddleware は簡単なレート制限を実装するミドルウェア
func RateLimitMiddleware() gin.HandlerFunc {
	// 実際の実装では、Redisやメモリストアを使用して制限を管理
	return func(c *gin.Context) {
		// 簡略化のため、この例では常に通す
		// 実際の実装では、IPアドレスやユーザーIDベースでレート制限を実装
		
		// レート制限に引っかかった場合の例
		/*
		if isRateLimited(c.ClientIP()) {
			appErr := NewAppError(
				SYS_RATE_LIMIT_EXCEEDED,
				GetErrorMessage(SYS_RATE_LIMIT_EXCEEDED),
				"1分間のリクエスト数が上限を超えました",
			)
			
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			c.Abort()
			return
		}
		*/
		
		c.Next()
	}
}

// NotFoundMiddleware は存在しないルートへのアクセスを処理するミドルウェア
func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appErr := NewAppError(
			RES_NOT_FOUND,
			GetErrorMessage(RES_NOT_FOUND),
			"指定されたエンドポイントが見つかりません",
		)
		
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
	}
}

// LoggingMiddleware はリクエスト/レスポンスをログに記録するミドルウェア
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// エラーが発生した場合は詳細をログに記録
		if param.StatusCode >= 400 {
			log.Printf("[ERROR] %s %s %d - %s",
				param.Method,
				param.Path,
				param.StatusCode,
				param.ErrorMessage,
			)
		}
		
		return ""
	})
}