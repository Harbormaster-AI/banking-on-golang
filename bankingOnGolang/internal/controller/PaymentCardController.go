
package controller

import (
    PaymentCardDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to PaymentCardDAO for database creation
// ----------------------------------------------------------------------------
func CreatePaymentCard(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty PaymentCard model
	// ----------------------------------------------------------------------------
	data := model.PaymentCard{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a PaymentCard model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object to create
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.CreatePaymentCard( data )
	
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
// Get controller, delegates to PaymentCardDAO to find the relevant PaymentCard
// ----------------------------------------------------------------------------
func GetPaymentCard(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.GetPaymentCard(data.Id)
	
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
// GetAll controller, delegates to PaymentCardDAO for database read of all PaymentCards
// ----------------------------------------------------------------------------
func GetAllPaymentCard(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.GetAllPaymentCard()
	
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
// Update controller, delegates to PaymentCardDAO for database save
// ----------------------------------------------------------------------------
func UpdatePaymentCard(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty PaymentCard model
	// ----------------------------------------------------------------------------
	var data = model.PaymentCard{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a PaymentCard model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UpdatePaymentCard(data)

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
// Delete controller, delegates to PaymentCardDAO for database deletion
// ----------------------------------------------------------------------------
func DeletePaymentCard(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := PaymentCardDAO.DeletePaymentCard(data.Id)

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
	// assigns a Bank on a PaymentCard
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignBankToPaymentCard(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AssignBankToPaymentCard(data.ParentId, data.ChildId)

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
	// unassigns a Bank on a PaymentCard
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignBankFromPaymentCard( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UnassignBankFromPaymentCard(data.ParentId)

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
	// assigns a Account on a PaymentCard
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignAccountToPaymentCard(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AssignAccountToPaymentCard(data.ParentId, data.ChildId)

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
	// unassigns a Account on a PaymentCard
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignAccountFromPaymentCard( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UnassignAccountFromPaymentCard(data.ParentId)

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
	// assigns a Customer on a PaymentCard
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignCustomerToPaymentCard(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AssignCustomerToPaymentCard(data.ParentId, data.ChildId)

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
	// unassigns a Customer on a PaymentCard
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignCustomerFromPaymentCard( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UnassignCustomerFromPaymentCard(data.ParentId)

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
	// adds one or more transactionsIds as a Transactions to a PaymentCard
	// ----------------------------------------------------------------------------
func AddTransactionsToPaymentCard(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty AddToRequest model
	// ----------------------------------------------------------------------------
	data := model.AddToRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AddToRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AddTransactionsToPaymentCard(data.ParentId, data.ChildIds)

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
	// removes one or more transactionsIds as a Transactions from a PaymentCard
	// delegates via URI to an ORM handler
	// ----------------------------------------------------------------------------
func RemoveTransactionsFromPaymentCard(w http.ResponseWriter, r *http.Request)  {

	// ----------------------------------------------------------------------------
	// Initialize an empty RemoveFromRequest model
	// ----------------------------------------------------------------------------
	data := model.RemoveFromRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a RemoveFromRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	// ----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.RemoveTransactionsFromPaymentCard(data.ParentId, data.ChildIds)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}
		
