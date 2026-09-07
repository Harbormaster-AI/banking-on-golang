package controller

import (
    ExchangeRateDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	utils.ParseBody(r, data)

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
	// Delegate to the ExchangeRate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.GetExchangeRate(ID)
	
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
	utils.ParseBody(r, data)

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
	// Delegate to the ExchangeRate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExchangeRateDAO.DeleteExchangeRate(ID)

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

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.AssignBankToExchangeRate(exchangeRateId, bankId)

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

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.UnassignBankFromExchangeRate(exchangeRateId)

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

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fxTradesIds,_ := vars["fxTradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.AddFxTradesToExchangeRate(exchangeRateId, fxTradesIds)

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

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fxTradesIds,_ := vars["fxTradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.RemoveFxTradesFromExchangeRate(exchangeRateId, fxTradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
