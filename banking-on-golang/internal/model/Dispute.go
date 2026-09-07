package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Dispute Declaration
//==============================================================
type Dispute struct {
    gorm.Model
     DisputeReference                                    string
    RaisedOn                                                            time.Time
    Reason                                    string
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    PaymentCardId         *uint
    PaymentCard           *PaymentCard `gorm:"foreignKey:PaymentCardId"`
    Status                      DisputeStatus

// parent associations as their child

}

