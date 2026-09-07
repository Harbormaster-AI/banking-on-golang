package controller

import (
    AccountDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AccountDAO for database creation
//----------------------------------------------------------------------------
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Account model
	//----------------------------------------------------------------------------
	data := model.Account{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Account model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object to create
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.CreateAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AccountDAO to find the relevant Account
//----------------------------------------------------------------------------
func GetAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Account data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.GetAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AccountDAO for database read of all Accounts
//----------------------------------------------------------------------------
func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Account data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.GetAllAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AccountDAO for database save
//----------------------------------------------------------------------------
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Account model
	//----------------------------------------------------------------------------
	var data = model.Account{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Account model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UpdateAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Account data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccountDAO.DeleteAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignBankToAccount(accountId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignBankFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Branch on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBranchToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchId,_ := strconv.ParseUint( vars["branchId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignBranchToAccount(accountId, branchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Branch on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBranchFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignBranchFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignProductToAccount(accountId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignProductFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more ownersIds as a Owners to a Account
	//----------------------------------------------------------------------------
func AddOwnersToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownersIds,_ := vars["ownersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddOwnersToAccount(accountId, ownersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownersIds as a Owners from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnersFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownersIds,_ := vars["ownersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveOwnersFromAccount(accountId, ownersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a Account
	//----------------------------------------------------------------------------
func AddTransactionsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddTransactionsToAccount(accountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveTransactionsFromAccount(accountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more statementsIds as a Statements to a Account
	//----------------------------------------------------------------------------
func AddStatementsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	statementsIds,_ := vars["statementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddStatementsToAccount(accountId, statementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more statementsIds as a Statements from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStatementsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	statementsIds,_ := vars["statementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveStatementsFromAccount(accountId, statementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more standingInstructionsIds as a StandingInstructions to a Account
	//----------------------------------------------------------------------------
func AddStandingInstructionsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	standingInstructionsIds,_ := vars["standingInstructionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddStandingInstructionsToAccount(accountId, standingInstructionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more standingInstructionsIds as a StandingInstructions from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStandingInstructionsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	standingInstructionsIds,_ := vars["standingInstructionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveStandingInstructionsFromAccount(accountId, standingInstructionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more feeChargesIds as a FeeCharges to a Account
	//----------------------------------------------------------------------------
func AddFeeChargesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeChargesIds,_ := vars["feeChargesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddFeeChargesToAccount(accountId, feeChargesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more feeChargesIds as a FeeCharges from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeeChargesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeChargesIds,_ := vars["feeChargesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveFeeChargesFromAccount(accountId, feeChargesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
