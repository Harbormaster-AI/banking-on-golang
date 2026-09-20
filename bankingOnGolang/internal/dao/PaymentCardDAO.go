
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PaymentCardDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePaymentCard - creates a new db entry
//----------------------------------------------------------------------------
func CreatePaymentCard(obj model.PaymentCard)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PaymentCard with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PaymentCard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePaymentCard", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPaymentCard - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPaymentCard(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PaymentCard

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PaymentCard with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PaymentCard using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PaymentCard using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPaymentCard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPaymentCard - returns all
//----------------------------------------------------------------------------
func GetAllPaymentCard()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PaymentCard

	//----------------------------------------------------------------------------
	// Request the ORM to find all PaymentCard
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PaymentCard" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PaymentCard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPaymentCard", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePaymentCard - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePaymentCard(obj model.PaymentCard)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PaymentCard using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PaymentCard using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePaymentCard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePaymentCard - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePaymentCard(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPaymentCard(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PaymentCard)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PaymentCard using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PaymentCard using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePaymentCard", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a PaymentCard
//----------------------------------------------------------------------------
func AssignBankToPaymentCard( paymentCardId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

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
			// assign the Bank	to the PaymentCard
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentCard
			//----------------------------------------------------------------------------
			return UpdatePaymentCard(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a PaymentCard
//----------------------------------------------------------------------------
func UnassignBankFromPaymentCard(paymentCardId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentCard
		//----------------------------------------------------------------------------
		return UpdatePaymentCard(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a PaymentCard
//----------------------------------------------------------------------------
func AssignAccountToPaymentCard( paymentCardId uuid.UUID, accountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the PaymentCard
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentCard
			//----------------------------------------------------------------------------
			return UpdatePaymentCard(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a PaymentCard
//----------------------------------------------------------------------------
func UnassignAccountFromPaymentCard(paymentCardId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentCard
		//----------------------------------------------------------------------------
		return UpdatePaymentCard(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a PaymentCard
//----------------------------------------------------------------------------
func AssignCustomerToPaymentCard( paymentCardId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

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
			// assign the Customer	to the PaymentCard
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentCard
			//----------------------------------------------------------------------------
			return UpdatePaymentCard(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a PaymentCard
//----------------------------------------------------------------------------
func UnassignCustomerFromPaymentCard(paymentCardId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentCard
		//----------------------------------------------------------------------------
		return UpdatePaymentCard(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a PaymentCard
//----------------------------------------------------------------------------
func AddTransactionsToPaymentCard ( paymentCardId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

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
				return utils.RequestResult{false, msg, "unassignTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentCard from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentCard(paymentCardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a PaymentCard
//----------------------------------------------------------------------------
func RemoveTransactionsFromPaymentCard( paymentCardId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentCard(paymentCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentCard)

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
				return utils.RequestResult{false, msg, "removeTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentCard from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentCard(paymentCardId)

	} else {
		return parentRequestResult
	}
}

