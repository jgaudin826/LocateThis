package models

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"pseudo"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.Email == "" {
		return errors.New("email must not be null")
	} else if a.Password == "" {
		return errors.New("password must not be null")
	} else if a.Username == "" {
		return errors.New("pseudo must not be null")
	}
	return nil
}

type UserResponse struct {
	ID       uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"pseudo"`
}
