package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/model"

type TransactionService interface {
	Run() error
	CreateTransaction() (*model.Transaction, error)
	PersistTransaction() error
	SendTransactionByEmail() error
}
