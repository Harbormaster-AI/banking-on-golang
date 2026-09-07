package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Bank Declaration
//==============================================================
type Bank struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    SwiftBic                                                            string
    HeadquartersCountry                                    string
    Website                                    string
     Branches           []Branch `gorm:"foreignKey:BranchesFromBankId"`
     Products           []BankingProduct `gorm:"foreignKey:ProductsFromBankId"`
     Customers           []Customer `gorm:"foreignKey:CustomersFromBankId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromBankId"`
     PaymentCards           []PaymentCard `gorm:"foreignKey:PaymentCardsFromBankId"`
     LoanAccounts           []LoanAccount `gorm:"foreignKey:LoanAccountsFromBankId"`
     ExchangeRates           []ExchangeRate `gorm:"foreignKey:ExchangeRatesFromBankId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromBankId"`
     ThirdPartyProviders           []ThirdPartyProvider `gorm:"foreignKey:ThirdPartyProvidersFromBankId"`

// parent associations as their child

}

