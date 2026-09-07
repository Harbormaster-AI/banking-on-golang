package controller

import (
    IdentityDocumentDAO "banking-on-golang/internal/dao"
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to IdentityDocumentDAO for database creation
//----------------------------------------------------------------------------
func CreateIdentityDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty IdentityDocument model
	//----------------------------------------------------------------------------
	data := model.IdentityDocument{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a IdentityDocument model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object to create
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.CreateIdentityDocument( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to IdentityDocumentDAO to find the relevant IdentityDocument
//----------------------------------------------------------------------------
func GetIdentityDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the IdentityDocument data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.GetIdentityDocument(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to IdentityDocumentDAO for database read of all IdentityDocuments
//----------------------------------------------------------------------------
func GetAllIdentityDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object to get all
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.GetAllIdentityDocument()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to IdentityDocumentDAO for database save
//----------------------------------------------------------------------------
func UpdateIdentityDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty IdentityDocument model
	//----------------------------------------------------------------------------
	var data = model.IdentityDocument{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a IdentityDocument model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the IdentityDocument data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.UpdateIdentityDocument(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to IdentityDocumentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteIdentityDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the IdentityDocument data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := IdentityDocumentDAO.DeleteIdentityDocument(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a KycProfile on a IdentityDocument
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignKycProfileToIdentityDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	identityDocumentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfileId,_ := strconv.ParseUint( vars["kycProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IdentityDocument DAO
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.AssignKycProfileToIdentityDocument(identityDocumentId, kycProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a KycProfile on a IdentityDocument
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignKycProfileFromIdentityDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	identityDocumentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IdentityDocument DAO
	//----------------------------------------------------------------------------
	requestResult := IdentityDocumentDAO.UnassignKycProfileFromIdentityDocument(identityDocumentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


