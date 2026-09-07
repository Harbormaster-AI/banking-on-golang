package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FXTradeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFXTrade - creates a new db entry
//----------------------------------------------------------------------------
func CreateFXTrade(obj model.FXTrade)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FXTrade with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FXTrade", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFXTrade", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFXTrade - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFXTrade(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FXTrade

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FXTrade with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FXTrade using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FXTrade using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFXTrade", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFXTrade - returns all
//----------------------------------------------------------------------------
func GetAllFXTrade()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FXTrade

	//----------------------------------------------------------------------------
	// Request the ORM to find all FXTrade
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FXTrade" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FXTrade", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFXTrade", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFXTrade - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFXTrade(obj model.FXTrade)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FXTrade using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FXTrade using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFXTrade", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFXTrade - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFXTrade(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFXTrade(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FXTrade using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FXTrade using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFXTrade", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a FXTrade
//----------------------------------------------------------------------------
func AssignCustomerToFXTrade( fXTradeId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the Customer	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.Customer = &CustomerObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", CustomerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a FXTrade
//----------------------------------------------------------------------------
func UnassignCustomerFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		FXTradeObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		FXTradeObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Bank on a FXTrade
//----------------------------------------------------------------------------
func AssignBankToFXTrade( fXTradeId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var BankObj model.Bank

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Bank with a
		// matching bankId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&BankObj, bankId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bank	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a FXTrade
//----------------------------------------------------------------------------
func UnassignBankFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		FXTradeObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		FXTradeObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExchangeRate on a FXTrade
//----------------------------------------------------------------------------
func AssignExchangeRateToFXTrade( fXTradeId uint64, exchangeRateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var ExchangeRateObj model.ExchangeRate

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExchangeRate with a
		// matching exchangeRateId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&ExchangeRateObj, exchangeRateId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExchangeRate	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.ExchangeRate = &ExchangeRateObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRate", exchangeRateId )
			return utils.RequestResult{false, msg, "assignExchangeRate", ExchangeRateObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExchangeRate on a FXTrade
//----------------------------------------------------------------------------
func UnassignExchangeRateFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty ExchangeRate to the ExchangeRate
		//----------------------------------------------------------------------------
		FXTradeObj.ExchangeRate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExchangeRate
		//----------------------------------------------------------------------------
		FXTradeObj.ExchangeRateId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SourceAccount on a FXTrade
//----------------------------------------------------------------------------
func AssignSourceAccountToFXTrade( fXTradeId uint64, sourceAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the SourceAccount	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.SourceAccount = &AccountObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceAccount", sourceAccountId )
			return utils.RequestResult{false, msg, "assignSourceAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceAccount on a FXTrade
//----------------------------------------------------------------------------
func UnassignSourceAccountFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Account to the SourceAccount
		//----------------------------------------------------------------------------
		FXTradeObj.SourceAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceAccount
		//----------------------------------------------------------------------------
		FXTradeObj.SourceAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a FXTrade
//----------------------------------------------------------------------------
func AssignDestinationAccountToFXTrade( fXTradeId uint64, destinationAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the DestinationAccount	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.DestinationAccount = &AccountObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )
			return utils.RequestResult{false, msg, "assignDestinationAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a FXTrade
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		FXTradeObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		FXTradeObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a FXTrade
//----------------------------------------------------------------------------
func AssignTransactionToFXTrade( fXTradeId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the Transaction	to the FXTrade
			//----------------------------------------------------------------------------
			FXTradeObj.Transaction = &TransactionObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(FXTradeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", TransactionObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a FXTrade
//----------------------------------------------------------------------------
func UnassignTransactionFromFXTrade(fXTradeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FXTradeObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		FXTradeObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		FXTradeObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(FXTradeObj)

	} else {
		return parentRequestResult
	}

}


