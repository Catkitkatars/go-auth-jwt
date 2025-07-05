package service

import (
	"authjwt/internal/models"
	"authjwt/internal/repositories"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	hashedPass, hashErr := hashPassword(user.Password)

	if hashErr != nil {
		return nil, fmt.Errorf("Srv.RegisterUser.hashPassword: %v", hashErr)
	}

	user.Password = hashedPass

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("s.repo.Create: %v", hashErr)
	}

	return createdUser, nil
}

func (s *UserService) AuthUser(user *models.User) (string, error) {
	hPass, err := hashPassword(user.Password)

	if err != nil {
		return "", fmt.Errorf("s.AuthUser: %v", err)
	}

	userRepo, err := s.repo.GetByEmail(user)

	if err != nil {
		return "", fmt.Errorf("s.repo.GetByEmail: %v", err)
	}

	if userRepo.Password != hPass {
		return "", fmt.Errorf("passwords do not match")
	}

	return "token", nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
