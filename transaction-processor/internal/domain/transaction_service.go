package domain

import "github.com/amaterazu7/transaction-processor/internal/domain/model"

type TransactionService interface {
	Run(accountId string) (int, error)
	ValidateAccount(accountId string) (bool, error)
	CreateTransactionList(transaction *model.Transaction) error
	PersistTransaction() error
	SendTransactionByEmail() error
}
