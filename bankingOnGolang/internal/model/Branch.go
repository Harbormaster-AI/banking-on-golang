
package model

import (
)

//==============================================================
// Branch Declaration
//==============================================================
type Branch struct {
    BaseModel
     Name            string
    BranchCode            string
    Address        Address `gorm:"embedded;embeddedPrefix:branch_address"`
    Phone            string
    OpeningHours            string
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromBranchId"`
     LoanAccounts           []LoanAccount `gorm:"foreignKey:LoanAccountsFromBranchId"`
     Atms           []ATM `gorm:"foreignKey:AtmsFromBranchId"`

// parent associations as their child

}

