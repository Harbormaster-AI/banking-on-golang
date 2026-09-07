package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RepaymentSchedule Declaration
//==============================================================
type RepaymentSchedule struct {
    gorm.Model
     InstallmentNumber                                    int32
    DueDate                                                            time.Time
    PrincipalDue                                                            string
    InterestDue                                                            string
    TotalDue                                                            string
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    PaymentId         *uint
    Payment           *LoanPayment `gorm:"foreignKey:PaymentId"`
    Status                      InstallmentStatus

// parent associations as their child

}

