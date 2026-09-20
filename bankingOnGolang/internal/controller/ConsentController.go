
package controller

import (
    ConsentDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConsentDAO for database creation
//----------------------------------------------------------------------------
func create(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Consent model
	//----------------------------------------------------------------------------
	data := model.Consent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Consent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object to create
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.CreateConsent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ConsentDAO to find the relevant Consent
//----------------------------------------------------------------------------
func get(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.GetConsent(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ConsentDAO for database read of all Consents
//----------------------------------------------------------------------------
func getAll(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.GetAllConsent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ConsentDAO for database save
//----------------------------------------------------------------------------
func update(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Consent model
	//----------------------------------------------------------------------------
	var data = model.Consent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Consent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UpdateConsent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ConsentDAO for database deletion
//----------------------------------------------------------------------------
func delete(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConsentDAO.DeleteConsent(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Consent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignCustomer(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignCustomerToConsent(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Consent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignCustomer( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignCustomerFromConsent(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Bank on a Consent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignBank(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignBankToConsent(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a Consent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignBank( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignBankFromConsent(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ThirdPartyProvider on a Consent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignThirdPartyProvider(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignThirdPartyProviderToConsent(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ThirdPartyProvider on a Consent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignThirdPartyProvider( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignThirdPartyProviderFromConsent(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more authorizedAccountsIds as a AuthorizedAccounts to a Consent
	//----------------------------------------------------------------------------
func addToAuthorizedAccounts(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AddAuthorizedAccountsToConsent(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more authorizedAccountsIds as a AuthorizedAccounts from a Consent
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func removeFromAuthorizedAccounts(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.RemoveAuthorizedAccountsFromConsent(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
