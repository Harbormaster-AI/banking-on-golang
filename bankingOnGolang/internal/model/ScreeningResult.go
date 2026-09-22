
package model

import (
    "time"
)

//==============================================================
// ScreeningResult Declaration
//==============================================================
type ScreeningResult struct {
    BaseModel
     ScreeningDate            time.Time
    Provider            string
    KycProfileId         *uint
    KycProfile           *KycProfile `gorm:"foreignKey:KycProfileId"`
    Outcome            ScreeningOutcome

// parent associations as their child

}

