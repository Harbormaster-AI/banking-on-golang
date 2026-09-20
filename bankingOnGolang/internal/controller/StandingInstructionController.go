
package controller

import (
    StandingInstructionDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	 "encoding/json"
)

//----------------------------------------------------------------------------
// Create controller, delegates to StandingInstructionDAO for database creation
//----------------------------------------------------------------------------
func CreateStandingInstruction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StandingInstruction model
	//----------------------------------------------------------------------------
	data := model.StandingInstruction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StandingInstruction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction data access object to create
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.CreateStandingInstruction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to StandingInstructionDAO to find the relevant StandingInstruction
//----------------------------------------------------------------------------
func GetStandingInstruction(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.GetStandingInstruction(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to StandingInstructionDAO for database read of all StandingInstructions
//----------------------------------------------------------------------------
func GetAllStandingInstruction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.GetAllStandingInstruction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to StandingInstructionDAO for database save
//----------------------------------------------------------------------------
func UpdateStandingInstruction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StandingInstruction model
	//----------------------------------------------------------------------------
	var data = model.StandingInstruction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StandingInstruction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.UpdateStandingInstruction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to StandingInstructionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteStandingInstruction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := StandingInstructionDAO.DeleteStandingInstruction(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a StandingInstruction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToStandingInstruction(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction DAO
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.AssignAccountToStandingInstruction(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a StandingInstruction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromStandingInstruction( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction DAO
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.UnassignAccountFromStandingInstruction(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Beneficiary on a StandingInstruction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBeneficiaryToStandingInstruction(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction DAO
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.AssignBeneficiaryToStandingInstruction(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Beneficiary on a StandingInstruction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBeneficiaryFromStandingInstruction( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the StandingInstruction DAO
	//----------------------------------------------------------------------------
	requestResult := StandingInstructionDAO.UnassignBeneficiaryFromStandingInstruction(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


