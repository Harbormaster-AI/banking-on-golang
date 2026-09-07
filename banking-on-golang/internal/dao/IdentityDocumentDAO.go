package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing IdentityDocumentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateIdentityDocument - creates a new db entry
//----------------------------------------------------------------------------
func CreateIdentityDocument(obj model.IdentityDocument)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a IdentityDocument with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a IdentityDocument", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateIdentityDocument", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetIdentityDocument - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetIdentityDocument(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.IdentityDocument

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a IdentityDocument with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a IdentityDocument using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a IdentityDocument using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetIdentityDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllIdentityDocument - returns all
//----------------------------------------------------------------------------
func GetAllIdentityDocument()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.IdentityDocument

	//----------------------------------------------------------------------------
	// Request the ORM to find all IdentityDocument
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all IdentityDocument" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all IdentityDocument", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllIdentityDocument", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateIdentityDocument - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateIdentityDocument(obj model.IdentityDocument)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a IdentityDocument using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a IdentityDocument using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateIdentityDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteIdentityDocument - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteIdentityDocument(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the IdentityDocument with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetIdentityDocument(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.IdentityDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.IdentityDocument)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a IdentityDocument using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a IdentityDocument using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteIdentityDocument", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a IdentityDocument
//----------------------------------------------------------------------------
func AssignKycProfileToIdentityDocument( identityDocumentId uint64, kycProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the IdentityDocument with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIdentityDocument(identityDocumentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.IdentityDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		IdentityDocumentObj,_ := parentRequestResult.Data. (model.IdentityDocument)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var KycProfileObj model.KycProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KycProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&KycProfileObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the IdentityDocument
			//----------------------------------------------------------------------------
			IdentityDocumentObj.KycProfile = &KycProfileObj

			//----------------------------------------------------------------------------
			// save the IdentityDocument
			//----------------------------------------------------------------------------
			return UpdateIdentityDocument(IdentityDocumentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )
			return utils.RequestResult{false, msg, "assignKycProfile", KycProfileObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a IdentityDocument
//----------------------------------------------------------------------------
func UnassignKycProfileFromIdentityDocument(identityDocumentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the IdentityDocument with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIdentityDocument(identityDocumentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.IdentityDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		IdentityDocumentObj,_ := parentRequestResult.Data. (model.IdentityDocument)

		//----------------------------------------------------------------------------
		// assign an empty KycProfile to the KycProfile
		//----------------------------------------------------------------------------
		IdentityDocumentObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		IdentityDocumentObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the IdentityDocument
		//----------------------------------------------------------------------------
		return UpdateIdentityDocument(IdentityDocumentObj)

	} else {
		return parentRequestResult
	}

}


