package models

import (
	"errors"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *LoginRequest) Bind(r *http.Request) error {
	if a.Email == "" {
		return errors.New("email must not be null")
	} else if a.Password == "" {
		return errors.New("password must not be null")
	}
	return nil
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (t *TokenRequest) Bind(r *http.Request) error {
	if t.RefreshToken == "" {
		return errors.New("refresh_token is required")
	}
	return nil
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
