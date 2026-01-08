package group_user

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

/*
Group_User:
- POST /group-user/
- GET /group-user/
- DELETE /group-user/{id}/users/{id}
*/

func Routes(configuration *config.Config) chi.Router {
	GroupUserConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", GroupUserConfig.PostUserToGroupHandler)
	router.Get("/", GroupUserConfig.GetAllGroupUserHandler) // FOR DEBUG ONLY
	router.Delete("/{id}/users/{userID}", GroupUserConfig.DeleteUserFromGroupHandler)
	return router
}
