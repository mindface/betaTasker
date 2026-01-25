package auth

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{}

func (b *BcryptService) Compare(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
