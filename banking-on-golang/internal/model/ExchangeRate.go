package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ExchangeRate Declaration
//==============================================================
type ExchangeRate struct {
    gorm.Model
     BaseCurrency                                    string
    CounterCurrency                                    string
    Rate                                                            string
    AsOf                                                            time.Time
    Source                                    string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     FxTrades           []FXTrade `gorm:"foreignKey:FxTradesFromExchangeRateId"`

// parent associations as their child

}

