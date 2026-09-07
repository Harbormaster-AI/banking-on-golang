package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LoanAccount Declaration
//==============================================================
type LoanAccount struct {
    gorm.Model
     LoanNumber                                    string
    PrincipalAmount                                                            string
    OutstandingPrincipal                                                            string
    InterestRate                                                            string
    OriginationDate                                                            time.Time
    MaturityDate                                                            time.Time
    PaymentDayOfMonth                                    int32
    Currency                                    string
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
    LoanType                      LoanType
    RateType                      RateType
    Compounding                      InterestCompounding
    Status                      LoanStatus

// parent associations as their child

}

