package application

import (
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

type TransactionResult struct {
	AccountId           string
	TotalBalance        float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	AverageBalance      float64
	MonthMap            map[string]int
}

type CsvTransactionService struct {
	TransactionResult      TransactionResult
	AccountRepository      domain.AccountRepository
	TransactionRepository  domain.TransactionRepository
	HandleBucketRepository domain.HandleBucketRepository
}

func NewCsvTransactionService(
	accountId string,
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
	handleBucketRepository domain.HandleBucketRepository,
) domain.TransactionService {
	return &CsvTransactionService{
		TransactionResult: TransactionResult{
			AccountId:           accountId,
			TotalBalance:        0,
			AverageDebitAmount:  0,
			AverageCreditAmount: 0,
			AverageBalance:      0,
			MonthMap:            make(map[string]int),
		},
		AccountRepository:      accountRepository,
		TransactionRepository:  transactionRepository,
		HandleBucketRepository: handleBucketRepository,
	}
}

func (cts CsvTransactionService) RunProcessor() (int, error) {
	isValid, err := cts.ValidateAccount()
	if !isValid || err != nil {
		var msg = ""
		if err != nil {
			msg = err.Error()
		}
		return 404, errors.New(fmt.Sprintf("Account ID is not valid, %s", msg))
	}

	// fileName := cts.TransactionResult.AccountId + "transaction_list.csv" // TODO: MOVE to EnV
	content, err := cts.HandleBucketRepository.FindFileByName("17d340fa-5bf5-4429-8167-bafe4c0af0a7_transaction_list.csv")
	if err != nil {
		return 502, errors.New(err.Error())
	}

	rows := 0
	for _, line := range strings.Fields(strings.ReplaceAll(content, "txId;txDate;transaction", "")) {
		rows = rows + 1
		tx := models.Transaction{}
		err = createTransactionFromString(line, cts.TransactionResult.AccountId, &tx)
		if err != nil {
			return 500, errors.New(fmt.Sprintf("Creating Transaction: %s", err.Error()))
		}

		// TODO:: REMOVE THIS ONLY to TESTS
		//  log.Printf("\n [INFO] :: ")
		//  log.Printf("%v", tx)

		err = cts.PersistTransaction()
		if err != nil {
			return 500, errors.New(fmt.Sprintf("Persisting Transaction: %s", err.Error()))
		}
	}

	return 200, nil
}

func (cts CsvTransactionService) ValidateAccount() (bool, error) {
	account, err := cts.AccountRepository.FindById(cts.TransactionResult.AccountId)
	log.Printf("\n [INFO] :: ")
	log.Printf("REF:: %v", account)
	log.Printf("ACCOUNT:: %v", &account)

	if err != nil || account == nil {
		return false, err
	}
	return true, nil
}

func (cts CsvTransactionService) PersistTransaction() error {
	// TODO: call repositories and save the transaction in the DB
	cts.TransactionRepository.Save()
	return nil
}

func createTransactionFromString(txCrud string, accountId string, transaction *models.Transaction) error {
	columns := strings.Split(txCrud, ";")
	txCrudId, uuidErr := uuid.Parse(columns[0])
	txAccountCrudId, accUuidErr := uuid.Parse(accountId)
	txCrudDate, dateErr := time.Parse("2006/01/02", columns[1])
	txCrudAmount, floatErr := strconv.ParseFloat(columns[2], 64)
	txType := models.DEBIT
	if uuidErr != nil || accUuidErr != nil || dateErr != nil || floatErr != nil {
		var errMsg string
		switch {
		case uuidErr != nil:
			errMsg = fmt.Sprintf(uuidErr.Error())
		case accUuidErr != nil:
			errMsg = fmt.Sprintf(accUuidErr.Error())
		case dateErr != nil:
			errMsg = fmt.Sprintf(dateErr.Error())
		default:
			errMsg = fmt.Sprintf(floatErr.Error())
		}
		return errors.New(fmt.Sprintf("Unable to parse the transaction values [ %s ]", errMsg))
	}

	if txCrudAmount > 0 {
		txType = models.CREDIT
	}
	transaction.Id = txCrudId
	transaction.AccountId = txAccountCrudId
	transaction.Amount = txCrudAmount
	transaction.TxType = txType
	transaction.CreatedAt = txCrudDate

	return nil
}
