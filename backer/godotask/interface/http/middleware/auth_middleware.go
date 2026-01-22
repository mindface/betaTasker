package middleware

import (
	"net/http"
	"strings"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/godotask/domain/entity"
	"github.com/godotask/domain/service"
)

func AuthMiddleware(tokenSvc service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := GetClaimsFromAuthorizationHeader(c, tokenSvc)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
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

func GetClaimsFromAuthorizationHeader(
	c *gin.Context,
	tokenSvc service.TokenService,
) (*entity.AuthClaims, error) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header required")
	}

	const bearer = "Bearer "
	if !strings.HasPrefix(authHeader, bearer) {
		return nil, errors.New("invalid authorization header")
	}

	token := strings.TrimPrefix(authHeader, bearer)
	return tokenSvc.Parse(token)
}
