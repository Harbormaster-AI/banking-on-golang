package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Transaction Declaration
//==============================================================
type Transaction struct {
    gorm.Model
     BookingDate                                                            time.Time
    ValueDate                                                            time.Time
    Amount                                                            string
    Description                                    string
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    ExternalCounterpartyId         *uint
    ExternalCounterparty           *ExternalAccount `gorm:"foreignKey:ExternalCounterpartyId"`
    PaymentCardId         *uint
    PaymentCard           *PaymentCard `gorm:"foreignKey:PaymentCardId"`
    FundsTransferId         *uint
    FundsTransfer           *FundsTransfer `gorm:"foreignKey:FundsTransferId"`
    FxTradeId         *uint
    FxTrade           *FXTrade `gorm:"foreignKey:FxTradeId"`
    DisputeId         *uint
    Dispute           *Dispute `gorm:"foreignKey:DisputeId"`
    Direction                      TransactionDirection
    TransactionType                      TransactionType
    Status                      TransactionStatus
    Channel                      ChannelType

// parent associations as their child

}

