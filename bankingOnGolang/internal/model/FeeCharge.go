
package model

import (
    "time"
)

//==============================================================
// FeeCharge Declaration
//==============================================================
type FeeCharge struct {
    BaseModel
     FeeCode              string
    Amount          Money `gorm:"embedded;embeddedPrefix:feeCharge_amount"`
    AppliedOn              time.Time
    AccountId           *uint
    Account             *Account `gorm:"foreignKey:AccountId"`
    LoanAccountId           *uint
    LoanAccount             *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    FeeType              FeeType

// parent associations as their child
    FeeChargesFromAccountId    *uint
    FeeChargesFromLoanAccountId    *uint

}

