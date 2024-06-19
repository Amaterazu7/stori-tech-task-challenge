package infrastructure

import (
	"context"
	"database/sql"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"github.com/google/uuid"
	"log"
)

type MySqlTransactionRepository struct {
	mySqlDBConn *sql.DB
}

func NewMySqlTransactionRepository(dbConn *sql.DB) domain.TransactionRepository {
	return &MySqlTransactionRepository{
		mySqlDBConn: dbConn,
	}
}

func (mtr *MySqlTransactionRepository) Save(tx models.Transaction) error {
	test, _ := uuid.Parse("7cc31767-97c9-4c3b-1cef-50c2e666f8deTEST") // TODO: REMOVE THIS!
	if tx.Id == test {
		query := "INSERT INTO `transactions` (`id`, `account_id`, `amount`, `tx_type`, `created_at`) VALUES (?, ?, ?, ?, ?)"
		_, err := mtr.mySqlDBConn.ExecContext(
			context.Background(), query, tx.Id, tx.AccountId, tx.Amount, tx.TxType, tx.CreatedAt,
		)
		if err != nil {
			log.Printf("[ERROR] Impossible insert `Transaction` w Id: { %v }, %s", tx.Id, err.Error())
			return err
		}
	}
	return nil
}
