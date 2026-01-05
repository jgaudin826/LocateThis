package models

import (
	"errors"
	"net/http"
)

type GroupRequest struct {
	Name string `json:"name"`
}

func (a *GroupRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("name must not be null")
	}
	return nil
}

type GroupResponse struct {
	ID        uint               `json:"group_id"`
	Name      string             `json:"name"`
	Users     []UserResponse     `json:"users"`
	Locations []LocationResponse `json:"locations"`
}
