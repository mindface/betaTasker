package usecase

import (
	"errors"

	"github.com/godotask/domain/entity"
	domainRepo "github.com/godotask/domain/repository"
	domainService "github.com/godotask/domain/auth"
)

type AuthUsecase struct {
	userRepo domainRepo.UserRepository
	password domainService.PasswordService
	token    domainService.TokenService
}

func NewAuthUsecase(
	userRepo domainRepo.UserRepository,
	password domainService.PasswordService,
	token domainService.TokenService,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		password: password,
		token:    token,
	}
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := u.password.Compare(user.PasswordHash, password); err != nil {
		return "", errors.New("invalid email or password")
	}

	return u.token.Generate(user)
}

func (u *AuthUsecase) Register(
	username, email, role, password string,
) (string, *entity.User, error) {

	hashed, err := u.password.Hash(password)
	if err != nil {
		return "", nil, err
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
