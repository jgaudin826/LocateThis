package authentication

import (
	"context"
	"locate-this/config"
	"locate-this/database/dbmodel"
    "strconv"
	"locate-this/pkg/user"
	"log"
	"net/http"
) 

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token",
				http.StatusUnauthorized)
				return
			}

			id, err := ParseToken(secret, authHeader)
			if err != nil {
				http.Error(w, "Invalid token",
				http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "id", id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}	
}

func GetUserFromContext(ctx context.Context) (*dbmodel.UserEntry, error) {
	id, _ := ctx.Value("id").(string)

	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}
	UserConfig := user.New(configuration)

	ID, _ := strconv.Atoi(id)

	return UserConfig.UserEntryRepository.FindById(uint(ID))
}