
package model

import (
    "time"
)

//==============================================================
// LoanPayment Declaration
//==============================================================
type LoanPayment struct {
    BaseModel
     PaymentReference            string
    Amount            Money
    PaymentDate            time.Time
    LoanAccountId         *uint
    LoanAccount           *LoanAccount `gorm:"foreignKey:LoanAccountId"`
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    Method            PaymentMethod
    Status            PaymentStatus

// parent associations as their child

}

