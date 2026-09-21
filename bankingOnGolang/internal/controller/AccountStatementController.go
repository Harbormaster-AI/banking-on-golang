
package controller

import (
    AccountStatementDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AccountStatementDAO for database creation
//----------------------------------------------------------------------------
func CreateAccountStatement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccountStatement model
	//----------------------------------------------------------------------------
	data := model.AccountStatement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccountStatement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement data access object to create
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.CreateAccountStatement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AccountStatementDAO to find the relevant AccountStatement
//----------------------------------------------------------------------------
func GetAccountStatement(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.GetAccountStatement(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AccountStatementDAO for database read of all AccountStatements
//----------------------------------------------------------------------------
func GetAllAccountStatement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.GetAllAccountStatement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AccountStatementDAO for database save
//----------------------------------------------------------------------------
func UpdateAccountStatement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccountStatement model
	//----------------------------------------------------------------------------
	var data = model.AccountStatement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccountStatement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.UpdateAccountStatement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AccountStatementDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAccountStatement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccountStatementDAO.DeleteAccountStatement(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a AccountStatement
	// delegates to an ORM handler
	//----------------------------------------------------------------------------
func AssignAccountToAccountStatement(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement DAO
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.AssignAccountToAccountStatement(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a AccountStatement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromAccountStatement( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the AccountStatement DAO
	//----------------------------------------------------------------------------
	requestResult := AccountStatementDAO.UnassignAccountFromAccountStatement(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


