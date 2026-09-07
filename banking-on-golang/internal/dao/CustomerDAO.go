package dao

import (
    "banking-on-golang/internal/model"
    "banking-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CustomerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCustomer - creates a new db entry
//----------------------------------------------------------------------------
func CreateCustomer(obj model.Customer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Customer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCustomer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCustomer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCustomer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Customer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Customer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Customer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Customer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCustomer - returns all
//----------------------------------------------------------------------------
func GetAllCustomer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Customer

	//----------------------------------------------------------------------------
	// Request the ORM to find all Customer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Customer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCustomer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCustomer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCustomer(obj model.Customer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Customer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Customer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCustomer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCustomer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCustomer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Customer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Customer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCustomer", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bank on a Customer
//----------------------------------------------------------------------------
func AssignBankToCustomer( customerId uint64, bankId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
			// assign the Bank	to the Customer
			//----------------------------------------------------------------------------
			CustomerObj.Bank = &BankObj

			//----------------------------------------------------------------------------
			// save the Customer
			//----------------------------------------------------------------------------
			return UpdateCustomer(CustomerObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", BankObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Customer
//----------------------------------------------------------------------------
func UnassignBankFromCustomer(customerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		CustomerObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		CustomerObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Customer
		//----------------------------------------------------------------------------
		return UpdateCustomer(CustomerObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Customer
//----------------------------------------------------------------------------
func AddAccountsToCustomer ( customerId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("Accounts").Append( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Customer
//----------------------------------------------------------------------------
func RemoveAccountsFromCustomer( customerId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("Accounts").Delete( &AccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", AccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loanAccountsIds as a LoanAccounts to a Customer
//----------------------------------------------------------------------------
func AddLoanAccountsToCustomer ( customerId uint64, loanAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("LoanAccounts").Append( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "unassignLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loanAccountsIds as a LoanAccounts from a Customer
//----------------------------------------------------------------------------
func RemoveLoanAccountsFromCustomer( customerId uint64, loanAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("LoanAccounts").Delete( &LoanAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanAccounts", loanAccountsId )
				return utils.RequestResult{false, msg, "removeLoanAccounts", LoanAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentCardsIds as a PaymentCards to a Customer
//----------------------------------------------------------------------------
func AddPaymentCardsToCustomer ( customerId uint64, paymentCardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("PaymentCards").Append( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "unassignPaymentCards", PaymentCardObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentCardsIds as a PaymentCards from a Customer
//----------------------------------------------------------------------------
func RemovePaymentCardsFromCustomer( customerId uint64, paymentCardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("PaymentCards").Delete( &PaymentCardObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "removePaymentCards", PaymentCardObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more externalAccountsIds as a ExternalAccounts to a Customer
//----------------------------------------------------------------------------
func AddExternalAccountsToCustomer ( customerId uint64, externalAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( externalAccountsIds, ",")

		for _, externalAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ExternalAccountObj model.ExternalAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExternalAccount
			// with a matching externalAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ExternalAccountObj , externalAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ExternalAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("ExternalAccounts").Append( &ExternalAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalAccounts", externalAccountsId )
				return utils.RequestResult{false, msg, "unassignExternalAccounts", ExternalAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more externalAccountsIds as a ExternalAccounts from a Customer
//----------------------------------------------------------------------------
func RemoveExternalAccountsFromCustomer( customerId uint64, externalAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( externalAccountsIds, ",")

		for _, externalAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var ExternalAccountObj model.ExternalAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExternalAccount
			// with a matching externalAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&ExternalAccountObj , externalAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExternalAccountObj from the ExternalAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("ExternalAccounts").Delete( &ExternalAccountObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalAccounts", externalAccountsId )
				return utils.RequestResult{false, msg, "removeExternalAccounts", ExternalAccountObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more fundsTransfersIds as a FundsTransfers to a Customer
//----------------------------------------------------------------------------
func AddFundsTransfersToCustomer ( customerId uint64, fundsTransfersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( fundsTransfersIds, ",")

		for _, fundsTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FundsTransferObj model.FundsTransfer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FundsTransfer
			// with a matching fundsTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FundsTransferObj , fundsTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FundsTransfers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("FundsTransfers").Append( &FundsTransferObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfers", fundsTransfersId )
				return utils.RequestResult{false, msg, "unassignFundsTransfers", FundsTransferObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more fundsTransfersIds as a FundsTransfers from a Customer
//----------------------------------------------------------------------------
func RemoveFundsTransfersFromCustomer( customerId uint64, fundsTransfersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( fundsTransfersIds, ",")

		for _, fundsTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var FundsTransferObj model.FundsTransfer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FundsTransfer
			// with a matching fundsTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&FundsTransferObj , fundsTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FundsTransferObj from the FundsTransfers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("FundsTransfers").Delete( &FundsTransferObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfers", fundsTransfersId )
				return utils.RequestResult{false, msg, "removeFundsTransfers", FundsTransferObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more disputesIds as a Disputes to a Customer
//----------------------------------------------------------------------------
func AddDisputesToCustomer ( customerId uint64, disputesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var DisputeObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&DisputeObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Disputes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("Disputes").Append( &DisputeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "unassignDisputes", DisputeObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more disputesIds as a Disputes from a Customer
//----------------------------------------------------------------------------
func RemoveDisputesFromCustomer( customerId uint64, disputesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var DisputeObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&DisputeObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DisputeObj from the Disputes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("Disputes").Delete( &DisputeObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "removeDisputes", DisputeObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more kycProfilesIds as a KycProfiles to a Customer
//----------------------------------------------------------------------------
func AddKycProfilesToCustomer ( customerId uint64, kycProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var KycProfileObj model.KycProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KycProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&KycProfileObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the KycProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("KycProfiles").Append( &KycProfileObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "unassignKycProfiles", KycProfileObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more kycProfilesIds as a KycProfiles from a Customer
//----------------------------------------------------------------------------
func RemoveKycProfilesFromCustomer( customerId uint64, kycProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var KycProfileObj model.KycProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KycProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&KycProfileObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove KycProfileObj from the KycProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&CustomerObj).Association("KycProfiles").Delete( &KycProfileObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "removeKycProfiles", KycProfileObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a Customer
//----------------------------------------------------------------------------
func AddConsentsToCustomer ( customerId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("Consents").Append( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", ConsentObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a Customer
//----------------------------------------------------------------------------
func RemoveConsentsFromCustomer( customerId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		CustomerObj,_ := parentRequestResult.Data. (model.Customer)

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
				utils.GetDB().Model(&CustomerObj).Association("Consents").Delete( &ConsentObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", ConsentObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

