package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
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
	var requestResult utils.RequestResult
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
		createMsg = fmt.Sprintf( "Failed trying to create a Bank", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBank", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBank - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBank(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
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

	requestResult = utils.RequestResult{success, getMsg, "GetBank", obj}

	return requestResult
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
	    getAllMsg = fmt.Sprintf( "Retrieved all Bank" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Bank", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBank", objs}
	return requestResult
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

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBank", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBank - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBank(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBank(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Bank)

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

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBank", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more branchesIds as a Branches to a Bank
//----------------------------------------------------------------------------
func AddBranchesToBank ( bankId uint64, branchesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( branchesIds, ",")

		for _, branchesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var BranchObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&BranchObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Branches using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Branches").Append( &BranchObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )
				return utils.RequestResult{false, msg, "unassignBranches", BranchObj}
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
func RemoveBranchesFromBank( bankId uint64, branchesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( branchesIds, ",")

		for _, branchesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var BranchObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&BranchObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BranchObj from the Branches array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Branches").Delete( &BranchObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )
				return utils.RequestResult{false, msg, "removeBranches", BranchObj}
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
func AddProductsToBank ( bankId uint64, productsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var BankingProductObj model.BankingProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BankingProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&BankingProductObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Products using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Products").Append( &BankingProductObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "unassignProducts", BankingProductObj}
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
func RemoveProductsFromBank( bankId uint64, productsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var BankingProductObj model.BankingProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BankingProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&BankingProductObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BankingProductObj from the Products array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Products").Delete( &BankingProductObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "removeProducts", BankingProductObj}
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
func AddCustomersToBank ( bankId uint64, customersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CustomerObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CustomerObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Customers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Customers").Append( &CustomerObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "unassignCustomers", CustomerObj}
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
func RemoveCustomersFromBank( bankId uint64, customersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var CustomerObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&CustomerObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Customers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Customers").Delete( &CustomerObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "removeCustomers", CustomerObj}
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
func AddAccountsToBank ( bankId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

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
				utils.GetDB().Model(&BankObj).Association("Accounts").Append( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", AccountObj}
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
func RemoveAccountsFromBank( bankId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

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
				utils.GetDB().Model(&BankObj).Association("Accounts").Delete( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", AccountObj}
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
func AddPaymentCardsToBank ( bankId uint64, paymentCardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentCardsIds, ",")

		for _, paymentCardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var PaymentCardObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching paymentCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&PaymentCardObj , paymentCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentCards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("PaymentCards").Append( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "unassignPaymentCards", PaymentCardObj}
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
func RemovePaymentCardsFromBank( bankId uint64, paymentCardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentCardsIds, ",")

		for _, paymentCardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var PaymentCardObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching paymentCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&PaymentCardObj , paymentCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentCardObj from the PaymentCards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("PaymentCards").Delete( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "removePaymentCards", PaymentCardObj}
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
func AddLoanAccountsToBank ( bankId uint64, loanAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

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
				utils.GetDB().Model(&BankObj).Association("LoanAccounts").Append( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "unassignLoanAccounts", LoanAccountObj}
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
func RemoveLoanAccountsFromBank( bankId uint64, loanAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

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
				utils.GetDB().Model(&BankObj).Association("LoanAccounts").Delete( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "removeLoanAccounts", LoanAccountObj}
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
func AddExchangeRatesToBank ( bankId uint64, exchangeRatesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( exchangeRatesIds, ",")

		for _, exchangeRatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ExchangeRateObj model.ExchangeRate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExchangeRate
			// with a matching exchangeRatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ExchangeRateObj , exchangeRatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ExchangeRates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("ExchangeRates").Append( &ExchangeRateObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRates", exchangeRatesId )
				return utils.RequestResult{false, msg, "unassignExchangeRates", ExchangeRateObj}
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
func RemoveExchangeRatesFromBank( bankId uint64, exchangeRatesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( exchangeRatesIds, ",")

		for _, exchangeRatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ExchangeRateObj model.ExchangeRate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExchangeRate
			// with a matching exchangeRatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ExchangeRateObj , exchangeRatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExchangeRateObj from the ExchangeRates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("ExchangeRates").Delete( &ExchangeRateObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExchangeRates", exchangeRatesId )
				return utils.RequestResult{false, msg, "removeExchangeRates", ExchangeRateObj}
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
func AddConsentsToBank ( bankId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ConsentObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ConsentObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Consents").Append( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", ConsentObj}
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
func RemoveConsentsFromBank( bankId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ConsentObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ConsentObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("Consents").Delete( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", ConsentObj}
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
func AddThirdPartyProvidersToBank ( bankId uint64, thirdPartyProvidersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartyProvidersIds, ",")

		for _, thirdPartyProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ThirdPartyProviderObj model.ThirdPartyProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider
			// with a matching thirdPartyProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ThirdPartyProviderObj , thirdPartyProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ThirdPartyProviders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("ThirdPartyProviders").Append( &ThirdPartyProviderObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProviders", thirdPartyProvidersId )
				return utils.RequestResult{false, msg, "unassignThirdPartyProviders", ThirdPartyProviderObj}
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
func RemoveThirdPartyProvidersFromBank( bankId uint64, thirdPartyProvidersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Bank with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBank(bankId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Bank so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankObj,_ := parentRequestResult.Data. (model.Bank)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartyProvidersIds, ",")

		for _, thirdPartyProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ThirdPartyProviderObj model.ThirdPartyProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyProvider
			// with a matching thirdPartyProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ThirdPartyProviderObj , thirdPartyProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ThirdPartyProviderObj from the ThirdPartyProviders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&BankObj).Association("ThirdPartyProviders").Delete( &ThirdPartyProviderObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdPartyProviders", thirdPartyProvidersId )
				return utils.RequestResult{false, msg, "removeThirdPartyProviders", ThirdPartyProviderObj}
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

