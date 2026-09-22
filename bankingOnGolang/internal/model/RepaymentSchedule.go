
package model

import (
    "time"
)

//==============================================================
// RepaymentSchedule Declaration
//==============================================================
type RepaymentSchedule struct {
    BaseModel
     InstallmentNumber              int32
    DueDate              time.Time
    PrincipalDue          Money `gorm:"embedded;embeddedPrefix:repaymentSchedule_principalDue"`
    InterestDue          Money `gorm:"embedded;embeddedPrefix:repaymentSchedule_interestDue"`
    TotalDue          Money `gorm:"embedded;embeddedPrefix:repaymentSchedule_totalDue"`
    LoanAccountId           *uint
    LoanAccount             *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    PaymentId           *uint
    Payment             *LoanPayment `gorm:"foreignKey:PaymentId"`
    Status              InstallmentStatus

// parent associations as their child
    RepaymentScheduleFromLoanAccountId    *uint

}

