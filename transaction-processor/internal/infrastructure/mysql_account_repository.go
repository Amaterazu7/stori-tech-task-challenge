package infrastructure

import (
	"database/sql"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"log"
)

type MySqlAccountRepository struct {
	mySqlDBConn *sql.DB
}

func NewMySqlAccountRepository(dbConn *sql.DB) domain.AccountRepository {
	return &MySqlAccountRepository{
		mySqlDBConn: dbConn,
	}
}

func scanRow(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)
	err := rows.Scan(
		&account.Id, &account.Name, &account.Asset, &account.Type, &account.UpdatedAt, &account.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return account, nil
}

func (mar MySqlAccountRepository) FindById(id string) (*models.Account, error) {
	rows, err := mar.mySqlDBConn.Query(
		"SELECT * FROM `account` WHERE id=?", id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	account, err := scanRow(rows)

	log.Printf("ACCOUNT FROM Repository :: %v", account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
