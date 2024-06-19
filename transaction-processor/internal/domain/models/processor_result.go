package models

import (
	"math"
)

type ProcessorResult struct {
	AccountId           string
	email               string
	TotalBalance        float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	MonthMap            map[string]int
}

func (pr *ProcessorResult) SetEmail(email string) {
	pr.email = email
}

func (pr *ProcessorResult) GetEmail() string {
	return pr.email
}

func (pr *ProcessorResult) SetTotalBalance(totalBalance float64) {
	pr.TotalBalance = totalBalance
}

func (pr *ProcessorResult) AddBalance(totalBalance float64) {
	pr.TotalBalance += totalBalance
}

func (pr *ProcessorResult) FillAvgValues(tx *Transaction, debitAmount, debitCount, creditAmount, creditCount *float64) {
	if tx.TxType == DEBIT {
		*debitAmount += tx.Amount
		*debitCount++
	} else {
		*creditAmount += tx.Amount
		*creditCount++
	}
}

func (pr *ProcessorResult) CalculateAverageDebit(debitAmount float64, debitCount float64) {
	pr.AverageDebitAmount = math.Round((debitAmount/debitCount)*100) / 100
}

func (pr *ProcessorResult) CalculateAverageCredit(creditAmount float64, creditCount float64) {
	pr.AverageCreditAmount = math.Round((creditAmount/creditCount)*100) / 100
}

func (pr *ProcessorResult) FillMap(tx *Transaction) {
	val, _ := pr.MonthMap[tx.CreatedAt.Month().String()]
	pr.MonthMap[tx.CreatedAt.Month().String()] = val + 1
}
