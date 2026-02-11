package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/godotask/domain/entity"
	"github.com/godotask/domain/auth"
	"github.com/godotask/errors"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware(tokenSvc auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := GetClaimsFromAuthorizationHeader(c, tokenSvc)
		if err != nil {
			appErr := errors.NewAppError(
				errors.AUTH_UNAUTHORIZED,
				errors.GetErrorMessage(errors.AUTH_UNAUTHORIZED),
				err.Error() + " | Failed to parse auth token",
			)
			c.AbortWithStatusJSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			return
		}

		// context に詰める（controller / usecase で利用）
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
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
			appErr := errors.NewAppError(
				errors.AUTH_UNAUTHORIZED,
				errors.GetErrorMessage(errors.AUTH_UNAUTHORIZED),
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
				appErr := errors.NewAppError(
					errors.VAL_INVALID_FORMAT,
					errors.GetErrorMessage(errors.VAL_INVALID_FORMAT),
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
			appErr := errors.NewAppError(
				errors.VAL_CONSTRAINT_FAILED,
				errors.GetErrorMessage(errors.VAL_CONSTRAINT_FAILED),
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
			appErr := errors.NewAppError(
				errors.SYS_RATE_LIMIT_EXCEEDED,
				errors.GetErrorMessage(errors.SYS_RATE_LIMIT_EXCEEDED),
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
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
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


func GetClaimsFromAuthorizationHeader(
	c *gin.Context,
	tokenSvc auth.TokenService,
) (*entity.AuthClaims, error) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.ErrAuthTokenInvalid
	}

	const bearer = "Bearer "
	if !strings.HasPrefix(authHeader, bearer) {
		return nil, errors.ErrAuthTokenInvalid
	}

	token := strings.TrimPrefix(authHeader, bearer)
	return tokenSvc.Parse(token)
}
