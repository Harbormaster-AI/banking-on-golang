package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BankingProduct Declaration
//==============================================================
type BankingProduct struct {
    gorm.Model
     ProductCode                                    string
    Name                                    string
    Description                                    string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromBankingProductId"`
     LoanAccounts           []LoanAccount `gorm:"foreignKey:LoanAccountsFromBankingProductId"`
     PaymentCards           []PaymentCard `gorm:"foreignKey:PaymentCardsFromBankingProductId"`
    ProductCategory                      ProductCategory

// parent associations as their child

}

