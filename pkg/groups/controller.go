package groups

import (
	"fmt"
	"locate-this/config"
	"locate-this/database/dbmodel"
	"locate-this/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GroupConfig struct {
	*config.Config
}

func New(configuration *config.Config) *GroupConfig {
	return &GroupConfig{configuration}
}

// @Summary		Create a new group
// @Description	Create a new group entry
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			request	body		models.GroupRequest	true	"Group data"
// @Success		200		{object}	models.GroupResponse
// @Router			/groups [post]
func (config *GroupConfig) PostGroupHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	groupEntry := &dbmodel.GroupEntry{Name: req.Name}
	res, err := config.GroupEntryRepository.Create(groupEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create group"})
		return
	}

	groupResponse := &models.GroupResponse{ID: res.ID, Name: res.Name}
	render.JSON(w, r, groupResponse)
}

// @Summary		Get all groups
// @Description	Retrieve a list of all groups
// @Tags			groups
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.GroupResponse
// @Router			/groups [get]
func (config *GroupConfig) GetAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.GroupEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve groups"})
		return
	}

	groupsResponse := make([]models.GroupResponse, 0)
	for _, group := range entries {
		groupsResponse = append(groupsResponse, models.GroupResponse{
			ID:   group.ID,
			Name: group.Name,
		})
	}

	render.JSON(w, r, groupsResponse)
}

// @Summary		Get group by ID
// @Description	Retrieve a group by its ID
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Group ID"
// @Success		200	{object}	models.GroupResponse
// @Router			/groups/{id} [get]
func (config *GroupConfig) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	entry, err := config.GroupEntryRepository.FindByID(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve group"})
		return
	}
	groupResponse := &models.GroupResponse{ID: entry.ID, Name: entry.Name, Users: entry.Users, Locations: entry.Locations}
	render.JSON(w, r, groupResponse)
}

// @Summary		Update a group
// @Description	Update an existing group entry
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Group ID"
// @Param			request	body		models.GroupRequest	true	"Group data"
// @Success		200		{object}	models.GroupResponse
// @Router			/groups/{id} [put]
func (config *GroupConfig) PutGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}

	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	groupEntry := &dbmodel.GroupEntry{Name: req.Name}
	updated, err := config.GroupEntryRepository.Update(groupEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update group"})
		return
	}

	groupResponse := &models.GroupResponse{ID: uint(id), Name: updated.Name}
	render.JSON(w, r, groupResponse)
}

// @Summary		Delete a group
// @Description	Delete a group by its ID
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Group ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Router			/groups/{id} [delete]
func (config *GroupConfig) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.GroupEntryRepository.Delete(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete group"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
