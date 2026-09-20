
package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Collateral Declaration
//==============================================================
type Collateral struct {
    gorm.Model
     CollateralIdentifier            string
    AppraisedValue            Money
    Description            string
    Location            Address
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    CollateralType            CollateralType

// parent associations as their child

}

