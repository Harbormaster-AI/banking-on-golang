package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Customer Declaration
//==============================================================
type Customer struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    LegalName                                    string
    DateOfBirth                                                            time.Time
    TaxId                                    string
    Email                                    string
    Phone                                    string
    Address                                                            string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromCustomerId"`
     LoanAccounts           []LoanAccount `gorm:"foreignKey:LoanAccountsFromCustomerId"`
     PaymentCards           []PaymentCard `gorm:"foreignKey:PaymentCardsFromCustomerId"`
     ExternalAccounts           []ExternalAccount `gorm:"foreignKey:ExternalAccountsFromCustomerId"`
     FundsTransfers           []FundsTransfer `gorm:"foreignKey:FundsTransfersFromCustomerId"`
     Disputes           []Dispute `gorm:"foreignKey:DisputesFromCustomerId"`
     KycProfiles           []KycProfile `gorm:"foreignKey:KycProfilesFromCustomerId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromCustomerId"`
    CustomerType                      CustomerType
    RiskRating                      RiskRating
    KycStatus                      KycStatus

// parent associations as their child

}

