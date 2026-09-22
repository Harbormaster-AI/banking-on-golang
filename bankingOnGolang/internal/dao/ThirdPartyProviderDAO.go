
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a ThirdPartyProvider. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateThirdPartyProvider",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetThirdPartyProvider - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetThirdPartyProvider(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetThirdPartyProvider",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all ThirdPartyProvider"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ThirdPartyProvider. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllThirdPartyProvider",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateThirdPartyProvider",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteThirdPartyProvider - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteThirdPartyProvider(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetThirdPartyProvider(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.ThirdPartyProvider)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteThirdPartyProvider",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a ThirdPartyProvider
//----------------------------------------------------------------------------
func AssignBankToThirdPartyProvider( thirdPartyProviderId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

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
			// assign the Bank	to the ThirdPartyProvider
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the ThirdPartyProvider
			//----------------------------------------------------------------------------
			return UpdateThirdPartyProvider(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignBank",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a ThirdPartyProvider
//----------------------------------------------------------------------------
func UnassignBankFromThirdPartyProvider(thirdPartyProviderId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the ThirdPartyProvider
		//----------------------------------------------------------------------------
		return UpdateThirdPartyProvider(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a ThirdPartyProvider
//----------------------------------------------------------------------------
func AddConsentsToThirdPartyProvider ( thirdPartyProviderId uuid.UUID, consentsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		for _, consentsId:= range consentsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addConsentsToThirdPartyProvider",
                    Data:       childObj,
                }
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
func RemoveConsentsFromThirdPartyProvider( thirdPartyProviderId uuid.UUID, consentsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyProvider(thirdPartyProviderId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyProvider)

		for _, consentsId:= range consentsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeConsentsFromThirdPartyProvider",
                    Data:       childObj,
                }
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

