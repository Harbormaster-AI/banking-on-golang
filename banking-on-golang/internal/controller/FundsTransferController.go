package controller

import (
    FundsTransferDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FundsTransferDAO for database creation
//----------------------------------------------------------------------------
func CreateFundsTransfer(w http.ResponseWriter, r *http.Request) {
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
func GetFundsTransfer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FundsTransfer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.GetFundsTransfer(ID)
	
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
func GetAllFundsTransfer(w http.ResponseWriter, r *http.Request) {
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
func UpdateFundsTransfer(w http.ResponseWriter, r *http.Request) {
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
func DeleteFundsTransfer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FundsTransfer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FundsTransferDAO.DeleteFundsTransfer(ID)

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
func AssignSourceAccountToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceAccountId,_ := strconv.ParseUint( vars["sourceAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignSourceAccountToFundsTransfer(fundsTransferId, sourceAccountId)

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
func UnassignSourceAccountFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignSourceAccountFromFundsTransfer(fundsTransferId)

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
func AssignDestinationAccountToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationAccountId,_ := strconv.ParseUint( vars["destinationAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignDestinationAccountToFundsTransfer(fundsTransferId, destinationAccountId)

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
func UnassignDestinationAccountFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignDestinationAccountFromFundsTransfer(fundsTransferId)

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
func AssignExternalBeneficiaryToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	externalBeneficiaryId,_ := strconv.ParseUint( vars["externalBeneficiaryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignExternalBeneficiaryToFundsTransfer(fundsTransferId, externalBeneficiaryId)

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
func UnassignExternalBeneficiaryFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignExternalBeneficiaryFromFundsTransfer(fundsTransferId)

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
func AssignInitiatedByToFundsTransfer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	initiatedById,_ := strconv.ParseUint( vars["initiatedById"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AssignInitiatedByToFundsTransfer(fundsTransferId, initiatedById)

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
func UnassignInitiatedByFromFundsTransfer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.UnassignInitiatedByFromFundsTransfer(fundsTransferId)

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
func AddTransactionsToFundsTransfer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.AddTransactionsToFundsTransfer(fundsTransferId, transactionsIds)

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
func RemoveTransactionsFromFundsTransfer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fundsTransferId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FundsTransfer DAO
	//----------------------------------------------------------------------------
	requestResult := FundsTransferDAO.RemoveTransactionsFromFundsTransfer(fundsTransferId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
