package controller

import (
    ConsentDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConsentDAO for database creation
//----------------------------------------------------------------------------
func CreateConsent(w http.ResponseWriter, r *http.Request) {
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
func GetConsent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Consent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.GetConsent(ID)
	
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
func GetAllConsent(w http.ResponseWriter, r *http.Request) {
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
func UpdateConsent(w http.ResponseWriter, r *http.Request) {
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
func DeleteConsent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Consent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConsentDAO.DeleteConsent(ID)

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
func AssignCustomerToConsent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignCustomerToConsent(consentId, customerId)

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
func UnassignCustomerFromConsent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignCustomerFromConsent(consentId)

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
func AssignBankToConsent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignBankToConsent(consentId, bankId)

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
func UnassignBankFromConsent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignBankFromConsent(consentId)

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
func AssignThirdPartyProviderToConsent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartyProviderId,_ := strconv.ParseUint( vars["thirdPartyProviderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignThirdPartyProviderToConsent(consentId, thirdPartyProviderId)

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
func UnassignThirdPartyProviderFromConsent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignThirdPartyProviderFromConsent(consentId)

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
func AddAuthorizedAccountsToConsent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizedAccountsIds,_ := vars["authorizedAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AddAuthorizedAccountsToConsent(consentId, authorizedAccountsIds)

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
func RemoveAuthorizedAccountsFromConsent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizedAccountsIds,_ := vars["authorizedAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.RemoveAuthorizedAccountsFromConsent(consentId, authorizedAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
