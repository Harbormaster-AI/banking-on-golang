
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a StandingInstruction. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateStandingInstruction",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetStandingInstruction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStandingInstruction(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetStandingInstruction",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all StandingInstruction"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StandingInstruction. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllStandingInstruction",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateStandingInstruction",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteStandingInstruction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStandingInstruction(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStandingInstruction(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.StandingInstruction)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteStandingInstruction",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a StandingInstruction
//----------------------------------------------------------------------------
func AssignAccountToStandingInstruction( standingInstructionId uuid.UUID, accountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StandingInstruction)

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
			// assign the Account	to the StandingInstruction
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the StandingInstruction
			//----------------------------------------------------------------------------
			return UpdateStandingInstruction(parentObj)
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
// unassigns a Account on a StandingInstruction
//----------------------------------------------------------------------------
func UnassignAccountFromStandingInstruction(standingInstructionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the StandingInstruction
		//----------------------------------------------------------------------------
		return UpdateStandingInstruction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Beneficiary on a StandingInstruction
//----------------------------------------------------------------------------
func AssignBeneficiaryToStandingInstruction( standingInstructionId uuid.UUID, beneficiaryId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ExternalAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ExternalAccount with a
		// matching beneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, beneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Beneficiary	to the StandingInstruction
			//----------------------------------------------------------------------------
			parentObj.Beneficiary = &childObj

			//----------------------------------------------------------------------------
			// save the StandingInstruction
			//----------------------------------------------------------------------------
			return UpdateStandingInstruction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiary", beneficiaryId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignBeneficiary",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Beneficiary on a StandingInstruction
//----------------------------------------------------------------------------
func UnassignBeneficiaryFromStandingInstruction(standingInstructionId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StandingInstruction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStandingInstruction(standingInstructionId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StandingInstruction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StandingInstruction)

		//----------------------------------------------------------------------------
		// assign an empty ExternalAccount to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.Beneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.BeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the StandingInstruction
		//----------------------------------------------------------------------------
		return UpdateStandingInstruction(parentObj)

	} else {
		return parentRequestResult
	}

}


