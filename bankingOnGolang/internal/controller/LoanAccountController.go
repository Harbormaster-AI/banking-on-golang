package controller

import (
	LoanAccountDAO "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to LoanAccountDAO for database creation
// ----------------------------------------------------------------------------
func CreateLoanAccount(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty LoanAccount model
	// ----------------------------------------------------------------------------
	data := model.LoanAccount{}

	// ----------------------------------------------------------------------------
	// Parse the body into a LoanAccount model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object to create
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.CreateLoanAccount(data)

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
// Get controller, delegates to LoanAccountDAO to find the relevant LoanAccount
// ----------------------------------------------------------------------------
func GetLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.GetLoanAccount(data.Id)

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
// GetAll controller, delegates to LoanAccountDAO for database read of all LoanAccounts
// ----------------------------------------------------------------------------
func GetAllLoanAccount(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.GetAllLoanAccount()

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
// Update controller, delegates to LoanAccountDAO for database save
// ----------------------------------------------------------------------------
func UpdateLoanAccount(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty LoanAccount model
	// ----------------------------------------------------------------------------
	var data = model.LoanAccount{}

	// ----------------------------------------------------------------------------
	// Parse the body into a LoanAccount model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UpdateLoanAccount(data)

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
// Delete controller, delegates to LoanAccountDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteLoanAccount(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.DeleteLoanAccount(data.Id)

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
// assigns a Bank on a LoanAccount
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignBankToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignBankToLoanAccount(data.ParentId, data.ChildId)

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
// unassigns a Bank on a LoanAccount
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignBankFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignBankFromLoanAccount(data.ParentId)

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
// assigns a Branch on a LoanAccount
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignBranchToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignBranchToLoanAccount(data.ParentId, data.ChildId)

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
// unassigns a Branch on a LoanAccount
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignBranchFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignBranchFromLoanAccount(data.ParentId)

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
// assigns a Product on a LoanAccount
// delegates to an ORM handler
// ----------------------------------------------------------------------------
func AssignProductToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AssignProductToLoanAccount(data.ParentId, data.ChildId)

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
// unassigns a Product on a LoanAccount
// delegates to the ORM handler
// ----------------------------------------------------------------------------
func UnassignProductFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.UnassignProductFromLoanAccount(data.ParentId)

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
// adds one or more borrowersIds as a Borrowers to a LoanAccount
// ----------------------------------------------------------------------------
func AddBorrowersToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddBorrowersToLoanAccount(data.ParentId, data.ChildIds)

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
// removes one or more borrowersIds as a Borrowers from a LoanAccount
// delegates via URI to an ORM handler
// ----------------------------------------------------------------------------
func RemoveBorrowersFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveBorrowersFromLoanAccount(data.ParentId, data.ChildIds)

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
// adds one or more repaymentScheduleIds as a RepaymentSchedule to a LoanAccount
// ----------------------------------------------------------------------------
func AddRepaymentScheduleToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddRepaymentScheduleToLoanAccount(data.ParentId, data.ChildIds)

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
// removes one or more repaymentScheduleIds as a RepaymentSchedule from a LoanAccount
// delegates via URI to an ORM handler
// ----------------------------------------------------------------------------
func RemoveRepaymentScheduleFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveRepaymentScheduleFromLoanAccount(data.ParentId, data.ChildIds)

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
// adds one or more paymentsIds as a Payments to a LoanAccount
// ----------------------------------------------------------------------------
func AddPaymentsToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddPaymentsToLoanAccount(data.ParentId, data.ChildIds)

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
// removes one or more paymentsIds as a Payments from a LoanAccount
// delegates via URI to an ORM handler
// ----------------------------------------------------------------------------
func RemovePaymentsFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemovePaymentsFromLoanAccount(data.ParentId, data.ChildIds)

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
// adds one or more collateralIds as a Collateral to a LoanAccount
// ----------------------------------------------------------------------------
func AddCollateralToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddCollateralToLoanAccount(data.ParentId, data.ChildIds)

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
// removes one or more collateralIds as a Collateral from a LoanAccount
// delegates via URI to an ORM handler
// ----------------------------------------------------------------------------
func RemoveCollateralFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveCollateralFromLoanAccount(data.ParentId, data.ChildIds)

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
// adds one or more feeChargesIds as a FeeCharges to a LoanAccount
// ----------------------------------------------------------------------------
func AddFeeChargesToLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.AddFeeChargesToLoanAccount(data.ParentId, data.ChildIds)

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
// removes one or more feeChargesIds as a FeeCharges from a LoanAccount
// delegates via URI to an ORM handler
// ----------------------------------------------------------------------------
func RemoveFeeChargesFromLoanAccount(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the LoanAccount DAO
	// ----------------------------------------------------------------------------
	requestResult := LoanAccountDAO.RemoveFeeChargesFromLoanAccount(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
