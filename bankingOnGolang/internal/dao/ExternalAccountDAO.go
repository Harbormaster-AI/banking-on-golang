
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExternalAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExternalAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateExternalAccount(obj model.ExternalAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ExternalAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ExternalAccount. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateExternalAccount",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetExternalAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExternalAccount(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ExternalAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ExternalAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ExternalAccount using ID=%v", id )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetExternalAccount",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// GetAllExternalAccount - returns all
//----------------------------------------------------------------------------
func GetAllExternalAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ExternalAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all ExternalAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all ExternalAccount"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExternalAccount. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllExternalAccount",
        Data:       objs,
    }

}

//----------------------------------------------------------------------------
// UpdateExternalAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExternalAccount(obj model.ExternalAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ExternalAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ExternalAccount using ID=%v", obj.ID )
		success = false
	}

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateExternalAccount",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteExternalAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExternalAccount(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExternalAccount(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.ExternalAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ExternalAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ExternalAccount using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteExternalAccount",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a ExternalAccount
//----------------------------------------------------------------------------
func AssignCustomerToExternalAccount( externalAccountId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExternalAccount)

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
			// assign the Customer	to the ExternalAccount
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the ExternalAccount
			//----------------------------------------------------------------------------
			return UpdateExternalAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignCustomer",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a ExternalAccount
//----------------------------------------------------------------------------
func UnassignCustomerFromExternalAccount(externalAccountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExternalAccount)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the ExternalAccount
		//----------------------------------------------------------------------------
		return UpdateExternalAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a ExternalAccount
//----------------------------------------------------------------------------
func AddTransactionsToExternalAccount ( externalAccountId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExternalAccount)

		for _, transactionsId:= range transactionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addTransactionsToExternalAccount",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExternalAccount from the gorm
		//----------------------------------------------------------------------------
		return GetExternalAccount(externalAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a ExternalAccount
//----------------------------------------------------------------------------
func RemoveTransactionsFromExternalAccount( externalAccountId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExternalAccount)

		for _, transactionsId:= range transactionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeTransactionsFromExternalAccount",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExternalAccount from the gorm
		//----------------------------------------------------------------------------
		return GetExternalAccount(externalAccountId)

	} else {
		return parentRequestResult
	}
}

