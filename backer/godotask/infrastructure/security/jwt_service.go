package security

import (
	"time"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/godotask/domain/entity"
	"github.com/godotask/domain/auth"
)

type JWTService struct {
	secret     []byte
	expiration time.Duration
}

func NewJWTService(secret []byte, exp time.Duration) auth.TokenService {
	return &JWTService{
		secret:     secret,
		expiration: exp,
	}
}

func (s *JWTService) Generate(user *entity.User) (string, error) {
	claims := entity.AuthClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.expiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}


func (s *JWTService) Parse(tokenString string) (*entity.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&entity.AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return s.secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.AuthClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

