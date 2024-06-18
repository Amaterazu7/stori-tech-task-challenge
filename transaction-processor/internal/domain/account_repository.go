package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/models"

type AccountRepository interface {
	FindById(id string) (*models.Account, error)
}
