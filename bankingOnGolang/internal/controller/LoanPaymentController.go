package controller

import (
	LoanPaymentDAO "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to LoanPaymentDAO for database creation
// ----------------------------------------------------------------------------
func CreateLoanPayment(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty LoanPayment model
	// ----------------------------------------------------------------------------
	data := model.LoanPayment{}

	// ----------------------------------------------------------------------------
	// Parse the body into a LoanPayment model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object to create
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.CreateLoanPayment(data)

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
// Get controller, delegates to LoanPaymentDAO to find the relevant LoanPayment
// ----------------------------------------------------------------------------
func GetLoanPayment(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.GetLoanPayment(data.Id)

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
// GetAll controller, delegates to LoanPaymentDAO for database read of all LoanPayments
// ----------------------------------------------------------------------------
func GetAllLoanPayment(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.GetAllLoanPayment()

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
// Update controller, delegates to LoanPaymentDAO for database save
// ----------------------------------------------------------------------------
func UpdateLoanPayment(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty LoanPayment model
	// ----------------------------------------------------------------------------
	var data = model.LoanPayment{}

	// ----------------------------------------------------------------------------
	// Parse the body into a LoanPayment model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UpdateLoanPayment(data)

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
// Delete controller, delegates to LoanPaymentDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteLoanPayment(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.DeleteLoanPayment(data.Id)

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
// assigns a LoanAccount on a LoanPayment
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignLoanAccountToLoanPayment(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.AssignLoanAccountToLoanPayment(data.ParentId, data.ChildId)

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
// unassigns a LoanAccount on a LoanPayment
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignLoanAccountFromLoanPayment(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UnassignLoanAccountFromLoanPayment(data.ParentId)

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
// assigns a Transaction on a LoanPayment
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignTransactionToLoanPayment(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.AssignTransactionToLoanPayment(data.ParentId, data.ChildId)

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
// unassigns a Transaction on a LoanPayment
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignTransactionFromLoanPayment(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UnassignTransactionFromLoanPayment(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
