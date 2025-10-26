package usecase

import (
	"backend-hackathon/internal/domain"
	"backend-hackathon/internal/repository"
)

// UserUsecase defines the interface for user use case operations
type UserUsecase interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new user use case
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// CreateUser creates a new user
func (uc *userUsecase) CreateUser(user *domain.User) error {
	return uc.userRepo.Create(user)
}

// GetUserByID retrieves a user by ID
func (uc *userUsecase) GetUserByID(id uint) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

// GetUserByUsername retrieves a user by username
func (uc *userUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return uc.userRepo.GetByUsername(username)
}

// GetAllUsers retrieves all users
func (uc *userUsecase) GetAllUsers() ([]domain.User, error) {
	return uc.userRepo.GetAll()
}

// UpdateUser updates an existing user
func (uc *userUsecase) UpdateUser(user *domain.User) error {
	return uc.userRepo.Update(user)
}

// DeleteUser deletes a user by ID
func (uc *userUsecase) DeleteUser(id uint) error {
	return uc.userRepo.Delete(id)
}
