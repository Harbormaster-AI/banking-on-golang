
package model

import (
    "time"
    "github.com/shopspring/decimal"
)

//==============================================================
// FXTrade Declaration
//==============================================================
type FXTrade struct {
    BaseModel
     TradeReference            string
    TradeDate            time.Time
    SettlementDate            time.Time
    AmountSold        Money `gorm:"embedded;embeddedPrefix:fXTrade_amountSold"`
    AmountBought        Money `gorm:"embedded;embeddedPrefix:fXTrade_amountBought"`
    Rate            decimal.Decimal
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
    ExchangeRateId         *uint
    ExchangeRate           *ExchangeRate `gorm:"foreignKey:ExchangeRateId"`
    SourceAccountId         *uint
    SourceAccount           *Account `gorm:"foreignKey:SourceAccountId"`
    DestinationAccountId         *uint
    DestinationAccount           *Account `gorm:"foreignKey:DestinationAccountId"`
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    Status            TradeStatus

// parent associations as their child

}

