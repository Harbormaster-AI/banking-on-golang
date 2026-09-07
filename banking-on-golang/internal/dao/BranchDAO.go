package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BranchDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBranch - creates a new db entry
//----------------------------------------------------------------------------
func CreateBranch(obj model.Branch)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Branch with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Branch", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBranch", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBranch - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBranch(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Branch

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Branch with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Branch using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Branch using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBranch", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBranch - returns all
//----------------------------------------------------------------------------
func GetAllBranch()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Branch

	//----------------------------------------------------------------------------
	// Request the ORM to find all Branch
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Branch" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Branch", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBranch", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBranch - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBranch(obj model.Branch)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Branch using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Branch using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBranch", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBranch - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBranch(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBranch(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Branch)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Branch using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Branch using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBranch", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a Branch
//----------------------------------------------------------------------------
func AssignBankToBranch( branchId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

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
			// assign the Bank	to the Branch
			//----------------------------------------------------------------------------
			BranchObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the Branch
			//----------------------------------------------------------------------------
			return UpdateBranch(BranchObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Branch
//----------------------------------------------------------------------------
func UnassignBankFromBranch(branchId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		BranchObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		BranchObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Branch
		//----------------------------------------------------------------------------
		return UpdateBranch(BranchObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Branch
//----------------------------------------------------------------------------
func AddAccountsToBranch ( branchId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var AccountObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&AccountObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("Accounts").Append( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Branch
//----------------------------------------------------------------------------
func RemoveAccountsFromBranch( branchId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var AccountObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&AccountObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("Accounts").Delete( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loanAccountsIds as a LoanAccounts to a Branch
//----------------------------------------------------------------------------
func AddLoanAccountsToBranch ( branchId uint64, loanAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanAccountsIds, ",")

		for _, loanAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var LoanAccountObj model.LoanAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanAccount
			// with a matching loanAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&LoanAccountObj , loanAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LoanAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("LoanAccounts").Append( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "unassignLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loanAccountsIds as a LoanAccounts from a Branch
//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBranch( branchId uint64, loanAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanAccountsIds, ",")

		for _, loanAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var LoanAccountObj model.LoanAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanAccount
			// with a matching loanAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&LoanAccountObj , loanAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanAccountObj from the LoanAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("LoanAccounts").Delete( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "removeLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more atmsIds as a Atms to a Branch
//----------------------------------------------------------------------------
func AddAtmsToBranch ( branchId uint64, atmsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( atmsIds, ",")

		for _, atmsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ATMObj model.ATM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ATM
			// with a matching atmsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ATMObj , atmsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Atms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("Atms").Append( &ATMObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Atms", atmsId )
				return utils.RequestResult{false, msg, "unassignAtms", ATMObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more atmsIds as a Atms from a Branch
//----------------------------------------------------------------------------
func RemoveAtmsFromBranch( branchId uint64, atmsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BranchObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( atmsIds, ",")

		for _, atmsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ATMObj model.ATM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ATM
			// with a matching atmsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ATMObj , atmsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ATMObj from the Atms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BranchObj).Association("Atms").Delete( &ATMObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Atms", atmsId )
				return utils.RequestResult{false, msg, "removeAtms", ATMObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Branch from the gorm
		//----------------------------------------------------------------------------
		return GetBranch(branchId)

	} else {
		return parentRequestResult
	}
}

