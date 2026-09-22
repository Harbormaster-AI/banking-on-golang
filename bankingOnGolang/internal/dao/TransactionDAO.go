
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TransactionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTransaction - creates a new db entry
//----------------------------------------------------------------------------
func CreateTransaction(obj model.Transaction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Transaction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Transaction. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateTransaction",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetTransaction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTransaction(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Transaction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Transaction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Transaction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Transaction using ID=%v", id )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetTransaction",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// GetAllTransaction - returns all
//----------------------------------------------------------------------------
func GetAllTransaction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Transaction

	//----------------------------------------------------------------------------
	// Request the ORM to find all Transaction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all Transaction"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Transaction. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllTransaction",
        Data:       objs,
    }

}

//----------------------------------------------------------------------------
// UpdateTransaction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTransaction(obj model.Transaction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Transaction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Transaction using ID=%v", obj.ID )
		success = false
	}

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateTransaction",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteTransaction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTransaction(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTransaction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.Transaction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Transaction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Transaction using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteTransaction",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a Transaction
//----------------------------------------------------------------------------
func AssignAccountToTransaction( transactionId uuid.UUID, accountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the Account	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
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
// unassigns a Account on a Transaction
//----------------------------------------------------------------------------
func UnassignAccountFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExternalCounterparty on a Transaction
//----------------------------------------------------------------------------
func AssignExternalCounterpartyToTransaction( transactionId uuid.UUID, externalCounterpartyId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching externalCounterpartyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, externalCounterpartyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExternalCounterparty	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.ExternalCounterparty = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalCounterparty", externalCounterpartyId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignExternalCounterparty",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExternalCounterparty on a Transaction
//----------------------------------------------------------------------------
func UnassignExternalCounterpartyFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the ExternalCounterparty
		//----------------------------------------------------------------------------
		parentObj.ExternalCounterparty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExternalCounterparty
		//----------------------------------------------------------------------------
		parentObj.ExternalCounterpartyId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PaymentCard on a Transaction
//----------------------------------------------------------------------------
func AssignPaymentCardToTransaction( transactionId uuid.UUID, paymentCardId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the PaymentCard	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.PaymentCard = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
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
// unassigns a PaymentCard on a Transaction
//----------------------------------------------------------------------------
func UnassignPaymentCardFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the PaymentCard
		//----------------------------------------------------------------------------
		parentObj.PaymentCard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentCard
		//----------------------------------------------------------------------------
		parentObj.PaymentCardId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FundsTransfer on a Transaction
//----------------------------------------------------------------------------
func AssignFundsTransferToTransaction( transactionId uuid.UUID, fundsTransferId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FundsTransfer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FundsTransfer with a
		// matching fundsTransferId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, fundsTransferId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FundsTransfer	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.FundsTransfer = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfer", fundsTransferId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignFundsTransfer",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FundsTransfer on a Transaction
//----------------------------------------------------------------------------
func UnassignFundsTransferFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty FundsTransfer to the FundsTransfer
		//----------------------------------------------------------------------------
		parentObj.FundsTransfer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FundsTransfer
		//----------------------------------------------------------------------------
		parentObj.FundsTransferId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FxTrade on a Transaction
//----------------------------------------------------------------------------
func AssignFxTradeToTransaction( transactionId uuid.UUID, fxTradeId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FXTrade

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FXTrade with a
		// matching fxTradeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, fxTradeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FxTrade	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.FxTrade = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrade", fxTradeId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignFxTrade",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FxTrade on a Transaction
//----------------------------------------------------------------------------
func UnassignFxTradeFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty FXTrade to the FxTrade
		//----------------------------------------------------------------------------
		parentObj.FxTrade = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FxTrade
		//----------------------------------------------------------------------------
		parentObj.FxTradeId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dispute on a Transaction
//----------------------------------------------------------------------------
func AssignDisputeToTransaction( transactionId uuid.UUID, disputeId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Dispute

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Dispute with a
		// matching disputeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, disputeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dispute	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Dispute = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dispute", disputeId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignDispute",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dispute on a Transaction
//----------------------------------------------------------------------------
func UnassignDisputeFromTransaction(transactionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Dispute to the Dispute
		//----------------------------------------------------------------------------
		parentObj.Dispute = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dispute
		//----------------------------------------------------------------------------
		parentObj.DisputeId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}


