package controller

import (
	ScreeningResultDAO "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to ScreeningResultDAO for database creation
// ----------------------------------------------------------------------------
func CreateScreeningResult(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty ScreeningResult model
	// ----------------------------------------------------------------------------
	data := model.ScreeningResult{}

	// ----------------------------------------------------------------------------
	// Parse the body into a ScreeningResult model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object to create
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.CreateScreeningResult(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Get controller, delegates to ScreeningResultDAO to find the relevant ScreeningResult
// ----------------------------------------------------------------------------
func GetScreeningResult(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.GetScreeningResult(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// GetAll controller, delegates to ScreeningResultDAO for database read of all ScreeningResults
// ----------------------------------------------------------------------------
func GetAllScreeningResult(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.GetAllScreeningResult()

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Update controller, delegates to ScreeningResultDAO for database save
// ----------------------------------------------------------------------------
func UpdateScreeningResult(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty ScreeningResult model
	// ----------------------------------------------------------------------------
	var data = model.ScreeningResult{}

	// ----------------------------------------------------------------------------
	// Parse the body into a ScreeningResult model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.UpdateScreeningResult(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to ScreeningResultDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteScreeningResult(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.DeleteScreeningResult(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// assigns a KycProfile on a ScreeningResult
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignKycProfileToScreeningResult(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult DAO
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.AssignKycProfileToScreeningResult(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// ----------------------------------------------------------------------------
// unassigns a KycProfile on a ScreeningResult
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignKycProfileFromScreeningResult(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the ScreeningResult DAO
	// ----------------------------------------------------------------------------
	requestResult := ScreeningResultDAO.UnassignKycProfileFromScreeningResult(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
