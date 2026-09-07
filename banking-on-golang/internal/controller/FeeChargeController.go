package controller

import (
    FeeChargeDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FeeChargeDAO for database creation
//----------------------------------------------------------------------------
func CreateFeeCharge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FeeCharge model
	//----------------------------------------------------------------------------
	data := model.FeeCharge{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FeeCharge model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object to create
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.CreateFeeCharge( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FeeChargeDAO to find the relevant FeeCharge
//----------------------------------------------------------------------------
func GetFeeCharge(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FeeCharge data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.GetFeeCharge(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FeeChargeDAO for database read of all FeeCharges
//----------------------------------------------------------------------------
func GetAllFeeCharge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.GetAllFeeCharge()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FeeChargeDAO for database save
//----------------------------------------------------------------------------
func UpdateFeeCharge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FeeCharge model
	//----------------------------------------------------------------------------
	var data = model.FeeCharge{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FeeCharge model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UpdateFeeCharge(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FeeChargeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFeeCharge(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FeeCharge data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FeeChargeDAO.DeleteFeeCharge(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a FeeCharge
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToFeeCharge(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	feeChargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.AssignAccountToFeeCharge(feeChargeId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a FeeCharge
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromFeeCharge( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	feeChargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UnassignAccountFromFeeCharge(feeChargeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LoanAccount on a FeeCharge
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanAccountToFeeCharge(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	feeChargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountId,_ := strconv.ParseUint( vars["loanAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.AssignLoanAccountToFeeCharge(feeChargeId, loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LoanAccount on a FeeCharge
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanAccountFromFeeCharge( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	feeChargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	//----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UnassignLoanAccountFromFeeCharge(feeChargeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


