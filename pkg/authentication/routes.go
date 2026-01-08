package authentication

import (
	"locate-this/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	UserConfig := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware("demo_key_refresh"))
		r.Get("/login", UserConfig.LoginHandler)
		r.Get("/refresh", UserConfig.RefreshHandler)
	})
	return router
}
