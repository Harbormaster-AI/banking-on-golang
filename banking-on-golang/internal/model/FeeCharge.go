package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// FeeCharge Declaration
//==============================================================
type FeeCharge struct {
    gorm.Model
     FeeCode                                    string
    Amount                                                            string
    AppliedOn                                                            time.Time
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    FeeType                      FeeType

// parent associations as their child

}

