package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
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
	var requestResult utils.RequestResult
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
		createMsg = fmt.Sprintf( "Failed trying to create a ATM", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateATM", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetATM - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetATM(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
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

	requestResult = utils.RequestResult{success, getMsg, "GetATM", obj}

	return requestResult
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
	    getAllMsg = fmt.Sprintf( "Retrieved all ATM" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ATM", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllATM", objs}
	return requestResult
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateATM", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteATM - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteATM(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetATM(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ATM)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteATM", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Branch on a ATM
//----------------------------------------------------------------------------
func AssignBranchToATM( aTMId uint64, branchId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetATM(aTMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ATMObj,_ := parentRequestResult.Data. (model.ATM)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var BranchObj model.Branch

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Branch with a
		// matching branchId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&BranchObj, branchId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Branch	to the ATM
			//----------------------------------------------------------------------------
			ATMObj.Branch = &BranchObj

			//----------------------------------------------------------------------------
			// save the ATM
			//----------------------------------------------------------------------------
			return UpdateATM(ATMObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branch", branchId )
			return utils.RequestResult{false, msg, "assignBranch", BranchObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Branch on a ATM
//----------------------------------------------------------------------------
func UnassignBranchFromATM(aTMId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ATM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetATM(aTMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ATM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ATMObj,_ := parentRequestResult.Data. (model.ATM)

		//----------------------------------------------------------------------------
		// assign an empty Branch to the Branch
		//----------------------------------------------------------------------------
		ATMObj.Branch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Branch
		//----------------------------------------------------------------------------
		ATMObj.BranchId = nil;

		//----------------------------------------------------------------------------
		// save the ATM
		//----------------------------------------------------------------------------
		return UpdateATM(ATMObj)

	} else {
		return parentRequestResult
	}

}


