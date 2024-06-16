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
	Value     float64
	Type      string
	Reference string
	CreatedAt time.Time
}
