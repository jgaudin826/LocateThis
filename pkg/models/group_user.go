package models

import (
	"errors"
	"net/http"
)

type GroupUserRequest struct {
	GroupID uint `json:"group_id"`
	UserID  uint `json:"user_id"`
}

func (req *GroupUserRequest) Bind(r *http.Request) error {
	if req == nil {
		return errors.New("empty request")
	} else if req.UserID == 0 {
		return errors.New("user_id is required")
	} else if req.GroupID == 0 {
		return errors.New("group_id is required")
	}
	return nil
}

type GroupUserResponse struct {
	GroupID uint `json:"group_id"`
	UserID  uint `json:"user_id"`
}
