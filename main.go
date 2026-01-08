package main

import (
	"locate-this/config"
	"locate-this/pkg/authentication"
	"locate-this/pkg/group"
	"locate-this/pkg/group_location"
	"locate-this/pkg/group_user"
	"locate-this/pkg/location"
	"locate-this/pkg/user"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	// Routeurs
	router.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware("demo_key"))
		r.Mount("/api", group.Routes(configuration))
		r.Mount("/api", group_location.Routes(configuration))
		r.Mount("/api", group_user.Routes(configuration))
		r.Mount("/api", location.Routes(configuration))
		r.Mount("/api", user.Routes(configuration))
	})

	return router
}

func main() {
	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}
	// Initialisation des routes
	router := Routes(configuration)
	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
