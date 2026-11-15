package model

import (
	"time"
)

// ==============================================================
// Account Declaration
// ==============================================================
type Account struct {
	BaseModel
	AccountNumber        AccountNumber `gorm:"embedded;embeddedPrefix:account_accountNumber"`
	Iban                 IBAN          `gorm:"embedded;embeddedPrefix:account_iban"`
	AccountName          string
	Currency             string
	OpenedOn             time.Time
	ClosedOn             time.Time
	BankId               *uint
	Bank                 *Bank `gorm:"foreignKey:BankId"`
	BranchId             *uint
	Branch               *Branch `gorm:"foreignKey:BranchId"`
	ProductId            *uint
	Product              *BankingProduct       `gorm:"foreignKey:ProductId"`
	Owners               []Customer            `gorm:"foreignKey:OwnersFromAccountId"`
	Transactions         []Transaction         `gorm:"foreignKey:TransactionsFromAccountId"`
	Statements           []AccountStatement    `gorm:"foreignKey:StatementsFromAccountId"`
	StandingInstructions []StandingInstruction `gorm:"foreignKey:StandingInstructionsFromAccountId"`
	FeeCharges           []FeeCharge           `gorm:"foreignKey:FeeChargesFromAccountId"`
	AccountType          AccountType
	OwnershipType        AccountOwnershipType
	Status               AccountStatus

	// parent associations as their child
	AccountsFromBankId              *uint
	AccountsFromBranchId            *uint
	AccountsFromCustomerId          *uint
	AccountsFromBankingProductId    *uint
	AuthorizedAccountsFromConsentId *uint
}
