package application

import (
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/amaterazu7/transaction-processor/internal/domain/models"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvTransactionService struct {
	ProcessorResult        models.ProcessorResult
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
		ProcessorResult: models.ProcessorResult{
			AccountId:           accountId,
			TotalBalance:        0,
			AverageDebitAmount:  0,
			AverageCreditAmount: 0,
			MonthMap:            make(map[string]int),
		},
		AccountRepository:      accountRepository,
		TransactionRepository:  transactionRepository,
		HandleBucketRepository: handleBucketRepository,
	}
}

func (cts *CsvTransactionService) RunProcessor() (int, *models.ProcessorResult, error) {
	isValid, err := cts.ValidateAccount()
	if !isValid || err != nil {
		var msg = ""
		if err != nil {
			msg = err.Error()
		}
		return 404, &models.ProcessorResult{}, errors.New(fmt.Sprintf("Account ID is not valid, %s", msg))
	}

	var fileName strings.Builder
	fileName.WriteString(cts.ProcessorResult.AccountId)
	fileName.WriteString("_")
	fileName.WriteString(os.Getenv("SOURCE_FILE_NAME"))

	content, err := cts.HandleBucketRepository.FindFileByName(fileName.String())
	if err != nil {
		return 502, &models.ProcessorResult{}, err
	}

	tx, err := cts.TransactionRepository.BeginTransaction()
	if err != nil {
		return 502, &models.ProcessorResult{}, errors.New(fmt.Sprintf(` - BeginTx, %s`, err.Error()))
	}
	defer func() {
		_ = tx.Rollback()
	}()

	txAccountCrudId, _ := uuid.Parse(cts.ProcessorResult.AccountId)
	_ = cts.PersistTransaction(models.Transaction{Id: uuid.New(), AccountId: txAccountCrudId, Amount: 200.3, TxType: models.CREDIT, CreatedAt: time.Now()})
	_ = cts.PersistTransaction(models.Transaction{Id: uuid.New(), AccountId: txAccountCrudId, Amount: -200.3, TxType: models.DEBIT, CreatedAt: time.Now()})

	rows, debitAmount, debitCount, creditAmount, creditCount := 0, 0.0, 0.0, 0.0, 0.0
	for _, line := range strings.Fields(strings.ReplaceAll(content, "txId;txDate;transaction", "")) {
		rows = rows + 1
		transaction := models.Transaction{}
		err = createTransactionFromString(line, cts.ProcessorResult.AccountId, &transaction)
		if err != nil {
			return 500, &models.ProcessorResult{}, errors.New(fmt.Sprintf("Creating Transaction: %s", err.Error()))
		}

		err = cts.PersistTransaction(transaction)
		if err != nil {
			return 500, &models.ProcessorResult{}, errors.New(fmt.Sprintf("Persisting Transaction: %s", err.Error()))
		}

		cts.ProcessorResult.FillAvgValues(&transaction, &debitAmount, &debitCount, &creditAmount, &creditCount)
		cts.ProcessorResult.FillMap(&transaction)
		cts.ProcessorResult.AddBalance(transaction.Amount)
	}

	err = tx.Commit()
	if err != nil {
		return 502, &models.ProcessorResult{}, errors.New(fmt.Sprintf(` - CommitTx, %s`, err.Error()))
	}

	cts.ProcessorResult.CalculateAverageDebit(debitAmount, debitCount)
	cts.ProcessorResult.CalculateAverageCredit(creditAmount, creditCount)

	return 200, &cts.ProcessorResult, nil
}

func (cts *CsvTransactionService) ValidateAccount() (bool, error) {
	account, err := cts.AccountRepository.FindById(cts.ProcessorResult.AccountId)
	if err != nil || account == nil {
		return false, err
	}
	cts.ProcessorResult.SetEmail(account.Email)
	return true, nil
}

func (cts *CsvTransactionService) PersistTransaction(transaction models.Transaction) error {
	err := cts.TransactionRepository.Save(transaction)
	if err != nil {
		return err
	}
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
	transaction.Build(txCrudId, txAccountCrudId, txCrudAmount, txType, txCrudDate)

	return nil
}
