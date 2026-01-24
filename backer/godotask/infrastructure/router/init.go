package router

import (
	"time"
	"github.com/godotask/interface/http/controller"
	"github.com/godotask/usecase"
	"github.com/godotask/infrastructure/security"
	"github.com/godotask/infrastructure/db/repository"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/interface/http/middleware"
)

func Init() {
	userRepo := repository.NewGormUserRepository(model.DB)
	passwordSvc := security.NewBcryptPasswordService()
	tokenSvc := security.NewJWTService(
		[]byte("secret"),
		24*time.Hour,
	)

	authUsecase := usecase.NewAuthUsecase(
		userRepo,
		passwordSvc,
		tokenSvc,
	)

	authController = controller.NewAuthController(authUsecase)

	authMiddleware = middleware.AuthMiddleware(tokenSvc)

	router = setupRouter()
}
