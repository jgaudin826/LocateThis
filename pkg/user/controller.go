package user

import (
	"fmt"
	"locate-this/config"
	"locate-this/database/dbmodel"
	"locate-this/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	*config.Config
}

func New(configuration *config.Config) *UserConfig {
	return &UserConfig{configuration}
}

// @Summary		Create a new user
// @Description	Create a new user entry
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			request	body		models.UserRequest	true	"User data"
// @Success		200		{object}	models.UserResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users [post]
func (config *UserConfig) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.UserEntry{Email: req.Email, Password: req.Password, Username: req.Username}
	res, err := config.UserEntryRepository.Create(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}

	userResponse := &models.UserResponse{ID: res.ID, Email: res.Email, Username: res.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Get all users
// @Description	Retrieve a list of all users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.UserResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users [get]
func (config *UserConfig) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.UserEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve users"})
		return
	}

	usersResponse := make([]models.UserResponse, 0)
	for _, user := range entries {
		usersResponse = append(usersResponse, models.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	}

	render.JSON(w, r, usersResponse)
}

func (config *UserConfig) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	entry, err := config.UserEntryRepository.FindById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
		return
	}
	userResponse := &models.UserResponse{ID: entry.ID, Email: entry.Email, Username: entry.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Get user by ID
// @Description	Retrieve a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	models.UserResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users/{id} [get]
func (config *UserConfig) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	entry, err := config.UserEntryRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
		return
	}
	userResponse := &models.UserResponse{ID: entry.ID, Email: entry.Email, Username: entry.Username}
	render.JSON(w, r, userResponse)
}

func (config *UserConfig) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	entry, err := config.UserEntryRepository.FindByUsername(username)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
		return
	}
	userResponse := &models.UserResponse{ID: entry.ID, Email: entry.Email, Username: entry.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Get locations for a user
// @Description	Retrieve all locations created by a user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{array}	models.LocationResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users/{id}/locations [get]
func (config *UserConfig) GetLocationsForUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	locations, err := config.UserEntryRepository.FindLocationsForUser(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve locations"})
		return
	}
	render.JSON(w, r, locations)
}

// @Summary		Get groups for a user
// @Description	Retrieve all groups a user belongs to
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{array}	models.GroupResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users/{id}/groups [get]
func (config *UserConfig) GetGroupsForUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	groups, err := config.UserEntryRepository.FindGroupsForUser(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve groups"})
		return
	}
	render.JSON(w, r, groups)
}

// @Summary		Update a user
// @Description	Update a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"User ID"
// @Param			request	body		models.UserRequest	true	"Updated user data"
// @Success		200		{object}	models.UserResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router			/users/{id} [put]
func (config *UserConfig) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}

	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	userEntry := &dbmodel.UserEntry{Email: req.Email, Password: req.Password, Username: req.Username}
	updated, err := config.UserEntryRepository.Update(userEntry, uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	userResponse := &models.UserResponse{ID: uint(id), Email: updated.Email, Username: updated.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Delete a user
// @Description	Delete a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"User ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Security BearerAuth
// @Router			/users/{id} [delete]
func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.UserEntryRepository.Delete(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
