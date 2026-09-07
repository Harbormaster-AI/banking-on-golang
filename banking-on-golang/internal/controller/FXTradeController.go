package controller

import (
    FXTradeDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FXTradeDAO for database creation
//----------------------------------------------------------------------------
func CreateFXTrade(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXTrade model
	//----------------------------------------------------------------------------
	data := model.FXTrade{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXTrade model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object to create
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.CreateFXTrade( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FXTradeDAO to find the relevant FXTrade
//----------------------------------------------------------------------------
func GetFXTrade(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXTrade data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.GetFXTrade(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FXTradeDAO for database read of all FXTrades
//----------------------------------------------------------------------------
func GetAllFXTrade(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.GetAllFXTrade()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FXTradeDAO for database save
//----------------------------------------------------------------------------
func UpdateFXTrade(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXTrade model
	//----------------------------------------------------------------------------
	var data = model.FXTrade{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXTrade model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UpdateFXTrade(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FXTradeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFXTrade(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXTrade data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FXTradeDAO.DeleteFXTrade(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignCustomerToFXTrade(fXTradeId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignCustomerFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Bank on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignBankToFXTrade(fXTradeId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignBankFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ExchangeRate on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExchangeRateToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exchangeRateId,_ := strconv.ParseUint( vars["exchangeRateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignExchangeRateToFXTrade(fXTradeId, exchangeRateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ExchangeRate on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExchangeRateFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignExchangeRateFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SourceAccount on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSourceAccountToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceAccountId,_ := strconv.ParseUint( vars["sourceAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignSourceAccountToFXTrade(fXTradeId, sourceAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SourceAccount on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSourceAccountFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignSourceAccountFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationAccount on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDestinationAccountToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationAccountId,_ := strconv.ParseUint( vars["destinationAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignDestinationAccountToFXTrade(fXTradeId, destinationAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationAccount on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDestinationAccountFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignDestinationAccountFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a FXTrade
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToFXTrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignTransactionToFXTrade(fXTradeId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a FXTrade
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXTradeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	//----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignTransactionFromFXTrade(fXTradeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


