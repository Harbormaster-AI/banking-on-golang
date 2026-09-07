package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Collateral Declaration
//==============================================================
type Collateral struct {
    gorm.Model
     AppraisedValue                                                            string
    Description                                    string
    Location                                                            string
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    CollateralType                      CollateralType

// parent associations as their child

}

