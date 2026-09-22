
package model

import (
    "time"
)

//==============================================================
// RiskAssessment Declaration
//==============================================================
type RiskAssessment struct {
    BaseModel
     Score            int32
    AssessedOn            time.Time
    KycProfileId         *uint
    KycProfile           *KycProfile `gorm:"foreignKey:KycProfileId"`
    Rating            RiskRating

// parent associations as their child

}

