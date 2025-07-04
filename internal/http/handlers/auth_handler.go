package handlers

import (
	"authjwt/internal/dto"
	model "authjwt/internal/models"
	"authjwt/internal/repositories"
	service "authjwt/internal/services"
	"authjwt/internal/store"
	"encoding/json"
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

func (a AuthHandler) Registration(r *http.Request) (any, error) {
	var uDto dto.UserDto

	err := json.NewDecoder(r.Body).Decode(&uDto)
	if err != nil {
		return false, err
	}

	user := model.User{
		Name:     uDto.Name,
		Email:    uDto.Email,
		Password: uDto.Password,
	}

	_, createErr := a.UserService.RegisterUser(&user)

	if createErr != nil {
		return nil, err
	}

	return map[string]bool{"success": true}, nil
}

func (a AuthHandler) SayHello(r *http.Request) (any, error) {
	return "Hi, from authJwt", nil
}
