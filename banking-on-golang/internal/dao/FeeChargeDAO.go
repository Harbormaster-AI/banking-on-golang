package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FeeChargeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFeeCharge - creates a new db entry
//----------------------------------------------------------------------------
func CreateFeeCharge(obj model.FeeCharge)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FeeCharge with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FeeCharge", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFeeCharge", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFeeCharge - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFeeCharge(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FeeCharge

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FeeCharge with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FeeCharge using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FeeCharge using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFeeCharge", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFeeCharge - returns all
//----------------------------------------------------------------------------
func GetAllFeeCharge()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FeeCharge

	//----------------------------------------------------------------------------
	// Request the ORM to find all FeeCharge
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FeeCharge" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FeeCharge", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFeeCharge", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFeeCharge - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFeeCharge(obj model.FeeCharge)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FeeCharge using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FeeCharge using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFeeCharge", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFeeCharge - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFeeCharge(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FeeCharge with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFeeCharge(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeCharge so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FeeCharge)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FeeCharge using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FeeCharge using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFeeCharge", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a FeeCharge
//----------------------------------------------------------------------------
func AssignAccountToFeeCharge( feeChargeId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FeeCharge with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeCharge(feeChargeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeCharge so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FeeChargeObj,_ := parentRequestResult.Data. (model.FeeCharge)

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
			// assign the Account	to the FeeCharge
			//----------------------------------------------------------------------------
			FeeChargeObj.Account = &AccountObj

			//----------------------------------------------------------------------------
			// save the FeeCharge
			//----------------------------------------------------------------------------
			return UpdateFeeCharge(FeeChargeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a FeeCharge
//----------------------------------------------------------------------------
func UnassignAccountFromFeeCharge(feeChargeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeeCharge with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeCharge(feeChargeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeCharge so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FeeChargeObj,_ := parentRequestResult.Data. (model.FeeCharge)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		FeeChargeObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		FeeChargeObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the FeeCharge
		//----------------------------------------------------------------------------
		return UpdateFeeCharge(FeeChargeObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LoanAccount on a FeeCharge
//----------------------------------------------------------------------------
func AssignLoanAccountToFeeCharge( feeChargeId uint64, loanAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FeeCharge with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeCharge(feeChargeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeCharge so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FeeChargeObj,_ := parentRequestResult.Data. (model.FeeCharge)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var LoanAccountObj model.LoanAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LoanAccount with a
		// matching loanAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&LoanAccountObj, loanAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LoanAccount	to the FeeCharge
			//----------------------------------------------------------------------------
			FeeChargeObj.LoanAccount = &LoanAccountObj

			//----------------------------------------------------------------------------
			// save the FeeCharge
			//----------------------------------------------------------------------------
			return UpdateFeeCharge(FeeChargeObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId )
			return utils.RequestResult{false, msg, "assignLoanAccount", LoanAccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LoanAccount on a FeeCharge
//----------------------------------------------------------------------------
func UnassignLoanAccountFromFeeCharge(feeChargeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeeCharge with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeCharge(feeChargeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeCharge so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		FeeChargeObj,_ := parentRequestResult.Data. (model.FeeCharge)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		FeeChargeObj.LoanAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		FeeChargeObj.LoanAccountId = nil;

		//----------------------------------------------------------------------------
		// save the FeeCharge
		//----------------------------------------------------------------------------
		return UpdateFeeCharge(FeeChargeObj)

	} else {
		return parentRequestResult
	}

}


