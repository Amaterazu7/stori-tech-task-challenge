package domain

import (
	"database/sql"
)

type TransactionalStarterRepository interface {
	Conn() *sql.DB
	BeginTransaction() (*sql.Tx, error)
}
