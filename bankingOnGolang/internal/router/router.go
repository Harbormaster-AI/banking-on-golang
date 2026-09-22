package router

import (
    "bankingOnGolang/internal/controller"
    jsonResponseFormatter "bankingOnGolang/internal/response"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // Bank Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Bank/get", jsonResponseFormatter.FormatToJSON(BankController.GetBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankgetAll", jsonResponseFormatter.FormatToJSON(BankController.GetAllBank)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Bank/create", jsonResponseFormatter.FormatToJSON(BankController.CreateBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Bank/update", jsonResponseFormatter.FormatToJSON(BankController.UpdateBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBank/delete", jsonResponseFormatter.FormatToJSON(BankController.DeleteBank)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Bank/addToBranches/", jsonResponseFormatter.FormatToJSON(BankController.AddBranchesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromBranches/", jsonResponseFormatter.FormatToJSON(BankController.RemoveBranchesFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToProducts/", jsonResponseFormatter.FormatToJSON(BankController.AddProductsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromProducts/", jsonResponseFormatter.FormatToJSON(BankController.RemoveProductsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToCustomers/", jsonResponseFormatter.FormatToJSON(BankController.AddCustomersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromCustomers/", jsonResponseFormatter.FormatToJSON(BankController.RemoveCustomersFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToAccounts/", jsonResponseFormatter.FormatToJSON(BankController.AddAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(BankController.RemoveAccountsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(BankController.AddPaymentCardsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(BankController.RemovePaymentCardsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(BankController.AddLoanAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(BankController.RemoveLoanAccountsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToExchangeRates/", jsonResponseFormatter.FormatToJSON(BankController.AddExchangeRatesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromExchangeRates/", jsonResponseFormatter.FormatToJSON(BankController.RemoveExchangeRatesFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToConsents/", jsonResponseFormatter.FormatToJSON(BankController.AddConsentsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromConsents/", jsonResponseFormatter.FormatToJSON(BankController.RemoveConsentsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(BankController.AddThirdPartyProvidersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(BankController.RemoveThirdPartyProvidersFromBank)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Branch Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/get", jsonResponseFormatter.FormatToJSON(BranchController.GetBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BranchgetAll", jsonResponseFormatter.FormatToJSON(BranchController.GetAllBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Branch/create", jsonResponseFormatter.FormatToJSON(BranchController.CreateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Branch/update", jsonResponseFormatter.FormatToJSON(BranchController.UpdateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBranch/delete", jsonResponseFormatter.FormatToJSON(BranchController.DeleteBranch)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/assignBank", jsonResponseFormatter.FormatToJSON(BranchController.AssignBankToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Branch/unassignBank", jsonResponseFormatter.FormatToJSON(BranchController.UnassignBankFromBranch)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/addToAccounts/", jsonResponseFormatter.FormatToJSON(BranchController.AddAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(BranchController.RemoveAccountsFromBranch)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Branch/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(BranchController.AddLoanAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(BranchController.RemoveLoanAccountsFromBranch)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Branch/addToAtms/", jsonResponseFormatter.FormatToJSON(BranchController.AddAtmsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromAtms/", jsonResponseFormatter.FormatToJSON(BranchController.RemoveAtmsFromBranch)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // ATM Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ATM/get", jsonResponseFormatter.FormatToJSON(ATMController.GetATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ATMgetAll", jsonResponseFormatter.FormatToJSON(ATMController.GetAllATM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ATM/create", jsonResponseFormatter.FormatToJSON(ATMController.CreateATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ATM/update", jsonResponseFormatter.FormatToJSON(ATMController.UpdateATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteATM/delete", jsonResponseFormatter.FormatToJSON(ATMController.DeleteATM)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ATM/assignBranch", jsonResponseFormatter.FormatToJSON(ATMController.AssignBranchToATM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ATM/unassignBranch", jsonResponseFormatter.FormatToJSON(ATMController.UnassignBranchFromATM)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Customer/get", jsonResponseFormatter.FormatToJSON(CustomerController.GetCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CustomergetAll", jsonResponseFormatter.FormatToJSON(CustomerController.GetAllCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Customer/create", jsonResponseFormatter.FormatToJSON(CustomerController.CreateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Customer/update", jsonResponseFormatter.FormatToJSON(CustomerController.UpdateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteCustomer/delete", jsonResponseFormatter.FormatToJSON(CustomerController.DeleteCustomer)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/assignBank", jsonResponseFormatter.FormatToJSON(CustomerController.AssignBankToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Customer/unassignBank", jsonResponseFormatter.FormatToJSON(CustomerController.UnassignBankFromCustomer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/addToAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.AddAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.AddLoanAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveLoanAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(CustomerController.AddPaymentCardsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(CustomerController.RemovePaymentCardsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToExternalAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.AddExternalAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromExternalAccounts/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveExternalAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToFundsTransfers/", jsonResponseFormatter.FormatToJSON(CustomerController.AddFundsTransfersToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromFundsTransfers/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveFundsTransfersFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToDisputes/", jsonResponseFormatter.FormatToJSON(CustomerController.AddDisputesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromDisputes/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveDisputesFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToKycProfiles/", jsonResponseFormatter.FormatToJSON(CustomerController.AddKycProfilesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromKycProfiles/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveKycProfilesFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToConsents/", jsonResponseFormatter.FormatToJSON(CustomerController.AddConsentsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromConsents/", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveConsentsFromCustomer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // KycProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/get", jsonResponseFormatter.FormatToJSON(KycProfileController.GetKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KycProfilegetAll", jsonResponseFormatter.FormatToJSON(KycProfileController.GetAllKycProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KycProfile/create", jsonResponseFormatter.FormatToJSON(KycProfileController.CreateKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KycProfile/update", jsonResponseFormatter.FormatToJSON(KycProfileController.UpdateKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteKycProfile/delete", jsonResponseFormatter.FormatToJSON(KycProfileController.DeleteKycProfile)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/assignCustomer", jsonResponseFormatter.FormatToJSON(KycProfileController.AssignCustomerToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfile/unassignCustomer", jsonResponseFormatter.FormatToJSON(KycProfileController.UnassignCustomerFromKycProfile)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/addToIdentityDocuments/", jsonResponseFormatter.FormatToJSON(KycProfileController.AddIdentityDocumentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromIdentityDocuments/", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveIdentityDocumentsFromKycProfile)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/KycProfile/addToRiskAssessments/", jsonResponseFormatter.FormatToJSON(KycProfileController.AddRiskAssessmentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromRiskAssessments/", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveRiskAssessmentsFromKycProfile)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/KycProfile/addToScreenings/", jsonResponseFormatter.FormatToJSON(KycProfileController.AddScreeningsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromScreenings/", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveScreeningsFromKycProfile)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // IdentityDocument Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/IdentityDocument/get", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.GetIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/IdentityDocumentgetAll", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.GetAllIdentityDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/create", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.CreateIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/update", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.UpdateIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteIdentityDocument/delete", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.DeleteIdentityDocument)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/IdentityDocument/assignKycProfile", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.AssignKycProfileToIdentityDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/unassignKycProfile", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.UnassignKycProfileFromIdentityDocument)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/RiskAssessment/get", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessmentgetAll", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetAllRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/create", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.CreateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/update", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UpdateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteRiskAssessment/delete", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.DeleteRiskAssessment)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RiskAssessment/assignKycProfile", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.AssignKycProfileToRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/unassignKycProfile", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UnassignKycProfileFromRiskAssessment)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ScreeningResult/get", jsonResponseFormatter.FormatToJSON(ScreeningResultController.GetScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScreeningResultgetAll", jsonResponseFormatter.FormatToJSON(ScreeningResultController.GetAllScreeningResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/create", jsonResponseFormatter.FormatToJSON(ScreeningResultController.CreateScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/update", jsonResponseFormatter.FormatToJSON(ScreeningResultController.UpdateScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteScreeningResult/delete", jsonResponseFormatter.FormatToJSON(ScreeningResultController.DeleteScreeningResult)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ScreeningResult/assignKycProfile", jsonResponseFormatter.FormatToJSON(ScreeningResultController.AssignKycProfileToScreeningResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/unassignKycProfile", jsonResponseFormatter.FormatToJSON(ScreeningResultController.UnassignKycProfileFromScreeningResult)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/BankingProduct/get", jsonResponseFormatter.FormatToJSON(BankingProductController.GetBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankingProductgetAll", jsonResponseFormatter.FormatToJSON(BankingProductController.GetAllBankingProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/create", jsonResponseFormatter.FormatToJSON(BankingProductController.CreateBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/update", jsonResponseFormatter.FormatToJSON(BankingProductController.UpdateBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBankingProduct/delete", jsonResponseFormatter.FormatToJSON(BankingProductController.DeleteBankingProduct)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BankingProduct/assignBank", jsonResponseFormatter.FormatToJSON(BankingProductController.AssignBankToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/unassignBank", jsonResponseFormatter.FormatToJSON(BankingProductController.UnassignBankFromBankingProduct)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BankingProduct/addToAccounts/", jsonResponseFormatter.FormatToJSON(BankingProductController.AddAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(BankingProductController.RemoveAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/BankingProduct/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(BankingProductController.AddLoanAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(BankingProductController.RemoveLoanAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/BankingProduct/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(BankingProductController.AddPaymentCardsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(BankingProductController.RemovePaymentCardsFromBankingProduct)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Account Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/get", jsonResponseFormatter.FormatToJSON(AccountController.GetAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountgetAll", jsonResponseFormatter.FormatToJSON(AccountController.GetAllAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Account/create", jsonResponseFormatter.FormatToJSON(AccountController.CreateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Account/update", jsonResponseFormatter.FormatToJSON(AccountController.UpdateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteAccount/delete", jsonResponseFormatter.FormatToJSON(AccountController.DeleteAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/assignBank", jsonResponseFormatter.FormatToJSON(AccountController.AssignBankToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignBank", jsonResponseFormatter.FormatToJSON(AccountController.UnassignBankFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/assignBranch", jsonResponseFormatter.FormatToJSON(AccountController.AssignBranchToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignBranch", jsonResponseFormatter.FormatToJSON(AccountController.UnassignBranchFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/assignProduct", jsonResponseFormatter.FormatToJSON(AccountController.AssignProductToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignProduct", jsonResponseFormatter.FormatToJSON(AccountController.UnassignProductFromAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/addToOwners/", jsonResponseFormatter.FormatToJSON(AccountController.AddOwnersToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromOwners/", jsonResponseFormatter.FormatToJSON(AccountController.RemoveOwnersFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToTransactions/", jsonResponseFormatter.FormatToJSON(AccountController.AddTransactionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(AccountController.RemoveTransactionsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToStatements/", jsonResponseFormatter.FormatToJSON(AccountController.AddStatementsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromStatements/", jsonResponseFormatter.FormatToJSON(AccountController.RemoveStatementsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToStandingInstructions/", jsonResponseFormatter.FormatToJSON(AccountController.AddStandingInstructionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromStandingInstructions/", jsonResponseFormatter.FormatToJSON(AccountController.RemoveStandingInstructionsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(AccountController.AddFeeChargesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(AccountController.RemoveFeeChargesFromAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // AccountStatement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/get", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatementgetAll", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAllAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/create", jsonResponseFormatter.FormatToJSON(AccountStatementController.CreateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/update", jsonResponseFormatter.FormatToJSON(AccountStatementController.UpdateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteAccountStatement/delete", jsonResponseFormatter.FormatToJSON(AccountStatementController.DeleteAccountStatement)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/assignAccount", jsonResponseFormatter.FormatToJSON(AccountStatementController.AssignAccountToAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/unassignAccount", jsonResponseFormatter.FormatToJSON(AccountStatementController.UnassignAccountFromAccountStatement)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Transaction/get", jsonResponseFormatter.FormatToJSON(TransactionController.GetTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TransactiongetAll", jsonResponseFormatter.FormatToJSON(TransactionController.GetAllTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Transaction/create", jsonResponseFormatter.FormatToJSON(TransactionController.CreateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Transaction/update", jsonResponseFormatter.FormatToJSON(TransactionController.UpdateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteTransaction/delete", jsonResponseFormatter.FormatToJSON(TransactionController.DeleteTransaction)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Transaction/assignAccount", jsonResponseFormatter.FormatToJSON(TransactionController.AssignAccountToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignAccount", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignAccountFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignExternalCounterparty", jsonResponseFormatter.FormatToJSON(TransactionController.AssignExternalCounterpartyToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignExternalCounterparty", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignExternalCounterpartyFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignPaymentCard", jsonResponseFormatter.FormatToJSON(TransactionController.AssignPaymentCardToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignPaymentCardFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignFundsTransfer", jsonResponseFormatter.FormatToJSON(TransactionController.AssignFundsTransferToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignFundsTransfer", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignFundsTransferFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignFxTrade", jsonResponseFormatter.FormatToJSON(TransactionController.AssignFxTradeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignFxTrade", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignFxTradeFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignDispute", jsonResponseFormatter.FormatToJSON(TransactionController.AssignDisputeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignDispute", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignDisputeFromTransaction)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ExternalAccount/get", jsonResponseFormatter.FormatToJSON(ExternalAccountController.GetExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExternalAccountgetAll", jsonResponseFormatter.FormatToJSON(ExternalAccountController.GetAllExternalAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/create", jsonResponseFormatter.FormatToJSON(ExternalAccountController.CreateExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/update", jsonResponseFormatter.FormatToJSON(ExternalAccountController.UpdateExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteExternalAccount/delete", jsonResponseFormatter.FormatToJSON(ExternalAccountController.DeleteExternalAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExternalAccount/assignCustomer", jsonResponseFormatter.FormatToJSON(ExternalAccountController.AssignCustomerToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/unassignCustomer", jsonResponseFormatter.FormatToJSON(ExternalAccountController.UnassignCustomerFromExternalAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExternalAccount/addToTransactions/", jsonResponseFormatter.FormatToJSON(ExternalAccountController.AddTransactionsToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExternalAccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(ExternalAccountController.RemoveTransactionsFromExternalAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // FundsTransfer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/get", jsonResponseFormatter.FormatToJSON(FundsTransferController.GetFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FundsTransfergetAll", jsonResponseFormatter.FormatToJSON(FundsTransferController.GetAllFundsTransfer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/create", jsonResponseFormatter.FormatToJSON(FundsTransferController.CreateFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/update", jsonResponseFormatter.FormatToJSON(FundsTransferController.UpdateFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFundsTransfer/delete", jsonResponseFormatter.FormatToJSON(FundsTransferController.DeleteFundsTransfer)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/assignSourceAccount", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignSourceAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignSourceAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignDestinationAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignDestinationAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignExternalBeneficiaryToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignExternalBeneficiaryFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignInitiatedBy", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignInitiatedByToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignInitiatedBy", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignInitiatedByFromFundsTransfer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/addToTransactions/", jsonResponseFormatter.FormatToJSON(FundsTransferController.AddTransactionsToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransferremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(FundsTransferController.RemoveTransactionsFromFundsTransfer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // StandingInstruction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StandingInstruction/get", jsonResponseFormatter.FormatToJSON(StandingInstructionController.GetStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StandingInstructiongetAll", jsonResponseFormatter.FormatToJSON(StandingInstructionController.GetAllStandingInstruction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/create", jsonResponseFormatter.FormatToJSON(StandingInstructionController.CreateStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/update", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UpdateStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteStandingInstruction/delete", jsonResponseFormatter.FormatToJSON(StandingInstructionController.DeleteStandingInstruction)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StandingInstruction/assignAccount", jsonResponseFormatter.FormatToJSON(StandingInstructionController.AssignAccountToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/unassignAccount", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UnassignAccountFromStandingInstruction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/StandingInstruction/assignBeneficiary", jsonResponseFormatter.FormatToJSON(StandingInstructionController.AssignBeneficiaryToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/unassignBeneficiary", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UnassignBeneficiaryFromStandingInstruction)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/PaymentCard/get", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetPaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCardgetAll", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetAllPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/create", jsonResponseFormatter.FormatToJSON(PaymentCardController.CreatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/update", jsonResponseFormatter.FormatToJSON(PaymentCardController.UpdatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentCard/delete", jsonResponseFormatter.FormatToJSON(PaymentCardController.DeletePaymentCard)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentCard/assignBank", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignBankToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignBank", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignBankFromPaymentCard)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/PaymentCard/assignAccount", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignAccountToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignAccount", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignAccountFromPaymentCard)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/PaymentCard/assignCustomer", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignCustomerToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignCustomer", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignCustomerFromPaymentCard)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentCard/addToTransactions/", jsonResponseFormatter.FormatToJSON(PaymentCardController.AddTransactionsToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCardremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(PaymentCardController.RemoveTransactionsFromPaymentCard)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // LoanAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/get", jsonResponseFormatter.FormatToJSON(LoanAccountController.GetLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanAccountgetAll", jsonResponseFormatter.FormatToJSON(LoanAccountController.GetAllLoanAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/create", jsonResponseFormatter.FormatToJSON(LoanAccountController.CreateLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/update", jsonResponseFormatter.FormatToJSON(LoanAccountController.UpdateLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanAccount/delete", jsonResponseFormatter.FormatToJSON(LoanAccountController.DeleteLoanAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/assignBank", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignBankToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignBank", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignBankFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/assignBranch", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignBranchToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignBranch", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignBranchFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/assignProduct", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignProductToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignProduct", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignProductFromLoanAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/addToBorrowers/", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddBorrowersToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromBorrowers/", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveBorrowersFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddRepaymentScheduleToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveRepaymentScheduleFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToPayments/", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddPaymentsToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromPayments/", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemovePaymentsFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToCollateral/", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddCollateralToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromCollateral/", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveCollateralFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddFeeChargesToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveFeeChargesFromLoanAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // RepaymentSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/get", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedulegetAll", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetAllRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/create", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.CreateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/update", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UpdateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteRepaymentSchedule/delete", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.DeleteRepaymentSchedule)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/assignLoanAccount", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AssignLoanAccountToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UnassignLoanAccountFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/RepaymentSchedule/assignPayment", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AssignPaymentToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/unassignPayment", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UnassignPaymentFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/LoanPayment/get", jsonResponseFormatter.FormatToJSON(LoanPaymentController.GetLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanPaymentgetAll", jsonResponseFormatter.FormatToJSON(LoanPaymentController.GetAllLoanPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/create", jsonResponseFormatter.FormatToJSON(LoanPaymentController.CreateLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/update", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UpdateLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanPayment/delete", jsonResponseFormatter.FormatToJSON(LoanPaymentController.DeleteLoanPayment)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanPayment/assignLoanAccount", jsonResponseFormatter.FormatToJSON(LoanPaymentController.AssignLoanAccountToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UnassignLoanAccountFromLoanPayment)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanPayment/assignTransaction", jsonResponseFormatter.FormatToJSON(LoanPaymentController.AssignTransactionToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/unassignTransaction", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UnassignTransactionFromLoanPayment)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Collateral/get", jsonResponseFormatter.FormatToJSON(CollateralController.GetCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CollateralgetAll", jsonResponseFormatter.FormatToJSON(CollateralController.GetAllCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Collateral/create", jsonResponseFormatter.FormatToJSON(CollateralController.CreateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Collateral/update", jsonResponseFormatter.FormatToJSON(CollateralController.UpdateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteCollateral/delete", jsonResponseFormatter.FormatToJSON(CollateralController.DeleteCollateral)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Collateral/assignLoanAccount", jsonResponseFormatter.FormatToJSON(CollateralController.AssignLoanAccountToCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Collateral/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(CollateralController.UnassignLoanAccountFromCollateral)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/FeeCharge/get", jsonResponseFormatter.FormatToJSON(FeeChargeController.GetFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeChargegetAll", jsonResponseFormatter.FormatToJSON(FeeChargeController.GetAllFeeCharge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/create", jsonResponseFormatter.FormatToJSON(FeeChargeController.CreateFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/update", jsonResponseFormatter.FormatToJSON(FeeChargeController.UpdateFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFeeCharge/delete", jsonResponseFormatter.FormatToJSON(FeeChargeController.DeleteFeeCharge)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FeeCharge/assignAccount", jsonResponseFormatter.FormatToJSON(FeeChargeController.AssignAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/unassignAccount", jsonResponseFormatter.FormatToJSON(FeeChargeController.UnassignAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FeeCharge/assignLoanAccount", jsonResponseFormatter.FormatToJSON(FeeChargeController.AssignLoanAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(FeeChargeController.UnassignLoanAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ExchangeRate/get", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRategetAll", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetAllExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/create", jsonResponseFormatter.FormatToJSON(ExchangeRateController.CreateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/update", jsonResponseFormatter.FormatToJSON(ExchangeRateController.UpdateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteExchangeRate/delete", jsonResponseFormatter.FormatToJSON(ExchangeRateController.DeleteExchangeRate)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExchangeRate/assignBank", jsonResponseFormatter.FormatToJSON(ExchangeRateController.AssignBankToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/unassignBank", jsonResponseFormatter.FormatToJSON(ExchangeRateController.UnassignBankFromExchangeRate)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExchangeRate/addToFxTrades/", jsonResponseFormatter.FormatToJSON(ExchangeRateController.AddFxTradesToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExchangeRateremoveFromFxTrades/", jsonResponseFormatter.FormatToJSON(ExchangeRateController.RemoveFxTradesFromExchangeRate)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // FXTrade Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXTrade/get", jsonResponseFormatter.FormatToJSON(FXTradeController.GetFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXTradegetAll", jsonResponseFormatter.FormatToJSON(FXTradeController.GetAllFXTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FXTrade/create", jsonResponseFormatter.FormatToJSON(FXTradeController.CreateFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXTrade/update", jsonResponseFormatter.FormatToJSON(FXTradeController.UpdateFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFXTrade/delete", jsonResponseFormatter.FormatToJSON(FXTradeController.DeleteFXTrade)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXTrade/assignCustomer", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignCustomerToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignCustomer", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignCustomerFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignBank", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignBankToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignBank", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignBankFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignExchangeRate", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignExchangeRateToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignExchangeRate", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignExchangeRateFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignSourceAccount", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignSourceAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignSourceAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignDestinationAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignDestinationAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignTransaction", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignTransactionToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignTransaction", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignTransactionFromFXTrade)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Dispute/get", jsonResponseFormatter.FormatToJSON(DisputeController.GetDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DisputegetAll", jsonResponseFormatter.FormatToJSON(DisputeController.GetAllDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dispute/create", jsonResponseFormatter.FormatToJSON(DisputeController.CreateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dispute/update", jsonResponseFormatter.FormatToJSON(DisputeController.UpdateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteDispute/delete", jsonResponseFormatter.FormatToJSON(DisputeController.DeleteDispute)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dispute/assignTransaction", jsonResponseFormatter.FormatToJSON(DisputeController.AssignTransactionToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignTransaction", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignTransactionFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignCustomer", jsonResponseFormatter.FormatToJSON(DisputeController.AssignCustomerToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignCustomer", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignCustomerFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignAccount", jsonResponseFormatter.FormatToJSON(DisputeController.AssignAccountToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignAccount", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignAccountFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignPaymentCard", jsonResponseFormatter.FormatToJSON(DisputeController.AssignPaymentCardToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignPaymentCardFromDispute)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Consent/get", jsonResponseFormatter.FormatToJSON(ConsentController.GetConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ConsentgetAll", jsonResponseFormatter.FormatToJSON(ConsentController.GetAllConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Consent/create", jsonResponseFormatter.FormatToJSON(ConsentController.CreateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Consent/update", jsonResponseFormatter.FormatToJSON(ConsentController.UpdateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteConsent/delete", jsonResponseFormatter.FormatToJSON(ConsentController.DeleteConsent)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Consent/assignCustomer", jsonResponseFormatter.FormatToJSON(ConsentController.AssignCustomerToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignCustomer", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignCustomerFromConsent)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Consent/assignBank", jsonResponseFormatter.FormatToJSON(ConsentController.AssignBankToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignBank", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignBankFromConsent)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Consent/assignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(ConsentController.AssignThirdPartyProviderToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignThirdPartyProviderFromConsent)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Consent/addToAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(ConsentController.AddAuthorizedAccountsToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ConsentremoveFromAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(ConsentController.RemoveAuthorizedAccountsFromConsent)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // ThirdPartyProvider Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/get", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.GetThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvidergetAll", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.GetAllThirdPartyProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/create", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.CreateThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/update", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.UpdateThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteThirdPartyProvider/delete", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.DeleteThirdPartyProvider)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/assignBank", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.AssignBankToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/unassignBank", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.UnassignBankFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/addToConsents/", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.AddConsentsToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProviderremoveFromConsents/", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.RemoveConsentsFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")


    return router
}
