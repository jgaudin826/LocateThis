package location

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

type LocationConfig struct {
	*config.Config
}

func New(configuration *config.Config) *LocationConfig {
	return &LocationConfig{configuration}
}

// @Summary		Create a new location
// @Description	Create a new location entry
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			request	body		models.LocationRequest	true	"Location data"
// @Success		200		{object}	models.LocationResponse
// @Router			/locations [post]
func (config *LocationConfig) PostLocationHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.LocationRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	locationEntry := &dbmodel.LocationEntry{Name: req.Name, Latitude: req.Latitude, Longitude: req.Longitude}
	res, err := config.LocationEntryRepository.Create(locationEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create location"})
		return
	}

	locationResponse := &models.LocationResponse{ID: res.ID, Name: res.Name, Latitude: res.Latitude, Longitude: res.Longitude}
	render.JSON(w, r, locationResponse)
}

// @Summary		Get all locations
// @Description	Retrieve a list of all locations
// @Tags			locations
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.LocationResponse
// @Router			/locations [get]
func (config *LocationConfig) GetAllLocationHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.LocationEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve locations"})
		return
	}

	locationsResponse := make([]models.LocationResponse, 0)
	for _, location := range entries {
		locationsResponse = append(locationsResponse, models.LocationResponse{
			ID:        location.ID,
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		})
	}

	render.JSON(w, r, locationsResponse)
}

// @Summary		Get location by ID
// @Description	Retrieve a location by its ID
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Location ID"
// @Success		200	{object}	models.LocationResponse
// @Router			/locations/{id} [get]
func (config *LocationConfig) GetLocationByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	entry, err := config.LocationEntryRepository.FindById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve location"})
		return
	}
	locationResponse := &models.LocationResponse{ID: entry.ID, Name: entry.Name, Latitude: entry.Latitude, Longitude: entry.Longitude}
	render.JSON(w, r, locationResponse)
}

// @Summary		Get groups for a location
// @Description	Retrieve all groups that contain this location
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Location ID"
// @Success		200	{array}	models.GroupResponse
// @Router			/locations/{id}/groups [get]
func (config *LocationConfig) GetGroupsForLocationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	groups, err := config.LocationEntryRepository.FindGroupsForLocation(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve groups"})
		return
	}

	groupsResponse := make([]models.GroupResponse, 0)
	for _, group := range groups {
		groupsResponse = append(groupsResponse, models.GroupResponse{
			ID:   group.ID,
			Name: group.Name,
		})
	}

	render.JSON(w, r, groupsResponse)
}

// @Summary		Update a location
// @Description	Update an existing location entry
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Location ID"
// @Param			request	body		models.LocationRequest	true	"Location data"
// @Success		200		{object}	models.LocationResponse
// @Router			/locations/{id} [put]
func (config *LocationConfig) PutLocationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}

	req := &models.LocationRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	locationEntry := &dbmodel.LocationEntry{Name: req.Name, Latitude: req.Latitude, Longitude: req.Longitude}
	updated, err := config.LocationEntryRepository.Update(locationEntry, uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update location"})
		return
	}

	locationResponse := &models.LocationResponse{ID: uint(id), Name: updated.Name, Latitude: updated.Latitude, Longitude: updated.Longitude}
	render.JSON(w, r, locationResponse)
}

// @Summary		Delete a location
// @Description	Delete a location by its ID
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Location ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Router			/locations/{id} [delete]
func (config *LocationConfig) DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.LocationEntryRepository.Delete(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete location"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
