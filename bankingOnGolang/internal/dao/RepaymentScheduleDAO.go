package dao

import (
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func init() {
	fmt.Println(strings.ToTitle("Initializing RepaymentScheduleDAO..."))
}

// ----------------------------------------------------------------------------
// CreateRepaymentSchedule - creates a new db entry
// ----------------------------------------------------------------------------
func CreateRepaymentSchedule(obj model.RepaymentSchedule) utils.RequestResult {
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
		createMsg = fmt.Sprintf("Created a RepaymentSchedule with ID=%v", obj.ID)
		success = true
	} else {
		createMsg = fmt.Sprintf("Failed trying to create a RepaymentSchedule. Result: %s", result)
		success = false
	}

	return utils.RequestResult{
		Success: success,
		Msg:     createMsg,
		Call:    "CreateRepaymentSchedule",
		Data:    obj,
	}
}

// ----------------------------------------------------------------------------
// GetRepaymentSchedule - returns the matching the provided identifier
// ----------------------------------------------------------------------------
func GetRepaymentSchedule(id uuid.UUID) utils.RequestResult {
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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
		getMsg = fmt.Sprintf("Retrieved a RepaymentSchedule using ID=%v", id)
		success = true
	} else {
		getMsg = fmt.Sprintf("Failed trying to retrieve a RepaymentSchedule using ID=%v", id)
		success = false
	}

	return utils.RequestResult{
		Success: success,
		Msg:     getMsg,
		Call:    "GetRepaymentSchedule",
		Data:    obj,
	}

}

// ----------------------------------------------------------------------------
// GetAllRepaymentSchedule - returns all
// ----------------------------------------------------------------------------
func GetAllRepaymentSchedule() (requestResult utils.RequestResult) {
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
		getAllMsg = "Retrieved all RepaymentSchedule"
		success = true
	} else {
		getAllMsg = fmt.Sprintf("Failed trying to retrieve all RepaymentSchedule. Result: %s", result)
		success = false
	}

	return utils.RequestResult{
		Success: success,
		Msg:     getAllMsg,
		Call:    "GetAllRepaymentSchedule",
		Data:    objs,
	}

}

// ----------------------------------------------------------------------------
// UpdateRepaymentSchedule - updates matching the provided identifier
// ----------------------------------------------------------------------------
func UpdateRepaymentSchedule(obj model.RepaymentSchedule) (requestResult utils.RequestResult) {
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
		updateMsg = fmt.Sprintf("Updated a RepaymentSchedule using ID=%v", obj.ID)
		success = true
	} else {
		updateMsg = fmt.Sprintf("Failed trying to update a RepaymentSchedule using ID=%v", obj.ID)
		success = false
	}

	return utils.RequestResult{
		Success: success,
		Msg:     updateMsg,
		Call:    "UpdateRepaymentSchedule",
		Data:    obj,
	}

}

// ----------------------------------------------------------------------------
// DeleteRepaymentSchedule - deletes matching the provided identifier
// ----------------------------------------------------------------------------
func DeleteRepaymentSchedule(id uuid.UUID) (requestResult utils.RequestResult) {
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRepaymentSchedule(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj, _ := requestResult.Data.(model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
			deleteMsg = fmt.Sprintf("Deleted a RepaymentSchedule using ID=%v", id)
			success = true
		} else {
			deleteMsg = fmt.Sprintf("Failed trying to delete a RepaymentSchedule using ID=%v", id)
			success = false
		}

		requestResult = utils.RequestResult{
			Success: success,
			Msg:     deleteMsg,
			Call:    "DeleteRepaymentSchedule",
			Data:    requestResult.Data,
		}

	}

	return requestResult
}

// ----------------------------------------------------------------------------
// assigns a LoanAccount on a RepaymentSchedule
// ----------------------------------------------------------------------------
func AssignLoanAccountToRepaymentSchedule(repaymentScheduleId uuid.UUID, loanAccountId uuid.UUID) utils.RequestResult {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj, _ := parentRequestResult.Data.(model.RepaymentSchedule)

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
			msg := fmt.Sprintf("Failed trying to read %s using ID=%v", "LoanAccount", loanAccountId)

			return utils.RequestResult{
				Success: false,
				Msg:     msg,
				Call:    "assignLoanAccount",
				Data:    childObj,
			}
		}
	} else {
		return parentRequestResult
	}
}

// ----------------------------------------------------------------------------
// unassigns a LoanAccount on a RepaymentSchedule
// ----------------------------------------------------------------------------
func UnassignLoanAccountFromRepaymentSchedule(repaymentScheduleId uuid.UUID) utils.RequestResult {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj, _ := parentRequestResult.Data.(model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// assign an empty LoanAccount to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccount = nil

		//----------------------------------------------------------------------------
		// assign  nil to the LoanAccount
		//----------------------------------------------------------------------------
		parentObj.LoanAccountId = nil

		//----------------------------------------------------------------------------
		// save the RepaymentSchedule
		//----------------------------------------------------------------------------
		return UpdateRepaymentSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}

// ----------------------------------------------------------------------------
// assigns a Payment on a RepaymentSchedule
// ----------------------------------------------------------------------------
func AssignPaymentToRepaymentSchedule(repaymentScheduleId uuid.UUID, paymentId uuid.UUID) utils.RequestResult {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj, _ := parentRequestResult.Data.(model.RepaymentSchedule)

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
			msg := fmt.Sprintf("Failed trying to read %s using ID=%v", "Payment", paymentId)

			return utils.RequestResult{
				Success: false,
				Msg:     msg,
				Call:    "assignPayment",
				Data:    childObj,
			}
		}
	} else {
		return parentRequestResult
	}
}

// ----------------------------------------------------------------------------
// unassigns a Payment on a RepaymentSchedule
// ----------------------------------------------------------------------------
func UnassignPaymentFromRepaymentSchedule(repaymentScheduleId uuid.UUID) utils.RequestResult {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj, _ := parentRequestResult.Data.(model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// assign an empty LoanPayment to the Payment
		//----------------------------------------------------------------------------
		parentObj.Payment = nil

		//----------------------------------------------------------------------------
		// assign  nil to the Payment
		//----------------------------------------------------------------------------
		parentObj.PaymentId = nil

		//----------------------------------------------------------------------------
		// save the RepaymentSchedule
		//----------------------------------------------------------------------------
		return UpdateRepaymentSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}
