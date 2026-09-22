
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a Consent. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateConsent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetConsent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetConsent(id uuid.UUID)(utils.RequestResult){
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
	    getAllMsg = "Retrieved all Consent"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Consent. Result: %s", result )
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

	requestResult = utils.RequestResult{
        Success: success,
        Msg: updateMsg,
        Call:  "UpdateConsent",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteConsent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteConsent(id uuid.UUID)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data.(model.Consent)

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

        requestResult = utils.RequestResult{
            Success: success,
            Msg: deleteMsg,
            Call:  "DeleteConsent",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Consent
//----------------------------------------------------------------------------
func AssignCustomerToConsent( consentId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the Consent
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignCustomer",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Consent
//----------------------------------------------------------------------------
func UnassignCustomerFromConsent(consentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Bank on a Consent
//----------------------------------------------------------------------------
func AssignBankToConsent( consentId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Bank

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Bank with a
		// matching bankId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, bankId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bank	to the Consent
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignBank",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Consent
//----------------------------------------------------------------------------
func UnassignBankFromConsent(consentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ThirdPartyProvider on a Consent
//----------------------------------------------------------------------------
func AssignThirdPartyProviderToConsent( consentId uuid.UUID, thirdPartyProviderId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ThirdPartyProvider

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider with a
		// matching thirdPartyProviderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, thirdPartyProviderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ThirdPartyProvider	to the Consent
			//----------------------------------------------------------------------------
			parentObj.ThirdPartyProvider = &childObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProvider", thirdPartyProviderId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignThirdPartyProvider",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ThirdPartyProvider on a Consent
//----------------------------------------------------------------------------
func UnassignThirdPartyProviderFromConsent(consentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty ThirdPartyProvider to the ThirdPartyProvider
		//----------------------------------------------------------------------------
		parentObj.ThirdPartyProvider = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ThirdPartyProvider
		//----------------------------------------------------------------------------
		parentObj.ThirdPartyProviderId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more authorizedAccountsIds as a AuthorizedAccounts to a Consent
//----------------------------------------------------------------------------
func AddAuthorizedAccountsToConsent ( consentId uuid.UUID, authorizedAccountsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		for _, authorizedAccountsId:= range authorizedAccountsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching authorizedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AuthorizedAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AuthorizedAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuthorizedAccounts", authorizedAccountsId )

                requestResult = utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignAuthorizedAccounts",
                    Data:    childObj,
                }
				return requestResult
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
func RemoveAuthorizedAccountsFromConsent( consentId uuid.UUID, authorizedAccountsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		for _, authorizedAccountsId:= range authorizedAccountsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching authorizedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the AuthorizedAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AuthorizedAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuthorizedAccounts", authorizedAccountsId )
                requestResult = utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeAuthorizedAccounts",
                                    Data:    childObj,
                                }
				return requestResult
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

