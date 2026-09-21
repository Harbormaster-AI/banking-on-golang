
package controller

import (
    BankDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to BankDAO for database creation
// ----------------------------------------------------------------------------
func CreateBank(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Bank model
	// ----------------------------------------------------------------------------
	data := model.Bank{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Bank model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank data access object to create
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.CreateBank( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Get controller, delegates to BankDAO to find the relevant Bank
// ----------------------------------------------------------------------------
func GetBank(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.GetBank(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to BankDAO for database read of all Banks
// ----------------------------------------------------------------------------
func GetAllBank(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the Bank data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.GetAllBank()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Update controller, delegates to BankDAO for database save
// ----------------------------------------------------------------------------
func UpdateBank(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Bank model
	// ----------------------------------------------------------------------------
	var data = model.Bank{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Bank model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.UpdateBank(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to BankDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteBank(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := BankDAO.DeleteBank(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	// ----------------------------------------------------------------------------
	// adds one or more branchesIds as a Branches to a Bank
	// ----------------------------------------------------------------------------
func AddBranchesToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddBranchesToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more branchesIds as a Branches from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveBranchesFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveBranchesFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more productsIds as a Products to a Bank
	// ----------------------------------------------------------------------------
func AddProductsToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddProductsToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more productsIds as a Products from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveProductsFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveProductsFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more customersIds as a Customers to a Bank
	// ----------------------------------------------------------------------------
func AddCustomersToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddCustomersToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more customersIds as a Customers from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveCustomersFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveCustomersFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Bank
	// ----------------------------------------------------------------------------
func AddAccountsToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddAccountsToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveAccountsFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveAccountsFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more paymentCardsIds as a PaymentCards to a Bank
	// ----------------------------------------------------------------------------
func AddPaymentCardsToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddPaymentCardsToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more paymentCardsIds as a PaymentCards from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemovePaymentCardsFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemovePaymentCardsFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more loanAccountsIds as a LoanAccounts to a Bank
	// ----------------------------------------------------------------------------
func AddLoanAccountsToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddLoanAccountsToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more loanAccountsIds as a LoanAccounts from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveLoanAccountsFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveLoanAccountsFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more exchangeRatesIds as a ExchangeRates to a Bank
	// ----------------------------------------------------------------------------
func AddExchangeRatesToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddExchangeRatesToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more exchangeRatesIds as a ExchangeRates from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveExchangeRatesFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveExchangeRatesFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a Bank
	// ----------------------------------------------------------------------------
func AddConsentsToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddConsentsToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveConsentsFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveConsentsFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	// ----------------------------------------------------------------------------
	// adds one or more thirdPartyProvidersIds as a ThirdPartyProviders to a Bank
	// ----------------------------------------------------------------------------
func AddThirdPartyProvidersToBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.AddThirdPartyProvidersToBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// removes one or more thirdPartyProvidersIds as a ThirdPartyProviders from a Bank
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveThirdPartyProvidersFromBank(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	// ----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveThirdPartyProvidersFromBank(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
