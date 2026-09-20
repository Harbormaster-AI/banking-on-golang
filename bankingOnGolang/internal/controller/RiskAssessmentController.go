
package controller

import (
    RiskAssessmentDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	 "encoding/json"
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.GetRiskAssessment(data.Id)
	
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RiskAssessmentDAO.DeleteRiskAssessment(data.Id)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.AssignKycProfileToRiskAssessment(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the RiskAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := RiskAssessmentDAO.UnassignKycProfileFromRiskAssessment(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


