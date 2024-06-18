package application

import (
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
)

type EmailSenderService struct{}

func NewEmailSenderService() domain.SenderService {
	return &EmailSenderService{}
}

func (es EmailSenderService) SendMessage() (int, error) {
	// TODO: create a email msg and send by email
	err := es.SendTransactionByEmail()

	if err != nil {
		return 500, errors.New(fmt.Sprintf("Sending Transaction: %s", err.Error()))
	}

	return 200, nil

}

func (es EmailSenderService) SendTransactionByEmail() error {
	return nil
}
