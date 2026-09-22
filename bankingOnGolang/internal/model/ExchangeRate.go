
package model

import (
    "time"
    "github.com/shopspring/decimal"
)

//==============================================================
// ExchangeRate Declaration
//==============================================================
type ExchangeRate struct {
    BaseModel
     BaseCurrency              string
    CounterCurrency              string
    Rate              decimal.Decimal
    AsOf              time.Time
    Source              string
    BankId           *uint
    Bank             *Bank `gorm:"foreignKey:BankId"`
     FxTrades             []FXTrade `gorm:"foreignKey:FxTradesFromExchangeRateId"`

// parent associations as their child
    ExchangeRatesFromBankId    *uint

}

