
package model

import (
)

//==============================================================
// ExternalAccount Declaration
//==============================================================
type ExternalAccount struct {
    BaseModel
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

