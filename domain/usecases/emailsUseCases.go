package usecases

import (
	"backend-go/mod/domain/entities"
	"backend-go/mod/domain/repository"
)

type EmailsUseCases struct {
	back repository.EmailsRespository
}

func NewEmailsUseCases(back repository.EmailsRespository) EmailsUseCases {
	return EmailsUseCases{back}
}

func (usecases EmailsUseCases) SearchEmails(to string, from string, subject string, messageId string, page uint16, size uint16) (*entities.Emails, error) {
	return usecases.back.SearchEmails(to, from, subject, messageId, page, size)
}

func (usecases EmailsUseCases) SearchEmail(messageId string) (*entities.Email, error) {
	return usecases.back.SearchEmail(messageId)
}
