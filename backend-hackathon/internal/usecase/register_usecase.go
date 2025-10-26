package usecase

import (
	"backend-hackathon/internal/domain"
	"backend-hackathon/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUsecase interface {
	Register(username, password string) (*domain.User, error)
}

type registerUsecase struct {
	userRepo repository.UserRepository
}

func NewRegisterUsecase(userRepo repository.UserRepository) RegisterUsecase {
	return &registerUsecase{userRepo: userRepo}
}

func (uc *registerUsecase) Register(username, password string) (*domain.User, error) {
	// Check if username already exists
	existingUser, err := uc.userRepo.GetByUsername(username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:     username,
		PasswordHash: string(passwordHash),
	}
	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
