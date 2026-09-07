package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BankingProductDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBankingProduct - creates a new db entry
//----------------------------------------------------------------------------
func CreateBankingProduct(obj model.BankingProduct)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BankingProduct with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BankingProduct", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBankingProduct", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBankingProduct - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBankingProduct(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BankingProduct

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BankingProduct with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BankingProduct using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BankingProduct using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBankingProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBankingProduct - returns all
//----------------------------------------------------------------------------
func GetAllBankingProduct()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BankingProduct

	//----------------------------------------------------------------------------
	// Request the ORM to find all BankingProduct
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BankingProduct" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BankingProduct", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBankingProduct", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBankingProduct - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBankingProduct(obj model.BankingProduct)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BankingProduct using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BankingProduct using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBankingProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBankingProduct - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBankingProduct(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBankingProduct(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BankingProduct)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BankingProduct using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BankingProduct using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBankingProduct", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a BankingProduct
//----------------------------------------------------------------------------
func AssignBankToBankingProduct( bankingProductId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
			// assign the Bank	to the BankingProduct
			//----------------------------------------------------------------------------
			BankingProductObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the BankingProduct
			//----------------------------------------------------------------------------
			return UpdateBankingProduct(BankingProductObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a BankingProduct
//----------------------------------------------------------------------------
func UnassignBankFromBankingProduct(bankingProductId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		BankingProductObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		BankingProductObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the BankingProduct
		//----------------------------------------------------------------------------
		return UpdateBankingProduct(BankingProductObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a BankingProduct
//----------------------------------------------------------------------------
func AddAccountsToBankingProduct ( bankingProductId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("Accounts").Append( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a BankingProduct
//----------------------------------------------------------------------------
func RemoveAccountsFromBankingProduct( bankingProductId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("Accounts").Delete( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loanAccountsIds as a LoanAccounts to a BankingProduct
//----------------------------------------------------------------------------
func AddLoanAccountsToBankingProduct ( bankingProductId uint64, loanAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("LoanAccounts").Append( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "unassignLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loanAccountsIds as a LoanAccounts from a BankingProduct
//----------------------------------------------------------------------------
func RemoveLoanAccountsFromBankingProduct( bankingProductId uint64, loanAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("LoanAccounts").Delete( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "removeLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentCardsIds as a PaymentCards to a BankingProduct
//----------------------------------------------------------------------------
func AddPaymentCardsToBankingProduct ( bankingProductId uint64, paymentCardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("PaymentCards").Append( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "unassignPaymentCards", PaymentCardObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentCardsIds as a PaymentCards from a BankingProduct
//----------------------------------------------------------------------------
func RemovePaymentCardsFromBankingProduct( bankingProductId uint64, paymentCardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BankingProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBankingProduct(bankingProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankingProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		BankingProductObj,_ := parentRequestResult.Data. (model.BankingProduct)

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
				utils.GetDB().Model(&BankingProductObj).Association("PaymentCards").Delete( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "removePaymentCards", PaymentCardObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BankingProduct from the gorm
		//----------------------------------------------------------------------------
		return GetBankingProduct(bankingProductId)

	} else {
		return parentRequestResult
	}
}

