package controller

import (
    BankDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BankDAO for database creation
//----------------------------------------------------------------------------
func CreateBank(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Bank model
	//----------------------------------------------------------------------------
	data := model.Bank{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Bank model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Bank data access object to create
	//----------------------------------------------------------------------------
	requestResult := BankDAO.CreateBank( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BankDAO to find the relevant Bank
//----------------------------------------------------------------------------
func GetBank(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BankDAO.GetBank(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BankDAO for database read of all Banks
//----------------------------------------------------------------------------
func GetAllBank(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Bank data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BankDAO.GetAllBank()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BankDAO for database save
//----------------------------------------------------------------------------
func UpdateBank(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Bank model
	//----------------------------------------------------------------------------
	var data = model.Bank{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Bank model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BankDAO.UpdateBank(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BankDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBank(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the Bank data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BankDAO.DeleteBank(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more branchesIds as a Branches to a Bank
	//----------------------------------------------------------------------------
func AddBranchesToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchesIds,_ := vars["branchesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddBranchesToBank(bankId, branchesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more branchesIds as a Branches from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBranchesFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchesIds,_ := vars["branchesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveBranchesFromBank(bankId, branchesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more productsIds as a Products to a Bank
	//----------------------------------------------------------------------------
func AddProductsToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddProductsToBank(bankId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productsIds as a Products from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductsFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveProductsFromBank(bankId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more customersIds as a Customers to a Bank
	//----------------------------------------------------------------------------
func AddCustomersToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddCustomersToBank(bankId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more customersIds as a Customers from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCustomersFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveCustomersFromBank(bankId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Bank
	//----------------------------------------------------------------------------
func AddAccountsToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddAccountsToBank(bankId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveAccountsFromBank(bankId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentCardsIds as a PaymentCards to a Bank
	//----------------------------------------------------------------------------
func AddPaymentCardsToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddPaymentCardsToBank(bankId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentCardsIds as a PaymentCards from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentCardsFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemovePaymentCardsFromBank(bankId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loanAccountsIds as a LoanAccounts to a Bank
	//----------------------------------------------------------------------------
func AddLoanAccountsToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddLoanAccountsToBank(bankId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loanAccountsIds as a LoanAccounts from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveLoanAccountsFromBank(bankId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more exchangeRatesIds as a ExchangeRates to a Bank
	//----------------------------------------------------------------------------
func AddExchangeRatesToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exchangeRatesIds,_ := vars["exchangeRatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddExchangeRatesToBank(bankId, exchangeRatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more exchangeRatesIds as a ExchangeRates from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExchangeRatesFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exchangeRatesIds,_ := vars["exchangeRatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveExchangeRatesFromBank(bankId, exchangeRatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a Bank
	//----------------------------------------------------------------------------
func AddConsentsToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddConsentsToBank(bankId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveConsentsFromBank(bankId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more thirdPartyProvidersIds as a ThirdPartyProviders to a Bank
	//----------------------------------------------------------------------------
func AddThirdPartyProvidersToBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartyProvidersIds,_ := vars["thirdPartyProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.AddThirdPartyProvidersToBank(bankId, thirdPartyProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more thirdPartyProvidersIds as a ThirdPartyProviders from a Bank
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveThirdPartyProvidersFromBank(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartyProvidersIds,_ := vars["thirdPartyProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Bank DAO
	//----------------------------------------------------------------------------
	requestResult := BankDAO.RemoveThirdPartyProvidersFromBank(bankId, thirdPartyProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
