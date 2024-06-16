package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	DEBIT  = "DEBIT"
	CREDIT = "CREDIT"
)

type Transaction struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	Amount    float64
	TxType    string
	Reference string
	CreatedAt time.Time
}

func (tx Transaction) Build(
	id uuid.UUID, accountId uuid.UUID, amount float64, txType string, reference string, createdAt time.Time,
) {
	tx.Id = id
	tx.AccountId = accountId
	tx.Amount = amount
	tx.TxType = txType
	tx.Reference = reference
	tx.CreatedAt = createdAt
}
