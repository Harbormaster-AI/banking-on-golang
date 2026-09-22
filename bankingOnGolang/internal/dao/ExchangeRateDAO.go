
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a ExchangeRate. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateExchangeRate",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetExchangeRate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExchangeRate(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetExchangeRate",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all ExchangeRate"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExchangeRate. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllExchangeRate",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateExchangeRate",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteExchangeRate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExchangeRate(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExchangeRate(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.ExchangeRate)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteExchangeRate",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a ExchangeRate
//----------------------------------------------------------------------------
func AssignBankToExchangeRate( exchangeRateId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

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
			// assign the Bank	to the ExchangeRate
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the ExchangeRate
			//----------------------------------------------------------------------------
			return UpdateExchangeRate(parentObj)
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
// unassigns a Bank on a ExchangeRate
//----------------------------------------------------------------------------
func UnassignBankFromExchangeRate(exchangeRateId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the ExchangeRate
		//----------------------------------------------------------------------------
		return UpdateExchangeRate(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more fxTradesIds as a FxTrades to a ExchangeRate
//----------------------------------------------------------------------------
func AddFxTradesToExchangeRate ( exchangeRateId uuid.UUID, fxTradesIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		for _, fxTradesId:= range fxTradesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FXTrade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXTrade
			// with a matching fxTradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , fxTradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FxTrades using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FxTrades").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrades", fxTradesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addFxTradesToExchangeRate",
                    Data:       childObj,
                }
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
func RemoveFxTradesFromExchangeRate( exchangeRateId uuid.UUID, fxTradesIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		for _, fxTradesId:= range fxTradesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FXTrade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXTrade
			// with a matching fxTradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , fxTradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FXTradeObj from the FxTrades array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FxTrades").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxTrades", fxTradesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeFxTradesFromExchangeRate",
                    Data:       childObj,
                }
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

