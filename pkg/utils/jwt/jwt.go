package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/services"
)

type jwtGenerator struct {
	secretKey string
}

func NewTokenGenerator(secretKey string) services.TokenGenerator {
	return &jwtGenerator{
		secretKey: secretKey,
	}
}

func (g *jwtGenerator) Generate(userID string, email string) (string, error) {

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(g.secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
