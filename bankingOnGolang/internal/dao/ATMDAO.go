
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ATMDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateATM - creates a new db entry
//----------------------------------------------------------------------------
func CreateATM(obj model.ATM)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ATM with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ATM. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateATM",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetATM - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetATM(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ATM

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ATM with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ATM using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ATM using ID=%v", id )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetATM",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// GetAllATM - returns all
//----------------------------------------------------------------------------
func GetAllATM()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ATM

	//----------------------------------------------------------------------------
	// Request the ORM to find all ATM
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all ATM"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ATM. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllATM",
        Data:       objs,
    }

}

//----------------------------------------------------------------------------
// UpdateATM - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateATM(obj model.ATM)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ATM using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ATM using ID=%v", obj.ID )
		success = false
	}

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateATM",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteATM - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteATM(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetATM(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.ATM)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ATM using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ATM using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteATM",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Branch on a ATM
//----------------------------------------------------------------------------
func AssignBranchToATM( aTMId uuid.UUID, branchId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetATM(aTMId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ATM)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Branch

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Branch with a
		// matching branchId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, branchId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Branch	to the ATM
			//----------------------------------------------------------------------------
			parentObj.Branch = &childObj

			//----------------------------------------------------------------------------
			// save the ATM
			//----------------------------------------------------------------------------
			return UpdateATM(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branch", branchId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignBranch",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Branch on a ATM
//----------------------------------------------------------------------------
func UnassignBranchFromATM(aTMId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetATM(aTMId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ATM)

		//----------------------------------------------------------------------------
		// assign an empty Branch to the Branch
		//----------------------------------------------------------------------------
		parentObj.Branch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Branch
		//----------------------------------------------------------------------------
		parentObj.BranchId = nil;

		//----------------------------------------------------------------------------
		// save the ATM
		//----------------------------------------------------------------------------
		return UpdateATM(parentObj)

	} else {
		return parentRequestResult
	}

}


