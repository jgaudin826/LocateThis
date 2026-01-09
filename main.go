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
	"os"

	_ "locate-this/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			LocateThis API
// @version			1.0
// @description		API for managing groups, users, locations, and authentication
// @host			localhost:8080
// @BasePath		/api
// @securityDefinitions.apikey	BearerAuth
// @in				header
// @name			Authorization
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Mount("/api/auth", authentication.Routes(configuration))

	router.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware(os.Getenv("JWT_SECRET")))
		r.Mount("/api/groups", group.Routes(configuration))
		r.Mount("/api/group-location", group_location.Routes(configuration))
		r.Mount("/api/group-user", group_user.Routes(configuration))
		r.Mount("/api/locations", location.Routes(configuration))
		r.Mount("/api/users", user.Routes(configuration))
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

	log.Println("Server running on http://localhost:" + os.Getenv("PORT"))
	log.Println("Swagger UI available at http://localhost:" + os.Getenv("PORT") + "/swagger/index.html")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
