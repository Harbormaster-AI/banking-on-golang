
package controller

import (
    FXTradeDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to FXTradeDAO for database creation
// ----------------------------------------------------------------------------
func CreateFXTrade(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FXTrade model
	// ----------------------------------------------------------------------------
	data := model.FXTrade{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a FXTrade model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object to create
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.CreateFXTrade( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Get controller, delegates to FXTradeDAO to find the relevant FXTrade
// ----------------------------------------------------------------------------
func GetFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.GetFXTrade(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to FXTradeDAO for database read of all FXTrades
// ----------------------------------------------------------------------------
func GetAllFXTrade(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.GetAllFXTrade()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Update controller, delegates to FXTradeDAO for database save
// ----------------------------------------------------------------------------
func UpdateFXTrade(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty FXTrade model
	// ----------------------------------------------------------------------------
	var data = model.FXTrade{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a FXTrade model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UpdateFXTrade(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to FXTradeDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteFXTrade(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := FXTradeDAO.DeleteFXTrade(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// assigns a Customer on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignCustomerToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignCustomerToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a Customer on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignCustomerFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignCustomerFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	// ----------------------------------------------------------------------------
	// assigns a Bank on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignBankToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignBankToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a Bank on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignBankFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignBankFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	// ----------------------------------------------------------------------------
	// assigns a ExchangeRate on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignExchangeRateToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignExchangeRateToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a ExchangeRate on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignExchangeRateFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignExchangeRateFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	// ----------------------------------------------------------------------------
	// assigns a SourceAccount on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignSourceAccountToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignSourceAccountToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a SourceAccount on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignSourceAccountFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignSourceAccountFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	// ----------------------------------------------------------------------------
	// assigns a DestinationAccount on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignDestinationAccountToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignDestinationAccountToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a DestinationAccount on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignDestinationAccountFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignDestinationAccountFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	// ----------------------------------------------------------------------------
	// assigns a Transaction on a FXTrade
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignTransactionToFXTrade(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.AssignTransactionToFXTrade(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a Transaction on a FXTrade
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignTransactionFromFXTrade( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the FXTrade DAO
	// ----------------------------------------------------------------------------
	requestResult := FXTradeDAO.UnassignTransactionFromFXTrade(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


