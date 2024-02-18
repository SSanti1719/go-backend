package rest

import (
	"backend-go/mod/domain/repository"
	"backend-go/mod/domain/usecases"
	"backend-go/mod/infrastructure/entrypoints/rest/middlewares"
	"backend-go/mod/infrastructure/entrypoints/rest/routes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Rest struct {
	router *chi.Mux
}

func NewRest(emailAdapter repository.EmailsRespository) Rest {
	r := chi.NewRouter()
	// Configuración de CORS
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Origin"},
	}
	corsHandler := cors.New(corsOptions)

	r.Use(corsHandler.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middlewares.ErrorHandler)
	r.MethodNotAllowed(middlewares.MethodNotAllowedHandler)
	r.NotFound(middlewares.NotFoundHandler)

	authUseCases := usecases.NewAuthUseCases()

	r.Group(func(r chi.Router) {
		emailsUseCases := usecases.NewEmailsUseCases(emailAdapter)
		r.Use(middlewares.ValidateJwt(authUseCases))
		r.Mount("/emails", routes.NewEmailsRoutes(emailsUseCases).Router)
	})

	r.Group(func(r chi.Router) {
		r.Use(middlewares.ValidateBasic(authUseCases))
		r.Mount("/auth", routes.NewAuthRoutes(authUseCases).Router)
	})

	return Rest{r}
}

func (rest Rest) Run(port string) {
	fmt.Println("Running on port " + port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), rest.router)
}
