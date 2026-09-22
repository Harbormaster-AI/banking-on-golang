
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a Branch. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateBranch",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetBranch - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBranch(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetBranch",
        Data:       obj,
    }

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
	    getAllMsg = "Retrieved all Branch"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Branch. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllBranch",
        Data:       objs,
    }

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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateBranch",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteBranch - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBranch(id uuid.UUID)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data.(model.Branch)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteBranch",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a Branch
//----------------------------------------------------------------------------
func AssignBankToBranch( branchId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

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

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignBank",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Branch
//----------------------------------------------------------------------------
func UnassignBankFromBranch(branchId uuid.UUID)(utils.RequestResult) {

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
func AddAccountsToBranch ( branchId uuid.UUID, accountsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, accountsId:= range accountsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addAccountsToBranch",
                    Data:       childObj,
                }
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
func RemoveAccountsFromBranch( branchId uuid.UUID, accountsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, accountsId:= range accountsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeAccountsFromBranch",
                    Data:       childObj,
                }
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
func AddLoanAccountsToBranch ( branchId uuid.UUID, loanAccountsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, loanAccountsId:= range loanAccountsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addLoanAccountsToBranch",
                    Data:       childObj,
                }
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
func RemoveLoanAccountsFromBranch( branchId uuid.UUID, loanAccountsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, loanAccountsId:= range loanAccountsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeLoanAccountsFromBranch",
                    Data:       childObj,
                }
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
func AddAtmsToBranch ( branchId uuid.UUID, atmsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, atmsId:= range atmsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addAtmsToBranch",
                    Data:       childObj,
                }
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
func RemoveAtmsFromBranch( branchId uuid.UUID, atmsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, atmsId:= range atmsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeAtmsFromBranch",
                    Data:       childObj,
                }
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

