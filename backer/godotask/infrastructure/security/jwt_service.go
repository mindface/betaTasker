package security

import (
	"time"
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"

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
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return s.secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
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
