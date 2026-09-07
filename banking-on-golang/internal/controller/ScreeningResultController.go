package controller

import (
    ScreeningResultDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ScreeningResultDAO for database creation
//----------------------------------------------------------------------------
func CreateScreeningResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ScreeningResult model
	//----------------------------------------------------------------------------
	data := model.ScreeningResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ScreeningResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object to create
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.CreateScreeningResult( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ScreeningResultDAO to find the relevant ScreeningResult
//----------------------------------------------------------------------------
func GetScreeningResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ScreeningResult data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.GetScreeningResult(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ScreeningResultDAO for database read of all ScreeningResults
//----------------------------------------------------------------------------
func GetAllScreeningResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.GetAllScreeningResult()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ScreeningResultDAO for database save
//----------------------------------------------------------------------------
func UpdateScreeningResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ScreeningResult model
	//----------------------------------------------------------------------------
	var data = model.ScreeningResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ScreeningResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.UpdateScreeningResult(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ScreeningResultDAO for database deletion
//----------------------------------------------------------------------------
func DeleteScreeningResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ScreeningResult data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ScreeningResultDAO.DeleteScreeningResult(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a KycProfile on a ScreeningResult
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignKycProfileToScreeningResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	screeningResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfileId,_ := strconv.ParseUint( vars["kycProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScreeningResult DAO
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.AssignKycProfileToScreeningResult(screeningResultId, kycProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a KycProfile on a ScreeningResult
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignKycProfileFromScreeningResult( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	screeningResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScreeningResult DAO
	//----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.UnassignKycProfileFromScreeningResult(screeningResultId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


