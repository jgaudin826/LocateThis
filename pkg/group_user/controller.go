package group_user

import (
	"locate-this/config"
	"locate-this/database/dbmodel"
	"locate-this/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GroupUserConfig struct {
	*config.Config
}

func New(configuration *config.Config) *GroupUserConfig {
	return &GroupUserConfig{configuration}
}

// @Summary		Add user to group
// @Description	Add a user to a group
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"Group ID"
// @Param			request	body		models.AddUserRequest	true	"User ID"
// @Success		200		{object}	models.UserResponse
// @Router			/groups/{id}/users [post]
func (config *GroupUserConfig) PostUserToGroupHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupUserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	groupUserEntry := &dbmodel.GroupUserEntry{GroupID: req.GroupID, UserID: req.UserID}
	_, err := config.GroupUserEntryRepository.Create(groupUserEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to add user to group"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User added to group successfully"})
}

// @Summary		Get all group-user associations
// @Description	Retrieve all group-user associations
// @Tags			groups
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.GroupUserResponse
// @Router			/groups/users/all [get]
func (config *GroupUserConfig) GetAllGroupUserHandler(w http.ResponseWriter, r *http.Request) {
	groupUsers, err := config.GroupUserEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve group-user associations"})
		return
	}

	groupUserResponse := make([]models.GroupUserResponse, 0)
	for _, gu := range groupUsers {
		groupUserResponse = append(groupUserResponse, models.GroupUserResponse{
			UserID:  gu.UserID,
			GroupID: gu.GroupID,
		})
	}

	render.JSON(w, r, groupUserResponse)
}

// @Summary		Delete user from group
// @Description	Remove a user from a group
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Group ID"
// @Param			userID	path		int	true	"User ID"
// @Success		200		{string}	string	"Successfully removed user from group"
// @Router			/groups/{id}/users/{userID} [delete]
func (config *GroupUserConfig) DeleteUserFromGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid group ID"})
		return
	}

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	err = config.GroupUserEntryRepository.Delete(uint(userID), uint(groupID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to remove user from group"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User removed from group successfully"})
}
