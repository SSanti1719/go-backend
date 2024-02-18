package middlewares

import (
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	panic(models.Error{
		Status:  http.StatusNotFound,
		Message: "Not found",
	})
}
