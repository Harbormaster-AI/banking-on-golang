package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExchangeRateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExchangeRate - creates a new db entry
//----------------------------------------------------------------------------
func CreateExchangeRate(obj model.ExchangeRate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ExchangeRate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ExchangeRate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExchangeRate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExchangeRate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExchangeRate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ExchangeRate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ExchangeRate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ExchangeRate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ExchangeRate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExchangeRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExchangeRate - returns all
//----------------------------------------------------------------------------
func GetAllExchangeRate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ExchangeRate

	//----------------------------------------------------------------------------
	// Request the ORM to find all ExchangeRate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ExchangeRate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExchangeRate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExchangeRate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExchangeRate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExchangeRate(obj model.ExchangeRate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ExchangeRate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ExchangeRate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExchangeRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExchangeRate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExchangeRate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExchangeRate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ExchangeRate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ExchangeRate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ExchangeRate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExchangeRate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a ExchangeRate
//----------------------------------------------------------------------------
func AssignBankToExchangeRate( exchangeRateId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExchangeRateObj,_ := parentRequestResult.Data. (model.ExchangeRate)

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
			// assign the Bank	to the ExchangeRate
			//----------------------------------------------------------------------------
			ExchangeRateObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the ExchangeRate
			//----------------------------------------------------------------------------
			return UpdateExchangeRate(ExchangeRateObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a ExchangeRate
//----------------------------------------------------------------------------
func UnassignBankFromExchangeRate(exchangeRateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExchangeRateObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		ExchangeRateObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		ExchangeRateObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the ExchangeRate
		//----------------------------------------------------------------------------
		return UpdateExchangeRate(ExchangeRateObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more fxTradesIds as a FxTrades to a ExchangeRate
//----------------------------------------------------------------------------
func AddFxTradesToExchangeRate ( exchangeRateId uint64, fxTradesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExchangeRateObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		// slice the ids on comma with no spaces
		ids := strings.Split( fxTradesIds, ",")

		for _, fxTradesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FXTradeObj model.FXTrade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXTrade
			// with a matching fxTradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FXTradeObj , fxTradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FxTrades using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ExchangeRateObj).Association("FxTrades").Append( &FXTradeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrades", fxTradesId )
				return utils.RequestResult{false, msg, "unassignFxTrades", FXTradeObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExchangeRate from the gorm
		//----------------------------------------------------------------------------
		return GetExchangeRate(exchangeRateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more fxTradesIds as a FxTrades from a ExchangeRate
//----------------------------------------------------------------------------
func RemoveFxTradesFromExchangeRate( exchangeRateId uint64, fxTradesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ExchangeRateObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		// slice the ids on comma with no spaces
		ids := strings.Split( fxTradesIds, ",")

		for _, fxTradesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FXTradeObj model.FXTrade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXTrade
			// with a matching fxTradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FXTradeObj , fxTradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FXTradeObj from the FxTrades array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&ExchangeRateObj).Association("FxTrades").Delete( &FXTradeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrades", fxTradesId )
				return utils.RequestResult{false, msg, "removeFxTrades", FXTradeObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExchangeRate from the gorm
		//----------------------------------------------------------------------------
		return GetExchangeRate(exchangeRateId)

	} else {
		return parentRequestResult
	}
}

