package infrastructure

import (
	"database/sql"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
)

type MySqlAccountRepository struct {
	mySqlDBConn *sql.DB
}

func NewMySqlAccountRepository(dbConn *sql.DB) domain.AccountRepository {
	return &MySqlAccountRepository{
		mySqlDBConn: dbConn,
	}
}

func scanRow(row *sql.Row) (*models.Account, error) {
	account := new(models.Account)
	err := row.Scan(&account.Id, &account.Name, &account.Email, &account.Asset, &account.Type, &account.UpdatedAt, &account.CreatedAt)

	if err != nil {
		return nil, err
	}
	return account, nil
}

func (mar *MySqlAccountRepository) FindById(id string) (*models.Account, error) {
	row := mar.mySqlDBConn.QueryRow("SELECT * FROM `account` WHERE id = ?", id)

	account, err := scanRow(row)
	if err != nil {
		return nil, err
	}

	return account, nil
}
