
package model

import (
    "time"
)

//==============================================================
// StandingInstruction Declaration
//==============================================================
type StandingInstruction struct {
    BaseModel
     InstructionId              string
    Amount          Money `gorm:"embedded;embeddedPrefix:standingInstruction_amount"`
    NextExecutionDate              time.Time
    AccountId           *uint
    Account             *Account `gorm:"foreignKey:AccountId"`
    BeneficiaryId           *uint
    Beneficiary             *ExternalAccount `gorm:"foreignKey:BeneficiaryId"`
    Frequency              StandingInstructionFrequency
    Status              StandingInstructionStatus

// parent associations as their child
    StandingInstructionsFromAccountId    *uint

}

