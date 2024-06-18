package domain

type TransactionService interface {
	RunProcessor() (int, error)
	ValidateAccount() (bool, error)
	PersistTransaction() error
}
