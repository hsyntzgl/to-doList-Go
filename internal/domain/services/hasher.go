package services

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password, passwordHash string) (bool, error)
}
