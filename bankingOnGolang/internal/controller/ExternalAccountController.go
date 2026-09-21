
package controller

import (
    ExternalAccountDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.GetExternalAccount(data.Id)
	
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExternalAccountDAO.DeleteExternalAccount(data.Id)

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
	//----------------------------------------------------------------------------
func AssignCustomerToExternalAccount(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.AssignCustomerToExternalAccount(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.UnassignCustomerFromExternalAccount(data.ParentId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.AddTransactionsToExternalAccount(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExternalAccount DAO
	//----------------------------------------------------------------------------
	requestResult := ExternalAccountDAO.RemoveTransactionsFromExternalAccount(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
