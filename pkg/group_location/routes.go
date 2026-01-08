package group_location

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	GroupLocationConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/groups/locations", GroupLocationConfig.PostLocationToGroupHandler)
	router.Get("/groups/locations", GroupLocationConfig.GetAllGroupLocationHandler) // FOR DEBUG ONLY
	router.Put("/groups/{id}/locations/{locationID}", GroupLocationConfig.PutLocationInGroupHandler)
	router.Delete("/groups/{id}/locations/{locationID}", GroupLocationConfig.DeleteLocationFromGroupHandler)
	return router
}
