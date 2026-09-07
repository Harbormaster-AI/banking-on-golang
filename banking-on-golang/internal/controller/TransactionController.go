package controller

import (
    TransactionDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TransactionDAO for database creation
//----------------------------------------------------------------------------
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	//----------------------------------------------------------------------------
	data := model.Transaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to create
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.CreateTransaction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TransactionDAO to find the relevant Transaction
//----------------------------------------------------------------------------
func GetTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Transaction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetTransaction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TransactionDAO for database read of all Transactions
//----------------------------------------------------------------------------
func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetAllTransaction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TransactionDAO for database save
//----------------------------------------------------------------------------
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	//----------------------------------------------------------------------------
	var data = model.Transaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UpdateTransaction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TransactionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Transaction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TransactionDAO.DeleteTransaction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignAccountToTransaction(transactionId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignAccountFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ExternalCounterparty on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExternalCounterpartyToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	externalCounterpartyId,_ := strconv.ParseUint( vars["externalCounterpartyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignExternalCounterpartyToTransaction(transactionId, externalCounterpartyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ExternalCounterparty on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExternalCounterpartyFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignExternalCounterpartyFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PaymentCard on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPaymentCardToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardId,_ := strconv.ParseUint( vars["paymentCardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignPaymentCardToTransaction(transactionId, paymentCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PaymentCard on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPaymentCardFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignPaymentCardFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a FundsTransfer on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFundsTransferToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fundsTransferId,_ := strconv.ParseUint( vars["fundsTransferId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignFundsTransferToTransaction(transactionId, fundsTransferId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a FundsTransfer on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFundsTransferFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignFundsTransferFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a FxTrade on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFxTradeToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fxTradeId,_ := strconv.ParseUint( vars["fxTradeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignFxTradeToTransaction(transactionId, fxTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a FxTrade on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFxTradeFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignFxTradeFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dispute on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDisputeToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputeId,_ := strconv.ParseUint( vars["disputeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignDisputeToTransaction(transactionId, disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dispute on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDisputeFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignDisputeFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


