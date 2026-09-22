
package model

import (
    "time"
)

//==============================================================
// LoanAccount Declaration
//==============================================================
type LoanAccount struct {
    BaseModel
     LoanNumber            string
    PrincipalAmount        Money `gorm:"embedded;embeddedPrefix:loanAccount_principalAmount"`
    OutstandingPrincipal        Money `gorm:"embedded;embeddedPrefix:loanAccount_outstandingPrincipal"`
    InterestRate        Percentage `gorm:"embedded;embeddedPrefix:loanAccount_interestRate"`
    OriginationDate            time.Time
    MaturityDate            time.Time
    PaymentDayOfMonth            int32
    Currency            string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
    BranchId         *uint
    Branch           *Branch `gorm:"foreignKey:BranchId"`
    ProductId         *uint
    Product           *BankingProduct `gorm:"foreignKey:ProductId"`
     Borrowers           []Customer `gorm:"foreignKey:BorrowersFromLoanAccountId"`
     RepaymentSchedule           []RepaymentSchedule `gorm:"foreignKey:RepaymentScheduleFromLoanAccountId"`
     Payments           []LoanPayment `gorm:"foreignKey:PaymentsFromLoanAccountId"`
     Collateral           []Collateral `gorm:"foreignKey:CollateralFromLoanAccountId"`
     FeeCharges           []FeeCharge `gorm:"foreignKey:FeeChargesFromLoanAccountId"`
    LoanType            LoanType
    RateType            RateType
    Compounding            InterestCompounding
    Status            LoanStatus

// parent associations as their child

}

