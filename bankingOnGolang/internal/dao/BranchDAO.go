
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
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
func GetBranch(id UUID)(utils.RequestResult){
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
func DeleteBranch(id UUID)(requestResult utils.RequestResult){
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
func AssignBankToBranch( branchId UUID, bankId UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

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
			// assign the Bank	to the Branch
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the Branch
			//----------------------------------------------------------------------------
			return UpdateBranch(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Branch
//----------------------------------------------------------------------------
func UnassignBankFromBranch(branchId UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Branch
		//----------------------------------------------------------------------------
		return UpdateBranch(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Branch
//----------------------------------------------------------------------------
func AddAccountsToBranch ( branchId UUID, accountsIds UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
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
func RemoveAccountsFromBranch( branchId UUID, accountsIds UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
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
func AddLoanAccountsToBranch ( branchId UUID, loanAccountsIds UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanAccountsIds, ",")

		for _, loanAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanAccount
			// with a matching loanAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loanAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LoanAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LoanAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "unassignLoanAccounts", childObj}
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
func RemoveLoanAccountsFromBranch( branchId UUID, loanAccountsIds UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanAccountsIds, ",")

		for _, loanAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanAccount
			// with a matching loanAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loanAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanAccountObj from the LoanAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LoanAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "removeLoanAccounts", childObj}
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
func AddAtmsToBranch ( branchId UUID, atmsIds UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( atmsIds, ",")

		for _, atmsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ATM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ATM
			// with a matching atmsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , atmsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Atms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Atms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Atms", atmsId )
				return utils.RequestResult{false, msg, "unassignAtms", childObj}
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
func RemoveAtmsFromBranch( branchId UUID, atmsIds UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Branch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBranch(branchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Branch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Branch)

		// slice the ids on comma with no spaces
		ids := strings.Split( atmsIds, ",")

		for _, atmsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ATM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ATM
			// with a matching atmsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , atmsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ATMObj from the Atms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Atms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Atms", atmsId )
				return utils.RequestResult{false, msg, "removeAtms", childObj}
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

