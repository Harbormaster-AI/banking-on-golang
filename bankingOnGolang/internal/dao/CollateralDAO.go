
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CollateralDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCollateral - creates a new db entry
//----------------------------------------------------------------------------
func CreateCollateral(obj model.Collateral)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Collateral with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Collateral. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCollateral", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCollateral - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCollateral(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Collateral

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Collateral with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Collateral using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Collateral using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCollateral", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCollateral - returns all
//----------------------------------------------------------------------------
func GetAllCollateral()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Collateral

	//----------------------------------------------------------------------------
	// Request the ORM to find all Collateral
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all Collateral"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Collateral. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCollateral", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCollateral - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCollateral(obj model.Collateral)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Collateral using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Collateral using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{
        Success: success,
        Message: updateMsg,
        Action:  "UpdateCollateral",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCollateral - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCollateral(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Collateral with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCollateral(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Collateral so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.Collateral)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Collateral using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Collateral using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success: success,
            Message: deleteMsg,
            Action:  "DeleteCollateral",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LoanAccount on a Collateral
//----------------------------------------------------------------------------
func AssignLoanAccountToCollateral( collateralId uuid.UUID, loanAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Collateral with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCollateral(collateralId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Collateral so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Collateral)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LoanAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LoanAccount with a
		// matching loanAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, loanAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LoanAccount	to the Collateral
			//----------------------------------------------------------------------------
			parentObj.LoanAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Collateral
			//----------------------------------------------------------------------------
			return UpdateCollateral(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignLoanAccount",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LoanAccount on a Collateral
//----------------------------------------------------------------------------
func UnassignLoanAccountFromCollateral(collateralId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Collateral with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCollateral(collateralId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Collateral so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Collateral)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Collateral
		//----------------------------------------------------------------------------
		return UpdateCollateral(parentObj)

	} else {
		return parentRequestResult
	}

}


