package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ExternalAccount Declaration
//==============================================================
type ExternalAccount struct {
    gorm.Model
     Name                                    string
    Iban                                                            string
    AccountNumber                                                            string
    Bic                                                            string
    BankName                                    string
    Country                                    string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromExternalAccountId"`

// parent associations as their child

}

