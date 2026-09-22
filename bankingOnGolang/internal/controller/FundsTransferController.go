
package controller

import (
    FundsTransferDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to FundsTransferDAO for database creation
// ----------------------------------------------------------------------------
func CreateFundsTransfer(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FundsTransfer model
	// ----------------------------------------------------------------------------
	data := model.FundsTransfer{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a FundsTransfer model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object to create
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.CreateFundsTransfer( data )
	
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
// Get controller, delegates to FundsTransferDAO to find the relevant FundsTransfer
// ----------------------------------------------------------------------------
func GetFundsTransfer(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.GetFundsTransfer(data.Id)
	
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
// GetAll controller, delegates to FundsTransferDAO for database read of all FundsTransfers
// ----------------------------------------------------------------------------
func GetAllFundsTransfer(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.GetAllFundsTransfer()
	
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
// Update controller, delegates to FundsTransferDAO for database save
// ----------------------------------------------------------------------------
func UpdateFundsTransfer(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FundsTransfer model
	// ----------------------------------------------------------------------------
	var data = model.FundsTransfer{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a FundsTransfer model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UpdateFundsTransfer(data)

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
// Delete controller, delegates to FundsTransferDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteFundsTransfer(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := FundsTransferDAO.DeleteFundsTransfer(data.Id)

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
	// assigns a SourceAccount on a FundsTransfer
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignSourceAccountToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignSourceAccountToFundsTransfer(data.ParentId, data.ChildId)

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
	// unassigns a SourceAccount on a FundsTransfer
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignSourceAccountFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignSourceAccountFromFundsTransfer(data.ParentId)

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
	// assigns a DestinationAccount on a FundsTransfer
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignDestinationAccountToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignDestinationAccountToFundsTransfer(data.ParentId, data.ChildId)

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
	// unassigns a DestinationAccount on a FundsTransfer
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignDestinationAccountFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignDestinationAccountFromFundsTransfer(data.ParentId)

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
	// assigns a ExternalBeneficiary on a FundsTransfer
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignExternalBeneficiaryToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignExternalBeneficiaryToFundsTransfer(data.ParentId, data.ChildId)

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
	// unassigns a ExternalBeneficiary on a FundsTransfer
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignExternalBeneficiaryFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignExternalBeneficiaryFromFundsTransfer(data.ParentId)

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
	// assigns a InitiatedBy on a FundsTransfer
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignInitiatedByToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignInitiatedByToFundsTransfer(data.ParentId, data.ChildId)

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
	// unassigns a InitiatedBy on a FundsTransfer
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignInitiatedByFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignInitiatedByFromFundsTransfer(data.ParentId)

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
	// adds one or more transactionsIds as a Transactions to a FundsTransfer
	// ----------------------------------------------------------------------------
func AddTransactionsToFundsTransfer(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AddTransactionsToFundsTransfer(data.ParentId, data.ChildIds)

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
	// removes one or more transactionsIds as a Transactions from a FundsTransfer
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveTransactionsFromFundsTransfer(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	// ----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.RemoveTransactionsFromFundsTransfer(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}
		
