
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AccountStatementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAccountStatement - creates a new db entry
//----------------------------------------------------------------------------
func CreateAccountStatement(obj model.AccountStatement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AccountStatement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AccountStatement. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAccountStatement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAccountStatement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAccountStatement(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AccountStatement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AccountStatement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AccountStatement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AccountStatement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAccountStatement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAccountStatement - returns all
//----------------------------------------------------------------------------
func GetAllAccountStatement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AccountStatement

	//----------------------------------------------------------------------------
	// Request the ORM to find all AccountStatement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all AccountStatement"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AccountStatement. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAccountStatement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAccountStatement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAccountStatement(obj model.AccountStatement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AccountStatement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AccountStatement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{
        Success: success,
        Message: updateMsg,
        Action:  "UpdateAccountStatement",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAccountStatement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAccountStatement(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AccountStatement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAccountStatement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccountStatement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.AccountStatement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AccountStatement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AccountStatement using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success: success,
            Message: deleteMsg,
            Action:  "DeleteAccountStatement",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a AccountStatement
//----------------------------------------------------------------------------
func AssignAccountToAccountStatement( accountStatementId uuid.UUID, accountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AccountStatement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccountStatement(accountStatementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccountStatement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccountStatement)

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
			// assign the Account	to the AccountStatement
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the AccountStatement
			//----------------------------------------------------------------------------
			return UpdateAccountStatement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )

            requestResult = utils.RequestResult{
                Success: false,
                Message: msg,
                Action:  "assignAccount",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a AccountStatement
//----------------------------------------------------------------------------
func UnassignAccountFromAccountStatement(accountStatementId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccountStatement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccountStatement(accountStatementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccountStatement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccountStatement)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the AccountStatement
		//----------------------------------------------------------------------------
		return UpdateAccountStatement(parentObj)

	} else {
		return parentRequestResult
	}

}


