
package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ATM Declaration
//==============================================================
type ATM struct {
    gorm.Model
     TerminalId            string
    Location            Address
    BranchId         *uint
    Branch           *Branch `gorm:"foreignKey:BranchId"`
    Status            ATMStatus

// parent associations as their child

}

