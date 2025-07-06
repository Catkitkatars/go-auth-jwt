package handlers

import (
	"authjwt/internal/dto"
	model "authjwt/internal/models"
	"authjwt/internal/repositories"
	service "authjwt/internal/services"
	"authjwt/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	service.UserService
}

func NewAuthHandler() *AuthHandler {
	repo := repositories.NewUserRepo(store.DB)
	serv := service.NewUserService(*repo)
	return &AuthHandler{
		UserService: *serv,
	}
}

func (h AuthHandler) Registration(r *http.Request) (any, error) {
	uDto, err := h.toUserDto(r)
	if err != nil {
		return false, fmt.Errorf("h.Registration.toUserDto: %v", err)
	}

	user := model.User{
		Name:     uDto.Name,
		Email:    uDto.Email,
		Password: uDto.Password,
	}

	_, createErr := h.UserService.RegisterUser(&user)

	if createErr != nil {
		return nil, fmt.Errorf("h.UserService.RegisterUser: %v", createErr)
	}

	return map[string]bool{"success": true}, nil
}

func (h AuthHandler) Login(r *http.Request) (any, error) {
	uDto, err := h.toUserDto(r)
	if err != nil {
		return nil, fmt.Errorf("h.Login.toUserDto: %v", err)
	}

	user := model.User{
		Name:     uDto.Name,
		Email:    uDto.Email,
		Password: uDto.Password,
	}

	jwt, err := h.UserService.AuthUser(&user)

	if err != nil {
		return nil, fmt.Errorf("h.Login.UserService.AuthUser: %v", err)
	}

	return jwt, nil
}

func (h AuthHandler) SayHello(r *http.Request) (any, error) {
	return "Hi, from authJwt", nil
}

func (h AuthHandler) SayByeBye(r *http.Request) (any, error) {
	return "Bye-Bye, from authJwt", nil
}
func (h AuthHandler) SaySomeThing(r *http.Request) (any, error) {
	return "SomeThing, from authJwt", nil
}

func (h AuthHandler) toUserDto(r *http.Request) (*dto.UserDto, error) {
	var uDto dto.UserDto

	err := json.NewDecoder(r.Body).Decode(&uDto)
	if err != nil {
		return nil, err
	}

	return &uDto, nil
}
