package usecase

import (
	"errors"

	"github.com/godotask/domain/entity"

	domainRepo "github.com/godotask/domain/repository"
	domainService "github.com/godotask/domain/service"
)

type RegisterUsecase struct {
	userRepo domainRepo.UserRepository
	password domainService.PasswordService
	token    domainService.TokenService
}

func (u *RegisterUsecase) Register(
	username, email, role, password string,
) (string, *entity.User, error) {

	hashed, err := u.password.Hash(password)
	if err != nil {
		return "", nil, errors.New("failed to hash password")
	}

	user := &entity.User{
		Username:     username,
		Email:        email,
		Role:         role,
		PasswordHash: hashed,
	}

	if err := u.userRepo.Create(user); err != nil {
		return "", nil, err
	}

	token, err := u.token.Generate(user)

	return token, user, err
}
