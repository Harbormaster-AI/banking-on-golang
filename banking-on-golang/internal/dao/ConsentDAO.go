package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ConsentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateConsent - creates a new db entry
//----------------------------------------------------------------------------
func CreateConsent(obj model.Consent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Consent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Consent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateConsent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetConsent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetConsent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Consent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Consent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Consent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Consent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetConsent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllConsent - returns all
//----------------------------------------------------------------------------
func GetAllConsent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Consent

	//----------------------------------------------------------------------------
	// Request the ORM to find all Consent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Consent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Consent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllConsent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateConsent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateConsent(obj model.Consent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Consent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Consent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateConsent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteConsent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteConsent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetConsent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Consent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Consent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteConsent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Consent
//----------------------------------------------------------------------------
func AssignCustomerToConsent( consentId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var CustomerObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&CustomerObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the Consent
			//----------------------------------------------------------------------------
			ConsentObj.Customer = &CustomerObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(ConsentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", CustomerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Consent
//----------------------------------------------------------------------------
func UnassignCustomerFromConsent(consentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		ConsentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		ConsentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(ConsentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Bank on a Consent
//----------------------------------------------------------------------------
func AssignBankToConsent( consentId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

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
			// assign the Bank	to the Consent
			//----------------------------------------------------------------------------
			ConsentObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(ConsentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Consent
//----------------------------------------------------------------------------
func UnassignBankFromConsent(consentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		ConsentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		ConsentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(ConsentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ThirdPartyProvider on a Consent
//----------------------------------------------------------------------------
func AssignThirdPartyProviderToConsent( consentId uint64, thirdPartyProviderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var ThirdPartyProviderObj model.ThirdPartyProvider

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider with a
		// matching thirdPartyProviderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&ThirdPartyProviderObj, thirdPartyProviderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ThirdPartyProvider	to the Consent
			//----------------------------------------------------------------------------
			ConsentObj.ThirdPartyProvider = &ThirdPartyProviderObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(ConsentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProvider", thirdPartyProviderId )
			return utils.RequestResult{false, msg, "assignThirdPartyProvider", ThirdPartyProviderObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ThirdPartyProvider on a Consent
//----------------------------------------------------------------------------
func UnassignThirdPartyProviderFromConsent(consentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty ThirdPartyProvider to the ThirdPartyProvider
		//----------------------------------------------------------------------------
		ConsentObj.ThirdPartyProvider = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ThirdPartyProvider
		//----------------------------------------------------------------------------
		ConsentObj.ThirdPartyProviderId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(ConsentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more authorizedAccountsIds as a AuthorizedAccounts to a Consent
//----------------------------------------------------------------------------
func AddAuthorizedAccountsToConsent ( consentId uint64, authorizedAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizedAccountsIds, ",")

		for _, authorizedAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var AccountObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching authorizedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&AccountObj , authorizedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AuthorizedAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ConsentObj).Association("AuthorizedAccounts").Append( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuthorizedAccounts", authorizedAccountsId )
				return utils.RequestResult{false, msg, "unassignAuthorizedAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Consent from the gorm
		//----------------------------------------------------------------------------
		return GetConsent(consentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more authorizedAccountsIds as a AuthorizedAccounts from a Consent
//----------------------------------------------------------------------------
func RemoveAuthorizedAccountsFromConsent( consentId uint64, authorizedAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ConsentObj,_ := parentRequestResult.Data. (model.Consent)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizedAccountsIds, ",")

		for _, authorizedAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var AccountObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching authorizedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&AccountObj , authorizedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the AuthorizedAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ConsentObj).Association("AuthorizedAccounts").Delete( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuthorizedAccounts", authorizedAccountsId )
				return utils.RequestResult{false, msg, "removeAuthorizedAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Consent from the gorm
		//----------------------------------------------------------------------------
		return GetConsent(consentId)

	} else {
		return parentRequestResult
	}
}

