package routes

import (
	"backend-go/mod/domain/entities"
	"backend-go/mod/domain/usecases"
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type AuthRoutes struct {
	Router  *chi.Mux
	usecase usecases.AuthUseCases
}

func NewAuthRoutes(usecases usecases.AuthUseCases) AuthRoutes {
	r := chi.NewRouter()
	routes := AuthRoutes{r, usecases}
	r.Post("/generate", routes.GenerateJwt)
	r.Post("/validate", routes.ValidateJwt)
	return routes
}

func (authRoutes AuthRoutes) GenerateJwt(w http.ResponseWriter, r *http.Request) {
	var request models.GenerateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	if len(request.Username) == 0 || strings.Trim(request.Username, "") == "" {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The 'username' field is required",
		})
	}
	if len(request.Password) == 0 || strings.Trim(request.Password, "") == "" {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The 'password' field is required",
		})
	}

	token, err := authRoutes.usecase.GenerateJwt(request.Username, request.Password)
	if err != nil {
		er, ok := err.(*entities.AuthError)
		if !ok {
			panic(er)
		} else {
			panic(models.Error{
				Status:  http.StatusUnauthorized,
				Message: models.Response401.Message + ": " + err.Error(),
			})
		}

	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})

	if err != nil {
		panic(err)
	}
}

func (authRoutes AuthRoutes) ValidateJwt(w http.ResponseWriter, r *http.Request) {
	var request models.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	if len(request.Token) == 0 || strings.Trim(request.Token, "") == "" {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The 'token' field is required",
		})
	}

	valid, err := authRoutes.usecase.ValidateJwt(request.Token)
	if err != nil {
		panic(models.Error{
			Status:  http.StatusUnauthorized,
			Message: models.Response401.Message + ": " + err.Error(),
		})
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]any{
		"token": request.Token,
		"valid": valid,
	})

	if err != nil {
		panic(err)
	}
}
