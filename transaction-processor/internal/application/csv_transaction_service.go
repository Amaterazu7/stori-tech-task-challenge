package application

import (
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"github.com/google/uuid"
	"log"
)

type CsvTransactionService struct {
	AccountRepository     domain.AccountRepository
	TransactionRepository domain.TransactionRepository
}

func NewCsvTransactionService(
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
) domain.TransactionService {
	return &CsvTransactionService{
		AccountRepository:     accountRepository,
		TransactionRepository: transactionRepository,
	}
}

func (cts CsvTransactionService) Run(accountId string) (int, error) {
	isValid, err := cts.ValidateAccount(accountId)
	if !isValid || err != nil {
		log.Printf("Account ID is not valid : %s", err.Error())
		return 404, errors.New("account ID is not valid")
	}

	tx := models.Transaction{}
	err = cts.CreateTransactionList(&tx)
	if err != nil {
		return 500, errors.New(fmt.Sprintf("creating Transaction: %s", err.Error()))
	}
	log.Printf("[INFO] :: Created Transaction: %s", &tx.Id) // TODO:: REMOVE THIS ONLY to TESTS

	err = cts.PersistTransaction()
	if err != nil {
		return 500, errors.New(fmt.Sprintf("persisting Transaction: %s", err.Error()))
	}

	return 200, nil
}
func (cts CsvTransactionService) ValidateAccount(accountId string) (bool, error) {
	// TODO: call repositories and validate if account exit
	cts.AccountRepository.FindById()
	return true, nil
}

func (cts CsvTransactionService) CreateTransactionList(transaction *models.Transaction) error {
	// TODO: convert to transaction
	// transaction.Build(uuid.New())
	transaction.Id = uuid.New()

	return nil
}

func (cts CsvTransactionService) PersistTransaction() error {
	// TODO: call repositories and save the transaction in the DB
	cts.TransactionRepository.Save()
	return nil
}
