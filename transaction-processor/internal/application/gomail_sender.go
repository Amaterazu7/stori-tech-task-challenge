package application

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"path/filepath"
)

type EmailSenderService struct {
	filePrefix string
}

func NewEmailSenderService() domain.SenderService {
	filePath, _ := filepath.Abs(os.Getenv("SOURCE_TEMPLATE_PATH"))
	return &EmailSenderService{
		filePrefix: filePath,
	}
}

func (es *EmailSenderService) SendMessage(processorResult *models.ProcessorResult) (int, error) {
	var body bytes.Buffer
	t := template.Must(template.ParseFiles(es.filePrefix + "/email_body.html"))
	err := t.Execute(&body, processorResult)
	if err != nil {
		return 500, errors.New(fmt.Sprintf("Executing: %s", err.Error()))
	}

	err = es.SendTransactionByGomail(body, processorResult.GetEmail())
	if err != nil {
		return 500, errors.New(fmt.Sprintf("Sending Transaction: %s", err.Error()))
	}

	return 200, nil

}

func (es *EmailSenderService) SendTransactionByGomail(body bytes.Buffer, emailTo string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL_FROM"))
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Last bank transfers summary")
	m.SetBody("text/html", body.String())
	m.Embed(es.filePrefix + "/img/stori.png")

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
