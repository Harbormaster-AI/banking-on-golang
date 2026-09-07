package controller

import (
    ThirdPartyProviderDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ThirdPartyProviderDAO for database creation
//----------------------------------------------------------------------------
func CreateThirdPartyProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdPartyProvider model
	//----------------------------------------------------------------------------
	data := model.ThirdPartyProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdPartyProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider data access object to create
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.CreateThirdPartyProvider( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ThirdPartyProviderDAO to find the relevant ThirdPartyProvider
//----------------------------------------------------------------------------
func GetThirdPartyProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdPartyProvider data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.GetThirdPartyProvider(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ThirdPartyProviderDAO for database read of all ThirdPartyProviders
//----------------------------------------------------------------------------
func GetAllThirdPartyProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.GetAllThirdPartyProvider()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ThirdPartyProviderDAO for database save
//----------------------------------------------------------------------------
func UpdateThirdPartyProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdPartyProvider model
	//----------------------------------------------------------------------------
	var data = model.ThirdPartyProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdPartyProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.UpdateThirdPartyProvider(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ThirdPartyProviderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteThirdPartyProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdPartyProvider data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ThirdPartyProviderDAO.DeleteThirdPartyProvider(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a ThirdPartyProvider
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToThirdPartyProvider(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.AssignBankToThirdPartyProvider(thirdPartyProviderId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a ThirdPartyProvider
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromThirdPartyProvider( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.UnassignBankFromThirdPartyProvider(thirdPartyProviderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a ThirdPartyProvider
	//----------------------------------------------------------------------------
func AddConsentsToThirdPartyProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.AddConsentsToThirdPartyProvider(thirdPartyProviderId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a ThirdPartyProvider
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromThirdPartyProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.RemoveConsentsFromThirdPartyProvider(thirdPartyProviderId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
