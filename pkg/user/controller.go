package user

import (
	"LocateThis/config"
	"LocateThis/database/dbmodel"
	"LocateThis/pkg/models"
	"fmt"
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
// @Router			/users [post]
func (config *UserConfig) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.UserEntry{Email: req.Email, Password: req.Password, Pseudo: req.Pseudo}
	res, err := config.UserEntryRepository.Create(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}

	userResponse := &models.UserResponse{ID: res.ID, Email: res.Email, Pseudo: res.Pseudo}
	render.JSON(w, r, userResponse)
}

// @Summary		Get all users
// @Description	Retrieve a list of all users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.UserResponse
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
			ID:     user.ID,
			Email:  user.Email,
			Pseudo: user.Pseudo,
		})
	}

	render.JSON(w, r, usersResponse)
}

// @Summary		Get user by ID
// @Description	Retrieve a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	models.UserResponse
// @Router			/users/{id} [get]
func (config *UserConfig) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	entry, err := config.UserEntryRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
		return
	}
	userResponse := &models.UserResponse{ID: entry.ID, Email: entry.Email, Pseudo: entry.Pseudo}
	render.JSON(w, r, userResponse)
}

// @Summary		Update a user
// @Description	Update a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"User ID"
// @Param			request	body		models.UserRequest	true	"Updated user data"
// @Success		200		{object}	models.UserResponse
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

	userEntry := &dbmodel.UserEntry{Email: req.Email, Password: req.Password, Pseudo: req.Pseudo}
	updated, err := config.UserEntryRepository.Update(id, userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	userResponse := &models.UserResponse{ID: uint(id), Email: updated.Email, Pseudo: updated.Pseudo}
	render.JSON(w, r, userResponse)
}

// @Summary		Delete a user
// @Description	Delete a user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"User ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Router			/users/{id} [delete]
func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	err = config.UserEntryRepository.Delete(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}

// @Summary		User login
// @Description	Authenticate user and return JWT token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		models.LoginRequest	true	"Login credentials"
// @Success		200		{object}	models.TokenResponse
// @Router			/login [post]

func (config *UserConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.LoginRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	user, err := config.UserEntryRepository.FindByEmail(req.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
		return
	}

	token, err := GenerateToken(config.JWT_Secret, req.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}

	render.JSON(w, r, models.TokenResponse{Token: token})
}

// @Summary		Refresh token
// @Description	Generate a new JWT token from an existing valid token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	models.TokenResponse
// @Router			/refresh [post]
func (config *UserConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		render.JSON(w, r, map[string]string{"error": "Missing token"})
		return
	}

	email, err := ParseToken(config.JWT_Secret, authHeader)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid token"})
		return
	}

	newToken, err := GenerateToken(config.JWT_Secret, email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}

	render.JSON(w, r, models.TokenResponse{Token: newToken})
}
