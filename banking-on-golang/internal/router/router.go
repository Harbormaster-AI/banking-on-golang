package router

import (

    BankController "banking-on-golang/internal/controller"
    BranchController "banking-on-golang/internal/controller"
    ATMController "banking-on-golang/internal/controller"
    CustomerController "banking-on-golang/internal/controller"
    KycProfileController "banking-on-golang/internal/controller"
    IdentityDocumentController "banking-on-golang/internal/controller"
    RiskAssessmentController "banking-on-golang/internal/controller"
    ScreeningResultController "banking-on-golang/internal/controller"
    BankingProductController "banking-on-golang/internal/controller"
    AccountController "banking-on-golang/internal/controller"
    AccountStatementController "banking-on-golang/internal/controller"
    TransactionController "banking-on-golang/internal/controller"
    ExternalAccountController "banking-on-golang/internal/controller"
    FundsTransferController "banking-on-golang/internal/controller"
    StandingInstructionController "banking-on-golang/internal/controller"
    PaymentCardController "banking-on-golang/internal/controller"
    LoanAccountController "banking-on-golang/internal/controller"
    RepaymentScheduleController "banking-on-golang/internal/controller"
    LoanPaymentController "banking-on-golang/internal/controller"
    CollateralController "banking-on-golang/internal/controller"
    FeeChargeController "banking-on-golang/internal/controller"
    ExchangeRateController "banking-on-golang/internal/controller"
    FXTradeController "banking-on-golang/internal/controller"
    DisputeController "banking-on-golang/internal/controller"
    ConsentController "banking-on-golang/internal/controller"
    ThirdPartyProviderController "banking-on-golang/internal/controller"
    jsonResponseFormatter "banking-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "banking-on-golang/internal/controller"

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
    router.HandleFunc("/api/Bank/{id}", jsonResponseFormatter.FormatToJSON(BankController.GetBank)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Bank", jsonResponseFormatter.FormatToJSON(BankController.GetAllBank)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBank", jsonResponseFormatter.FormatToJSON(BankController.CreateBank)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Bank/{id}", jsonResponseFormatter.FormatToJSON(BankController.UpdateBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBank/{id}", jsonResponseFormatter.FormatToJSON(BankController.DeleteBank)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBranchesToBank/{parentId}/branchesId", jsonResponseFormatter.FormatToJSON(BankController.AddBranchesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBranchesFromBank/{parentId}/branchesIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveBranchesFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProductsToBank/{parentId}/productsId", jsonResponseFormatter.FormatToJSON(BankController.AddProductsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductsFromBank/{parentId}/productsIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveProductsFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCustomersToBank/{parentId}/customersId", jsonResponseFormatter.FormatToJSON(BankController.AddCustomersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCustomersFromBank/{parentId}/customersIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveCustomersFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAccountsToBank/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(BankController.AddAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromBank/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveAccountsFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentCardsToBank/{parentId}/paymentCardsId", jsonResponseFormatter.FormatToJSON(BankController.AddPaymentCardsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentCardsFromBank/{parentId}/paymentCardsIds", jsonResponseFormatter.FormatToJSON(BankController.RemovePaymentCardsFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoanAccountsToBank/{parentId}/loanAccountsId", jsonResponseFormatter.FormatToJSON(BankController.AddLoanAccountsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoanAccountsFromBank/{parentId}/loanAccountsIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveLoanAccountsFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExchangeRatesToBank/{parentId}/exchangeRatesId", jsonResponseFormatter.FormatToJSON(BankController.AddExchangeRatesToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExchangeRatesFromBank/{parentId}/exchangeRatesIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveExchangeRatesFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConsentsToBank/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(BankController.AddConsentsToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromBank/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveConsentsFromBank)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddThirdPartyProvidersToBank/{parentId}/thirdPartyProvidersId", jsonResponseFormatter.FormatToJSON(BankController.AddThirdPartyProvidersToBank)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveThirdPartyProvidersFromBank/{parentId}/thirdPartyProvidersIds", jsonResponseFormatter.FormatToJSON(BankController.RemoveThirdPartyProvidersFromBank)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Branch Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.GetBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Branch", jsonResponseFormatter.FormatToJSON(BranchController.GetAllBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBranch", jsonResponseFormatter.FormatToJSON(BranchController.CreateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Branch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.UpdateBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBranch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.DeleteBranch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToBranch/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(BranchController.AssignBankToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromBranch/{parentId}", jsonResponseFormatter.FormatToJSON(BranchController.UnassignBankFromBranch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToBranch/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(BranchController.AddAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromBranch/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(BranchController.RemoveAccountsFromBranch)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoanAccountsToBranch/{parentId}/loanAccountsId", jsonResponseFormatter.FormatToJSON(BranchController.AddLoanAccountsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoanAccountsFromBranch/{parentId}/loanAccountsIds", jsonResponseFormatter.FormatToJSON(BranchController.RemoveLoanAccountsFromBranch)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAtmsToBranch/{parentId}/atmsId", jsonResponseFormatter.FormatToJSON(BranchController.AddAtmsToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAtmsFromBranch/{parentId}/atmsIds", jsonResponseFormatter.FormatToJSON(BranchController.RemoveAtmsFromBranch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ATM Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ATM/{id}", jsonResponseFormatter.FormatToJSON(ATMController.GetATM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ATM", jsonResponseFormatter.FormatToJSON(ATMController.GetAllATM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewATM", jsonResponseFormatter.FormatToJSON(ATMController.CreateATM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ATM/{id}", jsonResponseFormatter.FormatToJSON(ATMController.UpdateATM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteATM/{id}", jsonResponseFormatter.FormatToJSON(ATMController.DeleteATM)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBranchToATM/{parentId}/branchId", jsonResponseFormatter.FormatToJSON(ATMController.AssignBranchToATM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBranchFromATM/{parentId}", jsonResponseFormatter.FormatToJSON(ATMController.UnassignBranchFromATM)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.GetCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Customer", jsonResponseFormatter.FormatToJSON(CustomerController.GetAllCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCustomer", jsonResponseFormatter.FormatToJSON(CustomerController.CreateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.UpdateCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCustomer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.DeleteCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToCustomer/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(CustomerController.AssignBankToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromCustomer/{parentId}", jsonResponseFormatter.FormatToJSON(CustomerController.UnassignBankFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToCustomer/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromCustomer/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveAccountsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoanAccountsToCustomer/{parentId}/loanAccountsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddLoanAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoanAccountsFromCustomer/{parentId}/loanAccountsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveLoanAccountsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentCardsToCustomer/{parentId}/paymentCardsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddPaymentCardsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentCardsFromCustomer/{parentId}/paymentCardsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemovePaymentCardsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExternalAccountsToCustomer/{parentId}/externalAccountsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddExternalAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExternalAccountsFromCustomer/{parentId}/externalAccountsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveExternalAccountsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFundsTransfersToCustomer/{parentId}/fundsTransfersId", jsonResponseFormatter.FormatToJSON(CustomerController.AddFundsTransfersToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFundsTransfersFromCustomer/{parentId}/fundsTransfersIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveFundsTransfersFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDisputesToCustomer/{parentId}/disputesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddDisputesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDisputesFromCustomer/{parentId}/disputesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveDisputesFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddKycProfilesToCustomer/{parentId}/kycProfilesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddKycProfilesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveKycProfilesFromCustomer/{parentId}/kycProfilesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveKycProfilesFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConsentsToCustomer/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddConsentsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromCustomer/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveConsentsFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // KycProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KycProfile/{id}", jsonResponseFormatter.FormatToJSON(KycProfileController.GetKycProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KycProfile", jsonResponseFormatter.FormatToJSON(KycProfileController.GetAllKycProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewKycProfile", jsonResponseFormatter.FormatToJSON(KycProfileController.CreateKycProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KycProfile/{id}", jsonResponseFormatter.FormatToJSON(KycProfileController.UpdateKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteKycProfile/{id}", jsonResponseFormatter.FormatToJSON(KycProfileController.DeleteKycProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToKycProfile/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(KycProfileController.AssignCustomerToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromKycProfile/{parentId}", jsonResponseFormatter.FormatToJSON(KycProfileController.UnassignCustomerFromKycProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddIdentityDocumentsToKycProfile/{parentId}/identityDocumentsId", jsonResponseFormatter.FormatToJSON(KycProfileController.AddIdentityDocumentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveIdentityDocumentsFromKycProfile/{parentId}/identityDocumentsIds", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveIdentityDocumentsFromKycProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRiskAssessmentsToKycProfile/{parentId}/riskAssessmentsId", jsonResponseFormatter.FormatToJSON(KycProfileController.AddRiskAssessmentsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRiskAssessmentsFromKycProfile/{parentId}/riskAssessmentsIds", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveRiskAssessmentsFromKycProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddScreeningsToKycProfile/{parentId}/screeningsId", jsonResponseFormatter.FormatToJSON(KycProfileController.AddScreeningsToKycProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveScreeningsFromKycProfile/{parentId}/screeningsIds", jsonResponseFormatter.FormatToJSON(KycProfileController.RemoveScreeningsFromKycProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // IdentityDocument Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/IdentityDocument/{id}", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.GetIdentityDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.GetAllIdentityDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewIdentityDocument", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.CreateIdentityDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/IdentityDocument/{id}", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.UpdateIdentityDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteIdentityDocument/{id}", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.DeleteIdentityDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToIdentityDocument/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.AssignKycProfileToIdentityDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromIdentityDocument/{parentId}", jsonResponseFormatter.FormatToJSON(IdentityDocumentController.UnassignKycProfileFromIdentityDocument)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/RiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetAllRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRiskAssessment", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.CreateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UpdateRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.DeleteRiskAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToRiskAssessment/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.AssignKycProfileToRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromRiskAssessment/{parentId}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UnassignKycProfileFromRiskAssessment)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/ScreeningResult/{id}", jsonResponseFormatter.FormatToJSON(ScreeningResultController.GetScreeningResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult", jsonResponseFormatter.FormatToJSON(ScreeningResultController.GetAllScreeningResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewScreeningResult", jsonResponseFormatter.FormatToJSON(ScreeningResultController.CreateScreeningResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScreeningResult/{id}", jsonResponseFormatter.FormatToJSON(ScreeningResultController.UpdateScreeningResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteScreeningResult/{id}", jsonResponseFormatter.FormatToJSON(ScreeningResultController.DeleteScreeningResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToScreeningResult/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(ScreeningResultController.AssignKycProfileToScreeningResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromScreeningResult/{parentId}", jsonResponseFormatter.FormatToJSON(ScreeningResultController.UnassignKycProfileFromScreeningResult)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/BankingProduct/{id}", jsonResponseFormatter.FormatToJSON(BankingProductController.GetBankingProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BankingProduct", jsonResponseFormatter.FormatToJSON(BankingProductController.GetAllBankingProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBankingProduct", jsonResponseFormatter.FormatToJSON(BankingProductController.CreateBankingProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankingProduct/{id}", jsonResponseFormatter.FormatToJSON(BankingProductController.UpdateBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBankingProduct/{id}", jsonResponseFormatter.FormatToJSON(BankingProductController.DeleteBankingProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToBankingProduct/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(BankingProductController.AssignBankToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromBankingProduct/{parentId}", jsonResponseFormatter.FormatToJSON(BankingProductController.UnassignBankFromBankingProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToBankingProduct/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(BankingProductController.AddAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromBankingProduct/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(BankingProductController.RemoveAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoanAccountsToBankingProduct/{parentId}/loanAccountsId", jsonResponseFormatter.FormatToJSON(BankingProductController.AddLoanAccountsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoanAccountsFromBankingProduct/{parentId}/loanAccountsIds", jsonResponseFormatter.FormatToJSON(BankingProductController.RemoveLoanAccountsFromBankingProduct)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentCardsToBankingProduct/{parentId}/paymentCardsId", jsonResponseFormatter.FormatToJSON(BankingProductController.AddPaymentCardsToBankingProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentCardsFromBankingProduct/{parentId}/paymentCardsIds", jsonResponseFormatter.FormatToJSON(BankingProductController.RemovePaymentCardsFromBankingProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Account Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/{id}", jsonResponseFormatter.FormatToJSON(AccountController.GetAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Account", jsonResponseFormatter.FormatToJSON(AccountController.GetAllAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAccount", jsonResponseFormatter.FormatToJSON(AccountController.CreateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Account/{id}", jsonResponseFormatter.FormatToJSON(AccountController.UpdateAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAccount/{id}", jsonResponseFormatter.FormatToJSON(AccountController.DeleteAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToAccount/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(AccountController.AssignBankToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignBankFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBranchToAccount/{parentId}/branchId", jsonResponseFormatter.FormatToJSON(AccountController.AssignBranchToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBranchFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignBranchFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToAccount/{parentId}/productId", jsonResponseFormatter.FormatToJSON(AccountController.AssignProductToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignProductFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOwnersToAccount/{parentId}/ownersId", jsonResponseFormatter.FormatToJSON(AccountController.AddOwnersToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnersFromAccount/{parentId}/ownersIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveOwnersFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToAccount/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(AccountController.AddTransactionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromAccount/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveTransactionsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddStatementsToAccount/{parentId}/statementsId", jsonResponseFormatter.FormatToJSON(AccountController.AddStatementsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveStatementsFromAccount/{parentId}/statementsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveStatementsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddStandingInstructionsToAccount/{parentId}/standingInstructionsId", jsonResponseFormatter.FormatToJSON(AccountController.AddStandingInstructionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveStandingInstructionsFromAccount/{parentId}/standingInstructionsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveStandingInstructionsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeeChargesToAccount/{parentId}/feeChargesId", jsonResponseFormatter.FormatToJSON(AccountController.AddFeeChargesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeeChargesFromAccount/{parentId}/feeChargesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveFeeChargesFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AccountStatement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AccountStatement", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAllAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAccountStatement", jsonResponseFormatter.FormatToJSON(AccountStatementController.CreateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.UpdateAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.DeleteAccountStatement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToAccountStatement/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(AccountStatementController.AssignAccountToAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromAccountStatement/{parentId}", jsonResponseFormatter.FormatToJSON(AccountStatementController.UnassignAccountFromAccountStatement)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/Transaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.GetTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Transaction", jsonResponseFormatter.FormatToJSON(TransactionController.GetAllTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTransaction", jsonResponseFormatter.FormatToJSON(TransactionController.CreateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Transaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.UpdateTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTransaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.DeleteTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToTransaction/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignAccountToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignAccountFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExternalCounterpartyToTransaction/{parentId}/externalCounterpartyId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignExternalCounterpartyToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExternalCounterpartyFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignExternalCounterpartyFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPaymentCardToTransaction/{parentId}/paymentCardId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignPaymentCardToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPaymentCardFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignPaymentCardFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFundsTransferToTransaction/{parentId}/fundsTransferId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignFundsTransferToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFundsTransferFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignFundsTransferFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFxTradeToTransaction/{parentId}/fxTradeId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignFxTradeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFxTradeFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignFxTradeFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDisputeToTransaction/{parentId}/disputeId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignDisputeToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDisputeFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignDisputeFromTransaction)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/ExternalAccount/{id}", jsonResponseFormatter.FormatToJSON(ExternalAccountController.GetExternalAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount", jsonResponseFormatter.FormatToJSON(ExternalAccountController.GetAllExternalAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExternalAccount", jsonResponseFormatter.FormatToJSON(ExternalAccountController.CreateExternalAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExternalAccount/{id}", jsonResponseFormatter.FormatToJSON(ExternalAccountController.UpdateExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExternalAccount/{id}", jsonResponseFormatter.FormatToJSON(ExternalAccountController.DeleteExternalAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToExternalAccount/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ExternalAccountController.AssignCustomerToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromExternalAccount/{parentId}", jsonResponseFormatter.FormatToJSON(ExternalAccountController.UnassignCustomerFromExternalAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToExternalAccount/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(ExternalAccountController.AddTransactionsToExternalAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromExternalAccount/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(ExternalAccountController.RemoveTransactionsFromExternalAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FundsTransfer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FundsTransfer/{id}", jsonResponseFormatter.FormatToJSON(FundsTransferController.GetFundsTransfer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer", jsonResponseFormatter.FormatToJSON(FundsTransferController.GetAllFundsTransfer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFundsTransfer", jsonResponseFormatter.FormatToJSON(FundsTransferController.CreateFundsTransfer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FundsTransfer/{id}", jsonResponseFormatter.FormatToJSON(FundsTransferController.UpdateFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFundsTransfer/{id}", jsonResponseFormatter.FormatToJSON(FundsTransferController.DeleteFundsTransfer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSourceAccountToFundsTransfer/{parentId}/sourceAccountId", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignSourceAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSourceAccountFromFundsTransfer/{parentId}", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignSourceAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationAccountToFundsTransfer/{parentId}/destinationAccountId", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignDestinationAccountToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationAccountFromFundsTransfer/{parentId}", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignDestinationAccountFromFundsTransfer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExternalBeneficiaryToFundsTransfer/{parentId}/externalBeneficiaryId", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignExternalBeneficiaryToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExternalBeneficiaryFromFundsTransfer/{parentId}", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignExternalBeneficiaryFromFundsTransfer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInitiatedByToFundsTransfer/{parentId}/initiatedById", jsonResponseFormatter.FormatToJSON(FundsTransferController.AssignInitiatedByToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInitiatedByFromFundsTransfer/{parentId}", jsonResponseFormatter.FormatToJSON(FundsTransferController.UnassignInitiatedByFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToFundsTransfer/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(FundsTransferController.AddTransactionsToFundsTransfer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromFundsTransfer/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(FundsTransferController.RemoveTransactionsFromFundsTransfer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // StandingInstruction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StandingInstruction/{id}", jsonResponseFormatter.FormatToJSON(StandingInstructionController.GetStandingInstruction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction", jsonResponseFormatter.FormatToJSON(StandingInstructionController.GetAllStandingInstruction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewStandingInstruction", jsonResponseFormatter.FormatToJSON(StandingInstructionController.CreateStandingInstruction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StandingInstruction/{id}", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UpdateStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteStandingInstruction/{id}", jsonResponseFormatter.FormatToJSON(StandingInstructionController.DeleteStandingInstruction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToStandingInstruction/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(StandingInstructionController.AssignAccountToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromStandingInstruction/{parentId}", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UnassignAccountFromStandingInstruction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBeneficiaryToStandingInstruction/{parentId}/beneficiaryId", jsonResponseFormatter.FormatToJSON(StandingInstructionController.AssignBeneficiaryToStandingInstruction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBeneficiaryFromStandingInstruction/{parentId}", jsonResponseFormatter.FormatToJSON(StandingInstructionController.UnassignBeneficiaryFromStandingInstruction)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/PaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentCard", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetAllPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentCard", jsonResponseFormatter.FormatToJSON(PaymentCardController.CreatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UpdatePaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.DeletePaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToPaymentCard/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignBankToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromPaymentCard/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignBankFromPaymentCard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToPaymentCard/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignAccountToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromPaymentCard/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignAccountFromPaymentCard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToPaymentCard/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignCustomerToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromPaymentCard/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignCustomerFromPaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToPaymentCard/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AddTransactionsToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromPaymentCard/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(PaymentCardController.RemoveTransactionsFromPaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LoanAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanAccount/{id}", jsonResponseFormatter.FormatToJSON(LoanAccountController.GetLoanAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanAccount", jsonResponseFormatter.FormatToJSON(LoanAccountController.GetAllLoanAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLoanAccount", jsonResponseFormatter.FormatToJSON(LoanAccountController.CreateLoanAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanAccount/{id}", jsonResponseFormatter.FormatToJSON(LoanAccountController.UpdateLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanAccount/{id}", jsonResponseFormatter.FormatToJSON(LoanAccountController.DeleteLoanAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToLoanAccount/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignBankToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromLoanAccount/{parentId}", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignBankFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBranchToLoanAccount/{parentId}/branchId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignBranchToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBranchFromLoanAccount/{parentId}", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignBranchFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToLoanAccount/{parentId}/productId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AssignProductToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromLoanAccount/{parentId}", jsonResponseFormatter.FormatToJSON(LoanAccountController.UnassignProductFromLoanAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBorrowersToLoanAccount/{parentId}/borrowersId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddBorrowersToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBorrowersFromLoanAccount/{parentId}/borrowersIds", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveBorrowersFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRepaymentScheduleToLoanAccount/{parentId}/repaymentScheduleId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddRepaymentScheduleToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRepaymentScheduleFromLoanAccount/{parentId}/repaymentScheduleIds", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveRepaymentScheduleFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentsToLoanAccount/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddPaymentsToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromLoanAccount/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemovePaymentsFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCollateralToLoanAccount/{parentId}/collateralId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddCollateralToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCollateralFromLoanAccount/{parentId}/collateralIds", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveCollateralFromLoanAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeeChargesToLoanAccount/{parentId}/feeChargesId", jsonResponseFormatter.FormatToJSON(LoanAccountController.AddFeeChargesToLoanAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeeChargesFromLoanAccount/{parentId}/feeChargesIds", jsonResponseFormatter.FormatToJSON(LoanAccountController.RemoveFeeChargesFromLoanAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RepaymentSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetAllRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRepaymentSchedule", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.CreateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UpdateRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.DeleteRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanAccountToRepaymentSchedule/{parentId}/loanAccountId", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AssignLoanAccountToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanAccountFromRepaymentSchedule/{parentId}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UnassignLoanAccountFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPaymentToRepaymentSchedule/{parentId}/paymentId", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AssignPaymentToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPaymentFromRepaymentSchedule/{parentId}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UnassignPaymentFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/LoanPayment/{id}", jsonResponseFormatter.FormatToJSON(LoanPaymentController.GetLoanPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanPayment", jsonResponseFormatter.FormatToJSON(LoanPaymentController.GetAllLoanPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLoanPayment", jsonResponseFormatter.FormatToJSON(LoanPaymentController.CreateLoanPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanPayment/{id}", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UpdateLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanPayment/{id}", jsonResponseFormatter.FormatToJSON(LoanPaymentController.DeleteLoanPayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanAccountToLoanPayment/{parentId}/loanAccountId", jsonResponseFormatter.FormatToJSON(LoanPaymentController.AssignLoanAccountToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanAccountFromLoanPayment/{parentId}", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UnassignLoanAccountFromLoanPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransactionToLoanPayment/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(LoanPaymentController.AssignTransactionToLoanPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromLoanPayment/{parentId}", jsonResponseFormatter.FormatToJSON(LoanPaymentController.UnassignTransactionFromLoanPayment)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/Collateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.GetCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Collateral", jsonResponseFormatter.FormatToJSON(CollateralController.GetAllCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCollateral", jsonResponseFormatter.FormatToJSON(CollateralController.CreateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Collateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.UpdateCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCollateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.DeleteCollateral)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanAccountToCollateral/{parentId}/loanAccountId", jsonResponseFormatter.FormatToJSON(CollateralController.AssignLoanAccountToCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanAccountFromCollateral/{parentId}", jsonResponseFormatter.FormatToJSON(CollateralController.UnassignLoanAccountFromCollateral)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/FeeCharge/{id}", jsonResponseFormatter.FormatToJSON(FeeChargeController.GetFeeCharge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FeeCharge", jsonResponseFormatter.FormatToJSON(FeeChargeController.GetAllFeeCharge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFeeCharge", jsonResponseFormatter.FormatToJSON(FeeChargeController.CreateFeeCharge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeCharge/{id}", jsonResponseFormatter.FormatToJSON(FeeChargeController.UpdateFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFeeCharge/{id}", jsonResponseFormatter.FormatToJSON(FeeChargeController.DeleteFeeCharge)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToFeeCharge/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(FeeChargeController.AssignAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromFeeCharge/{parentId}", jsonResponseFormatter.FormatToJSON(FeeChargeController.UnassignAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLoanAccountToFeeCharge/{parentId}/loanAccountId", jsonResponseFormatter.FormatToJSON(FeeChargeController.AssignLoanAccountToFeeCharge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanAccountFromFeeCharge/{parentId}", jsonResponseFormatter.FormatToJSON(FeeChargeController.UnassignLoanAccountFromFeeCharge)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/ExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetAllExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExchangeRate", jsonResponseFormatter.FormatToJSON(ExchangeRateController.CreateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.UpdateExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.DeleteExchangeRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToExchangeRate/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(ExchangeRateController.AssignBankToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromExchangeRate/{parentId}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.UnassignBankFromExchangeRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFxTradesToExchangeRate/{parentId}/fxTradesId", jsonResponseFormatter.FormatToJSON(ExchangeRateController.AddFxTradesToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFxTradesFromExchangeRate/{parentId}/fxTradesIds", jsonResponseFormatter.FormatToJSON(ExchangeRateController.RemoveFxTradesFromExchangeRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FXTrade Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXTrade/{id}", jsonResponseFormatter.FormatToJSON(FXTradeController.GetFXTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FXTrade", jsonResponseFormatter.FormatToJSON(FXTradeController.GetAllFXTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFXTrade", jsonResponseFormatter.FormatToJSON(FXTradeController.CreateFXTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXTrade/{id}", jsonResponseFormatter.FormatToJSON(FXTradeController.UpdateFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFXTrade/{id}", jsonResponseFormatter.FormatToJSON(FXTradeController.DeleteFXTrade)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToFXTrade/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignCustomerToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignCustomerFromFXTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBankToFXTrade/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignBankToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignBankFromFXTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExchangeRateToFXTrade/{parentId}/exchangeRateId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignExchangeRateToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExchangeRateFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignExchangeRateFromFXTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSourceAccountToFXTrade/{parentId}/sourceAccountId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignSourceAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSourceAccountFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignSourceAccountFromFXTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationAccountToFXTrade/{parentId}/destinationAccountId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignDestinationAccountToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationAccountFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignDestinationAccountFromFXTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransactionToFXTrade/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(FXTradeController.AssignTransactionToFXTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromFXTrade/{parentId}", jsonResponseFormatter.FormatToJSON(FXTradeController.UnassignTransactionFromFXTrade)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/Dispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.GetDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dispute", jsonResponseFormatter.FormatToJSON(DisputeController.GetAllDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDispute", jsonResponseFormatter.FormatToJSON(DisputeController.CreateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.UpdateDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.DeleteDispute)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTransactionToDispute/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignTransactionToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignTransactionFromDispute)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToDispute/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignCustomerToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignCustomerFromDispute)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToDispute/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignAccountToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignAccountFromDispute)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPaymentCardToDispute/{parentId}/paymentCardId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignPaymentCardToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPaymentCardFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignPaymentCardFromDispute)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/Consent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.GetConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Consent", jsonResponseFormatter.FormatToJSON(ConsentController.GetAllConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewConsent", jsonResponseFormatter.FormatToJSON(ConsentController.CreateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Consent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.UpdateConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteConsent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.DeleteConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToConsent/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignCustomerToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignCustomerFromConsent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBankToConsent/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignBankToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignBankFromConsent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignThirdPartyProviderToConsent/{parentId}/thirdPartyProviderId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignThirdPartyProviderToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignThirdPartyProviderFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignThirdPartyProviderFromConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAuthorizedAccountsToConsent/{parentId}/authorizedAccountsId", jsonResponseFormatter.FormatToJSON(ConsentController.AddAuthorizedAccountsToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAuthorizedAccountsFromConsent/{parentId}/authorizedAccountsIds", jsonResponseFormatter.FormatToJSON(ConsentController.RemoveAuthorizedAccountsFromConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ThirdPartyProvider Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyProvider/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.GetThirdPartyProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.GetAllThirdPartyProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewThirdPartyProvider", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.CreateThirdPartyProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyProvider/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.UpdateThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteThirdPartyProvider/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.DeleteThirdPartyProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBankToThirdPartyProvider/{parentId}/bankId", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.AssignBankToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankFromThirdPartyProvider/{parentId}", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.UnassignBankFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddConsentsToThirdPartyProvider/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.AddConsentsToThirdPartyProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromThirdPartyProvider/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(ThirdPartyProviderController.RemoveConsentsFromThirdPartyProvider)).Methods("DELETE", "OPTIONS")

    return router
}
