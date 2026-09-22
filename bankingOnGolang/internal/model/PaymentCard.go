
package model

import (
)

//==============================================================
// PaymentCard Declaration
//==============================================================
type PaymentCard struct {
    BaseModel
     CardNumber        CardPAN `gorm:"embedded;embeddedPrefix:paymentCard_cardNumber"`
    EmbossedName            string
    ExpiryMonth            int32
    ExpiryYear            int32
    BankId         *uint
    Bank           *Bank `gorm:"foreignKey:BankId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromPaymentCardId"`
    CardType            CardType
    CardStatus            CardStatus
    Network            CardNetwork

// parent associations as their child

}

