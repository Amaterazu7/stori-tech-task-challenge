package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/models"

type TransactionService interface {
	Run(accountId string) (int, error)
	ValidateAccount(accountId string) (bool, error)
	CreateTransactionList(transaction *models.Transaction) error
	PersistTransaction() error
}
