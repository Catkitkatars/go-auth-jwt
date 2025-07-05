package service

import (
	"authjwt/internal/config"
	"authjwt/internal/models"
	"authjwt/internal/repositories"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
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
		return nil, fmt.Errorf("hashPassword: %v", hashErr)
	}

	user.Password = hashedPass

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("s.repo.Create: %v", hashErr)
	}

	return createdUser, nil
}

func (s *UserService) AuthUser(user *models.User) (map[string]string, error) {
	userRepo, err := s.repo.GetByEmail(user.Email)

	if err != nil {
		return nil, fmt.Errorf("s.repo.GetByEmail: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userRepo.Password), []byte(user.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	rToken, err := genToken(userRepo.ID, config.Cfg.JwtRefreshTime, config.Cfg.JwtRefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("refresh.genToken: %v", err)
	}
	aToken, err := genToken(userRepo.ID, config.Cfg.JwtAccessTime, config.Cfg.JwtAccessSecret)
	if err != nil {
		return nil, fmt.Errorf("access.genToken: %v", err)
	}

	r, err := s.repo.SaveTokenByUser(userRepo, rToken)

	if err != nil || r == false {
		return nil, fmt.Errorf("s.repo.SaveTokenByUser: %v", err)
	}

	return map[string]string{"access_token": aToken, "refresh_token": rToken}, nil
}

func genToken(userID int64, dur time.Duration, secret string) (string, error) {
	claims := jwt.MapClaims{
		"uid": userID,
		"exp": time.Now().Add(dur).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
