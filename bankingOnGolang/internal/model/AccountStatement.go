
package model

import (
    "time"
)

//==============================================================
// AccountStatement Declaration
//==============================================================
type AccountStatement struct {
    BaseModel
     StatementNumber            string
    PeriodStart            time.Time
    PeriodEnd            time.Time
    OpeningBalance            Money
    ClosingBalance            Money
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    DeliveryMethod            StatementDeliveryMethod

// parent associations as their child

}

