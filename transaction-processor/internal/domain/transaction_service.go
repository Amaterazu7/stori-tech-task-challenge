package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/models"

type TransactionService interface {
	RunProcessor() (int, error)
	ValidateAccount() (bool, error)
	PersistTransaction(tx models.Transaction) error
}
