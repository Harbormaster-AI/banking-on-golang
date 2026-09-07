package controller

import (
    BranchDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BranchDAO for database creation
//----------------------------------------------------------------------------
func CreateBranch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Branch model
	//----------------------------------------------------------------------------
	data := model.Branch{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Branch model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Branch data access object to create
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.CreateBranch( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BranchDAO to find the relevant Branch
//----------------------------------------------------------------------------
func GetBranch(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Branch data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.GetBranch(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BranchDAO for database read of all Branchs
//----------------------------------------------------------------------------
func GetAllBranch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Branch data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.GetAllBranch()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BranchDAO for database save
//----------------------------------------------------------------------------
func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Branch model
	//----------------------------------------------------------------------------
	var data = model.Branch{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Branch model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Branch data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.UpdateBranch(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BranchDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBranch(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Branch data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BranchDAO.DeleteBranch(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a Branch
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToBranch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.AssignBankToBranch(branchId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a Branch
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromBranch( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.UnassignBankFromBranch(branchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Branch
	//----------------------------------------------------------------------------
func AddAccountsToBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.AddAccountsToBranch(branchId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Branch
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.RemoveAccountsFromBranch(branchId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loanAccountsIds as a LoanAccounts to a Branch
	//----------------------------------------------------------------------------
func AddLoanAccountsToBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.AddLoanAccountsToBranch(branchId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loanAccountsIds as a LoanAccounts from a Branch
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.RemoveLoanAccountsFromBranch(branchId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more atmsIds as a Atms to a Branch
	//----------------------------------------------------------------------------
func AddAtmsToBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	atmsIds,_ := vars["atmsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.AddAtmsToBranch(branchId, atmsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more atmsIds as a Atms from a Branch
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAtmsFromBranch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	branchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	atmsIds,_ := vars["atmsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Branch DAO
	//----------------------------------------------------------------------------
	requestResult := BranchDAO.RemoveAtmsFromBranch(branchId, atmsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
