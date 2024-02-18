package middlewares

import (
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"net/http"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	panic(models.Error{
		Status:  http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	})
}
