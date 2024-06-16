package infrastructure

import (
	"github.com/amaterazu7/transaction-processor/internal/domain"
)

type PostgresTransactionRepository struct{}

func NewPostgresTransactionRepository() domain.TransactionRepository {
	return &PostgresTransactionRepository{}
}

func (postgres PostgresTransactionRepository) Save() error {
	return nil
}
