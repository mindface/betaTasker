package user

import (
	// "net/http"
	// "time"
	// "github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	// "github.com/godotask/model"
	// "golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	} `json:"user"`
}

// func Login(c *gin.Context) {
// 	var req LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエスト"})
// 		return
// 	}

// 	var user model.User
// 	if err := model.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが間違っています"})
// 		return
// 	}

// 	if !user.IsActive {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "アカウントが無効化されています"})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが間違っています"})
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id":  user.ID,
// 		"username": user.Username,
// 		"email":    user.Email,
// 		"role":     user.Role,
// 		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte("your-secret-key"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
// 		return
// 	}

// 	response := LoginResponse{
// 		Token: tokenString,
// 		User: struct {
// 			ID       uint   `json:"id"`
// 			Username string `json:"username"`
// 			Email    string `json:"email"`
// 			Role     string `json:"role"`
// 		}{
// 			ID:       user.ID,
// 			Username: user.Username,
// 			Email:    user.Email,
// 			Role:     user.Role,
// 		},
// 	}

// 	c.JSON(http.StatusOK, response)
// } 