package service

import (
	"authjwt/internal/models"
	"authjwt/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	hashedPass, hashErr := HashPassword(user.Password)

	if hashErr != nil {
		return nil, hashErr
	}

	user.Password = hashedPass

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
