package infrastructure

import (
	"context"
	"database/sql"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
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
	query := "INSERT INTO `transactions` (`id`, `account_id`, `amount`, `tx_type`, `created_at`) VALUES (?, ?, ?, ?, ?)"
	_, err := mtr.mySqlDBConn.ExecContext(
		context.Background(), query, tx.Id, tx.AccountId, tx.Amount, tx.TxType, tx.CreatedAt,
	)
	if err != nil {
		log.Printf("[ERROR] Impossible insert `Transaction` w Id: { %v }, %s", tx.Id, err.Error())
		return err
	}
	return nil
}
