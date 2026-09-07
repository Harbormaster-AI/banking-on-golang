package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Account Declaration
//==============================================================
type Account struct {
    gorm.Model
     AccountNumber                                                            string
    Iban                                                            string
    AccountName                                    string
    Currency                                    string
    OpenedOn                                                            time.Time
    ClosedOn                                                            time.Time
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
    BranchId         *uint
    Branch           *Branch `gorm:"foreignKey:BranchId"`
    ProductId         *uint
    Product           *BankingProduct `gorm:"foreignKey:ProductId"`
     Owners           []Customer `gorm:"foreignKey:OwnersFromAccountId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromAccountId"`
     Statements           []AccountStatement `gorm:"foreignKey:StatementsFromAccountId"`
     StandingInstructions           []StandingInstruction `gorm:"foreignKey:StandingInstructionsFromAccountId"`
     FeeCharges           []FeeCharge `gorm:"foreignKey:FeeChargesFromAccountId"`
    AccountType                      AccountType
    OwnershipType                      AccountOwnershipType
    Status                      AccountStatus

// parent associations as their child

}

