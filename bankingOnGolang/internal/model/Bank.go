
package model

import (
)

//==============================================================
// Bank Declaration
//==============================================================
type Bank struct {
    BaseModel
     Name              string
    LegalName              string
    SwiftBic          BIC `gorm:"embedded;embeddedPrefix:bank_swiftBic"`
    HeadquartersCountry              string
    Website              string
     Branches             []Branch `gorm:"foreignKey:BranchesFromBankId"`
     Products             []BankingProduct `gorm:"foreignKey:ProductsFromBankId"`
     Customers             []Customer `gorm:"foreignKey:CustomersFromBankId"`
     Accounts             []Account `gorm:"foreignKey:AccountsFromBankId"`
     PaymentCards             []PaymentCard `gorm:"foreignKey:PaymentCardsFromBankId"`
     LoanAccounts             []LoanAccount `gorm:"foreignKey:LoanAccountsFromBankId"`
     ExchangeRates             []ExchangeRate `gorm:"foreignKey:ExchangeRatesFromBankId"`
     Consents             []Consent `gorm:"foreignKey:ConsentsFromBankId"`
     ThirdPartyProviders             []ThirdPartyProvider `gorm:"foreignKey:ThirdPartyProvidersFromBankId"`

// parent associations as their child

}

