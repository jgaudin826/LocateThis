package models

import (
	"errors"
	"net/http"
)

type LocationRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserID    uint    `json:"user_id"`
}

func (a *LocationRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("name must not be null")
	}
	return nil
}

type LocationResponse struct {
	ID        uint    `json:"location_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserID    uint    `json:"user_id"`
}
