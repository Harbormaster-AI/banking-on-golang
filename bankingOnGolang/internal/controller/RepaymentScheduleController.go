
package controller

import (
    RepaymentScheduleDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to RepaymentScheduleDAO for database creation
// ----------------------------------------------------------------------------
func CreateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	// ----------------------------------------------------------------------------
	data := model.RepaymentSchedule{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to create
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.CreateRepaymentSchedule( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}

// ----------------------------------------------------------------------------
// Get controller, delegates to RepaymentScheduleDAO to find the relevant RepaymentSchedule
// ----------------------------------------------------------------------------
func GetRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetRepaymentSchedule(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to RepaymentScheduleDAO for database read of all RepaymentSchedules
// ----------------------------------------------------------------------------
func GetAllRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetAllRepaymentSchedule()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}

// ----------------------------------------------------------------------------
// Update controller, delegates to RepaymentScheduleDAO for database save
// ----------------------------------------------------------------------------
func UpdateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	// ----------------------------------------------------------------------------
	var data = model.RepaymentSchedule{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UpdateRepaymentSchedule(data)

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
// Delete controller, delegates to RepaymentScheduleDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := RepaymentScheduleDAO.DeleteRepaymentSchedule(data.Id)

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
	// assigns a LoanAccount on a RepaymentSchedule
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignLoanAccountToRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AssignLoanAccountToRepaymentSchedule(data.ParentId, data.ChildId)

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
	// unassigns a LoanAccount on a RepaymentSchedule
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignLoanAccountFromRepaymentSchedule( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UnassignLoanAccountFromRepaymentSchedule(data.ParentId)

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
	// assigns a Payment on a RepaymentSchedule
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignPaymentToRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AssignPaymentToRepaymentSchedule(data.ParentId, data.ChildId)

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
	// unassigns a Payment on a RepaymentSchedule
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignPaymentFromRepaymentSchedule( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	// ----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UnassignPaymentFromRepaymentSchedule(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}


