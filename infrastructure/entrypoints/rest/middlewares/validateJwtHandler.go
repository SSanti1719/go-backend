package middlewares

import (
	"backend-go/mod/domain/usecases"
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"net/http"
	"strings"
)

func ValidateJwt(usecases usecases.AuthUseCases) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				panic(models.Response401)
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				panic(models.Response401)
			}
			valid, err := usecases.ValidateJwt(parts[1])
			if err != nil {
				panic(models.Error{
					Status:  http.StatusUnauthorized,
					Message: models.Response401.Message + ": " + err.Error(),
				})
			}

			if !valid {
				panic(models.Response401)
			}
			next.ServeHTTP(w, r)
		})
	}
}
