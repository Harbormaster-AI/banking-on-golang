
package model

import (
)

//==============================================================
// ATM Declaration
//==============================================================
type ATM struct {
    BaseModel
     TerminalId            string
    Location            Address
    BranchId         *uint
    Branch           *Branch `gorm:"foreignKey:BranchId"`
    Status            ATMStatus

// parent associations as their child

}

