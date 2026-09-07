package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
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
		createMsg = fmt.Sprintf( "Failed trying to create a LoanPayment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoanPayment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoanPayment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoanPayment(id uint64)(utils.RequestResult){
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
	    getAllMsg = fmt.Sprintf( "Retrieved all LoanPayment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LoanPayment", result )
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLoanPayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoanPayment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoanPayment(id uint64)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data. (model.LoanPayment)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLoanPayment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LoanAccount on a LoanPayment
//----------------------------------------------------------------------------
func AssignLoanAccountToLoanPayment( loanPaymentId uint64, loanAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanPaymentObj,_ := parentRequestResult.Data. (model.LoanPayment)

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
			// assign the LoanAccount	to the LoanPayment
			//----------------------------------------------------------------------------
			LoanPaymentObj.LoanAccount = &LoanAccountObj

			//----------------------------------------------------------------------------
			// save the LoanPayment
			//----------------------------------------------------------------------------
			return UpdateLoanPayment(LoanPaymentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId )
			return utils.RequestResult{false, msg, "assignLoanAccount", LoanAccountObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LoanAccount on a LoanPayment
//----------------------------------------------------------------------------
func UnassignLoanAccountFromLoanPayment(loanPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanPaymentObj,_ := parentRequestResult.Data. (model.LoanPayment)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		LoanPaymentObj.LoanAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		LoanPaymentObj.LoanAccountId = nil;

		//----------------------------------------------------------------------------
		// save the LoanPayment
		//----------------------------------------------------------------------------
		return UpdateLoanPayment(LoanPaymentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a LoanPayment
//----------------------------------------------------------------------------
func AssignTransactionToLoanPayment( loanPaymentId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanPaymentObj,_ := parentRequestResult.Data. (model.LoanPayment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var TransactionObj model.Transaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Transaction with a
		// matching transactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&TransactionObj, transactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Transaction	to the LoanPayment
			//----------------------------------------------------------------------------
			LoanPaymentObj.Transaction = &TransactionObj

			//----------------------------------------------------------------------------
			// save the LoanPayment
			//----------------------------------------------------------------------------
			return UpdateLoanPayment(LoanPaymentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", TransactionObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a LoanPayment
//----------------------------------------------------------------------------
func UnassignTransactionFromLoanPayment(loanPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanPayment(loanPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanPaymentObj,_ := parentRequestResult.Data. (model.LoanPayment)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		LoanPaymentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		LoanPaymentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the LoanPayment
		//----------------------------------------------------------------------------
		return UpdateLoanPayment(LoanPaymentObj)

	} else {
		return parentRequestResult
	}

}


