package group_location

import (
	"locate-this/config"
	"locate-this/database/dbmodel"
	"locate-this/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GroupLocationConfig struct {
	*config.Config
}

func New(configuration *config.Config) *GroupLocationConfig {
	return &GroupLocationConfig{configuration}
}

// @Summary		Share location in group
// @Description	Add a location to a group
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id		path		int							true	"Group ID"
// @Param			request	body		models.AddLocationRequest	true	"Location ID and visibility"
// @Success		200		{object}	models.LocationResponse
// @Router			/groups/{id}/locations [post]
func (config *GroupLocationConfig) PostLocationToGroupHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupLocationRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	groupLocationEntry := &dbmodel.GroupLocationEntry{GroupID: req.GroupID, LocationID: req.LocationID, IsVisibleCoordinates: req.IsVisibleCoordinates}
	_, err := config.GroupLocationEntryRepository.Create(groupLocationEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to share location in group"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Location shared in group successfully"})
}

// @Summary		Get all group-location associations
// @Description	Retrieve all group-location associations
// @Tags			groups
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.GroupLocationResponse
// @Router			/groups/locations/all [get]
func (config *GroupLocationConfig) GetAllGroupLocationHandler(w http.ResponseWriter, r *http.Request) {
	groupLocations, err := config.GroupLocationEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve group-location associations"})
		return
	}

	groupLocationResponse := make([]models.GroupLocationResponse, 0)
	for _, gl := range groupLocations {
		groupLocationResponse = append(groupLocationResponse, models.GroupLocationResponse{
			GroupID:              gl.GroupID,
			LocationID:           gl.LocationID,
			IsVisibleCoordinates: gl.IsVisibleCoordinates,
		})
	}

	render.JSON(w, r, groupLocationResponse)
}

// @Summary		Update location sharing in group
// @Description	Update location visibility settings in a group
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id			path		int							true	"Group ID"
// @Param			locationID	path		int							true	"Location ID"
// @Param			request		body		models.AddLocationRequest	true	"Updated visibility settings"
// @Success		200			{object}	models.LocationResponse
// @Router			/groups/{id}/locations/{locationID} [put]
func (config *GroupLocationConfig) PutLocationInGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid group ID"})
		return
	}

	locationID, err := strconv.Atoi(chi.URLParam(r, "locationID"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid location ID"})
		return
	}

	req := &models.GroupLocationRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	groupLocationEntry := &dbmodel.GroupLocationEntry{
		GroupID:              uint(groupID),
		LocationID:           uint(locationID),
		IsVisibleCoordinates: req.IsVisibleCoordinates,
	}
	_, err = config.GroupLocationEntryRepository.Update(groupLocationEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update location in group"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Location updated in group successfully"})
}

// @Summary		Delete location from group
// @Description	Remove a location from a group
// @Tags			groups
// @Accept			json
// @Produce		json
// @Param			id			path		int	true	"Group ID"
// @Param			locationID	path		int	true	"Location ID"
// @Success		200			{string}	string	"Successfully removed location from group"
// @Router			/groups/{id}/locations/{locationID} [delete]
func (config *GroupLocationConfig) DeleteLocationFromGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid group ID"})
		return
	}

	locationID, err := strconv.Atoi(chi.URLParam(r, "locationID"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid location ID"})
		return
	}

	err = config.GroupLocationEntryRepository.Delete(uint(groupID), uint(locationID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to remove location from group"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Location removed from group successfully"})
}
