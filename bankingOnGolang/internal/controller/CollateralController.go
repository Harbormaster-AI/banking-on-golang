
package controller

import (
    CollateralDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to CollateralDAO for database creation
// ----------------------------------------------------------------------------
func CreateCollateral(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Collateral model
	// ----------------------------------------------------------------------------
	data := model.Collateral{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Collateral model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral data access object to create
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.CreateCollateral( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Get controller, delegates to CollateralDAO to find the relevant Collateral
// ----------------------------------------------------------------------------
func GetCollateral(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.GetCollateral(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to CollateralDAO for database read of all Collaterals
// ----------------------------------------------------------------------------
func GetAllCollateral(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the Collateral data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.GetAllCollateral()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Update controller, delegates to CollateralDAO for database save
// ----------------------------------------------------------------------------
func UpdateCollateral(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Collateral model
	// ----------------------------------------------------------------------------
	var data = model.Collateral{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Collateral model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.UpdateCollateral(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to CollateralDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteCollateral(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := CollateralDAO.DeleteCollateral(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// assigns a LoanAccount on a Collateral
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignLoanAccountToCollateral(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral DAO
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.AssignLoanAccountToCollateral(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a LoanAccount on a Collateral
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignLoanAccountFromCollateral( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Collateral DAO
	// ----------------------------------------------------------------------------
	requestResult := CollateralDAO.UnassignLoanAccountFromCollateral(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


