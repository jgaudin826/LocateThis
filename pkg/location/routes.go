package location

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

/*
Locations:
- POST /locations
- GET /locations
- GET /locations/{id}
- PUT /locations/{id}
- DELETE /locations/{id}
*/

func Routes(configuration *config.Config) chi.Router {
	LocationConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", LocationConfig.PostLocationHandler)
	router.Get("/", LocationConfig.GetAllLocationHandler) // FOR DEBUG ONLY
	router.Get("/{id}", LocationConfig.GetLocationByIDHandler)
	router.Put("/{id}", LocationConfig.PutLocationHandler)
	router.Delete("/{id}", LocationConfig.DeleteLocationHandler)
	router.Get("/{id}/groups", LocationConfig.GetGroupsForLocationHandler)
	return router
}
