package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// StandingInstruction Declaration
//==============================================================
type StandingInstruction struct {
    gorm.Model
     InstructionId                                    string
    Amount                                                            string
    NextExecutionDate                                                            time.Time
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    BeneficiaryId         *uint
    Beneficiary           *ExternalAccount `gorm:"foreignKey:BeneficiaryId"`
    Frequency                      StandingInstructionFrequency
    Status                      StandingInstructionStatus

// parent associations as their child

}

