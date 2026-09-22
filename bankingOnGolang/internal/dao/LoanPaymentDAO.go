
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LoanPaymentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLoanPayment - creates a new db entry
//----------------------------------------------------------------------------
func CreateLoanPayment(obj model.LoanPayment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LoanPayment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LoanPayment. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoanPayment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoanPayment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoanPayment(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LoanPayment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LoanPayment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LoanPayment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LoanPayment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLoanPayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLoanPayment - returns all
//----------------------------------------------------------------------------
func GetAllLoanPayment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LoanPayment

	//----------------------------------------------------------------------------
	// Request the ORM to find all LoanPayment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all LoanPayment"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LoanPayment. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLoanPayment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLoanPayment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLoanPayment(obj model.LoanPayment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LoanPayment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LoanPayment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{
        Success: success,
        Msg: updateMsg,
        Call:  "UpdateLoanPayment",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoanPayment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoanPayment(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLoanPayment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.LoanPayment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LoanPayment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LoanPayment using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success: success,
            Msg: deleteMsg,
            Call:  "DeleteLoanPayment",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LoanAccount on a LoanPayment
//----------------------------------------------------------------------------
func AssignLoanAccountToLoanPayment( loanPaymentId uuid.UUID, loanAccountId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanPayment)

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
			// assign the LoanAccount	to the LoanPayment
			//----------------------------------------------------------------------------
			parentObj.LoanAccount = &childObj

			//----------------------------------------------------------------------------
			// save the LoanPayment
			//----------------------------------------------------------------------------
			return UpdateLoanPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignLoanAccount",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LoanAccount on a LoanPayment
//----------------------------------------------------------------------------
func UnassignLoanAccountFromLoanPayment(loanPaymentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanPayment)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccountId = nil;

		//----------------------------------------------------------------------------
		// save the LoanPayment
		//----------------------------------------------------------------------------
		return UpdateLoanPayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a LoanPayment
//----------------------------------------------------------------------------
func AssignTransactionToLoanPayment( loanPaymentId uuid.UUID, transactionId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanPayment)

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
			// assign the Transaction	to the LoanPayment
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the LoanPayment
			//----------------------------------------------------------------------------
			return UpdateLoanPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignTransaction",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a LoanPayment
//----------------------------------------------------------------------------
func UnassignTransactionFromLoanPayment(loanPaymentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanPayment)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the LoanPayment
		//----------------------------------------------------------------------------
		return UpdateLoanPayment(parentObj)

	} else {
		return parentRequestResult
	}

}


