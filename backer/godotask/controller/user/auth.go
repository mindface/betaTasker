package user

import (
	"net/http"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/godotask/model"
	"github.com/godotask/errors"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var tokenBlacklist = make(map[string]struct{})

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Bearerトークン形式のチェック
		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with 'Bearer '"})
			c.Abort()
			return
		}
		tokenString := authHeader[len(bearerPrefix):]

		if _, blacklisted := tokenBlacklist[authHeader]; blacklisted {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
		})

		claims, ok := token.Claims.(*Claims)
		if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
						"code":    "AUTH_INVALID_CLAIMS",
						"message": "Invalid token claims",
						"detail":  "Failed to parse claims",
				})
				c.Abort()
				return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// トークンが有効であれば、ユーザー情報とロールをコンテキストに設定
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// GetUserIDFromContext: コンテキストから userID を安全に取得
func GetUserIDFromContext(c *gin.Context) (uint, bool) {
    userID, exists := c.Get("user_id")
    if !exists {
        return 0, false
    }
    uid, ok := userID.(uint)
    return uid, ok
}

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Role string `json:"role" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := model.User {
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := model.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// JWTトークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		// 必要に応じて他のクレームも
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
    "token": tokenString,
    "user": user,
	})
}

func Logout(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
    if tokenString != "" {
        tokenBlacklist[tokenString] = struct{}{}
    }
    c.JSON(200, gin.H{"message": "Logged out successfully"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
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
  fmt.Printf("Login input: %+v\n", input)

	var user model.User
	if err := model.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			appErr := errors.NewAppError(
				errors.RES_NOT_FOUND,
				errors.GetErrorMessage(errors.RES_NOT_FOUND),
				err.Error() + " | Invalid email or password",
			)
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
			return
		}
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
	fmt.Printf("Fetched user: %+v\n", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		appErr := errors.NewAppError(
			errors.AUTH_UNAUTHORIZED,
			"Invalid email or password",
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// expirationTime := time.Now().Add(24 * time.Hour)
	// expirationTime := time.Now().Add(10 * time.Minute)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			"Failed to generate token: "+err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	fmt.Printf("Generated token2: %s\n", tokenString)
	fmt.Printf("FetchGenerateded user: %+v\n", user)
	response := LoginResponse{
		Token: tokenString,
		User: struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	}

	c.JSON(http.StatusOK, response)
}

func Profile(c *gin.Context) {
	userName, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user model.User
	if err := model.DB.Where("username = ?", userName).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
			"isActive":  user.IsActive,
			"role":      user.Role,
		},
	})
}