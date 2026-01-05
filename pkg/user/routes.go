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
    router.Post("/users", UserConfig.)
    router.Get("/users", UserConfig.) 
    router.Get("/users/{id}", UserConfig.)
    router.Put("/users/{id}", UserConfig.) 
    router.Delete("/users/{id}", UserConfig.) 
    router.Group(func(r chi.Router) {
    	r.Use(authentication.AuthMiddleware("demo_key_refresh"))
		r.Get("/login", UserConfig.)
        r.Get("/refresh", UserConfig.) 
    })
	return router 
}