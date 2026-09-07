package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
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
		createMsg = fmt.Sprintf( "Failed trying to create a FundsTransfer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFundsTransfer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFundsTransfer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFundsTransfer(id uint64)(utils.RequestResult){
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
	    getAllMsg = fmt.Sprintf( "Retrieved all FundsTransfer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FundsTransfer", result )
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFundsTransfer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFundsTransfer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFundsTransfer(id uint64)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data. (model.FundsTransfer)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFundsTransfer", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SourceAccount on a FundsTransfer
//----------------------------------------------------------------------------
func AssignSourceAccountToFundsTransfer( fundsTransferId uint64, sourceAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var AccountObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching sourceAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&AccountObj, sourceAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SourceAccount	to the FundsTransfer
			//----------------------------------------------------------------------------
			FundsTransferObj.SourceAccount = &AccountObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(FundsTransferObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceAccount", sourceAccountId )
			return utils.RequestResult{false, msg, "assignSourceAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceAccount on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignSourceAccountFromFundsTransfer(fundsTransferId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Account to the SourceAccount
		//----------------------------------------------------------------------------
		FundsTransferObj.SourceAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceAccount
		//----------------------------------------------------------------------------
		FundsTransferObj.SourceAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(FundsTransferObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a FundsTransfer
//----------------------------------------------------------------------------
func AssignDestinationAccountToFundsTransfer( fundsTransferId uint64, destinationAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var AccountObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching destinationAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&AccountObj, destinationAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DestinationAccount	to the FundsTransfer
			//----------------------------------------------------------------------------
			FundsTransferObj.DestinationAccount = &AccountObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(FundsTransferObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )
			return utils.RequestResult{false, msg, "assignDestinationAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromFundsTransfer(fundsTransferId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		FundsTransferObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		FundsTransferObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(FundsTransferObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExternalBeneficiary on a FundsTransfer
//----------------------------------------------------------------------------
func AssignExternalBeneficiaryToFundsTransfer( fundsTransferId uint64, externalBeneficiaryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var ExternalAccountObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching externalBeneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&ExternalAccountObj, externalBeneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExternalBeneficiary	to the FundsTransfer
			//----------------------------------------------------------------------------
			FundsTransferObj.ExternalBeneficiary = &ExternalAccountObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(FundsTransferObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalBeneficiary", externalBeneficiaryId )
			return utils.RequestResult{false, msg, "assignExternalBeneficiary", ExternalAccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExternalBeneficiary on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignExternalBeneficiaryFromFundsTransfer(fundsTransferId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the ExternalBeneficiary
		//----------------------------------------------------------------------------
		FundsTransferObj.ExternalBeneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExternalBeneficiary
		//----------------------------------------------------------------------------
		FundsTransferObj.ExternalBeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(FundsTransferObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InitiatedBy on a FundsTransfer
//----------------------------------------------------------------------------
func AssignInitiatedByToFundsTransfer( fundsTransferId uint64, initiatedById uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var CustomerObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching initiatedById
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&CustomerObj, initiatedById).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InitiatedBy	to the FundsTransfer
			//----------------------------------------------------------------------------
			FundsTransferObj.InitiatedBy = &CustomerObj

			//----------------------------------------------------------------------------
			// save the FundsTransfer
			//----------------------------------------------------------------------------
			return UpdateFundsTransfer(FundsTransferObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InitiatedBy", initiatedById )
			return utils.RequestResult{false, msg, "assignInitiatedBy", CustomerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InitiatedBy on a FundsTransfer
//----------------------------------------------------------------------------
func UnassignInitiatedByFromFundsTransfer(fundsTransferId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the InitiatedBy
		//----------------------------------------------------------------------------
		FundsTransferObj.InitiatedBy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InitiatedBy
		//----------------------------------------------------------------------------
		FundsTransferObj.InitiatedById = nil;

		//----------------------------------------------------------------------------
		// save the FundsTransfer
		//----------------------------------------------------------------------------
		return UpdateFundsTransfer(FundsTransferObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a FundsTransfer
//----------------------------------------------------------------------------
func AddTransactionsToFundsTransfer ( fundsTransferId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var TransactionObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&TransactionObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&FundsTransferObj).Association("Transactions").Append( &TransactionObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "unassignTransactions", TransactionObj}
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
func RemoveTransactionsFromFundsTransfer( fundsTransferId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FundsTransfer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFundsTransfer(fundsTransferId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FundsTransfer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FundsTransferObj,_ := parentRequestResult.Data. (model.FundsTransfer)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var TransactionObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&TransactionObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&FundsTransferObj).Association("Transactions").Delete( &TransactionObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "removeTransactions", TransactionObj}
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

