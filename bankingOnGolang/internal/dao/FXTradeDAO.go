
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a FXTrade. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateFXTrade",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetFXTrade - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFXTrade(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetFXTrade",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all FXTrade"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FXTrade. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllFXTrade",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateFXTrade",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteFXTrade - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFXTrade(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFXTrade(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.FXTrade)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteFXTrade",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a FXTrade
//----------------------------------------------------------------------------
func AssignCustomerToFXTrade( fXTradeId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the Customer	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
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
// unassigns a Customer on a FXTrade
//----------------------------------------------------------------------------
func UnassignCustomerFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Bank on a FXTrade
//----------------------------------------------------------------------------
func AssignBankToFXTrade( fXTradeId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the Bank	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignBank",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a FXTrade
//----------------------------------------------------------------------------
func UnassignBankFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ExchangeRate on a FXTrade
//----------------------------------------------------------------------------
func AssignExchangeRateToFXTrade( fXTradeId uuid.UUID, exchangeRateId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ExchangeRate

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExchangeRate with a
		// matching exchangeRateId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, exchangeRateId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ExchangeRate	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.ExchangeRate = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRate", exchangeRateId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignExchangeRate",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ExchangeRate on a FXTrade
//----------------------------------------------------------------------------
func UnassignExchangeRateFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty ExchangeRate to the ExchangeRate
		//----------------------------------------------------------------------------
		parentObj.ExchangeRate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ExchangeRate
		//----------------------------------------------------------------------------
		parentObj.ExchangeRateId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SourceAccount on a FXTrade
//----------------------------------------------------------------------------
func AssignSourceAccountToFXTrade( fXTradeId uuid.UUID, sourceAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the SourceAccount	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.SourceAccount = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceAccount", sourceAccountId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignSourceAccount",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceAccount on a FXTrade
//----------------------------------------------------------------------------
func UnassignSourceAccountFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Account to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a FXTrade
//----------------------------------------------------------------------------
func AssignDestinationAccountToFXTrade( fXTradeId uuid.UUID, destinationAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the DestinationAccount	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.DestinationAccount = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignDestinationAccount",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a FXTrade
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a FXTrade
//----------------------------------------------------------------------------
func AssignTransactionToFXTrade( fXTradeId uuid.UUID, transactionId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

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
			// assign the Transaction	to the FXTrade
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the FXTrade
			//----------------------------------------------------------------------------
			return UpdateFXTrade(parentObj)
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
// unassigns a Transaction on a FXTrade
//----------------------------------------------------------------------------
func UnassignTransactionFromFXTrade(fXTradeId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXTrade with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXTrade(fXTradeId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXTrade so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXTrade)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the FXTrade
		//----------------------------------------------------------------------------
		return UpdateFXTrade(parentObj)

	} else {
		return parentRequestResult
	}

}


