package security

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/godotask/domain/service"
)

type BcryptPasswordService struct{}

func NewBcryptPasswordService() service.PasswordService {
	return &BcryptPasswordService{}
}

func (b *BcryptPasswordService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *BcryptPasswordService) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
