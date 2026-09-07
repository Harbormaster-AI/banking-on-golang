package controller

import (
    DisputeDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DisputeDAO for database creation
//----------------------------------------------------------------------------
func CreateDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	//----------------------------------------------------------------------------
	data := model.Dispute{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to create
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.CreateDispute( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DisputeDAO to find the relevant Dispute
//----------------------------------------------------------------------------
func GetDispute(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dispute data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetDispute(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DisputeDAO for database read of all Disputes
//----------------------------------------------------------------------------
func GetAllDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetAllDispute()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DisputeDAO for database save
//----------------------------------------------------------------------------
func UpdateDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	//----------------------------------------------------------------------------
	var data = model.Dispute{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UpdateDispute(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DisputeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDispute(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dispute data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DisputeDAO.DeleteDispute(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignTransactionToDispute(disputeId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignTransactionFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignCustomerToDispute(disputeId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignCustomerFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignAccountToDispute(disputeId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignAccountFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PaymentCard on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPaymentCardToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardId,_ := strconv.ParseUint( vars["paymentCardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignPaymentCardToDispute(disputeId, paymentCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PaymentCard on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPaymentCardFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignPaymentCardFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


