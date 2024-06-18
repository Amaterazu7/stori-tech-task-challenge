package models

import (
	"github.com/google/uuid"
	"time"
)

type TxTypes string

const (
	DEBIT  TxTypes = "DEBIT"
	CREDIT TxTypes = "CREDIT"
)

type Transaction struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	Amount    float64
	TxType    TxTypes
	CreatedAt time.Time
}

// TODO::  Add builderPatter
func (tx Transaction) Build(
	id uuid.UUID, accountId uuid.UUID, amount float64, txType TxTypes, createdAt time.Time,
) {
	tx.Id = id
	tx.AccountId = accountId
	tx.Amount = amount
	tx.TxType = txType
	tx.CreatedAt = createdAt
}
