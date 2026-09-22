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

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(controller.PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(controller.PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // Bank Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Bank/get", jsonResponseFormatter.FormatToJSON(controller.BankController.GetBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankgetAll", jsonResponseFormatter.FormatToJSON(controller.BankController.GetAllBank)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Bank/create", jsonResponseFormatter.FormatToJSON(controller.BankController.CreateBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Bank/update", jsonResponseFormatter.FormatToJSON(controller.BankController.UpdateBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBank/delete", jsonResponseFormatter.FormatToJSON(controller.BankController.DeleteBank)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Bank/addToBranches/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddBranchesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromBranches/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveBranchesFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToProducts/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddProductsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromProducts/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveProductsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToCustomers/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddCustomersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromCustomers/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveCustomersFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveAccountsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddPaymentCardsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemovePaymentCardsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddLoanAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveLoanAccountsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToExchangeRates/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddExchangeRatesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromExchangeRates/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveExchangeRatesFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddConsentsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveConsentsFromBank)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Bank/addToThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(controller.BankController.AddThirdPartyProvidersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankremoveFromThirdPartyProviders/", jsonResponseFormatter.FormatToJSON(controller.BankController.RemoveThirdPartyProvidersFromBank)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Branch Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/get", jsonResponseFormatter.FormatToJSON(controller.BranchController.GetBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BranchgetAll", jsonResponseFormatter.FormatToJSON(controller.BranchController.GetAllBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Branch/create", jsonResponseFormatter.FormatToJSON(controller.BranchController.CreateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Branch/update", jsonResponseFormatter.FormatToJSON(controller.BranchController.UpdateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBranch/delete", jsonResponseFormatter.FormatToJSON(controller.BranchController.DeleteBranch)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/assignBank", jsonResponseFormatter.FormatToJSON(controller.BranchController.AssignBankToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Branch/unassignBank", jsonResponseFormatter.FormatToJSON(controller.BranchController.UnassignBankFromBranch)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.BranchController.AddAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.BranchController.RemoveAccountsFromBranch)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Branch/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BranchController.AddLoanAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BranchController.RemoveLoanAccountsFromBranch)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Branch/addToAtms/", jsonResponseFormatter.FormatToJSON(controller.BranchController.AddAtmsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BranchremoveFromAtms/", jsonResponseFormatter.FormatToJSON(controller.BranchController.RemoveAtmsFromBranch)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // ATM Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ATM/get", jsonResponseFormatter.FormatToJSON(controller.ATMController.GetATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ATMgetAll", jsonResponseFormatter.FormatToJSON(controller.ATMController.GetAllATM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ATM/create", jsonResponseFormatter.FormatToJSON(controller.ATMController.CreateATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ATM/update", jsonResponseFormatter.FormatToJSON(controller.ATMController.UpdateATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteATM/delete", jsonResponseFormatter.FormatToJSON(controller.ATMController.DeleteATM)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ATM/assignBranch", jsonResponseFormatter.FormatToJSON(controller.ATMController.AssignBranchToATM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ATM/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.ATMController.UnassignBranchFromATM)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Customer/get", jsonResponseFormatter.FormatToJSON(controller.CustomerController.GetCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CustomergetAll", jsonResponseFormatter.FormatToJSON(controller.CustomerController.GetAllCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Customer/create", jsonResponseFormatter.FormatToJSON(controller.CustomerController.CreateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Customer/update", jsonResponseFormatter.FormatToJSON(controller.CustomerController.UpdateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteCustomer/delete", jsonResponseFormatter.FormatToJSON(controller.CustomerController.DeleteCustomer)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/assignBank", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AssignBankToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Customer/unassignBank", jsonResponseFormatter.FormatToJSON(controller.CustomerController.UnassignBankFromCustomer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddLoanAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveLoanAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddPaymentCardsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemovePaymentCardsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToExternalAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddExternalAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromExternalAccounts/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveExternalAccountsFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToFundsTransfers/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddFundsTransfersToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromFundsTransfers/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveFundsTransfersFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToDisputes/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddDisputesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromDisputes/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveDisputesFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToKycProfiles/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddKycProfilesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromKycProfiles/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveKycProfilesFromCustomer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Customer/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.AddConsentsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/CustomerremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.CustomerController.RemoveConsentsFromCustomer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // KycProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/get", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.GetKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KycProfilegetAll", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.GetAllKycProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KycProfile/create", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.CreateKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KycProfile/update", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.UpdateKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteKycProfile/delete", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.DeleteKycProfile)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.AssignCustomerToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfile/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.UnassignCustomerFromKycProfile)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/addToIdentityDocuments/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.AddIdentityDocumentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromIdentityDocuments/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.RemoveIdentityDocumentsFromKycProfile)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/KycProfile/addToRiskAssessments/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.AddRiskAssessmentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromRiskAssessments/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.RemoveRiskAssessmentsFromKycProfile)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/KycProfile/addToScreenings/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.AddScreeningsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/KycProfileremoveFromScreenings/", jsonResponseFormatter.FormatToJSON(controller.KycProfileController.RemoveScreeningsFromKycProfile)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // IdentityDocument Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/IdentityDocument/get", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.GetIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/IdentityDocumentgetAll", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.GetAllIdentityDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/create", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.CreateIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/update", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.UpdateIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteIdentityDocument/delete", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.DeleteIdentityDocument)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/IdentityDocument/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.AssignKycProfileToIdentityDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.IdentityDocumentController.UnassignKycProfileFromIdentityDocument)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/RiskAssessment/get", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.GetRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessmentgetAll", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.GetAllRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/create", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.CreateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/update", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.UpdateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteRiskAssessment/delete", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.DeleteRiskAssessment)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RiskAssessment/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.AssignKycProfileToRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.RiskAssessmentController.UnassignKycProfileFromRiskAssessment)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ScreeningResult/get", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.GetScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScreeningResultgetAll", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.GetAllScreeningResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/create", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.CreateScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/update", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.UpdateScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteScreeningResult/delete", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.DeleteScreeningResult)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ScreeningResult/assignKycProfile", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.AssignKycProfileToScreeningResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/unassignKycProfile", jsonResponseFormatter.FormatToJSON(controller.ScreeningResultController.UnassignKycProfileFromScreeningResult)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/BankingProduct/get", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.GetBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankingProductgetAll", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.GetAllBankingProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/create", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.CreateBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/update", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.UpdateBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteBankingProduct/delete", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.DeleteBankingProduct)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BankingProduct/assignBank", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.AssignBankToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/unassignBank", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.UnassignBankFromBankingProduct)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BankingProduct/addToAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.AddAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.RemoveAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/BankingProduct/addToLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.AddLoanAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromLoanAccounts/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.RemoveLoanAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/BankingProduct/addToPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.AddPaymentCardsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/BankingProductremoveFromPaymentCards/", jsonResponseFormatter.FormatToJSON(controller.BankingProductController.RemovePaymentCardsFromBankingProduct)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Account Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/get", jsonResponseFormatter.FormatToJSON(controller.AccountController.GetAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountgetAll", jsonResponseFormatter.FormatToJSON(controller.AccountController.GetAllAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Account/create", jsonResponseFormatter.FormatToJSON(controller.AccountController.CreateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Account/update", jsonResponseFormatter.FormatToJSON(controller.AccountController.UpdateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteAccount/delete", jsonResponseFormatter.FormatToJSON(controller.AccountController.DeleteAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/assignBank", jsonResponseFormatter.FormatToJSON(controller.AccountController.AssignBankToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignBank", jsonResponseFormatter.FormatToJSON(controller.AccountController.UnassignBankFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/assignBranch", jsonResponseFormatter.FormatToJSON(controller.AccountController.AssignBranchToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.AccountController.UnassignBranchFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/assignProduct", jsonResponseFormatter.FormatToJSON(controller.AccountController.AssignProductToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Account/unassignProduct", jsonResponseFormatter.FormatToJSON(controller.AccountController.UnassignProductFromAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/addToOwners/", jsonResponseFormatter.FormatToJSON(controller.AccountController.AddOwnersToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromOwners/", jsonResponseFormatter.FormatToJSON(controller.AccountController.RemoveOwnersFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.AccountController.AddTransactionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.AccountController.RemoveTransactionsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToStatements/", jsonResponseFormatter.FormatToJSON(controller.AccountController.AddStatementsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromStatements/", jsonResponseFormatter.FormatToJSON(controller.AccountController.RemoveStatementsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToStandingInstructions/", jsonResponseFormatter.FormatToJSON(controller.AccountController.AddStandingInstructionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromStandingInstructions/", jsonResponseFormatter.FormatToJSON(controller.AccountController.RemoveStandingInstructionsFromAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Account/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.AccountController.AddFeeChargesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.AccountController.RemoveFeeChargesFromAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // AccountStatement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/get", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.GetAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatementgetAll", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.GetAllAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/create", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.CreateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/update", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.UpdateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteAccountStatement/delete", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.DeleteAccountStatement)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/assignAccount", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.AssignAccountToAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.AccountStatementController.UnassignAccountFromAccountStatement)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Transaction/get", jsonResponseFormatter.FormatToJSON(controller.TransactionController.GetTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TransactiongetAll", jsonResponseFormatter.FormatToJSON(controller.TransactionController.GetAllTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Transaction/create", jsonResponseFormatter.FormatToJSON(controller.TransactionController.CreateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Transaction/update", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UpdateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteTransaction/delete", jsonResponseFormatter.FormatToJSON(controller.TransactionController.DeleteTransaction)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Transaction/assignAccount", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignAccountToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignAccountFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignExternalCounterparty", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignExternalCounterpartyToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignExternalCounterparty", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignExternalCounterpartyFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignPaymentCardToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignPaymentCardFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignFundsTransfer", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignFundsTransferToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignFundsTransfer", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignFundsTransferFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignFxTrade", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignFxTradeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignFxTrade", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignFxTradeFromTransaction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Transaction/assignDispute", jsonResponseFormatter.FormatToJSON(controller.TransactionController.AssignDisputeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Transaction/unassignDispute", jsonResponseFormatter.FormatToJSON(controller.TransactionController.UnassignDisputeFromTransaction)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ExternalAccount/get", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.GetExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExternalAccountgetAll", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.GetAllExternalAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/create", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.CreateExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/update", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.UpdateExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteExternalAccount/delete", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.DeleteExternalAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExternalAccount/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.AssignCustomerToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.UnassignCustomerFromExternalAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExternalAccount/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.AddTransactionsToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExternalAccountremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.ExternalAccountController.RemoveTransactionsFromExternalAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // FundsTransfer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/get", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.GetFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FundsTransfergetAll", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.GetAllFundsTransfer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/create", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.CreateFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/update", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.UpdateFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFundsTransfer/delete", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.DeleteFundsTransfer)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/assignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.AssignSourceAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.UnassignSourceAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.AssignDestinationAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.UnassignDestinationAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.AssignExternalBeneficiaryToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignExternalBeneficiary", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.UnassignExternalBeneficiaryFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FundsTransfer/assignInitiatedBy", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.AssignInitiatedByToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/unassignInitiatedBy", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.UnassignInitiatedByFromFundsTransfer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.AddTransactionsToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FundsTransferremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.FundsTransferController.RemoveTransactionsFromFundsTransfer)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // StandingInstruction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StandingInstruction/get", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.GetStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StandingInstructiongetAll", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.GetAllStandingInstruction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/create", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.CreateStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/update", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.UpdateStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteStandingInstruction/delete", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.DeleteStandingInstruction)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StandingInstruction/assignAccount", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.AssignAccountToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.UnassignAccountFromStandingInstruction)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/StandingInstruction/assignBeneficiary", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.AssignBeneficiaryToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/unassignBeneficiary", jsonResponseFormatter.FormatToJSON(controller.StandingInstructionController.UnassignBeneficiaryFromStandingInstruction)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/PaymentCard/get", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.GetPaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCardgetAll", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.GetAllPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/create", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.CreatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/update", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.UpdatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentCard/delete", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.DeletePaymentCard)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentCard/assignBank", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.AssignBankToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignBank", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.UnassignBankFromPaymentCard)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/PaymentCard/assignAccount", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.AssignAccountToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.UnassignAccountFromPaymentCard)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/PaymentCard/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.AssignCustomerToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.UnassignCustomerFromPaymentCard)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentCard/addToTransactions/", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.AddTransactionsToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/PaymentCardremoveFromTransactions/", jsonResponseFormatter.FormatToJSON(controller.PaymentCardController.RemoveTransactionsFromPaymentCard)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // LoanAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/get", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.GetLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanAccountgetAll", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.GetAllLoanAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/create", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.CreateLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/update", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.UpdateLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanAccount/delete", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.DeleteLoanAccount)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/assignBank", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AssignBankToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignBank", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.UnassignBankFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/assignBranch", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AssignBranchToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignBranch", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.UnassignBranchFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/assignProduct", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AssignProductToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/unassignProduct", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.UnassignProductFromLoanAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/addToBorrowers/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AddBorrowersToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromBorrowers/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.RemoveBorrowersFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AddRepaymentScheduleToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromRepaymentSchedule/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.RemoveRepaymentScheduleFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToPayments/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AddPaymentsToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromPayments/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.RemovePaymentsFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToCollateral/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AddCollateralToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromCollateral/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.RemoveCollateralFromLoanAccount)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanAccount/addToFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.AddFeeChargesToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanAccountremoveFromFeeCharges/", jsonResponseFormatter.FormatToJSON(controller.LoanAccountController.RemoveFeeChargesFromLoanAccount)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // RepaymentSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/get", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.GetRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedulegetAll", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.GetAllRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/create", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.CreateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/update", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.UpdateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteRepaymentSchedule/delete", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.DeleteRepaymentSchedule)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.AssignLoanAccountToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.UnassignLoanAccountFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/RepaymentSchedule/assignPayment", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.AssignPaymentToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/unassignPayment", jsonResponseFormatter.FormatToJSON(controller.RepaymentScheduleController.UnassignPaymentFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/LoanPayment/get", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.GetLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanPaymentgetAll", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.GetAllLoanPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/create", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.CreateLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/update", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.UpdateLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanPayment/delete", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.DeleteLoanPayment)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanPayment/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.AssignLoanAccountToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.UnassignLoanAccountFromLoanPayment)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/LoanPayment/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.AssignTransactionToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.LoanPaymentController.UnassignTransactionFromLoanPayment)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Collateral/get", jsonResponseFormatter.FormatToJSON(controller.CollateralController.GetCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CollateralgetAll", jsonResponseFormatter.FormatToJSON(controller.CollateralController.GetAllCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Collateral/create", jsonResponseFormatter.FormatToJSON(controller.CollateralController.CreateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Collateral/update", jsonResponseFormatter.FormatToJSON(controller.CollateralController.UpdateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteCollateral/delete", jsonResponseFormatter.FormatToJSON(controller.CollateralController.DeleteCollateral)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Collateral/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.CollateralController.AssignLoanAccountToCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Collateral/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.CollateralController.UnassignLoanAccountFromCollateral)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/FeeCharge/get", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.GetFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeChargegetAll", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.GetAllFeeCharge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/create", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.CreateFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/update", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.UpdateFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFeeCharge/delete", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.DeleteFeeCharge)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FeeCharge/assignAccount", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.AssignAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.UnassignAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FeeCharge/assignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.AssignLoanAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/unassignLoanAccount", jsonResponseFormatter.FormatToJSON(controller.FeeChargeController.UnassignLoanAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/ExchangeRate/get", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.GetExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRategetAll", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.GetAllExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/create", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.CreateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/update", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.UpdateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteExchangeRate/delete", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.DeleteExchangeRate)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExchangeRate/assignBank", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.AssignBankToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/unassignBank", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.UnassignBankFromExchangeRate)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExchangeRate/addToFxTrades/", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.AddFxTradesToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ExchangeRateremoveFromFxTrades/", jsonResponseFormatter.FormatToJSON(controller.ExchangeRateController.RemoveFxTradesFromExchangeRate)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // FXTrade Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXTrade/get", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.GetFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXTradegetAll", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.GetAllFXTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FXTrade/create", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.CreateFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXTrade/update", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UpdateFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteFXTrade/delete", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.DeleteFXTrade)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXTrade/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignCustomerToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignCustomerFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignBank", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignBankToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignBank", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignBankFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignExchangeRate", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignExchangeRateToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignExchangeRate", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignExchangeRateFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignSourceAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignSourceAccount", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignSourceAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignDestinationAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignDestinationAccount", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignDestinationAccountFromFXTrade)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/FXTrade/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.AssignTransactionToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/FXTrade/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.FXTradeController.UnassignTransactionFromFXTrade)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Dispute/get", jsonResponseFormatter.FormatToJSON(controller.DisputeController.GetDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DisputegetAll", jsonResponseFormatter.FormatToJSON(controller.DisputeController.GetAllDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dispute/create", jsonResponseFormatter.FormatToJSON(controller.DisputeController.CreateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dispute/update", jsonResponseFormatter.FormatToJSON(controller.DisputeController.UpdateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteDispute/delete", jsonResponseFormatter.FormatToJSON(controller.DisputeController.DeleteDispute)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dispute/assignTransaction", jsonResponseFormatter.FormatToJSON(controller.DisputeController.AssignTransactionToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignTransaction", jsonResponseFormatter.FormatToJSON(controller.DisputeController.UnassignTransactionFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.DisputeController.AssignCustomerToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.DisputeController.UnassignCustomerFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignAccount", jsonResponseFormatter.FormatToJSON(controller.DisputeController.AssignAccountToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignAccount", jsonResponseFormatter.FormatToJSON(controller.DisputeController.UnassignAccountFromDispute)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Dispute/assignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.DisputeController.AssignPaymentCardToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Dispute/unassignPaymentCard", jsonResponseFormatter.FormatToJSON(controller.DisputeController.UnassignPaymentCardFromDispute)).Methods("DELETE", "OPTIONS")


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
    router.HandleFunc("/api/Consent/get", jsonResponseFormatter.FormatToJSON(controller.ConsentController.GetConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ConsentgetAll", jsonResponseFormatter.FormatToJSON(controller.ConsentController.GetAllConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Consent/create", jsonResponseFormatter.FormatToJSON(controller.ConsentController.CreateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Consent/update", jsonResponseFormatter.FormatToJSON(controller.ConsentController.UpdateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteConsent/delete", jsonResponseFormatter.FormatToJSON(controller.ConsentController.DeleteConsent)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Consent/assignCustomer", jsonResponseFormatter.FormatToJSON(controller.ConsentController.AssignCustomerToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignCustomer", jsonResponseFormatter.FormatToJSON(controller.ConsentController.UnassignCustomerFromConsent)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Consent/assignBank", jsonResponseFormatter.FormatToJSON(controller.ConsentController.AssignBankToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignBank", jsonResponseFormatter.FormatToJSON(controller.ConsentController.UnassignBankFromConsent)).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/Consent/assignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(controller.ConsentController.AssignThirdPartyProviderToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/Consent/unassignThirdPartyProvider", jsonResponseFormatter.FormatToJSON(controller.ConsentController.UnassignThirdPartyProviderFromConsent)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Consent/addToAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(controller.ConsentController.AddAuthorizedAccountsToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ConsentremoveFromAuthorizedAccounts/", jsonResponseFormatter.FormatToJSON(controller.ConsentController.RemoveAuthorizedAccountsFromConsent)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // ThirdPartyProvider Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/get", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.GetThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvidergetAll", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.GetAllThirdPartyProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/create", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.CreateThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/update", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.UpdateThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeleteThirdPartyProvider/delete", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.DeleteThirdPartyProvider)).Methods("POST", "OPTIONS")



    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/assignBank", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.AssignBankToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/unassignBank", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.UnassignBankFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")


    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/addToConsents/", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.AddConsentsToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProviderremoveFromConsents/", jsonResponseFormatter.FormatToJSON(controller.ThirdPartyProviderController.RemoveConsentsFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")


    return router
}
