package service

import "github.com/godotask/domain/entity"

type TokenService interface {
	Generate(user *entity.User) (string, error)
	Parse(token string) (*entity.AuthClaims, error)
}
