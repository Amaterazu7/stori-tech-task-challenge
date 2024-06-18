package models

import (
	"github.com/google/uuid"
	"time"
)

type AssetTypes string

const (
	CASH   AssetTypes = "CASH"
	CRYPTO AssetTypes = "CRYPTO"
)

type AccountTypes string

const (
	USD AccountTypes = "USD"
	GBP AccountTypes = "GBP"
	EUR AccountTypes = "EUR"
	ETH AccountTypes = "ETH"
	BTC AccountTypes = "BTC"
)

type Account struct {
	Id        uuid.UUID
	Name      string
	Asset     AssetTypes
	Type      AccountTypes
	UpdatedAt time.Time
	CreatedAt time.Time
}
