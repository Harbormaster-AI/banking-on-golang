
package controller

import (
    ExchangeRateDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	 "encoding/json"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExchangeRateDAO for database creation
//----------------------------------------------------------------------------
func CreateExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExchangeRate model
	//----------------------------------------------------------------------------
	data := model.ExchangeRate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExchangeRate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.CreateExchangeRate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExchangeRateDAO to find the relevant ExchangeRate
//----------------------------------------------------------------------------
func GetExchangeRate(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	//----------------------------------------------------------------------------
	data := model.GetRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.GetExchangeRate(data.Id)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExchangeRateDAO for database read of all ExchangeRates
//----------------------------------------------------------------------------
func GetAllExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.GetAllExchangeRate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExchangeRateDAO for database save
//----------------------------------------------------------------------------
func UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExchangeRate model
	//----------------------------------------------------------------------------
	var data = model.ExchangeRate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExchangeRate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.UpdateExchangeRate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExchangeRateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	//----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExchangeRateDAO.DeleteExchangeRate(data.Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a ExchangeRate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToExchangeRate(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	//----------------------------------------------------------------------------
	data := model.AssignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.AssignBankToExchangeRate(data.ParentId, data.ChildId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a ExchangeRate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromExchangeRate( w http.ResponseWriter, r *http.Request ) {

	//----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	//----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.UnassignBankFromExchangeRate(data.ParentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more fxTradesIds as a FxTrades to a ExchangeRate
	//----------------------------------------------------------------------------
func AddFxTradesToExchangeRate(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	//----------------------------------------------------------------------------
	data := model.AddToRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.AddFxTradesToExchangeRate(data.ParentId, data.childIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more fxTradesIds as a FxTrades from a ExchangeRate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFxTradesFromExchangeRate(w http.ResponseWriter, r *http.Request)  {

	//----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	//----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	//----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.RemoveFxTradesFromExchangeRate(data.ParentId, data.ChildIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
