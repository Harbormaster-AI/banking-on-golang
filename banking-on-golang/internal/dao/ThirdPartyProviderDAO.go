package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ThirdPartyProviderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateThirdPartyProvider - creates a new db entry
//----------------------------------------------------------------------------
func CreateThirdPartyProvider(obj model.ThirdPartyProvider)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ThirdPartyProvider with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ThirdPartyProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateThirdPartyProvider", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetThirdPartyProvider - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetThirdPartyProvider(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ThirdPartyProvider

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ThirdPartyProvider using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ThirdPartyProvider using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetThirdPartyProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllThirdPartyProvider - returns all
//----------------------------------------------------------------------------
func GetAllThirdPartyProvider()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ThirdPartyProvider

	//----------------------------------------------------------------------------
	// Request the ORM to find all ThirdPartyProvider
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ThirdPartyProvider" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ThirdPartyProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllThirdPartyProvider", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateThirdPartyProvider - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateThirdPartyProvider(obj model.ThirdPartyProvider)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ThirdPartyProvider using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ThirdPartyProvider using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateThirdPartyProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteThirdPartyProvider - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteThirdPartyProvider(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetThirdPartyProvider(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ThirdPartyProvider)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ThirdPartyProvider using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ThirdPartyProvider using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteThirdPartyProvider", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a ThirdPartyProvider
//----------------------------------------------------------------------------
func AssignBankToThirdPartyProvider( thirdPartyProviderId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var BankObj model.Bank

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Bank with a
		// matching bankId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&BankObj, bankId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bank	to the ThirdPartyProvider
			//----------------------------------------------------------------------------
			ThirdPartyProviderObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the ThirdPartyProvider
			//----------------------------------------------------------------------------
			return UpdateThirdPartyProvider(ThirdPartyProviderObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a ThirdPartyProvider
//----------------------------------------------------------------------------
func UnassignBankFromThirdPartyProvider(thirdPartyProviderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the ThirdPartyProvider
		//----------------------------------------------------------------------------
		return UpdateThirdPartyProvider(ThirdPartyProviderObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a ThirdPartyProvider
//----------------------------------------------------------------------------
func AddConsentsToThirdPartyProvider ( thirdPartyProviderId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ConsentObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ConsentObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ThirdPartyProviderObj).Association("Consents").Append( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", ConsentObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdPartyProvider from the gorm
		//----------------------------------------------------------------------------
		return GetThirdPartyProvider(thirdPartyProviderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a ThirdPartyProvider
//----------------------------------------------------------------------------
func RemoveConsentsFromThirdPartyProvider( thirdPartyProviderId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ThirdPartyProviderObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ConsentObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ConsentObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ThirdPartyProviderObj).Association("Consents").Delete( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", ConsentObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdPartyProvider from the gorm
		//----------------------------------------------------------------------------
		return GetThirdPartyProvider(thirdPartyProviderId)

	} else {
		return parentRequestResult
	}
}

