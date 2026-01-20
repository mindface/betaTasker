package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/usecase"
)

type AuthController struct {
	usecase *usecase.AuthUsecase
}

var Auth *AuthController

func InitAuthController(u *usecase.AuthUsecase) {
	Auth = &AuthController{usecase: u}
}

func NewAuthController(u *usecase.AuthUsecase) *AuthController {
	return &AuthController{usecase: u}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := c.usecase.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Logoutの実装についての評価を考慮する
func (c *AuthController) Logout(ctx *gin.Context) {
	// JWT方式ではサーバ側で破棄する状態は持たない
	// クライアント側でトークンを削除させる

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	token, user, err := c.usecase.Register(
		req.Username,
		req.Email,
		req.Role,
		req.Password,
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}

