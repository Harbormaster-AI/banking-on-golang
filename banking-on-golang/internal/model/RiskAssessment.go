package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RiskAssessment Declaration
//==============================================================
type RiskAssessment struct {
    gorm.Model
     Score                                    int32
    AssessedOn                                                            time.Time
    KycProfileId         *uint
    KycProfile           *KycProfile `gorm:"foreignKey:KycProfileId"`
    Rating                      RiskRating

// parent associations as their child

}

