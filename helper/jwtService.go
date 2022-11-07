package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userId int, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

var SECRETE_KEY = "rahasia"

type jwtService struct{}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int, email string) (string, error) {
	ttl := 60 * time.Second

	claim := jwt.MapClaims{
		"id":    userId,
		"email": email,
		"exp":   time.Now().UTC().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SECRETE_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRETE_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
