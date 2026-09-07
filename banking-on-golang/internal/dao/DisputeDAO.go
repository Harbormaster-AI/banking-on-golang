package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DisputeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDispute - creates a new db entry
//----------------------------------------------------------------------------
func CreateDispute(obj model.Dispute)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Dispute with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Dispute", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDispute", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDispute - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDispute(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Dispute

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Dispute with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Dispute using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Dispute using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDispute", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDispute - returns all
//----------------------------------------------------------------------------
func GetAllDispute()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Dispute

	//----------------------------------------------------------------------------
	// Request the ORM to find all Dispute
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Dispute" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dispute", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDispute", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDispute - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDispute(obj model.Dispute)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Dispute using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Dispute using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDispute", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDispute - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDispute(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDispute(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Dispute using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Dispute using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDispute", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func AssignTransactionToDispute( disputeId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var TransactionObj model.Transaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Transaction with a
		// matching transactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&TransactionObj, transactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Transaction	to the Dispute
			//----------------------------------------------------------------------------
			DisputeObj.Transaction = &TransactionObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(DisputeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", TransactionObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func UnassignTransactionFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		DisputeObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		DisputeObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(DisputeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a Dispute
//----------------------------------------------------------------------------
func AssignCustomerToDispute( disputeId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

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
			// assign the Customer	to the Dispute
			//----------------------------------------------------------------------------
			DisputeObj.Customer = &CustomerObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(DisputeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", CustomerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Dispute
//----------------------------------------------------------------------------
func UnassignCustomerFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		DisputeObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		DisputeObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(DisputeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Dispute
//----------------------------------------------------------------------------
func AssignAccountToDispute( disputeId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var AccountObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&AccountObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the Dispute
			//----------------------------------------------------------------------------
			DisputeObj.Account = &AccountObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(DisputeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Dispute
//----------------------------------------------------------------------------
func UnassignAccountFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		DisputeObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		DisputeObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(DisputeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PaymentCard on a Dispute
//----------------------------------------------------------------------------
func AssignPaymentCardToDispute( disputeId uint64, paymentCardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var PaymentCardObj model.PaymentCard

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentCard with a
		// matching paymentCardId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&PaymentCardObj, paymentCardId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PaymentCard	to the Dispute
			//----------------------------------------------------------------------------
			DisputeObj.PaymentCard = &PaymentCardObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(DisputeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCard", paymentCardId )
			return utils.RequestResult{false, msg, "assignPaymentCard", PaymentCardObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PaymentCard on a Dispute
//----------------------------------------------------------------------------
func UnassignPaymentCardFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		DisputeObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the PaymentCard
		//----------------------------------------------------------------------------
		DisputeObj.PaymentCard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentCard
		//----------------------------------------------------------------------------
		DisputeObj.PaymentCardId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(DisputeObj)

	} else {
		return parentRequestResult
	}

}


