package controller

import (
    RepaymentScheduleDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RepaymentScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	//----------------------------------------------------------------------------
	data := model.RepaymentSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.CreateRepaymentSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RepaymentScheduleDAO to find the relevant RepaymentSchedule
//----------------------------------------------------------------------------
func GetRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RepaymentSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetRepaymentSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RepaymentScheduleDAO for database read of all RepaymentSchedules
//----------------------------------------------------------------------------
func GetAllRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetAllRepaymentSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RepaymentScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	//----------------------------------------------------------------------------
	var data = model.RepaymentSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UpdateRepaymentSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RepaymentScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RepaymentSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RepaymentScheduleDAO.DeleteRepaymentSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a LoanAccount on a RepaymentSchedule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanAccountToRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountId,_ := strconv.ParseUint( vars["loanAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AssignLoanAccountToRepaymentSchedule(repaymentScheduleId, loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LoanAccount on a RepaymentSchedule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanAccountFromRepaymentSchedule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UnassignLoanAccountFromRepaymentSchedule(repaymentScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Payment on a RepaymentSchedule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPaymentToRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentId,_ := strconv.ParseUint( vars["paymentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AssignPaymentToRepaymentSchedule(repaymentScheduleId, paymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Payment on a RepaymentSchedule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPaymentFromRepaymentSchedule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UnassignPaymentFromRepaymentSchedule(repaymentScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


