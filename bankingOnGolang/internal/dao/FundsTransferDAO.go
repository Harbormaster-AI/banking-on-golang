
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FundsTransferDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFundsTransfer - creates a new db entry
//----------------------------------------------------------------------------
func CreateFundsTransfer(obj model.FundsTransfer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FundsTransfer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FundsTransfer. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFundsTransfer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFundsTransfer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFundsTransfer(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FundsTransfer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FundsTransfer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FundsTransfer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FundsTransfer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFundsTransfer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFundsTransfer - returns all
//----------------------------------------------------------------------------
func GetAllFundsTransfer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FundsTransfer

	//----------------------------------------------------------------------------
	// Request the ORM to find all FundsTransfer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all FundsTransfer"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FundsTransfer. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFundsTransfer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFundsTransfer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFundsTransfer(obj model.FundsTransfer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FundsTransfer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FundsTransfer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{
        Success: success,
        Message: updateMsg,
        Action:  "UpdateFundsTransfer",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFundsTransfer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFundsTransfer(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFundsTransfer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FundsTransfer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FundsTransfer using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success: success,
            Message: deleteMsg,
            Action:  "DeleteFundsTransfer",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SourceAccount on a FundsTransfer
//----------------------------------------------------------------------------
func AssignSourceAccountToFundsTransfer( fundsTransferId uuid.UUID, sourceAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching sourceAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, sourceAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SourceAccount	to the FundsTransfer
			//----------------------------------------------------------------------------
			parentObj.SourceAccount = &childObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceAccount", sourceAccountId )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignSourceAccount",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceAccount on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignSourceAccountFromFundsTransfer(fundsTransferId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Account to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a FundsTransfer
//----------------------------------------------------------------------------
func AssignDestinationAccountToFundsTransfer( fundsTransferId uuid.UUID, destinationAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching destinationAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, destinationAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DestinationAccount	to the FundsTransfer
			//----------------------------------------------------------------------------
			parentObj.DestinationAccount = &childObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignDestinationAccount",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromFundsTransfer(fundsTransferId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExternalBeneficiary on a FundsTransfer
//----------------------------------------------------------------------------
func AssignExternalBeneficiaryToFundsTransfer( fundsTransferId uuid.UUID, externalBeneficiaryId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching externalBeneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, externalBeneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExternalBeneficiary	to the FundsTransfer
			//----------------------------------------------------------------------------
			parentObj.ExternalBeneficiary = &childObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalBeneficiary", externalBeneficiaryId )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignExternalBeneficiary",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExternalBeneficiary on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignExternalBeneficiaryFromFundsTransfer(fundsTransferId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the ExternalBeneficiary
		//----------------------------------------------------------------------------
		parentObj.ExternalBeneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExternalBeneficiary
		//----------------------------------------------------------------------------
		parentObj.ExternalBeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InitiatedBy on a FundsTransfer
//----------------------------------------------------------------------------
func AssignInitiatedByToFundsTransfer( fundsTransferId uuid.UUID, initiatedById uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching initiatedById
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, initiatedById).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InitiatedBy	to the FundsTransfer
			//----------------------------------------------------------------------------
			parentObj.InitiatedBy = &childObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InitiatedBy", initiatedById )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignInitiatedBy",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InitiatedBy on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignInitiatedByFromFundsTransfer(fundsTransferId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the InitiatedBy
		//----------------------------------------------------------------------------
		parentObj.InitiatedBy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InitiatedBy
		//----------------------------------------------------------------------------
		parentObj.InitiatedById = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a FundsTransfer
//----------------------------------------------------------------------------
func AddTransactionsToFundsTransfer ( fundsTransferId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

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

                requestResult = utils.RequestResult{
                    Success: false,
                    Message: msg,
                    Action:  "unassignTransactions",
                    Data:    childObj,
                }
				return requestResult
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FundsTransfer from the gorm
		//----------------------------------------------------------------------------
		return GetFundsTransfer(fundsTransferId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a FundsTransfer
//----------------------------------------------------------------------------
func RemoveTransactionsFromFundsTransfer( fundsTransferId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FundsTransfer)

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
                requestResult = utils.RequestResult{
                                    Success: false,
                                    Message: msg,
                                    Action:  "removeTransactions",
                                    Data:    childObj,
                                }
				return requestResult
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FundsTransfer from the gorm
		//----------------------------------------------------------------------------
		return GetFundsTransfer(fundsTransferId)

	} else {
		return parentRequestResult
	}
}

