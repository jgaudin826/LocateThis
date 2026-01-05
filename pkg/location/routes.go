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
	LocationsConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/locations", LocationsConfig.PostLocationHandler)
	router.Get("/locations", LocationsConfig.GetAllLocationHandler)
	router.Get("/locations/{id}", LocationsConfig.GetLocationByIDHandler)
	router.Put("/locations/{id}", LocationsConfig.PutLocationHandler)
	router.Delete("/locations/{id}", LocationsConfig.DeleteLocationHandler)
	router.Get("/locations/{id}/groups", LocationsConfig.GetGroupsForLocationHandler)
	return router
}
