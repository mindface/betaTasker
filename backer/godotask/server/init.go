package server

import (
	"time"
	"github.com/godotask/interface/http/controller"
	"github.com/godotask/usecase"
	"github.com/godotask/infrastructure/repository"
	"github.com/godotask/infrastructure/security"
	"github.com/godotask/model"
)

// authに対するDI対応をしている
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

	controller.InitAuthController(authUsecase)
}
