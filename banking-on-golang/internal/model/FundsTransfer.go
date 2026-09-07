package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// FundsTransfer Declaration
//==============================================================
type FundsTransfer struct {
    gorm.Model
     TransferReference                                    string
    Amount                                                            string
    RequestedDate                                                            time.Time
    ExecutionDate                                                            time.Time
    Purpose                                    string
    FeeAmount                                                            string
    SourceAccountId         *uint
    SourceAccount           *Account `gorm:"foreignKey:SourceAccountId"`
    DestinationAccountId         *uint
    DestinationAccount           *Account `gorm:"foreignKey:DestinationAccountId"`
    ExternalBeneficiaryId         *uint
    ExternalBeneficiary           *ExternalAccount `gorm:"foreignKey:ExternalBeneficiaryId"`
    InitiatedById         *uint
    InitiatedBy           *Customer `gorm:"foreignKey:InitiatedById"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromFundsTransferId"`
    Method                      PaymentMethod
    Status                      PaymentStatus

// parent associations as their child

}

