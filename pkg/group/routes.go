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

- POST /groups/{id}/users ({"user_id" : id})
- GET /groups/{id}/users
- DELETE /groups/{id}/users/{id}

- POST /groups/{id}/locations ({"location_id" : id})
- GET /groups/{id}/locations
- PUT /groups/{id}/locations/{id}
- DELETE /groups/{id}/locations/{id}
*/

func Routes(configuration *config.Config) chi.Router {
	GroupConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/groups", GroupConfig.PostGroupHandler)
	router.Get("/groups", GroupConfig.GetAllGroupHandler) // FOR DEBUG ONLY
	router.Get("/groups/{id}", GroupConfig.GetGroupByIDHandler)
	router.Put("/groups/{id}", GroupConfig.PutGroupHandler)
	router.Delete("/groups/{id}", GroupConfig.DeleteGroupHandler)
	router.Get("/groups/{id}/locations", GroupConfig.GetLocationsForGroupHandler)
	router.Get("/groups/{id}/users", GroupConfig.GetUsersForGroupHandler)
	return router
}
