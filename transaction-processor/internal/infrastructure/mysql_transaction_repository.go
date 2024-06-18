package infrastructure

import (
	"database/sql"
	"github.com/amaterazu7/transaction-processor/internal/domain"
)

type MySqlTransactionRepository struct {
	mySqlDBConn *sql.DB
}

func NewMySqlTransactionRepository(dbConn *sql.DB) domain.TransactionRepository {
	return &MySqlTransactionRepository{
		mySqlDBConn: dbConn,
	}
}

func (mtr MySqlTransactionRepository) Save() error {
	return nil
}
