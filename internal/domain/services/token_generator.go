package services

type TokenGenerator interface {
	Generate(userID string, email string) (string, error)
}
