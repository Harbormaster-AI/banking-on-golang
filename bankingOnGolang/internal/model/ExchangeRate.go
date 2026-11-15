package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// ==============================================================
// ExchangeRate Declaration
// ==============================================================
type ExchangeRate struct {
	BaseModel
	BaseCurrency    string
	CounterCurrency string
	Rate            decimal.Decimal
	AsOf            time.Time
	Source          string
	BankId          *uint
	Bank            *Bank     `gorm:"foreignKey:BankId"`
	FxTrades        []FXTrade `gorm:"foreignKey:FxTradesFromExchangeRateId"`

	// parent associations as their child
	ExchangeRatesFromBankId *uint
}
