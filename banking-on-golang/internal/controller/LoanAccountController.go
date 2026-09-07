package controller

import (
    LoanAccountDAO "bankingongolang/internal/dao"
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LoanAccountDAO for database creation
//----------------------------------------------------------------------------
func CreateLoanAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanAccount model
	//----------------------------------------------------------------------------
	data := model.LoanAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object to create
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.CreateLoanAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LoanAccountDAO to find the relevant LoanAccount
//----------------------------------------------------------------------------
func GetLoanAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.GetLoanAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LoanAccountDAO for database read of all LoanAccounts
//----------------------------------------------------------------------------
func GetAllLoanAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.GetAllLoanAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LoanAccountDAO for database save
//----------------------------------------------------------------------------
func UpdateLoanAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanAccount model
	//----------------------------------------------------------------------------
	var data = model.LoanAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UpdateLoanAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LoanAccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLoanAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LoanAccountDAO.DeleteLoanAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a LoanAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToLoanAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignBankToLoanAccount(loanAccountId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a LoanAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromLoanAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignBankFromLoanAccount(loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Branch on a LoanAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBranchToLoanAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchId,_ := strconv.ParseUint( vars["branchId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignBranchToLoanAccount(loanAccountId, branchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Branch on a LoanAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBranchFromLoanAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignBranchFromLoanAccount(loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a LoanAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToLoanAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignProductToLoanAccount(loanAccountId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a LoanAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromLoanAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignProductFromLoanAccount(loanAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more borrowersIds as a Borrowers to a LoanAccount
	//----------------------------------------------------------------------------
func AddBorrowersToLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	borrowersIds,_ := vars["borrowersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddBorrowersToLoanAccount(loanAccountId, borrowersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more borrowersIds as a Borrowers from a LoanAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBorrowersFromLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	borrowersIds,_ := vars["borrowersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveBorrowersFromLoanAccount(loanAccountId, borrowersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more repaymentScheduleIds as a RepaymentSchedule to a LoanAccount
	//----------------------------------------------------------------------------
func AddRepaymentScheduleToLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repaymentScheduleIds,_ := vars["repaymentScheduleIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddRepaymentScheduleToLoanAccount(loanAccountId, repaymentScheduleIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more repaymentScheduleIds as a RepaymentSchedule from a LoanAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRepaymentScheduleFromLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repaymentScheduleIds,_ := vars["repaymentScheduleIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveRepaymentScheduleFromLoanAccount(loanAccountId, repaymentScheduleIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentsIds as a Payments to a LoanAccount
	//----------------------------------------------------------------------------
func AddPaymentsToLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddPaymentsToLoanAccount(loanAccountId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentsIds as a Payments from a LoanAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentsFromLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemovePaymentsFromLoanAccount(loanAccountId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more collateralIds as a Collateral to a LoanAccount
	//----------------------------------------------------------------------------
func AddCollateralToLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	collateralIds,_ := vars["collateralIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddCollateralToLoanAccount(loanAccountId, collateralIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more collateralIds as a Collateral from a LoanAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCollateralFromLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	collateralIds,_ := vars["collateralIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveCollateralFromLoanAccount(loanAccountId, collateralIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more feeChargesIds as a FeeCharges to a LoanAccount
	//----------------------------------------------------------------------------
func AddFeeChargesToLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeChargesIds,_ := vars["feeChargesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddFeeChargesToLoanAccount(loanAccountId, feeChargesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more feeChargesIds as a FeeCharges from a LoanAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeeChargesFromLoanAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeChargesIds,_ := vars["feeChargesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	//----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveFeeChargesFromLoanAccount(loanAccountId, feeChargesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
