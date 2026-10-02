package router

import (
	controller "bankingOnGolang/internal/controller"
	jsonResponseFormatter "bankingOnGolang/internal/response"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	//----------------------------------------------------------------------------
	// default controllers for health and availability checking
	//----------------------------------------------------------------------------

	router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(controller.Default__)).Methods("GET", "OPTIONS")
	router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(controller.Health__)).Methods("GET", "OPTIONS")

	//----------------------------------------------------------------------------
	// Bank Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Bank/get", jsonResponseFormatter.FormatToJSON(controller.GetBank)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/BankgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllBank)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Bank/create", jsonResponseFormatter.FormatToJSON(controller.CreateBank)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Bank/update", jsonResponseFormatter.FormatToJSON(controller.UpdateBank)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteBank/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteBank)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Bank/addToBranches/", jsonResponseFormatter.FormatToJSON(controller.AddBranchesToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromBranches/", jsonResponseFormatter.FormatToJSON(controller.RemoveBranchesFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToProducts/", jsonResponseFormatter.FormatToJSON(controller.AddProductsToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromProducts/", jsonResponseFormatter.FormatToJSON(controller.RemoveProductsFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToCustomers/", jsonResponseFormatter.FormatToJSON(controller.AddCustomersToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromCustomers/", jsonResponseFormatter.FormatToJSON(controller.RemoveCustomersFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddAccountsToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveAccountsFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.AddPaymentCardsToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.RemovePaymentCardsFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddLoanAccountsToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveLoanAccountsFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToExchangeRates/", jsonResponseFormatter.FormatToJSON(controller.AddExchangeRatesToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromExchangeRates/", jsonResponseFormatter.FormatToJSON(controller.RemoveExchangeRatesFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.AddConsentsToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.RemoveConsentsFromBank)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Bank/addToThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(controller.AddThirdPartyProvidersToBank)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankremoveFromThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(controller.RemoveThirdPartyProvidersFromBank)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Branch Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Branch/get", jsonResponseFormatter.FormatToJSON(controller.GetBranch)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/BranchgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllBranch)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Branch/create", jsonResponseFormatter.FormatToJSON(controller.CreateBranch)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Branch/update", jsonResponseFormatter.FormatToJSON(controller.UpdateBranch)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteBranch/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteBranch)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Branch/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToBranch)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Branch/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromBranch)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Branch/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddAccountsToBranch)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BranchremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveAccountsFromBranch)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Branch/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddLoanAccountsToBranch)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BranchremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveLoanAccountsFromBranch)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Branch/addToAtms/", jsonResponseFormatter.FormatToJSON(controller.AddAtmsToBranch)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BranchremoveFromAtms/", jsonResponseFormatter.FormatToJSON(controller.RemoveAtmsFromBranch)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// ATM Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ATM/get", jsonResponseFormatter.FormatToJSON(controller.GetATM)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ATMgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllATM)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ATM/create", jsonResponseFormatter.FormatToJSON(controller.CreateATM)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ATM/update", jsonResponseFormatter.FormatToJSON(controller.UpdateATM)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteATM/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteATM)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ATM/assignBranch", jsonResponseFormatter.FormatToJSON(controller.AssignBranchToATM)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ATM/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.UnassignBranchFromATM)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Customer Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Customer/get", jsonResponseFormatter.FormatToJSON(controller.GetCustomer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/CustomergetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllCustomer)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Customer/create", jsonResponseFormatter.FormatToJSON(controller.CreateCustomer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Customer/update", jsonResponseFormatter.FormatToJSON(controller.UpdateCustomer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteCustomer/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteCustomer)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Customer/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Customer/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromCustomer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Customer/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddAccountsToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddLoanAccountsToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveLoanAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.AddPaymentCardsToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.RemovePaymentCardsFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToExternalAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddExternalAccountsToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromExternalAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveExternalAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToFundsTransfers/", jsonResponseFormatter.FormatToJSON(controller.AddFundsTransfersToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromFundsTransfers/", jsonResponseFormatter.FormatToJSON(controller.RemoveFundsTransfersFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToDisputes/", jsonResponseFormatter.FormatToJSON(controller.AddDisputesToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromDisputes/", jsonResponseFormatter.FormatToJSON(controller.RemoveDisputesFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToKycProfiles/", jsonResponseFormatter.FormatToJSON(controller.AddKycProfilesToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromKycProfiles/", jsonResponseFormatter.FormatToJSON(controller.RemoveKycProfilesFromCustomer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Customer/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.AddConsentsToCustomer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/CustomerremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.RemoveConsentsFromCustomer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// KycProfile Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/KycProfile/get", jsonResponseFormatter.FormatToJSON(controller.GetKycProfile)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/KycProfilegetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllKycProfile)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/KycProfile/create", jsonResponseFormatter.FormatToJSON(controller.CreateKycProfile)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/KycProfile/update", jsonResponseFormatter.FormatToJSON(controller.UpdateKycProfile)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteKycProfile/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteKycProfile)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/KycProfile/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToKycProfile)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/KycProfile/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromKycProfile)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/KycProfile/addToIdentityDocuments/", jsonResponseFormatter.FormatToJSON(controller.AddIdentityDocumentsToKycProfile)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/KycProfileremoveFromIdentityDocuments/", jsonResponseFormatter.FormatToJSON(controller.RemoveIdentityDocumentsFromKycProfile)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/KycProfile/addToRiskAssessments/", jsonResponseFormatter.FormatToJSON(controller.AddRiskAssessmentsToKycProfile)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/KycProfileremoveFromRiskAssessments/", jsonResponseFormatter.FormatToJSON(controller.RemoveRiskAssessmentsFromKycProfile)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/KycProfile/addToScreenings/", jsonResponseFormatter.FormatToJSON(controller.AddScreeningsToKycProfile)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/KycProfileremoveFromScreenings/", jsonResponseFormatter.FormatToJSON(controller.RemoveScreeningsFromKycProfile)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// IdentityDocument Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/IdentityDocument/get", jsonResponseFormatter.FormatToJSON(controller.GetIdentityDocument)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/IdentityDocumentgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllIdentityDocument)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/IdentityDocument/create", jsonResponseFormatter.FormatToJSON(controller.CreateIdentityDocument)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/IdentityDocument/update", jsonResponseFormatter.FormatToJSON(controller.UpdateIdentityDocument)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteIdentityDocument/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteIdentityDocument)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/IdentityDocument/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.AssignKycProfileToIdentityDocument)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/IdentityDocument/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.UnassignKycProfileFromIdentityDocument)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// RiskAssessment Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/RiskAssessment/get", jsonResponseFormatter.FormatToJSON(controller.GetRiskAssessment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/RiskAssessmentgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllRiskAssessment)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/RiskAssessment/create", jsonResponseFormatter.FormatToJSON(controller.CreateRiskAssessment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/RiskAssessment/update", jsonResponseFormatter.FormatToJSON(controller.UpdateRiskAssessment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteRiskAssessment/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteRiskAssessment)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/RiskAssessment/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.AssignKycProfileToRiskAssessment)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/RiskAssessment/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.UnassignKycProfileFromRiskAssessment)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// ScreeningResult Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ScreeningResult/get", jsonResponseFormatter.FormatToJSON(controller.GetScreeningResult)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ScreeningResultgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllScreeningResult)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ScreeningResult/create", jsonResponseFormatter.FormatToJSON(controller.CreateScreeningResult)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ScreeningResult/update", jsonResponseFormatter.FormatToJSON(controller.UpdateScreeningResult)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteScreeningResult/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteScreeningResult)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ScreeningResult/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.AssignKycProfileToScreeningResult)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ScreeningResult/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.UnassignKycProfileFromScreeningResult)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// BankingProduct Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/BankingProduct/get", jsonResponseFormatter.FormatToJSON(controller.GetBankingProduct)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/BankingProductgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllBankingProduct)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/BankingProduct/create", jsonResponseFormatter.FormatToJSON(controller.CreateBankingProduct)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/BankingProduct/update", jsonResponseFormatter.FormatToJSON(controller.UpdateBankingProduct)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteBankingProduct/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteBankingProduct)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/BankingProduct/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToBankingProduct)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankingProduct/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromBankingProduct)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/BankingProduct/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankingProductremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/BankingProduct/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddLoanAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankingProductremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveLoanAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/BankingProduct/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.AddPaymentCardsToBankingProduct)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/BankingProductremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.RemovePaymentCardsFromBankingProduct)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Account Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Account/get", jsonResponseFormatter.FormatToJSON(controller.GetAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/AccountgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllAccount)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Account/create", jsonResponseFormatter.FormatToJSON(controller.CreateAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Account/update", jsonResponseFormatter.FormatToJSON(controller.UpdateAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteAccount/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteAccount)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Account/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Account/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/assignBranch", jsonResponseFormatter.FormatToJSON(controller.AssignBranchToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Account/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.UnassignBranchFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/assignProduct", jsonResponseFormatter.FormatToJSON(controller.AssignProductToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Account/unassignProduct", jsonResponseFormatter.FormatToJSON(controller.UnassignProductFromAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Account/addToOwners/", jsonResponseFormatter.FormatToJSON(controller.AddOwnersToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountremoveFromOwners/", jsonResponseFormatter.FormatToJSON(controller.RemoveOwnersFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.AddTransactionsToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.RemoveTransactionsFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/addToStatements/", jsonResponseFormatter.FormatToJSON(controller.AddStatementsToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountremoveFromStatements/", jsonResponseFormatter.FormatToJSON(controller.RemoveStatementsFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/addToStandingInstructions/", jsonResponseFormatter.FormatToJSON(controller.AddStandingInstructionsToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountremoveFromStandingInstructions/", jsonResponseFormatter.FormatToJSON(controller.RemoveStandingInstructionsFromAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Account/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.AddFeeChargesToAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.RemoveFeeChargesFromAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// AccountStatement Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/AccountStatement/get", jsonResponseFormatter.FormatToJSON(controller.GetAccountStatement)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/AccountStatementgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllAccountStatement)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/AccountStatement/create", jsonResponseFormatter.FormatToJSON(controller.CreateAccountStatement)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/AccountStatement/update", jsonResponseFormatter.FormatToJSON(controller.UpdateAccountStatement)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteAccountStatement/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteAccountStatement)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/AccountStatement/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToAccountStatement)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/AccountStatement/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromAccountStatement)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Transaction Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Transaction/get", jsonResponseFormatter.FormatToJSON(controller.GetTransaction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/TransactiongetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllTransaction)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Transaction/create", jsonResponseFormatter.FormatToJSON(controller.CreateTransaction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Transaction/update", jsonResponseFormatter.FormatToJSON(controller.UpdateTransaction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteTransaction/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteTransaction)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Transaction/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromTransaction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Transaction/assignExternalCounterparty", jsonResponseFormatter.FormatToJSON(controller.AssignExternalCounterpartyToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignExternalCounterparty", jsonResponseFormatter.FormatToJSON(controller.UnassignExternalCounterpartyFromTransaction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Transaction/assignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.AssignPaymentCardToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.UnassignPaymentCardFromTransaction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Transaction/assignFundsTransfer", jsonResponseFormatter.FormatToJSON(controller.AssignFundsTransferToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignFundsTransfer", jsonResponseFormatter.FormatToJSON(controller.UnassignFundsTransferFromTransaction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Transaction/assignFxTrade", jsonResponseFormatter.FormatToJSON(controller.AssignFxTradeToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignFxTrade", jsonResponseFormatter.FormatToJSON(controller.UnassignFxTradeFromTransaction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Transaction/assignDispute", jsonResponseFormatter.FormatToJSON(controller.AssignDisputeToTransaction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Transaction/unassignDispute", jsonResponseFormatter.FormatToJSON(controller.UnassignDisputeFromTransaction)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// ExternalAccount Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExternalAccount/get", jsonResponseFormatter.FormatToJSON(controller.GetExternalAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ExternalAccountgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllExternalAccount)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ExternalAccount/create", jsonResponseFormatter.FormatToJSON(controller.CreateExternalAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ExternalAccount/update", jsonResponseFormatter.FormatToJSON(controller.UpdateExternalAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteExternalAccount/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteExternalAccount)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExternalAccount/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToExternalAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ExternalAccount/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromExternalAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExternalAccount/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.AddTransactionsToExternalAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ExternalAccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.RemoveTransactionsFromExternalAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// FundsTransfer Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FundsTransfer/get", jsonResponseFormatter.FormatToJSON(controller.GetFundsTransfer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FundsTransfergetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllFundsTransfer)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/create", jsonResponseFormatter.FormatToJSON(controller.CreateFundsTransfer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/update", jsonResponseFormatter.FormatToJSON(controller.UpdateFundsTransfer)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteFundsTransfer/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteFundsTransfer)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FundsTransfer/assignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.AssignSourceAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignSourceAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FundsTransfer/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.AssignDestinationAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignDestinationAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FundsTransfer/assignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(controller.AssignExternalBeneficiaryToFundsTransfer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/unassignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(controller.UnassignExternalBeneficiaryFromFundsTransfer)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FundsTransfer/assignInitiatedBy", jsonResponseFormatter.FormatToJSON(controller.AssignInitiatedByToFundsTransfer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FundsTransfer/unassignInitiatedBy", jsonResponseFormatter.FormatToJSON(controller.UnassignInitiatedByFromFundsTransfer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FundsTransfer/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.AddTransactionsToFundsTransfer)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FundsTransferremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.RemoveTransactionsFromFundsTransfer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// StandingInstruction Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/StandingInstruction/get", jsonResponseFormatter.FormatToJSON(controller.GetStandingInstruction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/StandingInstructiongetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllStandingInstruction)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/StandingInstruction/create", jsonResponseFormatter.FormatToJSON(controller.CreateStandingInstruction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/StandingInstruction/update", jsonResponseFormatter.FormatToJSON(controller.UpdateStandingInstruction)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteStandingInstruction/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteStandingInstruction)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/StandingInstruction/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToStandingInstruction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/StandingInstruction/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromStandingInstruction)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/StandingInstruction/assignBeneficiary", jsonResponseFormatter.FormatToJSON(controller.AssignBeneficiaryToStandingInstruction)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/StandingInstruction/unassignBeneficiary", jsonResponseFormatter.FormatToJSON(controller.UnassignBeneficiaryFromStandingInstruction)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// PaymentCard Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/PaymentCard/get", jsonResponseFormatter.FormatToJSON(controller.GetPaymentCard)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/PaymentCardgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllPaymentCard)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/PaymentCard/create", jsonResponseFormatter.FormatToJSON(controller.CreatePaymentCard)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/PaymentCard/update", jsonResponseFormatter.FormatToJSON(controller.UpdatePaymentCard)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeletePaymentCard/delete", jsonResponseFormatter.FormatToJSON(controller.DeletePaymentCard)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/PaymentCard/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToPaymentCard)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/PaymentCard/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromPaymentCard)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/PaymentCard/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToPaymentCard)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/PaymentCard/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromPaymentCard)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/PaymentCard/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToPaymentCard)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/PaymentCard/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromPaymentCard)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/PaymentCard/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.AddTransactionsToPaymentCard)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/PaymentCardremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.RemoveTransactionsFromPaymentCard)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// LoanAccount Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/LoanAccount/get", jsonResponseFormatter.FormatToJSON(controller.GetLoanAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/LoanAccountgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllLoanAccount)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/LoanAccount/create", jsonResponseFormatter.FormatToJSON(controller.CreateLoanAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/LoanAccount/update", jsonResponseFormatter.FormatToJSON(controller.UpdateLoanAccount)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteLoanAccount/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteLoanAccount)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/LoanAccount/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccount/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/assignBranch", jsonResponseFormatter.FormatToJSON(controller.AssignBranchToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccount/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.UnassignBranchFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/assignProduct", jsonResponseFormatter.FormatToJSON(controller.AssignProductToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccount/unassignProduct", jsonResponseFormatter.FormatToJSON(controller.UnassignProductFromLoanAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/LoanAccount/addToBorrowers/", jsonResponseFormatter.FormatToJSON(controller.AddBorrowersToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccountremoveFromBorrowers/", jsonResponseFormatter.FormatToJSON(controller.RemoveBorrowersFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/addToRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(controller.AddRepaymentScheduleToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccountremoveFromRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(controller.RemoveRepaymentScheduleFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/addToPayments/", jsonResponseFormatter.FormatToJSON(controller.AddPaymentsToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccountremoveFromPayments/", jsonResponseFormatter.FormatToJSON(controller.RemovePaymentsFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/addToCollateral/", jsonResponseFormatter.FormatToJSON(controller.AddCollateralToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccountremoveFromCollateral/", jsonResponseFormatter.FormatToJSON(controller.RemoveCollateralFromLoanAccount)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanAccount/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.AddFeeChargesToLoanAccount)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanAccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.RemoveFeeChargesFromLoanAccount)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// RepaymentSchedule Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/RepaymentSchedule/get", jsonResponseFormatter.FormatToJSON(controller.GetRepaymentSchedule)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/RepaymentSchedulegetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllRepaymentSchedule)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/RepaymentSchedule/create", jsonResponseFormatter.FormatToJSON(controller.CreateRepaymentSchedule)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/RepaymentSchedule/update", jsonResponseFormatter.FormatToJSON(controller.UpdateRepaymentSchedule)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteRepaymentSchedule/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteRepaymentSchedule)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/RepaymentSchedule/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.AssignLoanAccountToRepaymentSchedule)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/RepaymentSchedule/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignLoanAccountFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/RepaymentSchedule/assignPayment", jsonResponseFormatter.FormatToJSON(controller.AssignPaymentToRepaymentSchedule)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/RepaymentSchedule/unassignPayment", jsonResponseFormatter.FormatToJSON(controller.UnassignPaymentFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// LoanPayment Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/LoanPayment/get", jsonResponseFormatter.FormatToJSON(controller.GetLoanPayment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/LoanPaymentgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllLoanPayment)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/LoanPayment/create", jsonResponseFormatter.FormatToJSON(controller.CreateLoanPayment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/LoanPayment/update", jsonResponseFormatter.FormatToJSON(controller.UpdateLoanPayment)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteLoanPayment/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteLoanPayment)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/LoanPayment/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.AssignLoanAccountToLoanPayment)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanPayment/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignLoanAccountFromLoanPayment)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/LoanPayment/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.AssignTransactionToLoanPayment)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/LoanPayment/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.UnassignTransactionFromLoanPayment)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Collateral Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Collateral/get", jsonResponseFormatter.FormatToJSON(controller.GetCollateral)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/CollateralgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllCollateral)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Collateral/create", jsonResponseFormatter.FormatToJSON(controller.CreateCollateral)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Collateral/update", jsonResponseFormatter.FormatToJSON(controller.UpdateCollateral)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteCollateral/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteCollateral)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Collateral/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.AssignLoanAccountToCollateral)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Collateral/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignLoanAccountFromCollateral)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// FeeCharge Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FeeCharge/get", jsonResponseFormatter.FormatToJSON(controller.GetFeeCharge)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FeeChargegetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllFeeCharge)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/FeeCharge/create", jsonResponseFormatter.FormatToJSON(controller.CreateFeeCharge)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FeeCharge/update", jsonResponseFormatter.FormatToJSON(controller.UpdateFeeCharge)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteFeeCharge/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteFeeCharge)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FeeCharge/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToFeeCharge)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FeeCharge/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FeeCharge/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.AssignLoanAccountToFeeCharge)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FeeCharge/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignLoanAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// ExchangeRate Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExchangeRate/get", jsonResponseFormatter.FormatToJSON(controller.GetExchangeRate)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ExchangeRategetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllExchangeRate)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ExchangeRate/create", jsonResponseFormatter.FormatToJSON(controller.CreateExchangeRate)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ExchangeRate/update", jsonResponseFormatter.FormatToJSON(controller.UpdateExchangeRate)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteExchangeRate/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteExchangeRate)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExchangeRate/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToExchangeRate)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ExchangeRate/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromExchangeRate)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ExchangeRate/addToFxTrades/", jsonResponseFormatter.FormatToJSON(controller.AddFxTradesToExchangeRate)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ExchangeRateremoveFromFxTrades/", jsonResponseFormatter.FormatToJSON(controller.RemoveFxTradesFromExchangeRate)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// FXTrade Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FXTrade/get", jsonResponseFormatter.FormatToJSON(controller.GetFXTrade)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FXTradegetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllFXTrade)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/FXTrade/create", jsonResponseFormatter.FormatToJSON(controller.CreateFXTrade)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/FXTrade/update", jsonResponseFormatter.FormatToJSON(controller.UpdateFXTrade)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteFXTrade/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteFXTrade)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/FXTrade/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromFXTrade)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FXTrade/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromFXTrade)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FXTrade/assignExchangeRate", jsonResponseFormatter.FormatToJSON(controller.AssignExchangeRateToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignExchangeRate", jsonResponseFormatter.FormatToJSON(controller.UnassignExchangeRateFromFXTrade)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FXTrade/assignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.AssignSourceAccountToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignSourceAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FXTrade/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.AssignDestinationAccountToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignDestinationAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/FXTrade/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.AssignTransactionToFXTrade)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/FXTrade/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.UnassignTransactionFromFXTrade)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Dispute Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Dispute/get", jsonResponseFormatter.FormatToJSON(controller.GetDispute)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DisputegetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllDispute)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Dispute/create", jsonResponseFormatter.FormatToJSON(controller.CreateDispute)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Dispute/update", jsonResponseFormatter.FormatToJSON(controller.UpdateDispute)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteDispute/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteDispute)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Dispute/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.AssignTransactionToDispute)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Dispute/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.UnassignTransactionFromDispute)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Dispute/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToDispute)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Dispute/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromDispute)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Dispute/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AssignAccountToDispute)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Dispute/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.UnassignAccountFromDispute)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Dispute/assignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.AssignPaymentCardToDispute)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Dispute/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.UnassignPaymentCardFromDispute)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Consent Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Consent/get", jsonResponseFormatter.FormatToJSON(controller.GetConsent)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ConsentgetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllConsent)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Consent/create", jsonResponseFormatter.FormatToJSON(controller.CreateConsent)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Consent/update", jsonResponseFormatter.FormatToJSON(controller.UpdateConsent)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteConsent/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteConsent)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Consent/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.AssignCustomerToConsent)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Consent/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.UnassignCustomerFromConsent)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Consent/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToConsent)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Consent/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromConsent)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/Consent/assignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(controller.AssignThirdPartyProviderToConsent)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Consent/unassignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(controller.UnassignThirdPartyProviderFromConsent)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/Consent/addToAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(controller.AddAuthorizedAccountsToConsent)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ConsentremoveFromAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(controller.RemoveAuthorizedAccountsFromConsent)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// ThirdPartyProvider Routes to JSON response formatter first
	// then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard Lifecycle Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ThirdPartyProvider/get", jsonResponseFormatter.FormatToJSON(controller.GetThirdPartyProvider)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ThirdPartyProvidergetAll", jsonResponseFormatter.FormatToJSON(controller.GetAllThirdPartyProvider)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ThirdPartyProvider/create", jsonResponseFormatter.FormatToJSON(controller.CreateThirdPartyProvider)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ThirdPartyProvider/update", jsonResponseFormatter.FormatToJSON(controller.UpdateThirdPartyProvider)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/DeleteThirdPartyProvider/delete", jsonResponseFormatter.FormatToJSON(controller.DeleteThirdPartyProvider)).Methods("POST", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ThirdPartyProvider/assignBank", jsonResponseFormatter.FormatToJSON(controller.AssignBankToThirdPartyProvider)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ThirdPartyProvider/unassignBank", jsonResponseFormatter.FormatToJSON(controller.UnassignBankFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------
	router.HandleFunc("/api/ThirdPartyProvider/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.AddConsentsToThirdPartyProvider)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ThirdPartyProviderremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.RemoveConsentsFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")

	return router
}
