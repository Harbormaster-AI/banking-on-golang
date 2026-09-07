package controller

import (
    RiskAssessmentDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RiskAssessmentDAO for database creation
//----------------------------------------------------------------------------
func CreateRiskAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RiskAssessment model
	//----------------------------------------------------------------------------
	data := model.RiskAssessment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RiskAssessment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment data access object to create
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.CreateRiskAssessment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RiskAssessmentDAO to find the relevant RiskAssessment
//----------------------------------------------------------------------------
func GetRiskAssessment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RiskAssessment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.GetRiskAssessment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RiskAssessmentDAO for database read of all RiskAssessments
//----------------------------------------------------------------------------
func GetAllRiskAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.GetAllRiskAssessment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RiskAssessmentDAO for database save
//----------------------------------------------------------------------------
func UpdateRiskAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RiskAssessment model
	//----------------------------------------------------------------------------
	var data = model.RiskAssessment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RiskAssessment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.UpdateRiskAssessment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RiskAssessmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRiskAssessment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RiskAssessment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RiskAssessmentDAO.DeleteRiskAssessment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a KycProfile on a RiskAssessment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignKycProfileToRiskAssessment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	riskAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfileId,_ := strconv.ParseUint( vars["kycProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.AssignKycProfileToRiskAssessment(riskAssessmentId, kycProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a KycProfile on a RiskAssessment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignKycProfileFromRiskAssessment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	riskAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.UnassignKycProfileFromRiskAssessment(riskAssessmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


