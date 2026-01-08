package group_location

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

/*
Group_Location:
- POST /group-location/
- GET /group-location/
- PUT /group-location/{id}/locations/{id}
- DELETE /group-location/{id}/locations/{id}
*/
func Routes(configuration *config.Config) chi.Router {
	GroupLocationConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", GroupLocationConfig.PostLocationToGroupHandler)
	router.Get("/", GroupLocationConfig.GetAllGroupLocationHandler) // FOR DEBUG ONLY
	router.Put("/{id}/locations/{locationID}", GroupLocationConfig.PutLocationInGroupHandler)
	router.Delete("/{id}/locations/{locationID}", GroupLocationConfig.DeleteLocationFromGroupHandler)
	return router
}
