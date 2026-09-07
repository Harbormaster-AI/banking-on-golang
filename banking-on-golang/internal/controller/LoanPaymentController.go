package controller

import (
    LoanPaymentDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LoanPaymentDAO for database creation
//----------------------------------------------------------------------------
func CreateLoanPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanPayment model
	//----------------------------------------------------------------------------
	data := model.LoanPayment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanPayment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object to create
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.CreateLoanPayment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LoanPaymentDAO to find the relevant LoanPayment
//----------------------------------------------------------------------------
func GetLoanPayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanPayment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.GetLoanPayment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LoanPaymentDAO for database read of all LoanPayments
//----------------------------------------------------------------------------
func GetAllLoanPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.GetAllLoanPayment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LoanPaymentDAO for database save
//----------------------------------------------------------------------------
func UpdateLoanPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanPayment model
	//----------------------------------------------------------------------------
	var data = model.LoanPayment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanPayment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UpdateLoanPayment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LoanPaymentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLoanPayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanPayment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LoanPaymentDAO.DeleteLoanPayment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a LoanAccount on a LoanPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanAccountToLoanPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountId,_ := strconv.ParseUint( vars["loanAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.AssignLoanAccountToLoanPayment(loanPaymentId, loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LoanAccount on a LoanPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanAccountFromLoanPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UnassignLoanAccountFromLoanPayment(loanPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a LoanPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToLoanPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.AssignTransactionToLoanPayment(loanPaymentId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a LoanPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromLoanPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanPayment DAO
	//----------------------------------------------------------------------------
	requestResult := LoanPaymentDAO.UnassignTransactionFromLoanPayment(loanPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


