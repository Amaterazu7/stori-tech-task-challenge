package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/models"

type SenderService interface {
	SendMessage(processorResult *models.ProcessorResult) (int, error)
}
