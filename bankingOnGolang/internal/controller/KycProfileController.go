
package controller

import (
    KycProfileDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to KycProfileDAO for database creation
//----------------------------------------------------------------------------
func create(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KycProfile model
	//----------------------------------------------------------------------------
	data := model.KycProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KycProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.CreateKycProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to KycProfileDAO to find the relevant KycProfile
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
	// Delegate to the KycProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.GetKycProfile(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to KycProfileDAO for database read of all KycProfiles
//----------------------------------------------------------------------------
func getAll(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the KycProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.GetAllKycProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to KycProfileDAO for database save
//----------------------------------------------------------------------------
func update(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KycProfile model
	//----------------------------------------------------------------------------
	var data = model.KycProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KycProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.UpdateKycProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to KycProfileDAO for database deletion
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
	// Delegate to the KycProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := KycProfileDAO.DeleteKycProfile(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a KycProfile
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
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AssignCustomerToKycProfile(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a KycProfile
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
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.UnassignCustomerFromKycProfile(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more identityDocumentsIds as a IdentityDocuments to a KycProfile
	//----------------------------------------------------------------------------
func addToIdentityDocuments(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddIdentityDocumentsToKycProfile(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more identityDocumentsIds as a IdentityDocuments from a KycProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func removeFromIdentityDocuments(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveIdentityDocumentsFromKycProfile(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more riskAssessmentsIds as a RiskAssessments to a KycProfile
	//----------------------------------------------------------------------------
func addToRiskAssessments(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddRiskAssessmentsToKycProfile(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more riskAssessmentsIds as a RiskAssessments from a KycProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func removeFromRiskAssessments(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveRiskAssessmentsFromKycProfile(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more screeningsIds as a Screenings to a KycProfile
	//----------------------------------------------------------------------------
func addToScreenings(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddScreeningsToKycProfile(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more screeningsIds as a Screenings from a KycProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func removeFromScreenings(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveScreeningsFromKycProfile(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
