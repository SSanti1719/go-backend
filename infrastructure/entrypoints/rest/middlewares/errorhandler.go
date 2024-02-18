package middlewares

import (
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.Header().Set("Content-Type", "application/json")

				response := models.Error{}
				err, ok := r.(models.Error)
				if ok {
					w.WriteHeader(err.Status)
					response = err
				} else {
					response.Status = http.StatusInternalServerError
					w.WriteHeader(response.Status)
					err, ok := r.(error)
					if ok {
						response.Message = err.Error()
					} else {
						response.Message = fmt.Sprint(r)
					}
				}
				json.NewEncoder(w).Encode(response)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
