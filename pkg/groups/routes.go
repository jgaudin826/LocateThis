package groups

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
    GroupsConfig := New(configuration)
    router := chi.NewRouter()  
    router.Group(func(r chi.Router) {
    	router.Post("/groups", GroupsConfig.PostGroupHandler)
    	router.Get("/groups", GroupsConfig.GetAllGroupHandler) 
    	router.Get("/groups/{id}", GroupsConfig.GetGroupByIDHandler)
    	router.Put("/groups/{id}", GroupsConfig.PutGroupHandler) 
    	router.Delete("/groups/{id}", GroupsConfig.DeleteGroupHandler) 
    })
    router.Group(func(r chi.Router) {
    	router.Post("/groups/{id}/users", GroupsConfig.)
    	router.Get("/groups/{id}/users", GroupsConfig.GetUsersForGroupHandler)
    	router.Delete("/groups/{id}/users/{id}", GroupsConfig.)
    })
    router.Group(func(r chi.Router) {
    	router.Post("/groups/{id}/locations", GroupsConfig.)
    	router.Get("/groups/{id}/locations", GroupsConfig.GetLocationsForGroupHandler) 
    	router.Put("/groups/{id}/locations/{id} ", GroupsConfig.) 
    	router.Delete("/groups/{id}/locations/{id} ", GroupsConfig.) 
    })
	return router 
}