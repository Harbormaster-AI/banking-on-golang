
package model

import (
    "time"
)

//==============================================================
// IdentityDocument Declaration
//==============================================================
type IdentityDocument struct {
    BaseModel
     DocumentNumber              string
    IssuingCountry              string
    ExpirationDate              time.Time
    KycProfileId           *uint
    KycProfile             *KycProfile `gorm:"foreignKey:KycProfileId"`
    DocumentType              IdentityDocumentType

// parent associations as their child
    IdentityDocumentsFromKycProfileId    *uint

}

