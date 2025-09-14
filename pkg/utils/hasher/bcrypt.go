package hasher

import (
	"github.com/hsyntzgl/to-doList-Go/internal/domain/services"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct{}

func NewBcryptHasher() services.PasswordHasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Hash(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (h *bcryptHasher) Verify(password, passwordHash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err == nil {
		return true, nil
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return false, err
}
