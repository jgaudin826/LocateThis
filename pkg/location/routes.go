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
	router.Post("/locations", LocationConfig.PostLocationHandler)
	router.Get("/locations", LocationConfig.GetAllLocationHandler) // FOR DEBUG ONLY
	router.Get("/locations/{id}", LocationConfig.GetLocationByIDHandler)
	router.Put("/locations/{id}", LocationConfig.PutLocationHandler)
	router.Delete("/locations/{id}", LocationConfig.DeleteLocationHandler)
	router.Get("/locations/{id}/groups", LocationConfig.GetGroupsForLocationHandler)
	return router
}
