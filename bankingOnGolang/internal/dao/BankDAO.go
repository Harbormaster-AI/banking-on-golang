
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BankDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBank - creates a new db entry
//----------------------------------------------------------------------------
func CreateBank(obj model.Bank)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Bank with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Bank. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        createMsg,
        Call:       "CreateBank",
        Data:       obj,
    }
}


//----------------------------------------------------------------------------
// GetBank - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBank(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Bank

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Bank with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Bank using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Bank using ID=%v", id )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getMsg,
        Call:       "GetBank",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// GetAllBank - returns all
//----------------------------------------------------------------------------
func GetAllBank()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Bank

	//----------------------------------------------------------------------------
	// Request the ORM to find all Bank
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all Bank"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Bank. Result: %s", result )
		success = false
	}

    return utils.RequestResult{
        Success:    success,
        Msg:        getAllMsg,
        Call:       "GetAllBank",
        Data:       objs,
    }

}

//----------------------------------------------------------------------------
// UpdateBank - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBank(obj model.Bank)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Bank using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Bank using ID=%v", obj.ID )
		success = false
	}

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateBank",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteBank - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBank(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBank(id)

	if requestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.Bank)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Bank using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Bank using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteBank",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more branchesIds as a Branches to a Bank
//----------------------------------------------------------------------------
func AddBranchesToBank ( bankId uuid.UUID, branchesIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, branchesId:= range branchesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Branches using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Branches").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Branches").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append Branches association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addBranchesToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more branchesIds as a Branches from a Bank
//----------------------------------------------------------------------------
func RemoveBranchesFromBank( bankId uuid.UUID, branchesIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, branchesId:= range branchesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BranchObj from the Branches array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Branches").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Branches").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove Branches association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeBranchesFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more productsIds as a Products to a Bank
//----------------------------------------------------------------------------
func AddProductsToBank ( bankId uuid.UUID, productsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, productsId:= range productsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BankingProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BankingProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Products using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Products").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append Products association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addProductsToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productsIds as a Products from a Bank
//----------------------------------------------------------------------------
func RemoveProductsFromBank( bankId uuid.UUID, productsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, productsId:= range productsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BankingProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BankingProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BankingProductObj from the Products array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Products").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove Products association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeProductsFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more customersIds as a Customers to a Bank
//----------------------------------------------------------------------------
func AddCustomersToBank ( bankId uuid.UUID, customersIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, customersId:= range customersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Customers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Customers").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append Customers association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addCustomersToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more customersIds as a Customers from a Bank
//----------------------------------------------------------------------------
func RemoveCustomersFromBank( bankId uuid.UUID, customersIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, customersId:= range customersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Customers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Customers").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove Customers association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeCustomersFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Bank
//----------------------------------------------------------------------------
func AddAccountsToBank ( bankId uuid.UUID, accountsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

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

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Accounts").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append Accounts association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addAccountsToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Bank
//----------------------------------------------------------------------------
func RemoveAccountsFromBank( bankId uuid.UUID, accountsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

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
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Accounts").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove Accounts association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeAccountsFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentCardsIds as a PaymentCards to a Bank
//----------------------------------------------------------------------------
func AddPaymentCardsToBank ( bankId uuid.UUID, paymentCardsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, paymentCardsId:= range paymentCardsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching paymentCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentCards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentCards").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("PaymentCards").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append PaymentCards association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addPaymentCardsToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentCardsIds as a PaymentCards from a Bank
//----------------------------------------------------------------------------
func RemovePaymentCardsFromBank( bankId uuid.UUID, paymentCardsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, paymentCardsId:= range paymentCardsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching paymentCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentCardObj from the PaymentCards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentCards").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("PaymentCards").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove PaymentCards association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removePaymentCardsFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loanAccountsIds as a LoanAccounts to a Bank
//----------------------------------------------------------------------------
func AddLoanAccountsToBank ( bankId uuid.UUID, loanAccountsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

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

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("LoanAccounts").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append LoanAccounts association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addLoanAccountsToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loanAccountsIds as a LoanAccounts from a Bank
//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBank( bankId uuid.UUID, loanAccountsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

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
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("LoanAccounts").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove LoanAccounts association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeLoanAccountsFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more exchangeRatesIds as a ExchangeRates to a Bank
//----------------------------------------------------------------------------
func AddExchangeRatesToBank ( bankId uuid.UUID, exchangeRatesIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, exchangeRatesId:= range exchangeRatesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExchangeRate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExchangeRate
			// with a matching exchangeRatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exchangeRatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ExchangeRates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ExchangeRates").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("ExchangeRates").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append ExchangeRates association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRates", exchangeRatesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addExchangeRatesToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more exchangeRatesIds as a ExchangeRates from a Bank
//----------------------------------------------------------------------------
func RemoveExchangeRatesFromBank( bankId uuid.UUID, exchangeRatesIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, exchangeRatesId:= range exchangeRatesIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExchangeRate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExchangeRate
			// with a matching exchangeRatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exchangeRatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExchangeRateObj from the ExchangeRates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ExchangeRates").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("ExchangeRates").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove ExchangeRates association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRates", exchangeRatesId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeExchangeRatesFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a Bank
//----------------------------------------------------------------------------
func AddConsentsToBank ( bankId uuid.UUID, consentsIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, consentsId:= range consentsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Consents").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append Consents association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addConsentsToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a Bank
//----------------------------------------------------------------------------
func RemoveConsentsFromBank( bankId uuid.UUID, consentsIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, consentsId:= range consentsIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("Consents").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove Consents association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeConsentsFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more thirdPartyProvidersIds as a ThirdPartyProviders to a Bank
//----------------------------------------------------------------------------
func AddThirdPartyProvidersToBank ( bankId uuid.UUID, thirdPartyProvidersIds []uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, thirdPartyProvidersId:= range thirdPartyProvidersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdPartyProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider
			// with a matching thirdPartyProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartyProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ThirdPartyProviders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdPartyProviders").Append( &childObj )

                if err := utils.GetDB().
                    Model(&parentObj).
                    Association("ThirdPartyProviders").
                    Append(&childObj); err != nil {
                        log.Printf("Failed to append ThirdPartyProviders association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProviders", thirdPartyProvidersId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "addThirdPartyProvidersToBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more thirdPartyProvidersIds as a ThirdPartyProviders from a Bank
//----------------------------------------------------------------------------
func RemoveThirdPartyProvidersFromBank( bankId uuid.UUID, thirdPartyProvidersIds []uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Bank)

		for _, thirdPartyProvidersId:= range thirdPartyProvidersIds {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdPartyProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider
			// with a matching thirdPartyProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartyProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ThirdPartyProviderObj from the ThirdPartyProviders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdPartyProviders").Delete( &childObj )
				if err := utils.GetDB().
                    Model(&parentObj).
                    Association("ThirdPartyProviders").
                    Delete(&childObj); err != nil {
                        log.Printf("Failed to remove ThirdPartyProviders association from Bank: %v", err)
                        return err
                }
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProviders", thirdPartyProvidersId )

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeThirdPartyProvidersFromBank",
                    Data:       childObj,
                }
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Bank from the gorm
		//----------------------------------------------------------------------------
		return GetBank(bankId)

	} else {
		return parentRequestResult
	}
}

