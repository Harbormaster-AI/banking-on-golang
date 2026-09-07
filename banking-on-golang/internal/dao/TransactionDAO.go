package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
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
	var requestResult utils.RequestResult
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
		createMsg = fmt.Sprintf( "Failed trying to create a Transaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTransaction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTransaction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTransaction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
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

	requestResult = utils.RequestResult{success, getMsg, "GetTransaction", obj}

	return requestResult
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
	    getAllMsg = fmt.Sprintf( "Retrieved all Transaction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Transaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTransaction", objs}
	return requestResult
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTransaction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTransaction(id uint64)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data. (model.Transaction)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTransaction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a Transaction
//----------------------------------------------------------------------------
func AssignAccountToTransaction( transactionId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the Account	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.Account = &AccountObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Transaction
//----------------------------------------------------------------------------
func UnassignAccountFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		TransactionObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		TransactionObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExternalCounterparty on a Transaction
//----------------------------------------------------------------------------
func AssignExternalCounterpartyToTransaction( transactionId uint64, externalCounterpartyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var ExternalAccountObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching externalCounterpartyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&ExternalAccountObj, externalCounterpartyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExternalCounterparty	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.ExternalCounterparty = &ExternalAccountObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalCounterparty", externalCounterpartyId )
			return utils.RequestResult{false, msg, "assignExternalCounterparty", ExternalAccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExternalCounterparty on a Transaction
//----------------------------------------------------------------------------
func UnassignExternalCounterpartyFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the ExternalCounterparty
		//----------------------------------------------------------------------------
		TransactionObj.ExternalCounterparty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExternalCounterparty
		//----------------------------------------------------------------------------
		TransactionObj.ExternalCounterpartyId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PaymentCard on a Transaction
//----------------------------------------------------------------------------
func AssignPaymentCardToTransaction( transactionId uint64, paymentCardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the PaymentCard	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.PaymentCard = &PaymentCardObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCard", paymentCardId )
			return utils.RequestResult{false, msg, "assignPaymentCard", PaymentCardObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PaymentCard on a Transaction
//----------------------------------------------------------------------------
func UnassignPaymentCardFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the PaymentCard
		//----------------------------------------------------------------------------
		TransactionObj.PaymentCard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentCard
		//----------------------------------------------------------------------------
		TransactionObj.PaymentCardId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FundsTransfer on a Transaction
//----------------------------------------------------------------------------
func AssignFundsTransferToTransaction( transactionId uint64, fundsTransferId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var FundsTransferObj model.FundsTransfer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FundsTransfer with a
		// matching fundsTransferId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&FundsTransferObj, fundsTransferId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FundsTransfer	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.FundsTransfer = &FundsTransferObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfer", fundsTransferId )
			return utils.RequestResult{false, msg, "assignFundsTransfer", FundsTransferObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FundsTransfer on a Transaction
//----------------------------------------------------------------------------
func UnassignFundsTransferFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty FundsTransfer to the FundsTransfer
		//----------------------------------------------------------------------------
		TransactionObj.FundsTransfer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FundsTransfer
		//----------------------------------------------------------------------------
		TransactionObj.FundsTransferId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FxTrade on a Transaction
//----------------------------------------------------------------------------
func AssignFxTradeToTransaction( transactionId uint64, fxTradeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var FXTradeObj model.FXTrade

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FXTrade with a
		// matching fxTradeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&FXTradeObj, fxTradeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FxTrade	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.FxTrade = &FXTradeObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrade", fxTradeId )
			return utils.RequestResult{false, msg, "assignFxTrade", FXTradeObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FxTrade on a Transaction
//----------------------------------------------------------------------------
func UnassignFxTradeFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty FXTrade to the FxTrade
		//----------------------------------------------------------------------------
		TransactionObj.FxTrade = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FxTrade
		//----------------------------------------------------------------------------
		TransactionObj.FxTradeId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dispute on a Transaction
//----------------------------------------------------------------------------
func AssignDisputeToTransaction( transactionId uint64, disputeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var DisputeObj model.Dispute

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Dispute with a
		// matching disputeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&DisputeObj, disputeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dispute	to the Transaction
			//----------------------------------------------------------------------------
			TransactionObj.Dispute = &DisputeObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(TransactionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dispute", disputeId )
			return utils.RequestResult{false, msg, "assignDispute", DisputeObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dispute on a Transaction
//----------------------------------------------------------------------------
func UnassignDisputeFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TransactionObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Dispute to the Dispute
		//----------------------------------------------------------------------------
		TransactionObj.Dispute = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dispute
		//----------------------------------------------------------------------------
		TransactionObj.DisputeId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(TransactionObj)

	} else {
		return parentRequestResult
	}

}


