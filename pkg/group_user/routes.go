package group_user

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	GroupUserConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/groups/users", GroupUserConfig.PostUserToGroupHandler)
	router.Get("/groups/users", GroupUserConfig.GetAllGroupUserHandler) // FOR DEBUG ONLY
	router.Delete("/groups/{id}/users/{userID}", GroupUserConfig.DeleteUserFromGroupHandler)
	return router
}
