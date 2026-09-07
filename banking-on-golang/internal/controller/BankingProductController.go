package controller

import (
    BankingProductDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BankingProductDAO for database creation
//----------------------------------------------------------------------------
func CreateBankingProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BankingProduct model
	//----------------------------------------------------------------------------
	data := model.BankingProduct{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BankingProduct model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct data access object to create
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.CreateBankingProduct( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BankingProductDAO to find the relevant BankingProduct
//----------------------------------------------------------------------------
func GetBankingProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BankingProduct data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.GetBankingProduct(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BankingProductDAO for database read of all BankingProducts
//----------------------------------------------------------------------------
func GetAllBankingProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.GetAllBankingProduct()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BankingProductDAO for database save
//----------------------------------------------------------------------------
func UpdateBankingProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BankingProduct model
	//----------------------------------------------------------------------------
	var data = model.BankingProduct{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BankingProduct model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.UpdateBankingProduct(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BankingProductDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBankingProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BankingProduct data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BankingProductDAO.DeleteBankingProduct(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bank on a BankingProduct
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankToBankingProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankId,_ := strconv.ParseUint( vars["bankId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.AssignBankToBankingProduct(bankingProductId, bankId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bank on a BankingProduct
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankFromBankingProduct( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.UnassignBankFromBankingProduct(bankingProductId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a BankingProduct
	//----------------------------------------------------------------------------
func AddAccountsToBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.AddAccountsToBankingProduct(bankingProductId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a BankingProduct
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.RemoveAccountsFromBankingProduct(bankingProductId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loanAccountsIds as a LoanAccounts to a BankingProduct
	//----------------------------------------------------------------------------
func AddLoanAccountsToBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.AddLoanAccountsToBankingProduct(bankingProductId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loanAccountsIds as a LoanAccounts from a BankingProduct
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanAccountsIds,_ := vars["loanAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.RemoveLoanAccountsFromBankingProduct(bankingProductId, loanAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentCardsIds as a PaymentCards to a BankingProduct
	//----------------------------------------------------------------------------
func AddPaymentCardsToBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.AddPaymentCardsToBankingProduct(bankingProductId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentCardsIds as a PaymentCards from a BankingProduct
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentCardsFromBankingProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bankingProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentCardsIds,_ := vars["paymentCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BankingProduct DAO
	//----------------------------------------------------------------------------
	requestResult := BankingProductDAO.RemovePaymentCardsFromBankingProduct(bankingProductId, paymentCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
