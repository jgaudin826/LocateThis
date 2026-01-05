package location

import (
	"fmt"
	"locate-this/config"
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

	locationEntry := &dbmodel.LocationEntry{Name: req.Name, Age: req.Age, Breed: req.Breed, Weight: req.Weight}
	res, err := config.LocationEntryRepository.Create(locationEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create location"})
		return
	}

	locationResponse := &models.LocationResponse{ID: res.ID, Name: res.Name, Age: res.Age, Breed: res.Breed, Weight: res.Weight}
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
			ID:     location.ID,
			Name:   location.Name,
			Age:    location.Age,
			Breed:  location.Breed,
			Weight: location.Weight,
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
	entry, err := config.LocationEntryRepository.FindByID(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve location"})
		return
	}
	locationResponse := &models.LocationResponse{ID: entry.ID, Name: entry.Name, Age: entry.Age, Breed: entry.Breed, Weight: entry.Weight}
	render.JSON(w, r, locationResponse)
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

	locationEntry := &dbmodel.LocationEntry{Name: req.Name, Age: req.Age, Breed: req.Breed, Weight: req.Weight}
	updated, err := config.LocationEntryRepository.Update(id, locationEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update location"})
		return
	}

	locationResponse := &models.LocationResponse{ID: uint(id), Name: updated.Name, Age: updated.Age, Breed: updated.Breed, Weight: updated.Weight}
	render.JSON(w, r, locationResponse)
}

// @Summary		Delete a location
// @Description	Delete a location by its ID
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Location ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Router			/locations/{id} [delete]
func (config *LocationConfig) DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	err = config.LocationEntryRepository.Delete(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete location"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
