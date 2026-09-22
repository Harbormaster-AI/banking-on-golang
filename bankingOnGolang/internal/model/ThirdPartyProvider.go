
package model

import (
)

//==============================================================
// ThirdPartyProvider Declaration
//==============================================================
type ThirdPartyProvider struct {
    BaseModel
     Name              string
    RegistrationId              string
    Website              string
    BankId           *uint
    Bank             *Bank `gorm:"foreignKey:BankId"`
     Consents             []Consent `gorm:"foreignKey:ConsentsFromThirdPartyProviderId"`

// parent associations as their child
    ThirdPartyProvidersFromBankId    *uint

}

