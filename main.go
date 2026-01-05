package main


import (
	"log"
	"net/http"
	"LocateThis/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
    router := chi.NewRouter()
	
	// Routeurs
	router.Mount("/api/authentification", user.Routes(configuration))
	router.Group(func(r chi.Router) {
    	r.Use(authentification.AuthMiddleware("demo_key"))
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