package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/entities"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/repositories"
	"github.com/hsyntzgl/to-doList-Go/internal/domain/services"
)

var ErrEmailAllreadyExists = errors.New("Bu mail adresi zaten kullanılır")
var ErrInvalidCredentials = errors.New("Mail adresi veya şifre yanlış")
var ErrForbidden = errors.New("Bu işlemi yapmaya yetkiniz bulunmamaktadır")
var ErrUserNotFound = errors.New("Kullanıcı bulunamadı")
var ErrValidation = errors.New("Lütfen bilgilerinizi doğru formatta girin")

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*entities.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	UpdateUser(ctx context.Context, actorID, targetID string, req UpdateUser) error
	ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error
	Delete(ctx context.Context, actorID, targetID string) error
}

type userService struct {
	userRepo       repositories.UserRepository
	hasher         services.PasswordHasher
	tokenGenerator services.TokenGenerator
}

func NewUserService(repo repositories.UserRepository, hasher services.PasswordHasher, tokenGenerator services.TokenGenerator) UserService {
	return &userService{
		userRepo:       repo,
		hasher:         hasher,
		tokenGenerator: tokenGenerator,
	}
}

func (s *userService) Register(ctx context.Context, username, email, password string) (*entities.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, ErrEmailAllreadyExists
	}

	hashedPassword, err := s.hasher.Hash(password)

	if err != nil {
		return nil, err
	}

	newUser := &entities.User{
		ID:           uuid.New().String(),
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	match, err := s.hasher.Verify(password, user.PasswordHash)

	if err != nil {
		return "", err
	}

	if !match {
		return "", ErrInvalidCredentials
	}

	token, err := s.tokenGenerator.Generate(user.ID, user.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) UpdateUser(ctx context.Context, actorID, targetID string, req UpdateUser) error {
	user, err := s.userRepo.GetByID(ctx, actorID)

	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrForbidden
		}
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	if actorID != targetID {
		return ErrForbidden
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	username := strings.TrimSpace(req.Username)

	if email == "" || username == "" {
		return ErrValidation
	}

	checkEmail, err := s.userRepo.GetByEmail(ctx, email)

	if err != nil {
		if !errors.Is(err, repositories.ErrNotFound) {
			return err
		}
	}

	if checkEmail != nil {
		return ErrEmailAllreadyExists
	}

	user.Email = req.Email
	user.Username = req.Username

	if err := s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
func (s *userService) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrInvalidCredentials
		}
		return err
	}

	result, err := s.hasher.Verify(oldPassword, user.PasswordHash)

	if err != nil {
		return err
	}

	if !result {
		return ErrInvalidCredentials
	}

	newPasswordHash, err := s.hasher.Hash(newPassword)

	if err != nil {
		return err
	}

	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now().UTC()

	if err := s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
func (s *userService) Delete(ctx context.Context, actorID, targetID string) error {
	user, err := s.userRepo.GetByID(ctx, actorID)

	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	if actorID != targetID {
		return ErrForbidden
	}

	err = s.userRepo.Delete(ctx, user.ID)

	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}
