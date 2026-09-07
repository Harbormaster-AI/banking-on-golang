package controller

import (
    CollateralDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CollateralDAO for database creation
//----------------------------------------------------------------------------
func CreateCollateral(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Collateral model
	//----------------------------------------------------------------------------
	data := model.Collateral{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Collateral model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Collateral data access object to create
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.CreateCollateral( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CollateralDAO to find the relevant Collateral
//----------------------------------------------------------------------------
func GetCollateral(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Collateral data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.GetCollateral(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CollateralDAO for database read of all Collaterals
//----------------------------------------------------------------------------
func GetAllCollateral(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Collateral data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.GetAllCollateral()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CollateralDAO for database save
//----------------------------------------------------------------------------
func UpdateCollateral(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Collateral model
	//----------------------------------------------------------------------------
	var data = model.Collateral{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Collateral model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Collateral data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.UpdateCollateral(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CollateralDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCollateral(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Collateral data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CollateralDAO.DeleteCollateral(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a LoanAccount on a Collateral
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanAccountToCollateral(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	collateralId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountId,_ := strconv.ParseUint( vars["loanAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Collateral DAO
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.AssignLoanAccountToCollateral(collateralId, loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LoanAccount on a Collateral
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanAccountFromCollateral( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	collateralId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Collateral DAO
	//----------------------------------------------------------------------------
	requestResult := CollateralDAO.UnassignLoanAccountFromCollateral(collateralId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


