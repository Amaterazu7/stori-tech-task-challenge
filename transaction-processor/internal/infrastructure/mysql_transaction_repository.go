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
	tx          *sql.Tx
	ctx         context.Context
}

func NewMySqlTransactionRepository(dbConn *sql.DB, txContext context.Context) domain.TransactionRepository {
	return &MySqlTransactionRepository{
		mySqlDBConn: dbConn,
		ctx:         txContext,
	}
}

func (mtr *MySqlTransactionRepository) Conn() *sql.DB {
	return mtr.mySqlDBConn
}

func (mtr *MySqlTransactionRepository) BeginTransaction() (*sql.Tx, error) {
	tx, err := mtr.mySqlDBConn.BeginTx(mtr.ctx, nil)
	if err != nil {
		return nil, err
	}
	mtr.tx = tx
	return tx, nil
}

func (mtr *MySqlTransactionRepository) Save(tx models.Transaction) error {
	query := "INSERT INTO `transactions` (`id`, `account_id`, `amount`, `tx_type`, `created_at`) VALUES (?, ?, ?, ?, ?)"
	_, err := mtr.tx.ExecContext(
		mtr.ctx, query, tx.Id, tx.AccountId, tx.Amount, tx.TxType, tx.CreatedAt,
	)
	if err != nil {
		log.Printf("[ERROR] Impossible insert `Transaction` w Id: { %v }, %s", tx.Id, err.Error())
		return err
	}
	return nil
}
