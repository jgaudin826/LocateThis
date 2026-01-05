package location

import (
	"locate-this/config"
	"locate-this/pkg/authentication"

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
    router.Post("/locations", LocationsConfig.)
    router.Get("/locations", LocationsConfig.) 
    router.Get("/locations/{id}", locationsConfig.) 
    router.Put("/locations/{id}", LocationsConfig.) 
    router.Delete("/locations/{id}", LocationsConfig.) 
	return router 
}