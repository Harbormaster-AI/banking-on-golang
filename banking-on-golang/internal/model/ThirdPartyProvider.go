package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ThirdPartyProvider Declaration
//==============================================================
type ThirdPartyProvider struct {
    gorm.Model
     Name                                    string
    RegistrationId                                    string
    Website                                    string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromThirdPartyProviderId"`

// parent associations as their child

}

