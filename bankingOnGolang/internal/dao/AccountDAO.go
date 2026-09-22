
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateAccount(obj model.Account)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Account with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Account. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAccount(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Account

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Account with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Account using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Account using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAccount - returns all
//----------------------------------------------------------------------------
func GetAllAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Account

	//----------------------------------------------------------------------------
	// Request the ORM to find all Account
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all Account"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Account. Result: %s", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAccount(obj model.Account)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Account using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Account using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{
        Success: success,
        Msg: updateMsg,
        Call:  "UpdateAccount",
        Data:    obj,
    }

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAccount(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.Account)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Account using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Account using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success: success,
            Msg: deleteMsg,
            Call:  "DeleteAccount",
            Data:    requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a Account
//----------------------------------------------------------------------------
func AssignBankToAccount( accountId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
			// assign the Bank	to the Account
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignBank",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Account
//----------------------------------------------------------------------------
func UnassignBankFromAccount(accountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Branch on a Account
//----------------------------------------------------------------------------
func AssignBranchToAccount( accountId uuid.UUID, branchId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
			// assign the Branch	to the Account
			//----------------------------------------------------------------------------
			parentObj.Branch = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branch", branchId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignBranch",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Branch on a Account
//----------------------------------------------------------------------------
func UnassignBranchFromAccount(accountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty Branch to the Branch
		//----------------------------------------------------------------------------
		parentObj.Branch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Branch
		//----------------------------------------------------------------------------
		parentObj.BranchId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a Account
//----------------------------------------------------------------------------
func AssignProductToAccount( accountId uuid.UUID, productId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
			// assign the Product	to the Account
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )

            requestResult = utils.RequestResult{
                Success: false,
                Msg: msg,
                Call:  "assignProduct",
                Data:    childObj,
            }
            return requestResult;
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a Account
//----------------------------------------------------------------------------
func UnassignProductFromAccount(accountId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty BankingProduct to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more ownersIds as a Owners to a Account
//----------------------------------------------------------------------------
func AddOwnersToAccount ( accountId uuid.UUID, ownersIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, ownersId:= range ownersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching ownersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Owners using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Owners").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owners", ownersId )

                return utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignOwners",
                    Data:    childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ownersIds as a Owners from a Account
//----------------------------------------------------------------------------
func RemoveOwnersFromAccount( accountId uuid.UUID, ownersIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, ownersId:= range ownersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching ownersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Owners array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Owners").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owners", ownersId )
                return utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeOwners",
                                    Data:    childObj,
                                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a Account
//----------------------------------------------------------------------------
func AddTransactionsToAccount ( accountId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, transactionsId:= range transactionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )

                return utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignTransactions",
                    Data:    childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a Account
//----------------------------------------------------------------------------
func RemoveTransactionsFromAccount( accountId uuid.UUID, transactionsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, transactionsId:= range transactionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
                return utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeTransactions",
                                    Data:    childObj,
                                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more statementsIds as a Statements to a Account
//----------------------------------------------------------------------------
func AddStatementsToAccount ( accountId uuid.UUID, statementsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, statementsId:= range statementsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AccountStatement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AccountStatement
			// with a matching statementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , statementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Statements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Statements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Statements", statementsId )

                return utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignStatements",
                    Data:    childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more statementsIds as a Statements from a Account
//----------------------------------------------------------------------------
func RemoveStatementsFromAccount( accountId uuid.UUID, statementsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, statementsId:= range statementsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AccountStatement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AccountStatement
			// with a matching statementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , statementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountStatementObj from the Statements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Statements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Statements", statementsId )
                return utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeStatements",
                                    Data:    childObj,
                                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more standingInstructionsIds as a StandingInstructions to a Account
//----------------------------------------------------------------------------
func AddStandingInstructionsToAccount ( accountId uuid.UUID, standingInstructionsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, standingInstructionsId:= range standingInstructionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StandingInstruction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StandingInstruction
			// with a matching standingInstructionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , standingInstructionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the StandingInstructions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StandingInstructions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StandingInstructions", standingInstructionsId )

                return utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignStandingInstructions",
                    Data:    childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more standingInstructionsIds as a StandingInstructions from a Account
//----------------------------------------------------------------------------
func RemoveStandingInstructionsFromAccount( accountId uuid.UUID, standingInstructionsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, standingInstructionsId:= range standingInstructionsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StandingInstruction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StandingInstruction
			// with a matching standingInstructionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , standingInstructionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove StandingInstructionObj from the StandingInstructions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StandingInstructions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StandingInstructions", standingInstructionsId )
                return utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeStandingInstructions",
                                    Data:    childObj,
                                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more feeChargesIds as a FeeCharges to a Account
//----------------------------------------------------------------------------
func AddFeeChargesToAccount ( accountId uuid.UUID, feeChargesIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, feeChargesId:= range feeChargesIds {
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

                return utils.RequestResult{
                    Success: false,
                    Msg: msg,
                    Call:  "unassignFeeCharges",
                    Data:    childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more feeChargesIds as a FeeCharges from a Account
//----------------------------------------------------------------------------
func RemoveFeeChargesFromAccount( accountId uuid.UUID, feeChargesIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		for _, feeChargesId:= range feeChargesIds {
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
                return utils.RequestResult{
                                    Success: false,
                                    Msg: msg,
                                    Call:  "removeFeeCharges",
                                    Data:    childObj,
                                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

