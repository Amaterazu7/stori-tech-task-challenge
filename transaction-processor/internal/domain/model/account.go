package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	USD = "USD"
	GBP = "GBP"
	EUR = "EUR"
	ETH = "CRYPTO"
	BTC = "CRYPTO"
)

type Account struct {
	Id        uuid.UUID
	Name      string
	Type      string
	Asset     string
	CreatedAt time.Time
}
