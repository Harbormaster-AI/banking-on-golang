package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing StandingInstructionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateStandingInstruction - creates a new db entry
//----------------------------------------------------------------------------
func CreateStandingInstruction(obj model.StandingInstruction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a StandingInstruction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a StandingInstruction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateStandingInstruction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetStandingInstruction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStandingInstruction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.StandingInstruction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a StandingInstruction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a StandingInstruction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a StandingInstruction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetStandingInstruction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllStandingInstruction - returns all
//----------------------------------------------------------------------------
func GetAllStandingInstruction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.StandingInstruction

	//----------------------------------------------------------------------------
	// Request the ORM to find all StandingInstruction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all StandingInstruction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StandingInstruction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllStandingInstruction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateStandingInstruction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateStandingInstruction(obj model.StandingInstruction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a StandingInstruction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a StandingInstruction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateStandingInstruction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteStandingInstruction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStandingInstruction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStandingInstruction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a StandingInstruction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a StandingInstruction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteStandingInstruction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a StandingInstruction
//----------------------------------------------------------------------------
func AssignAccountToStandingInstruction( standingInstructionId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		StandingInstructionObj,_ := parentRequestResult.Data. (model.StandingInstruction)

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
			// assign the Account	to the StandingInstruction
			//----------------------------------------------------------------------------
			StandingInstructionObj.Account = &AccountObj

			//----------------------------------------------------------------------------
			// save the StandingInstruction
			//----------------------------------------------------------------------------
			return UpdateStandingInstruction(StandingInstructionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", AccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a StandingInstruction
//----------------------------------------------------------------------------
func UnassignAccountFromStandingInstruction(standingInstructionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		StandingInstructionObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		StandingInstructionObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		StandingInstructionObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the StandingInstruction
		//----------------------------------------------------------------------------
		return UpdateStandingInstruction(StandingInstructionObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Beneficiary on a StandingInstruction
//----------------------------------------------------------------------------
func AssignBeneficiaryToStandingInstruction( standingInstructionId uint64, beneficiaryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		StandingInstructionObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var ExternalAccountObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching beneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&ExternalAccountObj, beneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Beneficiary	to the StandingInstruction
			//----------------------------------------------------------------------------
			StandingInstructionObj.Beneficiary = &ExternalAccountObj

			//----------------------------------------------------------------------------
			// save the StandingInstruction
			//----------------------------------------------------------------------------
			return UpdateStandingInstruction(StandingInstructionObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiary", beneficiaryId )
			return utils.RequestResult{false, msg, "assignBeneficiary", ExternalAccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Beneficiary on a StandingInstruction
//----------------------------------------------------------------------------
func UnassignBeneficiaryFromStandingInstruction(standingInstructionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		StandingInstructionObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the Beneficiary
		//----------------------------------------------------------------------------
		StandingInstructionObj.Beneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Beneficiary
		//----------------------------------------------------------------------------
		StandingInstructionObj.BeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the StandingInstruction
		//----------------------------------------------------------------------------
		return UpdateStandingInstruction(StandingInstructionObj)

	} else {
		return parentRequestResult
	}

}


