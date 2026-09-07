package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
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
	var requestResult utils.RequestResult
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
		createMsg = fmt.Sprintf( "Failed trying to create a ExternalAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExternalAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExternalAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExternalAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
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

	requestResult = utils.RequestResult{success, getMsg, "GetExternalAccount", obj}

	return requestResult
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
	    getAllMsg = fmt.Sprintf( "Retrieved all ExternalAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExternalAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExternalAccount", objs}
	return requestResult
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExternalAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExternalAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExternalAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExternalAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ExternalAccount)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExternalAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a ExternalAccount
//----------------------------------------------------------------------------
func AssignCustomerToExternalAccount( externalAccountId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExternalAccountObj,_ := parentRequestResult.Data. (model.ExternalAccount)

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
			// assign the Customer	to the ExternalAccount
			//----------------------------------------------------------------------------
			ExternalAccountObj.Customer = &CustomerObj

			//----------------------------------------------------------------------------
			// save the ExternalAccount
			//----------------------------------------------------------------------------
			return UpdateExternalAccount(ExternalAccountObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", CustomerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a ExternalAccount
//----------------------------------------------------------------------------
func UnassignCustomerFromExternalAccount(externalAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExternalAccountObj,_ := parentRequestResult.Data. (model.ExternalAccount)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		ExternalAccountObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		ExternalAccountObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the ExternalAccount
		//----------------------------------------------------------------------------
		return UpdateExternalAccount(ExternalAccountObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a ExternalAccount
//----------------------------------------------------------------------------
func AddTransactionsToExternalAccount ( externalAccountId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExternalAccountObj,_ := parentRequestResult.Data. (model.ExternalAccount)

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
				utils.GetDB().Model(&ExternalAccountObj).Association("Transactions").Append( &TransactionObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "unassignTransactions", TransactionObj}
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
func RemoveTransactionsFromExternalAccount( externalAccountId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ExternalAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExternalAccount(externalAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExternalAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExternalAccountObj,_ := parentRequestResult.Data. (model.ExternalAccount)

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
				utils.GetDB().Model(&ExternalAccountObj).Association("Transactions").Delete( &TransactionObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "removeTransactions", TransactionObj}
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

