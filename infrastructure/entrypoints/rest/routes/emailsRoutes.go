package routes

import (
	"backend-go/mod/domain/usecases"
	"backend-go/mod/infrastructure/entrypoints/rest/models"
	"backend-go/mod/infrastructure/entrypoints/rest/validators"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type EmailsRoutes struct {
	Router  *chi.Mux
	usecase usecases.EmailsUseCases
}

func NewEmailsRoutes(usecases usecases.EmailsUseCases) EmailsRoutes {
	r := chi.NewRouter()
	routes := EmailsRoutes{r, usecases}
	r.Get("/", routes.ListEmails)
	r.Get("/search/{id}", routes.GetEmail)
	return routes
}

func (emailRoutes EmailsRoutes) ListEmails(w http.ResponseWriter, r *http.Request) {
	toEmail := r.URL.Query().Get("toEmail")
	fromEmail := r.URL.Query().Get("fromEmail")
	subjectKeyword := r.URL.Query().Get("subjectKeyword")
	messageId := r.URL.Query().Get("messageId")
	pageS := r.URL.Query().Get("page")
	sizeS := r.URL.Query().Get("size")

	if len(toEmail) > 0 && !validators.EmailValidator(toEmail) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'toEmail' parameter is not in email format",
		})
	}
	if len(fromEmail) > 0 && !validators.EmailValidator(fromEmail) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'fromEmail' parameter is not in email format",
		})
	}
	if len(messageId) > 0 && !validators.MessageIdValidator(messageId) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The value of the 'messageId' parameter does not have the expected format",
		})
	}

	if len(pageS) > 0 && !validators.IsNumber(pageS) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'page' parameter is not in number format",
		})
	}
	if len(sizeS) > 0 && !validators.IsNumber(sizeS) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'size' parameter is not in number format",
		})
	}
	page, err := strconv.ParseUint(pageS, 10, 16)
	if err != nil || uint16(page) > uint16(10000) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'page' parameter is exceeds the size of 10000",
		})
	}
	size, err := strconv.ParseUint(sizeS, 10, 16)
	if err != nil || uint16(size) > uint16(10000) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "'size' parameter is exceeds the size of 10000",
		})
	}

	data, err := emailRoutes.usecase.SearchEmails(toEmail, fromEmail, subjectKeyword, messageId, uint16(page), uint16(size))
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func (emailRoutes EmailsRoutes) GetEmail(w http.ResponseWriter, r *http.Request) {
	messageId := chi.URLParam(r, "id")

	if len(messageId) == 0 || strings.Trim(messageId, "") == "" {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The 'messageId' parameter is required",
		})
	}
	if !validators.MessageIdValidator(messageId) {
		panic(models.Error{
			Status:  http.StatusBadRequest,
			Message: "The value of the 'messageId' parameter does not have the expected format",
		})
	}

	decoded, err := url.QueryUnescape(messageId)
	if err != nil {
		panic(err)
	}

	email, err := emailRoutes.usecase.SearchEmail(decoded)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(email)
	if err != nil {
		panic(err)
	}
}
