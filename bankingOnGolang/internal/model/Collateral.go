package model

import ()

// ==============================================================
// Collateral Declaration
// ==============================================================
type Collateral struct {
	BaseModel
	CollateralIdentifier string
	AppraisedValue       Money `gorm:"embedded;embeddedPrefix:collateral_appraisedValue"`
	Description          string
	Location             Address `gorm:"embedded;embeddedPrefix:collateral_location"`
	LoanAccountId        *uint
	LoanAccount          *LoanAccount `gorm:"foreignKey:LoanAccountId"`
	CollateralType       CollateralType

	// parent associations as their child
	CollateralFromLoanAccountId *uint
}
