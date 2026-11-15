package model

import ()

// ==============================================================
// ATM Declaration
// ==============================================================
type ATM struct {
	BaseModel
	TerminalId string
	Location   Address `gorm:"embedded;embeddedPrefix:aTM_location"`
	BranchId   *uint
	Branch     *Branch `gorm:"foreignKey:BranchId"`
	Status     ATMStatus

	// parent associations as their child
	AtmsFromBranchId *uint
}
