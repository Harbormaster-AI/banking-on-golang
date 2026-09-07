package model


//==============================================================
// CustomerType Declaration
//==============================================================
type CustomerType int
const (
    CustomerTypeIndividual CustomerType = iota
	CustomerTypeBusiness
	CustomerTypeNonProfit
	CustomerTypeGovernment
)


//==============================================================
// AccountType Declaration
//==============================================================
type AccountType int
const (
    AccountTypeChecking AccountType = iota
	AccountTypeSavings
	AccountTypeMoneyMarket
	AccountTypeTimeDeposit
)


//==============================================================
// AccountStatus Declaration
//==============================================================
type AccountStatus int
const (
    AccountStatusOpen AccountStatus = iota
	AccountStatusFrozen
	AccountStatusDormant
	AccountStatusClosed
)


//==============================================================
// AccountOwnershipType Declaration
//==============================================================
type AccountOwnershipType int
const (
    AccountOwnershipTypeSole AccountOwnershipType = iota
	AccountOwnershipTypeJoint
	AccountOwnershipTypeCorporate
	AccountOwnershipTypeTrust
)


//==============================================================
// StatementDeliveryMethod Declaration
//==============================================================
type StatementDeliveryMethod int
const (
    StatementDeliveryMethodElectronic StatementDeliveryMethod = iota
	StatementDeliveryMethodPaper
)


//==============================================================
// TransactionType Declaration
//==============================================================
type TransactionType int
const (
    TransactionTypeDeposit TransactionType = iota
	TransactionTypeWithdrawal
	TransactionTypeTransfer
	TransactionTypePayment
	TransactionTypeFee
	TransactionTypeInterest
	TransactionTypeAdjustment
	TransactionTypeChargeback
	TransactionTypeRefund
	TransactionTypeFXConversion
)


//==============================================================
// TransactionStatus Declaration
//==============================================================
type TransactionStatus int
const (
    TransactionStatusPending TransactionStatus = iota
	TransactionStatusPosted
	TransactionStatusReversed
	TransactionStatusFailed
	TransactionStatusCancelled
)


//==============================================================
// TransactionDirection Declaration
//==============================================================
type TransactionDirection int
const (
    TransactionDirectionCredit TransactionDirection = iota
	TransactionDirectionDebit
)


//==============================================================
// ChannelType Declaration
//==============================================================
type ChannelType int
const (
    ChannelTypeBranch ChannelType = iota
	ChannelTypeOnline
	ChannelTypeMobile
	ChannelTypeATM
	ChannelTypeAPI
	ChannelTypeCallCenter
)


//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod int
const (
    PaymentMethodInternalTransfer PaymentMethod = iota
	PaymentMethodACH
	PaymentMethodWire
	PaymentMethodSEPA
	PaymentMethodSWIFT
	PaymentMethodCard
	PaymentMethodCash
	PaymentMethodCheck
	PaymentMethodMobileWallet
)


//==============================================================
// PaymentStatus Declaration
//==============================================================
type PaymentStatus int
const (
    PaymentStatusInitiated PaymentStatus = iota
	PaymentStatusInProcess
	PaymentStatusSettled
	PaymentStatusFailed
	PaymentStatusReversed
	PaymentStatusCancelled
)


//==============================================================
// StandingInstructionFrequency Declaration
//==============================================================
type StandingInstructionFrequency int
const (
    StandingInstructionFrequencyOneTime StandingInstructionFrequency = iota
	StandingInstructionFrequencyWeekly
	StandingInstructionFrequencyBiWeekly
	StandingInstructionFrequencyMonthly
	StandingInstructionFrequencyQuarterly
	StandingInstructionFrequencyAnnually
)


//==============================================================
// StandingInstructionStatus Declaration
//==============================================================
type StandingInstructionStatus int
const (
    StandingInstructionStatusActive StandingInstructionStatus = iota
	StandingInstructionStatusPaused
	StandingInstructionStatusCancelled
	StandingInstructionStatusCompleted
)


//==============================================================
// CardType Declaration
//==============================================================
type CardType int
const (
    CardTypeDebit CardType = iota
	CardTypeCredit
	CardTypePrepaid
	CardTypeVirtual
)


//==============================================================
// CardStatus Declaration
//==============================================================
type CardStatus int
const (
    CardStatusActive CardStatus = iota
	CardStatusBlocked
	CardStatusLostStolen
	CardStatusExpired
	CardStatusClosed
)


//==============================================================
// CardNetwork Declaration
//==============================================================
type CardNetwork int
const (
    CardNetworkVisa CardNetwork = iota
	CardNetworkMastercard
	CardNetworkAmex
	CardNetworkDiscover
	CardNetworkUnionPay
	CardNetworkOther
)


