package domain

type TransactionRepository interface {
	Save() error
}
