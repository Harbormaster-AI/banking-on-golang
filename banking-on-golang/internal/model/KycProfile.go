package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// KycProfile Declaration
//==============================================================
type KycProfile struct {
    gorm.Model
     ProfileId                                    string
    LastReviewedOn                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     IdentityDocuments           []IdentityDocument `gorm:"foreignKey:IdentityDocumentsFromKycProfileId"`
     RiskAssessments           []RiskAssessment `gorm:"foreignKey:RiskAssessmentsFromKycProfileId"`
     Screenings           []ScreeningResult `gorm:"foreignKey:ScreeningsFromKycProfileId"`
    Status                      KycStatus

// parent associations as their child

}

