package controller

import (
	FeeChargeDAO "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to FeeChargeDAO for database creation
// ----------------------------------------------------------------------------
func CreateFeeCharge(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FeeCharge model
	// ----------------------------------------------------------------------------
	data := model.FeeCharge{}

	// ----------------------------------------------------------------------------
	// Parse the body into a FeeCharge model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object to create
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.CreateFeeCharge(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Get controller, delegates to FeeChargeDAO to find the relevant FeeCharge
// ----------------------------------------------------------------------------
func GetFeeCharge(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.GetFeeCharge(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// GetAll controller, delegates to FeeChargeDAO for database read of all FeeCharges
// ----------------------------------------------------------------------------
func GetAllFeeCharge(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.GetAllFeeCharge()

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Update controller, delegates to FeeChargeDAO for database save
// ----------------------------------------------------------------------------
func UpdateFeeCharge(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FeeCharge model
	// ----------------------------------------------------------------------------
	var data = model.FeeCharge{}

	// ----------------------------------------------------------------------------
	// Parse the body into a FeeCharge model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UpdateFeeCharge(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to FeeChargeDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteFeeCharge(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.DeleteFeeCharge(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// assigns a Account on a FeeCharge
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignAccountToFeeCharge(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.AssignAccountToFeeCharge(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// unassigns a Account on a FeeCharge
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignAccountFromFeeCharge(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UnassignAccountFromFeeCharge(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// assigns a LoanAccount on a FeeCharge
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignLoanAccountToFeeCharge(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.AssignLoanAccountToFeeCharge(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// unassigns a LoanAccount on a FeeCharge
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignLoanAccountFromFeeCharge(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FeeCharge DAO
	// ----------------------------------------------------------------------------
	requestResult := FeeChargeDAO.UnassignLoanAccountFromFeeCharge(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
