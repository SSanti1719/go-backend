package repository

import "backend-go/mod/domain/entities"

type EmailsRespository interface {
	SearchEmails(toEmail string, fromEmail string, subjectKeyword string, messageId string, page uint16, size uint16) (*entities.Emails, error)
	SearchEmail(messageId string) (*entities.Email, error)
}
