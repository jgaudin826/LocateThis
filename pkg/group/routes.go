package group

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

/*
Groups:
- POST /groups
- GET /groups
- GET /groups/{id}
- PUT /groups/{id}
- DELETE /groups/{id}

- GET /groups/{id}/locations
- GET /groups/{id}/users

*/

func Routes(configuration *config.Config) chi.Router {
	GroupConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", GroupConfig.PostGroupHandler)
	router.Get("/", GroupConfig.GetAllGroupHandler) // FOR DEBUG ONLY
	router.Get("/{id}", GroupConfig.GetGroupByIDHandler)
	router.Put("/{id}", GroupConfig.PutGroupHandler)
	router.Delete("/{id}", GroupConfig.DeleteGroupHandler)
	router.Get("/{id}/locations", GroupConfig.GetLocationsForGroupHandler)
	router.Get("/{id}/users", GroupConfig.GetUsersForGroupHandler)
	return router
}
