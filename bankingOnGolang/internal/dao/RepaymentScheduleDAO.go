
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RepaymentScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRepaymentSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateRepaymentSchedule(obj model.RepaymentSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RepaymentSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RepaymentSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRepaymentSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRepaymentSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRepaymentSchedule(id UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RepaymentSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RepaymentSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RepaymentSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRepaymentSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRepaymentSchedule - returns all
//----------------------------------------------------------------------------
func GetAllRepaymentSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RepaymentSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all RepaymentSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RepaymentSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RepaymentSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRepaymentSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRepaymentSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRepaymentSchedule(obj model.RepaymentSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RepaymentSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RepaymentSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRepaymentSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRepaymentSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRepaymentSchedule(id UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRepaymentSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RepaymentSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RepaymentSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRepaymentSchedule", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LoanAccount on a RepaymentSchedule
//----------------------------------------------------------------------------
func AssignLoanAccountToRepaymentSchedule( repaymentScheduleId UUID, loanAccountId UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

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
			// assign the LoanAccount	to the RepaymentSchedule
			//----------------------------------------------------------------------------
			parentObj.LoanAccount = &childObj

			//----------------------------------------------------------------------------
			// save the RepaymentSchedule
			//----------------------------------------------------------------------------
			return UpdateRepaymentSchedule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId )
			return utils.RequestResult{false, msg, "assignLoanAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LoanAccount on a RepaymentSchedule
//----------------------------------------------------------------------------
func UnassignLoanAccountFromRepaymentSchedule(repaymentScheduleId UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccountId = nil;

		//----------------------------------------------------------------------------
		// save the RepaymentSchedule
		//----------------------------------------------------------------------------
		return UpdateRepaymentSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Payment on a RepaymentSchedule
//----------------------------------------------------------------------------
func AssignPaymentToRepaymentSchedule( repaymentScheduleId UUID, paymentId UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LoanPayment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LoanPayment with a
		// matching paymentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, paymentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Payment	to the RepaymentSchedule
			//----------------------------------------------------------------------------
			parentObj.Payment = &childObj

			//----------------------------------------------------------------------------
			// save the RepaymentSchedule
			//----------------------------------------------------------------------------
			return UpdateRepaymentSchedule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payment", paymentId )
			return utils.RequestResult{false, msg, "assignPayment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Payment on a RepaymentSchedule
//----------------------------------------------------------------------------
func UnassignPaymentFromRepaymentSchedule(repaymentScheduleId UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// assign an empty LoanPayment to the Payment
		//----------------------------------------------------------------------------
		parentObj.Payment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Payment
		//----------------------------------------------------------------------------
		parentObj.PaymentId = nil;

		//----------------------------------------------------------------------------
		// save the RepaymentSchedule
		//----------------------------------------------------------------------------
		return UpdateRepaymentSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}


