package application

import (
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/model"
	"github.com/google/uuid"
	"log"
)

type CsvTransactionService struct{}

func NewCsvTransactionService() domain.TransactionService {
	return &CsvTransactionService{}
}

func (cts CsvTransactionService) Run() error {
	tx, err := cts.CreateTransaction()
	if err != nil {
		message := fmt.Sprintf("Creating Transaction: %s", err.Error())
		return errors.New(message)
	}

	log.Printf("[INFO] :: Created Transaction: %s", &tx.Id)

	err = cts.PersistTransaction()
	if err != nil {
		message := fmt.Sprintf("Persisting Transaction: %s", err.Error())
		return errors.New(message)
	}

	err = cts.SendTransactionByEmail()
	if err != nil {
		return errors.New(fmt.Sprintf("Sending Transaction: %s", err.Error()))
	}

	return nil
}

func (cts CsvTransactionService) CreateTransaction() (*model.Transaction, error) {
	// TODO: convert to transaction
	return &model.Transaction{Id: uuid.New()}, nil
}

func (cts CsvTransactionService) PersistTransaction() error {
	// TODO: call repository and save the transaction in the DB
	return nil
}

func (cts CsvTransactionService) SendTransactionByEmail() error {
	// TODO: create a email msg and send by email
	return nil
}
