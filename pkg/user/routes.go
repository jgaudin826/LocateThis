package user

import (
	"locate-this/config"
	"locate-this/pkg/authentication"

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
	router.Post("/users", UserConfig.PostUserHandler)
	router.Get("/users", UserConfig.GetAllUserHandler) // FOR DEBUG ONLY
	router.Get("/users/{id}", UserConfig.GetUserByEmailHandler)
	router.Put("/users/{id}", UserConfig.PutUserHandler)
	router.Delete("/users/{id}", UserConfig.DeleteUserHandler)
	router.Get("/users/{id}/locations", UserConfig.GetLocationsForUserHandler)
	router.Get("/users/{id}/groups", UserConfig.GetGroupsForUserHandler)
	router.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware("demo_key_refresh"))
		r.Get("/login", UserConfig.LoginHandler)
		r.Get("/refresh", UserConfig.RefreshHandler)
	})
	return router
}
