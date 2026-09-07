package controller

import (
    ExternalAccountDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExternalAccountDAO for database creation
//----------------------------------------------------------------------------
func CreateExternalAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExternalAccount model
	//----------------------------------------------------------------------------
	data := model.ExternalAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExternalAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.CreateExternalAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExternalAccountDAO to find the relevant ExternalAccount
//----------------------------------------------------------------------------
func GetExternalAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExternalAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.GetExternalAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExternalAccountDAO for database read of all ExternalAccounts
//----------------------------------------------------------------------------
func GetAllExternalAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.GetAllExternalAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExternalAccountDAO for database save
//----------------------------------------------------------------------------
func UpdateExternalAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExternalAccount model
	//----------------------------------------------------------------------------
	var data = model.ExternalAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExternalAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.UpdateExternalAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExternalAccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExternalAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExternalAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExternalAccountDAO.DeleteExternalAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a ExternalAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToExternalAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	externalAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.AssignCustomerToExternalAccount(externalAccountId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a ExternalAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromExternalAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	externalAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.UnassignCustomerFromExternalAccount(externalAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a ExternalAccount
	//----------------------------------------------------------------------------
func AddTransactionsToExternalAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	externalAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.AddTransactionsToExternalAccount(externalAccountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a ExternalAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromExternalAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	externalAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.RemoveTransactionsFromExternalAccount(externalAccountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
