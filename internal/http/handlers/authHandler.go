package handlers

import (
	"authjwt/internal/store"
	"database/sql"
	"net/http"
)

type AuthHandler struct {
	Store *sql.DB
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		Store: store.DB,
	}
}

func (a AuthHandler) SayHello(r *http.Request) (any, error) {
	return "Hi, from authJwt", nil
}
