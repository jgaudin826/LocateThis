package models

import (
	"errors"
	"net/http"
)

type GroupLocationRequest struct {
	GroupID              uint `json:"group_id"`
	LocationID           uint `json:"location_id"`
	IsVisibleCoordinates bool `json:"is_visible_coordinates"`
}

func (req *GroupLocationRequest) Bind(r *http.Request) error {
	if req == nil {
		return errors.New("empty request")
	} else if req.GroupID == 0 {
		return errors.New("group_id is required")
	} else if req.LocationID == 0 {
		return errors.New("location_id is required")
	}
	return nil
}

type GroupLocationResponse struct {
	GroupID              uint `json:"group_id"`
	LocationID           uint `json:"location_id"`
	IsVisibleCoordinates bool `json:"is_visible_coordinates"`
}
