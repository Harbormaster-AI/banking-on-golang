
package model

import (
)

//==============================================================
// ExternalAccount Declaration
//==============================================================
type ExternalAccount struct {
    BaseModel
     Name              string
    Iban          IBAN `gorm:"embedded;embeddedPrefix:externalAccount_iban"`
    AccountNumber          AccountNumber `gorm:"embedded;embeddedPrefix:externalAccount_accountNumber"`
    Bic          BIC `gorm:"embedded;embeddedPrefix:externalAccount_bic"`
    BankName              string
    Country              string
    CustomerId           *uint
    Customer             *Customer `gorm:"foreignKey:CustomerId"`
     Transactions             []Transaction `gorm:"foreignKey:TransactionsFromExternalAccountId"`

// parent associations as their child
    ExternalAccountsFromCustomerId    *uint

}

