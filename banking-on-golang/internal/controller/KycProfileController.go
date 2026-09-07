package controller

import (
    KycProfileDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to KycProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateKycProfile(w http.ResponseWriter, r *http.Request) {
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
func GetKycProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KycProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.GetKycProfile(ID)
	
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
func GetAllKycProfile(w http.ResponseWriter, r *http.Request) {
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
func UpdateKycProfile(w http.ResponseWriter, r *http.Request) {
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
func DeleteKycProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KycProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := KycProfileDAO.DeleteKycProfile(ID)

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
func AssignCustomerToKycProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AssignCustomerToKycProfile(kycProfileId, customerId)

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
func UnassignCustomerFromKycProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.UnassignCustomerFromKycProfile(kycProfileId)

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
func AddIdentityDocumentsToKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	identityDocumentsIds,_ := vars["identityDocumentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddIdentityDocumentsToKycProfile(kycProfileId, identityDocumentsIds)

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
func RemoveIdentityDocumentsFromKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	identityDocumentsIds,_ := vars["identityDocumentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveIdentityDocumentsFromKycProfile(kycProfileId, identityDocumentsIds)

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
func AddRiskAssessmentsToKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	riskAssessmentsIds,_ := vars["riskAssessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddRiskAssessmentsToKycProfile(kycProfileId, riskAssessmentsIds)

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
func RemoveRiskAssessmentsFromKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	riskAssessmentsIds,_ := vars["riskAssessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveRiskAssessmentsFromKycProfile(kycProfileId, riskAssessmentsIds)

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
func AddScreeningsToKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.AddScreeningsToKycProfile(kycProfileId, screeningsIds)

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
func RemoveScreeningsFromKycProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kycProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KycProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KycProfileDAO.RemoveScreeningsFromKycProfile(kycProfileId, screeningsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
