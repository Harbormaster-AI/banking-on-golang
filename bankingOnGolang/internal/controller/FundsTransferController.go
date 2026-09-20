
package controller

import (
    FundsTransferDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FundsTransferDAO for database creation
//----------------------------------------------------------------------------
func create(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FundsTransfer model
	//----------------------------------------------------------------------------
	data := model.FundsTransfer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FundsTransfer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object to create
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.CreateFundsTransfer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FundsTransferDAO to find the relevant FundsTransfer
//----------------------------------------------------------------------------
func get(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.GetFundsTransfer(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FundsTransferDAO for database read of all FundsTransfers
//----------------------------------------------------------------------------
func getAll(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.GetAllFundsTransfer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FundsTransferDAO for database save
//----------------------------------------------------------------------------
func update(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FundsTransfer model
	//----------------------------------------------------------------------------
	var data = model.FundsTransfer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FundsTransfer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UpdateFundsTransfer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FundsTransferDAO for database deletion
//----------------------------------------------------------------------------
func delete(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FundsTransferDAO.DeleteFundsTransfer(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SourceAccount on a FundsTransfer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignSourceAccount(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignSourceAccountToFundsTransfer(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SourceAccount on a FundsTransfer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignSourceAccount( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignSourceAccountFromFundsTransfer(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationAccount on a FundsTransfer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignDestinationAccount(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignDestinationAccountToFundsTransfer(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationAccount on a FundsTransfer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignDestinationAccount( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignDestinationAccountFromFundsTransfer(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ExternalBeneficiary on a FundsTransfer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignExternalBeneficiary(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignExternalBeneficiaryToFundsTransfer(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ExternalBeneficiary on a FundsTransfer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignExternalBeneficiary( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignExternalBeneficiaryFromFundsTransfer(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InitiatedBy on a FundsTransfer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func assignInitiatedBy(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignInitiatedByToFundsTransfer(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InitiatedBy on a FundsTransfer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func unassignInitiatedBy( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignInitiatedByFromFundsTransfer(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a FundsTransfer
	//----------------------------------------------------------------------------
func addToTransactions(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AddTransactionsToFundsTransfer(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a FundsTransfer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func removeFromTransactions(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.RemoveTransactionsFromFundsTransfer(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
