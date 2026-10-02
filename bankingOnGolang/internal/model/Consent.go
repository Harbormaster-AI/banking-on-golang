package model

import (
	"time"
)

// ==============================================================
// Consent Declaration
// ==============================================================
type Consent struct {
	BaseModel
	GrantedOn            time.Time
	ExpiresOn            time.Time
	CustomerId           *uint
	Customer             *Customer `gorm:"foreignKey:CustomerId"`
	BankId               *uint
	Bank                 *Bank     `gorm:"foreignKey:BankId"`
	AuthorizedAccounts   []Account `gorm:"foreignKey:AuthorizedAccountsFromConsentId"`
	ThirdPartyProviderId *uint
	ThirdPartyProvider   *ThirdPartyProvider `gorm:"foreignKey:ThirdPartyProviderId"`
	ConsentType          ConsentType
	Status               ConsentStatus

	// parent associations as their child
	ConsentsFromBankId               *uint
	ConsentsFromCustomerId           *uint
	ConsentsFromThirdPartyProviderId *uint
}
