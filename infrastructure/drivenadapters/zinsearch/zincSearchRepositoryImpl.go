package zinsearch

import (
	"backend-go/mod/domain/entities"
	"backend-go/mod/infrastructure/drivenadapters/zinsearch/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type zincSearchRepositoryImpl struct {
	client *http.Client
}

func NewZincSearchRepositoryImpl() zincSearchRepositoryImpl {
	var client = &http.Client{}
	return zincSearchRepositoryImpl{client}
}

func (z zincSearchRepositoryImpl) SearchEmails(toEmail string, fromEmail string, subjectKeyword string, messageId string, page uint16, size uint16) (*entities.Emails, error) {

	jsonDataToSend, err := z.getQueryEmail(toEmail, fromEmail, subjectKeyword, messageId, page, size)
	if err != nil {
		return nil, err
	}

	body, err := z.httpExec(z.client, jsonDataToSend)
	if err != nil {
		return nil, err
	}

	var respuesta models.Response
	if err := json.Unmarshal([]byte(body), &respuesta); err != nil {
		return nil, err
	}

	emails := make([]entities.Email, len(respuesta.Hits.Hits))
	for i, hit := range respuesta.Hits.Hits {
		emails[i] = entities.Email{
			MessageId: hit.Source.MessageId,
			Date:      hit.Source.Date,
			From:      hit.Source.From,
			To:        hit.Source.To,
			Subject:   hit.Source.Subject,
			Content:   hit.Source.Content,
		}
	}

	result := entities.Emails{
		Length: len(respuesta.Hits.Hits),
		Data:   emails,
	}

	return &result, nil
}

func (z zincSearchRepositoryImpl) SearchEmail(messageId string) (*entities.Email, error) {

	query := models.Query{
		SearchType: "querystring",
		QueryString: models.Term{
			Term: "+messageId:\"" + messageId + "\"",
		},
		MaxResults: 20,
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	body, err := z.httpExec(z.client, string(jsonData))
	if err != nil {
		return nil, err
	}

	var respuesta models.Response
	if err := json.Unmarshal([]byte(body), &respuesta); err != nil {
		return nil, err
	}

	var email entities.Email

	if len(respuesta.Hits.Hits) > 0 {
		email = entities.Email{
			MessageId: respuesta.Hits.Hits[0].Source.MessageId,
			Date:      respuesta.Hits.Hits[0].Source.Date,
			From:      respuesta.Hits.Hits[0].Source.From,
			To:        respuesta.Hits.Hits[0].Source.To,
			Subject:   respuesta.Hits.Hits[0].Source.Subject,
			Content:   respuesta.Hits.Hits[0].Source.Content,
		}
	} else {
		return nil, err
	}
	return &email, nil
}

func (z zincSearchRepositoryImpl) httpExec(client *http.Client, query string) ([]byte, error) {

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%s:%s%s%s%s",
			os.Getenv("ZINCSEARCH_IP"),
			os.Getenv("ZINCSEARCH_PORT"),
			"/api/", os.Getenv("ZINCSEARCH_INDEX"), "/_search"),
		strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(os.Getenv("ZINC_FIRST_ADMIN_USER"), os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("dear user, the service is currently unavailable. please try again later")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var zincResponse models.ErrorResponse
		if err := json.Unmarshal([]byte(body), &zincResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(zincResponse.Error)
	}
	return body, nil
}

func (z zincSearchRepositoryImpl) getQueryEmail(toEmail string, fromEmail string, subjectKeyword string, messageId string, page uint16, size uint16) (string, error) {

	queryString := ""
	searchType := "querystring"

	if fromEmail != "" {
		queryString += "+from:\"" + fromEmail + "\" "
	}
	if toEmail != "" {
		queryString += "+to:\"" + toEmail + "\" "
	}
	if subjectKeyword != "" {
		queryString += "+subject:\"" + subjectKeyword + "\""
	}
	if messageId != "" {
		queryString += "+messageId:\"" + messageId + "\""
	}
	if queryString == "" {
		searchType = "alldocuments"
	}

	query := models.Query{
		SearchType: searchType,
		QueryString: models.Term{
			Term: queryString,
		},
		From:       page * size,
		MaxResults: size,
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
