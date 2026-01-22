package server

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/controller"
)

var (
	router         *gin.Engine
	authController *controller.AuthController
	authMiddleware gin.HandlerFunc
)
