package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// IdentityDocument Declaration
//==============================================================
type IdentityDocument struct {
    gorm.Model
     DocumentNumber                                    string
    IssuingCountry                                    string
    ExpirationDate                                                            time.Time
    KycProfileId         *uint
    KycProfile           *KycProfile `gorm:"foreignKey:KycProfileId"`
    DocumentType                      IdentityDocumentType

// parent associations as their child

}

