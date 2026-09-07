package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ScreeningResult Declaration
//==============================================================
type ScreeningResult struct {
    gorm.Model
     ScreeningDate                                                            time.Time
    Provider                                    string
    KycProfileId         *uint
    KycProfile           *KycProfile `gorm:"foreignKey:KycProfileId"`
    Outcome                      ScreeningOutcome

// parent associations as their child

}

