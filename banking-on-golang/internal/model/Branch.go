package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Branch Declaration
//==============================================================
type Branch struct {
    gorm.Model
     Name                                    string
    BranchCode                                    string
    Address                                                            string
    Phone                                    string
    OpeningHours                                    string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromBranchId"`
     LoanAccounts           []LoanAccount `gorm:"foreignKey:LoanAccountsFromBranchId"`
     Atms           []ATM `gorm:"foreignKey:AtmsFromBranchId"`

// parent associations as their child

}

