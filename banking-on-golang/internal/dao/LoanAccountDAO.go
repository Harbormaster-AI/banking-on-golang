package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
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
func GetLoanAccount(id uint64)(utils.RequestResult){
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
func DeleteLoanAccount(id uint64)(requestResult utils.RequestResult){
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
func AssignBankToLoanAccount( loanAccountId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var BankObj model.Bank

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Bank with a
		// matching bankId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&BankObj, bankId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bank	to the LoanAccount
			//----------------------------------------------------------------------------
			LoanAccountObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(LoanAccountObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a LoanAccount
//----------------------------------------------------------------------------
func UnassignBankFromLoanAccount(loanAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		LoanAccountObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		LoanAccountObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(LoanAccountObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Branch on a LoanAccount
//----------------------------------------------------------------------------
func AssignBranchToLoanAccount( loanAccountId uint64, branchId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

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
			// assign the Branch	to the LoanAccount
			//----------------------------------------------------------------------------
			LoanAccountObj.Branch = &BranchObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(LoanAccountObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branch", branchId )
			return utils.RequestResult{false, msg, "assignBranch", BranchObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Branch on a LoanAccount
//----------------------------------------------------------------------------
func UnassignBranchFromLoanAccount(loanAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty Branch to the Branch
		//----------------------------------------------------------------------------
		LoanAccountObj.Branch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Branch
		//----------------------------------------------------------------------------
		LoanAccountObj.BranchId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(LoanAccountObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a LoanAccount
//----------------------------------------------------------------------------
func AssignProductToLoanAccount( loanAccountId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var BankingProductObj model.BankingProduct

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BankingProduct with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&BankingProductObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the LoanAccount
			//----------------------------------------------------------------------------
			LoanAccountObj.Product = &BankingProductObj

			//----------------------------------------------------------------------------
			// save the LoanAccount
			//----------------------------------------------------------------------------
			return UpdateLoanAccount(LoanAccountObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", BankingProductObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a LoanAccount
//----------------------------------------------------------------------------
func UnassignProductFromLoanAccount(loanAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		//----------------------------------------------------------------------------
		// assign an empty BankingProduct to the Product
		//----------------------------------------------------------------------------
		LoanAccountObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		LoanAccountObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the LoanAccount
		//----------------------------------------------------------------------------
		return UpdateLoanAccount(LoanAccountObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more borrowersIds as a Borrowers to a LoanAccount
//----------------------------------------------------------------------------
func AddBorrowersToLoanAccount ( loanAccountId uint64, borrowersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( borrowersIds, ",")

		for _, borrowersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CustomerObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching borrowersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CustomerObj , borrowersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Borrowers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Borrowers").Append( &CustomerObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Borrowers", borrowersId )
				return utils.RequestResult{false, msg, "unassignBorrowers", CustomerObj}
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
func RemoveBorrowersFromLoanAccount( loanAccountId uint64, borrowersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( borrowersIds, ",")

		for _, borrowersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CustomerObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching borrowersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CustomerObj , borrowersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Borrowers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Borrowers").Delete( &CustomerObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Borrowers", borrowersId )
				return utils.RequestResult{false, msg, "removeBorrowers", CustomerObj}
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
func AddRepaymentScheduleToLoanAccount ( loanAccountId uint64, repaymentScheduleIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( repaymentScheduleIds, ",")

		for _, repaymentScheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var RepaymentScheduleObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching repaymentScheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&RepaymentScheduleObj , repaymentScheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RepaymentSchedule using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("RepaymentSchedule").Append( &RepaymentScheduleObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RepaymentSchedule", repaymentScheduleId )
				return utils.RequestResult{false, msg, "unassignRepaymentSchedule", RepaymentScheduleObj}
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
func RemoveRepaymentScheduleFromLoanAccount( loanAccountId uint64, repaymentScheduleIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( repaymentScheduleIds, ",")

		for _, repaymentScheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var RepaymentScheduleObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching repaymentScheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&RepaymentScheduleObj , repaymentScheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RepaymentScheduleObj from the RepaymentSchedule array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("RepaymentSchedule").Delete( &RepaymentScheduleObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RepaymentSchedule", repaymentScheduleId )
				return utils.RequestResult{false, msg, "removeRepaymentSchedule", RepaymentScheduleObj}
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
func AddPaymentsToLoanAccount ( loanAccountId uint64, paymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var LoanPaymentObj model.LoanPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&LoanPaymentObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Payments").Append( &LoanPaymentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "unassignPayments", LoanPaymentObj}
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
func RemovePaymentsFromLoanAccount( loanAccountId uint64, paymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var LoanPaymentObj model.LoanPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&LoanPaymentObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanPaymentObj from the Payments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Payments").Delete( &LoanPaymentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "removePayments", LoanPaymentObj}
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
func AddCollateralToLoanAccount ( loanAccountId uint64, collateralIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CollateralObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CollateralObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Collateral using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Collateral").Append( &CollateralObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "unassignCollateral", CollateralObj}
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
func RemoveCollateralFromLoanAccount( loanAccountId uint64, collateralIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CollateralObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CollateralObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CollateralObj from the Collateral array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("Collateral").Delete( &CollateralObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "removeCollateral", CollateralObj}
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
func AddFeeChargesToLoanAccount ( loanAccountId uint64, feeChargesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeChargesIds, ",")

		for _, feeChargesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FeeChargeObj model.FeeCharge

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeCharge
			// with a matching feeChargesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FeeChargeObj , feeChargesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeeCharges using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("FeeCharges").Append( &FeeChargeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeCharges", feeChargesId )
				return utils.RequestResult{false, msg, "unassignFeeCharges", FeeChargeObj}
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
func RemoveFeeChargesFromLoanAccount( loanAccountId uint64, feeChargesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LoanAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanAccount(loanAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LoanAccountObj,_ := parentRequestResult.Data. (model.LoanAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeChargesIds, ",")

		for _, feeChargesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FeeChargeObj model.FeeCharge

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeCharge
			// with a matching feeChargesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FeeChargeObj , feeChargesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeeChargeObj from the FeeCharges array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LoanAccountObj).Association("FeeCharges").Delete( &FeeChargeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeCharges", feeChargesId )
				return utils.RequestResult{false, msg, "removeFeeCharges", FeeChargeObj}
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

