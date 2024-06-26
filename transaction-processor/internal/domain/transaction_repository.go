package domain

import (
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
)

type TransactionRepository interface {
	TransactionalStarterRepository
	Save(tx models.Transaction) error
}
