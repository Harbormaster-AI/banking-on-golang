
package controller

import (
    TransactionDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to TransactionDAO for database creation
// ----------------------------------------------------------------------------
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	// ----------------------------------------------------------------------------
	data := model.Transaction{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to create
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.CreateTransaction( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}

// ----------------------------------------------------------------------------
// Get controller, delegates to TransactionDAO to find the relevant Transaction
// ----------------------------------------------------------------------------
func GetTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetTransaction(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to TransactionDAO for database read of all Transactions
// ----------------------------------------------------------------------------
func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetAllTransaction()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}

// ----------------------------------------------------------------------------
// Update controller, delegates to TransactionDAO for database save
// ----------------------------------------------------------------------------
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	// ----------------------------------------------------------------------------
	var data = model.Transaction{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UpdateTransaction(data)

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
// Delete controller, delegates to TransactionDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := TransactionDAO.DeleteTransaction(data.Id)

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
	// assigns a Account on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignAccountToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignAccountToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a Account on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignAccountFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignAccountFromTransaction(data.ParentId)

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
	// assigns a ExternalCounterparty on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignExternalCounterpartyToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignExternalCounterpartyToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a ExternalCounterparty on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignExternalCounterpartyFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignExternalCounterpartyFromTransaction(data.ParentId)

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
	// assigns a PaymentCard on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignPaymentCardToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignPaymentCardToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a PaymentCard on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignPaymentCardFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignPaymentCardFromTransaction(data.ParentId)

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
	// assigns a FundsTransfer on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignFundsTransferToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignFundsTransferToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a FundsTransfer on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignFundsTransferFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignFundsTransferFromTransaction(data.ParentId)

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
	// assigns a FxTrade on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignFxTradeToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignFxTradeToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a FxTrade on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignFxTradeFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignFxTradeFromTransaction(data.ParentId)

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
	// assigns a Dispute on a Transaction
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignDisputeToTransaction(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignDisputeToTransaction(data.ParentId, data.ChildId)

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
	// unassigns a Dispute on a Transaction
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignDisputeFromTransaction( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	// ----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignDisputeFromTransaction(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}


