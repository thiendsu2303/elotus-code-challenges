package service

import (
    "backend-hackathon/internal/domain"
    "backend-hackathon/internal/repository"
    "errors"

    "golang.org/x/crypto/bcrypt"
)

// AuthService defines authentication related operations
type AuthService interface {
    Register(username, password string) (*domain.User, error)
}

type authService struct {
    userRepo repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo repository.UserRepository) AuthService {
    return &authService{userRepo: userRepo}
}

// Register creates a new user with hashed password after validating username uniqueness
func (s *authService) Register(username, password string) (*domain.User, error) {
    // Check if username already exists
    existingUser, err := s.userRepo.GetByUsername(username)
    if err == nil && existingUser != nil {
        return nil, errors.New("username already exists")
    }

    // Hash password
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &domain.User{
        Username:     username,
        PasswordHash: string(passwordHash),
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}