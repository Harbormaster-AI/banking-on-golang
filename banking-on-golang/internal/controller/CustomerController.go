package controller

import (
    CustomerDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CustomerDAO for database creation
//----------------------------------------------------------------------------
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	data := model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to create
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.CreateCustomer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CustomerDAO to find the relevant Customer
//----------------------------------------------------------------------------
func GetCustomer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Customer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetCustomer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CustomerDAO for database read of all Customers
//----------------------------------------------------------------------------
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetAllCustomer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CustomerDAO for database save
//----------------------------------------------------------------------------
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	var data = model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.UpdateCustomer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CustomerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Customer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CustomerDAO.DeleteCustomer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a Customer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AssignBankToCustomer(customerId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a Customer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromCustomer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.UnassignBankFromCustomer(customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Customer
	//----------------------------------------------------------------------------
func AddAccountsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddAccountsToCustomer(customerId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveAccountsFromCustomer(customerId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loanAccountsIds as a LoanAccounts to a Customer
	//----------------------------------------------------------------------------
func AddLoanAccountsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddLoanAccountsToCustomer(customerId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loanAccountsIds as a LoanAccounts from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoanAccountsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveLoanAccountsFromCustomer(customerId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentCardsIds as a PaymentCards to a Customer
	//----------------------------------------------------------------------------
func AddPaymentCardsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddPaymentCardsToCustomer(customerId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentCardsIds as a PaymentCards from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentCardsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemovePaymentCardsFromCustomer(customerId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more externalAccountsIds as a ExternalAccounts to a Customer
	//----------------------------------------------------------------------------
func AddExternalAccountsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	externalAccountsIds,_ := vars["externalAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddExternalAccountsToCustomer(customerId, externalAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more externalAccountsIds as a ExternalAccounts from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExternalAccountsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	externalAccountsIds,_ := vars["externalAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveExternalAccountsFromCustomer(customerId, externalAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more fundsTransfersIds as a FundsTransfers to a Customer
	//----------------------------------------------------------------------------
func AddFundsTransfersToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fundsTransfersIds,_ := vars["fundsTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddFundsTransfersToCustomer(customerId, fundsTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more fundsTransfersIds as a FundsTransfers from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFundsTransfersFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fundsTransfersIds,_ := vars["fundsTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveFundsTransfersFromCustomer(customerId, fundsTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more disputesIds as a Disputes to a Customer
	//----------------------------------------------------------------------------
func AddDisputesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddDisputesToCustomer(customerId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more disputesIds as a Disputes from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDisputesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveDisputesFromCustomer(customerId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more kycProfilesIds as a KycProfiles to a Customer
	//----------------------------------------------------------------------------
func AddKycProfilesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfilesIds,_ := vars["kycProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddKycProfilesToCustomer(customerId, kycProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more kycProfilesIds as a KycProfiles from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveKycProfilesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfilesIds,_ := vars["kycProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveKycProfilesFromCustomer(customerId, kycProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a Customer
	//----------------------------------------------------------------------------
func AddConsentsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddConsentsToCustomer(customerId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveConsentsFromCustomer(customerId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
