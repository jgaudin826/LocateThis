package authentication

import (
	"locate-this/config"
	"locate-this/database/dbmodel"
	"locate-this/pkg/models"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AuthConfig {
	return &AuthConfig{configuration}
}

// @Summary		User login
// @Description	Authenticate user and return JWT token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		models.UserRequest	true	"Login credentials"
// @Success		200		{object}	models.TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/login [post]
func (config *AuthConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}

	user, err := config.UserEntryRepository.FindByEmail(req.Email)
	if err != nil {
		user, err = config.UserEntryRepository.FindByUsername(req.Username)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
		return
	}

	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary		User register
// @Description	Create a new user and return JWT tokens
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		models.UserRequest	true	"Register credentials"
// @Success		200		{object}	models.TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/register [post]
func (config *AuthConfig) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	_, err := config.UserEntryRepository.FindByEmail(req.Email)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": " email or pseudo already in use"})
		return
	}
	_, err = config.UserEntryRepository.FindByUsername(req.Username)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": " email or pseudo already in use"})
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
	user := &models.UserResponse{ID: res.ID, Email: res.Email, Username: res.Username}

	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}
	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary		Refresh token
// @Description	Generate a new JWT token from an existing valid refresh token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		models.TokenRequest	true	"Refresh token"
// @Success		200		{object}	models.TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/refresh [post]
func (config *AuthConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}

	email, err := ParseToken(os.Getenv("REFRESH_SECRET"), req.RefreshToken)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid refresh token"})
		return
	}

	user, err := config.UserEntryRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}
