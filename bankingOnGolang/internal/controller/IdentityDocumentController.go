
package controller

import (
    IdentityDocumentDAO "bankingOnGolang/internal/dao"
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
	"net/http"
	 "encoding/json"
)

// ----------------------------------------------------------------------------
// Create controller, delegates to IdentityDocumentDAO for database creation
// ----------------------------------------------------------------------------
func CreateIdentityDocument(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty IdentityDocument model
	// ----------------------------------------------------------------------------
	data := model.IdentityDocument{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a IdentityDocument model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object to create
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.CreateIdentityDocument( data )
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Get controller, delegates to IdentityDocumentDAO to find the relevant IdentityDocument
// ----------------------------------------------------------------------------
func GetIdentityDocument(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty GetRequest model
	// ----------------------------------------------------------------------------
	data := model.GetRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a GetRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object
	// find the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.GetIdentityDocument(data.Id)
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// ----------------------------------------------------------------------------
// GetAll controller, delegates to IdentityDocumentDAO for database read of all IdentityDocuments
// ----------------------------------------------------------------------------
func GetAllIdentityDocument(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object to get all
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.GetAllIdentityDocument()
	
	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Update controller, delegates to IdentityDocumentDAO for database save
// ----------------------------------------------------------------------------
func UpdateIdentityDocument(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty IdentityDocument model
	// ----------------------------------------------------------------------------
	var data = model.IdentityDocument{}
	
	// ----------------------------------------------------------------------------
	// Parse the body into a IdentityDocument model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object
	// update the one with the matching identifier
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.UpdateIdentityDocument(data)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ----------------------------------------------------------------------------
// Delete controller, delegates to IdentityDocumentDAO for database deletion
// ----------------------------------------------------------------------------
func DeleteIdentityDocument(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------------
	// Initialize an empty DeleteRequest model
	// ----------------------------------------------------------------------------
	data := model.DeleteRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a DeleteRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object
	// delete the one with the matching identifier
	// ----------------------------------------------------------------------------	
	requestResult := IdentityDocumentDAO.DeleteIdentityDocument(data.Id)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// assigns a KycProfile on a IdentityDocument
	// delegates to an ORM handler
	// ----------------------------------------------------------------------------
func AssignKycProfileToIdentityDocument(w http.ResponseWriter, r *http.Request) {

	// ----------------------------------------------------------------------------
	// Initialize an empty AssignRequest model
	// ----------------------------------------------------------------------------
	data := model.AssignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a AssignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument DAO
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.AssignKycProfileToIdentityDocument(data.ParentId, data.ChildId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	// ----------------------------------------------------------------------------
	// unassigns a KycProfile on a IdentityDocument
	// delegates to the ORM handler
	// ----------------------------------------------------------------------------
func UnassignKycProfileFromIdentityDocument( w http.ResponseWriter, r *http.Request ) {

	// ----------------------------------------------------------------------------
	// Initialize an empty UnassignRequest model
	// ----------------------------------------------------------------------------
	data := model.UnassignRequest{}

	// ----------------------------------------------------------------------------
	// Parse the body into a UnassignRequest model structure
	// ----------------------------------------------------------------------------
	utils.ParseBody(r, &data)

	// ----------------------------------------------------------------------------
	// Delegate to the IdentityDocument DAO
	// ----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.UnassignKycProfileFromIdentityDocument(data.ParentId)

	// ----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	// ----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


