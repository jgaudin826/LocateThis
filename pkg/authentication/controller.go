package authentication

import (
	"encoding/json"
	"locate-this/config"
	"locate-this/pkg/models"
	"net/http"

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
// @Param			request	body		models.LoginRequest	true	"Login credentials"
// @Success		200		{object}	models.TokenResponse
// @Router			/login [post]
func (config *AuthConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.LoginRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid req", http.StatusBadRequest)
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

	token, err := GenerateToken(config.SecretJWT, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(config.SecretRefreshJWT, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	Tokens := &models.TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}

	render.JSON(w, r, Tokens)
}

// @Summary		Refresh token
// @Description	Generate a new JWT token from an existing valid token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	models.TokenResponse
// @Router			/refresh [post]
func (config *AuthConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid req", http.StatusBadRequest)
		return
	}

	email, err := ParseToken(config.SecretRefreshJWT, req.RefreshToken)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid refresh token"})
		return
	}

	user, err := config.UserEntryRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	token, err := GenerateToken(config.SecretJWT, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(config.SecretRefreshJWT, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	Tokens := &models.TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}

	render.JSON(w, r, Tokens)
}
