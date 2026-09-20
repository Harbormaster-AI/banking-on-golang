
package model

import (
)

//==============================================================
// Collateral Declaration
//==============================================================
type Collateral struct {
    BaseModel
     CollateralIdentifier            string
    AppraisedValue            Money
    Description            string
    Location            Address
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    CollateralType            CollateralType

// parent associations as their child

}

