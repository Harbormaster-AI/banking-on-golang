
package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ExternalAccount Declaration
//==============================================================
type ExternalAccount struct {
    gorm.Model
     Name            string
    Iban            IBAN
    AccountNumber            AccountNumber
    Bic            BIC
    BankName            string
    Country            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromExternalAccountId"`

// parent associations as their child

}

