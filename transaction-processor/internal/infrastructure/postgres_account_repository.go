package infrastructure

import (
	"github.com/amaterazu7/transaction-processor/internal/domain"
)

type PostgresAccountRepository struct{}

func NewPostgresAccountRepository() domain.AccountRepository {
	return &PostgresAccountRepository{}
}

func (postgres PostgresAccountRepository) FindById() error {
	return nil
}
