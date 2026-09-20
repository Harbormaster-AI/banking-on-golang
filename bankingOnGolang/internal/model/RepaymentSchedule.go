
package model

import (
    "time"
)

//==============================================================
// RepaymentSchedule Declaration
//==============================================================
type RepaymentSchedule struct {
    BaseModel
     InstallmentNumber            int32
    DueDate            time.Time
    PrincipalDue            Money
    InterestDue            Money
    TotalDue            Money
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    PaymentId         *uint
    Payment           *LoanPayment `gorm:"foreignKey:PaymentId"`
    Status            InstallmentStatus

// parent associations as their child

}

