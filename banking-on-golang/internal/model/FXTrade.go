package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// FXTrade Declaration
//==============================================================
type FXTrade struct {
    gorm.Model
     TradeReference                                    string
    TradeDate                                                            time.Time
    SettlementDate                                                            time.Time
    AmountSold                                                            string
    AmountBought                                                            string
    Rate                                                            string
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
    Status                      TradeStatus

// parent associations as their child

}

