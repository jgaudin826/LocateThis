package user

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

/*
Users:
- POST /users
- GET /users
- GET /users/{id}
- PUT /users/{id}
- DELETE /users/{id}
*/

func Routes(configuration *config.Config) chi.Router {
	UserConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", UserConfig.PostUserHandler)
	router.Get("/", UserConfig.GetAllUserHandler) // FOR DEBUG ONLY
	router.Get("/{id}", UserConfig.GetUserByEmailHandler)
	router.Put("/{id}", UserConfig.PutUserHandler)
	router.Delete("/{id}", UserConfig.DeleteUserHandler)
	router.Get("/{id}/locations", UserConfig.GetLocationsForUserHandler)
	router.Get("/{id}/groups", UserConfig.GetGroupsForUserHandler)
	return router
}
