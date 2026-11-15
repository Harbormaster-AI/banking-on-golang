package controller

import (
	DisputeDAO "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to DisputeDAO for database creation
// ----------------------------------------------------------------------------
func CreateDispute(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	// ----------------------------------------------------------------------------
	data := model.Dispute{}

	// ----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to create
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.CreateDispute(data)

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
// Get controller, delegates to DisputeDAO to find the relevant Dispute
// ----------------------------------------------------------------------------
func GetDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetDispute(data.Id)

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
// GetAll controller, delegates to DisputeDAO for database read of all Disputes
// ----------------------------------------------------------------------------
func GetAllDispute(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetAllDispute()

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
// Update controller, delegates to DisputeDAO for database save
// ----------------------------------------------------------------------------
func UpdateDispute(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	// ----------------------------------------------------------------------------
	var data = model.Dispute{}

	// ----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.UpdateDispute(data)

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
// Delete controller, delegates to DisputeDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteDispute(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.DeleteDispute(data.Id)

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
// assigns a Transaction on a Dispute
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignTransactionToDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignTransactionToDispute(data.ParentId, data.ChildId)

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
// unassigns a Transaction on a Dispute
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignTransactionFromDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignTransactionFromDispute(data.ParentId)

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
// assigns a Customer on a Dispute
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignCustomerToDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignCustomerToDispute(data.ParentId, data.ChildId)

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
// unassigns a Customer on a Dispute
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignCustomerFromDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignCustomerFromDispute(data.ParentId)

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
// assigns a Account on a Dispute
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignAccountToDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignAccountToDispute(data.ParentId, data.ChildId)

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
// unassigns a Account on a Dispute
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignAccountFromDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignAccountFromDispute(data.ParentId)

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
// assigns a PaymentCard on a Dispute
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignPaymentCardToDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignPaymentCardToDispute(data.ParentId, data.ChildId)

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
// unassigns a PaymentCard on a Dispute
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignPaymentCardFromDispute(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	// ----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignPaymentCardFromDispute(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
