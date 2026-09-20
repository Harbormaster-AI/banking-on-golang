
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LoanAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLoanAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateLoanAccount(obj model.LoanAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LoanAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LoanAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoanAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoanAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoanAccount(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LoanAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LoanAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LoanAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LoanAccount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLoanAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLoanAccount - returns all
//----------------------------------------------------------------------------
func GetAllLoanAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LoanAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all LoanAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LoanAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LoanAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLoanAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLoanAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLoanAccount(obj model.LoanAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LoanAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LoanAccount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLoanAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoanAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoanAccount(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLoanAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LoanAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LoanAccount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLoanAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a LoanAccount
//----------------------------------------------------------------------------
func AssignBankToLoanAccount( loanAccountId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Bank

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Bank with a
		// matching bankId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, bankId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bank	to the LoanAccount
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a LoanAccount
//----------------------------------------------------------------------------
func UnassignBankFromLoanAccount(loanAccountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Branch on a LoanAccount
//----------------------------------------------------------------------------
func AssignBranchToLoanAccount( loanAccountId uuid.UUID, branchId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

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
			// assign the Branch	to the LoanAccount
			//----------------------------------------------------------------------------
			parentObj.Branch = &childObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branch", branchId )
			return utils.RequestResult{false, msg, "assignBranch", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Branch on a LoanAccount
//----------------------------------------------------------------------------
func UnassignBranchFromLoanAccount(loanAccountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty Branch to the Branch
		//----------------------------------------------------------------------------
		parentObj.Branch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Branch
		//----------------------------------------------------------------------------
		parentObj.BranchId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a LoanAccount
//----------------------------------------------------------------------------
func AssignProductToLoanAccount( loanAccountId uuid.UUID, productId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BankingProduct

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BankingProduct with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the LoanAccount
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a LoanAccount
//----------------------------------------------------------------------------
func UnassignProductFromLoanAccount(loanAccountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty BankingProduct to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more borrowersIds as a Borrowers to a LoanAccount
//----------------------------------------------------------------------------
func AddBorrowersToLoanAccount ( loanAccountId uuid.UUID, borrowersIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( borrowersIds, ",")

		for _, borrowersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching borrowersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , borrowersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Borrowers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Borrowers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Borrowers", borrowersId )
				return utils.RequestResult{false, msg, "unassignBorrowers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more borrowersIds as a Borrowers from a LoanAccount
//----------------------------------------------------------------------------
func RemoveBorrowersFromLoanAccount( loanAccountId uuid.UUID, borrowersIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( borrowersIds, ",")

		for _, borrowersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching borrowersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , borrowersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Borrowers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Borrowers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Borrowers", borrowersId )
				return utils.RequestResult{false, msg, "removeBorrowers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more repaymentScheduleIds as a RepaymentSchedule to a LoanAccount
//----------------------------------------------------------------------------
func AddRepaymentScheduleToLoanAccount ( loanAccountId uuid.UUID, repaymentScheduleIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( repaymentScheduleIds, ",")

		for _, repaymentScheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching repaymentScheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repaymentScheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RepaymentSchedule using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RepaymentSchedule").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RepaymentSchedule", repaymentScheduleId )
				return utils.RequestResult{false, msg, "unassignRepaymentSchedule", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more repaymentScheduleIds as a RepaymentSchedule from a LoanAccount
//----------------------------------------------------------------------------
func RemoveRepaymentScheduleFromLoanAccount( loanAccountId uuid.UUID, repaymentScheduleIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( repaymentScheduleIds, ",")

		for _, repaymentScheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching repaymentScheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repaymentScheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RepaymentScheduleObj from the RepaymentSchedule array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RepaymentSchedule").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RepaymentSchedule", repaymentScheduleId )
				return utils.RequestResult{false, msg, "removeRepaymentSchedule", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentsIds as a Payments to a LoanAccount
//----------------------------------------------------------------------------
func AddPaymentsToLoanAccount ( loanAccountId uuid.UUID, paymentsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "unassignPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentsIds as a Payments from a LoanAccount
//----------------------------------------------------------------------------
func RemovePaymentsFromLoanAccount( loanAccountId uuid.UUID, paymentsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanPaymentObj from the Payments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "removePayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more collateralIds as a Collateral to a LoanAccount
//----------------------------------------------------------------------------
func AddCollateralToLoanAccount ( loanAccountId uuid.UUID, collateralIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Collateral using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Collateral").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "unassignCollateral", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more collateralIds as a Collateral from a LoanAccount
//----------------------------------------------------------------------------
func RemoveCollateralFromLoanAccount( loanAccountId uuid.UUID, collateralIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CollateralObj from the Collateral array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Collateral").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "removeCollateral", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more feeChargesIds as a FeeCharges to a LoanAccount
//----------------------------------------------------------------------------
func AddFeeChargesToLoanAccount ( loanAccountId uuid.UUID, feeChargesIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeChargesIds, ",")

		for _, feeChargesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeeCharge

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeCharge
			// with a matching feeChargesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feeChargesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeeCharges using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeeCharges").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeCharges", feeChargesId )
				return utils.RequestResult{false, msg, "unassignFeeCharges", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more feeChargesIds as a FeeCharges from a LoanAccount
//----------------------------------------------------------------------------
func RemoveFeeChargesFromLoanAccount( loanAccountId uuid.UUID, feeChargesIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeChargesIds, ",")

		for _, feeChargesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeeCharge

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeCharge
			// with a matching feeChargesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feeChargesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeeChargeObj from the FeeCharges array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeeCharges").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeCharges", feeChargesId )
				return utils.RequestResult{false, msg, "removeFeeCharges", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LoanAccount from the gorm
		//----------------------------------------------------------------------------
		return GetLoanAccount(loanAccountId)

	} else {
		return parentRequestResult
	}
}