//==============================================================
// LoanType Declaration
//==============================================================
type LoanType int
const (
    LoanTypeMortgage LoanType = iota
	LoanTypePersonal
	LoanTypeAuto
	LoanTypeSmallBusiness
	LoanTypeCreditLine
	LoanTypeStudent
)


//==============================================================
// LoanStatus Declaration
//==============================================================
type LoanStatus int
const (
    LoanStatusApplied LoanStatus = iota
	LoanStatusApproved
	LoanStatusActive
	LoanStatusDelinquent
	LoanStatusDefaulted
	LoanStatusClosed
)


//==============================================================
// RateType Declaration
//==============================================================
type RateType int
const (
    RateTypeFixed RateType = iota
	RateTypeVariable
)


//==============================================================
// InterestCompounding Declaration
//==============================================================
type InterestCompounding int
const (
    InterestCompoundingDaily InterestCompounding = iota
	InterestCompoundingMonthly
	InterestCompoundingQuarterly
	InterestCompoundingAnnually
)


//==============================================================
// InstallmentStatus Declaration
//==============================================================
type InstallmentStatus int
const (
    InstallmentStatusDue InstallmentStatus = iota
	InstallmentStatusPaid
	InstallmentStatusOverdue
	InstallmentStatusDeferred
)


//==============================================================
// FeeType Declaration
//==============================================================
type FeeType int
const (
    FeeTypeMaintenance FeeType = iota
	FeeTypeOverdraft
	FeeTypeWire
	FeeTypeATM
	FeeTypeCardAnnual
	FeeTypeLatePayment
	FeeTypeEarlyWithdrawal
	FeeTypeReplacementCard
)


//==============================================================
// RiskRating Declaration
//==============================================================
type RiskRating int
const (
    RiskRatingLow RiskRating = iota
	RiskRatingMedium
	RiskRatingHigh
)


//==============================================================
// KycStatus Declaration
//==============================================================
type KycStatus int
const (
    KycStatusPending KycStatus = iota
	KycStatusVerified
	KycStatusRejected
	KycStatusExpired
)


//==============================================================
// IdentityDocumentType Declaration
//==============================================================
type IdentityDocumentType int
const (
    IdentityDocumentTypePassport IdentityDocumentType = iota
	IdentityDocumentTypeNationalID
	IdentityDocumentTypeDriverLicense
	IdentityDocumentTypeResidencePermit
	IdentityDocumentTypeBusinessRegistration
	IdentityDocumentTypeTaxCertificate
)


//==============================================================
// ScreeningOutcome Declaration
//==============================================================
type ScreeningOutcome int
const (
    ScreeningOutcomeClear ScreeningOutcome = iota
	ScreeningOutcomeMatch
	ScreeningOutcomeReview
)


//==============================================================
// TradeStatus Declaration
//==============================================================
type TradeStatus int
const (
    TradeStatusBooked TradeStatus = iota
	TradeStatusSettled
	TradeStatusCancelled
)


//==============================================================
// ATMStatus Declaration
//==============================================================
type ATMStatus int
const (
    ATMStatusInService ATMStatus = iota
	ATMStatusOutOfService
	ATMStatusMaintenance
)


//==============================================================
// ConsentType Declaration
//==============================================================
type ConsentType int
const (
    ConsentTypeOpenBanking ConsentType = iota
	ConsentTypePaymentInitiation
	ConsentTypeAccountInformation
	ConsentTypeMarketing
	ConsentTypeDataSharing
)


//==============================================================
// ConsentStatus Declaration
//==============================================================
type ConsentStatus int
const (
    ConsentStatusActive ConsentStatus = iota
	ConsentStatusRevoked
	ConsentStatusExpired
)


//==============================================================
// DisputeStatus Declaration
//==============================================================
type DisputeStatus int
const (
    DisputeStatusOpen DisputeStatus = iota
	DisputeStatusUnderReview
	DisputeStatusResolved
	DisputeStatusRejected
	DisputeStatusWithdrawn
)


//==============================================================
// ProductCategory Declaration
//==============================================================
type ProductCategory int
const (
    ProductCategoryDeposit ProductCategory = iota
	ProductCategoryLoan
	ProductCategoryCard
	ProductCategoryPaymentService
	ProductCategoryInvestment
)


//==============================================================
// CollateralType Declaration
//==============================================================
type CollateralType int
const (
    CollateralTypeRealEstate CollateralType = iota
	CollateralTypeVehicle
	CollateralTypeCash
	CollateralTypeSecurities
	CollateralTypeGuarantee
	CollateralTypeEquipment
)

