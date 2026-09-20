
package controller

import (
    AccountDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.GetAccount(data.Id)
	
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
	utils.ParseBody(r, &data)

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
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccountDAO.DeleteAccount(data.Id)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignBankToAccount(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignBankFromAccount(data.ParentId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignBranchToAccount(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignBranchFromAccount(data.ParentId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignProductToAccount(data.ParentId, data.ChildId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignProductFromAccount(data.ParentId)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddOwnersToAccount(data.ParentId, data.childIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveOwnersFromAccount(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddTransactionsToAccount(data.ParentId, data.childIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveTransactionsFromAccount(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddStatementsToAccount(data.ParentId, data.childIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveStatementsFromAccount(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddStandingInstructionsToAccount(data.ParentId, data.childIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveStandingInstructionsFromAccount(data.ParentId, data.ChildIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddFeeChargesToAccount(data.ParentId, data.childIds)

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

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveFeeChargesFromAccount(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
