package services

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("jwt_secret_key"))

type JWTService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtKey)
	return t, err
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtKey, nil
	})
	return token, err
}
