
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a Dispute. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateDispute",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetDispute - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDispute(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetDispute",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all Dispute"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dispute. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllDispute",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateDispute",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteDispute - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDispute(id uuid.UUID)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data.(model.Dispute)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteDispute",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func AssignTransactionToDispute( disputeId uuid.UUID, transactionId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Transaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Transaction with a
		// matching transactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, transactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Transaction	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignTransaction",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func UnassignTransactionFromDispute(disputeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a Dispute
//----------------------------------------------------------------------------
func AssignCustomerToDispute( disputeId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

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
			// assign the Customer	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
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
// unassigns a Customer on a Dispute
//----------------------------------------------------------------------------
func UnassignCustomerFromDispute(disputeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Dispute
//----------------------------------------------------------------------------
func AssignAccountToDispute( disputeId uuid.UUID, accountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

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
			// assign the Account	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignAccount",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Dispute
//----------------------------------------------------------------------------
func UnassignAccountFromDispute(disputeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PaymentCard on a Dispute
//----------------------------------------------------------------------------
func AssignPaymentCardToDispute( disputeId uuid.UUID, paymentCardId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PaymentCard

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentCard with a
		// matching paymentCardId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, paymentCardId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PaymentCard	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.PaymentCard = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCard", paymentCardId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignPaymentCard",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PaymentCard on a Dispute
//----------------------------------------------------------------------------
func UnassignPaymentCardFromDispute(disputeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the PaymentCard
		//----------------------------------------------------------------------------
		parentObj.PaymentCard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentCard
		//----------------------------------------------------------------------------
		parentObj.PaymentCardId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}


