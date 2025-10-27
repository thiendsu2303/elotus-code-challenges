package service

import (
	"backend-hackathon/internal/domain"
	"backend-hackathon/internal/repository"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password string) (*domain.User, error)
	Login(username, password string) (string, time.Time, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
	accessTTL time.Duration
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string, accessTTL time.Duration) AuthService {
	return &authService{userRepo: userRepo, jwtSecret: jwtSecret, accessTTL: accessTTL}
}

func (s *authService) Register(username, password string) (*domain.User, error) {
	existingUser, err := s.userRepo.GetByUsername(username)
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

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(username, password string) (string, time.Time, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil || user == nil {
		return "", time.Time{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", time.Time{}, errors.New("invalid credentials")
	}

	now := time.Now().UTC()
	exp := now.Add(s.accessTTL)

	claims := jwt.MapClaims{
		"sub": user.ID,
		"iat": now.Unix(),
		"exp": exp.Unix(),
		"iss": "backend-hackathon",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, exp, nil
}
