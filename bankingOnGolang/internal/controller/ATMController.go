
package controller

import (
    ATMDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ATMDAO for database creation
//----------------------------------------------------------------------------
func CreateATM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ATM model
	//----------------------------------------------------------------------------
	data := model.ATM{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ATM model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM data access object to create
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.CreateATM( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ATMDAO to find the relevant ATM
//----------------------------------------------------------------------------
func GetATM(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.GetATM(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ATMDAO for database read of all ATMs
//----------------------------------------------------------------------------
func GetAllATM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ATM data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.GetAllATM()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ATMDAO for database save
//----------------------------------------------------------------------------
func UpdateATM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ATM model
	//----------------------------------------------------------------------------
	var data = model.ATM{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ATM model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.UpdateATM(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ATMDAO for database deletion
//----------------------------------------------------------------------------
func DeleteATM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ATMDAO.DeleteATM(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Branch on a ATM
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBranchToATM(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM DAO
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.AssignBranchToATM(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Branch on a ATM
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBranchFromATM( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ATM DAO
	//----------------------------------------------------------------------------
	requestResult := ATMDAO.UnassignBranchFromATM(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


