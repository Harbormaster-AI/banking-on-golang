
package controller

import (
    ThirdPartyProviderDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.GetThirdPartyProvider(data.Id)
	
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ThirdPartyProviderDAO.DeleteThirdPartyProvider(data.Id)

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
	//----------------------------------------------------------------------------
func AssignBankToThirdPartyProvider(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.AssignBankToThirdPartyProvider(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.UnassignBankFromThirdPartyProvider(data.ParentId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.AddConsentsToThirdPartyProvider(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyProviderDAO.RemoveConsentsFromThirdPartyProvider(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
