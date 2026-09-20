
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
func GetCustomer(id uuid.UUID)(utils.RequestResult){
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
func DeleteCustomer(id uuid.UUID)(requestResult utils.RequestResult){
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
func AssignBankToCustomer( customerId uuid.UUID, bankId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

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
			// assign the Bank	to the Customer
			//----------------------------------------------------------------------------
			parentObj.Bank = &childObj

			//----------------------------------------------------------------------------
			// save the Customer
			//----------------------------------------------------------------------------
			return UpdateCustomer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bank", bankId )
			return utils.RequestResult{false, msg, "assignBank", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bank on a Customer
//----------------------------------------------------------------------------
func UnassignBankFromCustomer(customerId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// assign an empty Bank to the Bank
		//----------------------------------------------------------------------------
		parentObj.Bank = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bank
		//----------------------------------------------------------------------------
		parentObj.BankId = nil;

		//----------------------------------------------------------------------------
		// save the Customer
		//----------------------------------------------------------------------------
		return UpdateCustomer(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Customer
//----------------------------------------------------------------------------
func AddAccountsToCustomer ( customerId uuid.UUID, accountsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

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
func RemoveAccountsFromCustomer( customerId uuid.UUID, accountsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

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
func AddLoanAccountsToCustomer ( customerId uuid.UUID, loanAccountsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

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
func RemoveLoanAccountsFromCustomer( customerId uuid.UUID, loanAccountsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

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
func AddPaymentCardsToCustomer ( customerId uuid.UUID, paymentCardsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentCardsIds, ",")

		for _, paymentCardsId:= range ids {
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

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "unassignPaymentCards", childObj}
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
func RemovePaymentCardsFromCustomer( customerId uuid.UUID, paymentCardsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentCardsIds, ",")

		for _, paymentCardsId:= range ids {
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

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentCards", paymentCardsId )
				return utils.RequestResult{false, msg, "removePaymentCards", childObj}
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
func AddExternalAccountsToCustomer ( customerId uuid.UUID, externalAccountsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( externalAccountsIds, ",")

		for _, externalAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExternalAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExternalAccount
			// with a matching externalAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , externalAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ExternalAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ExternalAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalAccounts", externalAccountsId )
				return utils.RequestResult{false, msg, "unassignExternalAccounts", childObj}
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
func RemoveExternalAccountsFromCustomer( customerId uuid.UUID, externalAccountsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( externalAccountsIds, ",")

		for _, externalAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExternalAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExternalAccount
			// with a matching externalAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , externalAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExternalAccountObj from the ExternalAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ExternalAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ExternalAccounts", externalAccountsId )
				return utils.RequestResult{false, msg, "removeExternalAccounts", childObj}
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
func AddFundsTransfersToCustomer ( customerId uuid.UUID, fundsTransfersIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( fundsTransfersIds, ",")

		for _, fundsTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FundsTransfer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FundsTransfer
			// with a matching fundsTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , fundsTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FundsTransfers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FundsTransfers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfers", fundsTransfersId )
				return utils.RequestResult{false, msg, "unassignFundsTransfers", childObj}
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
func RemoveFundsTransfersFromCustomer( customerId uuid.UUID, fundsTransfersIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( fundsTransfersIds, ",")

		for _, fundsTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FundsTransfer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FundsTransfer
			// with a matching fundsTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , fundsTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FundsTransferObj from the FundsTransfers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FundsTransfers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FundsTransfers", fundsTransfersId )
				return utils.RequestResult{false, msg, "removeFundsTransfers", childObj}
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
func AddDisputesToCustomer ( customerId uuid.UUID, disputesIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Disputes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "unassignDisputes", childObj}
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
func RemoveDisputesFromCustomer( customerId uuid.UUID, disputesIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DisputeObj from the Disputes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "removeDisputes", childObj}
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
func AddKycProfilesToCustomer ( customerId uuid.UUID, kycProfilesIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KycProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KycProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the KycProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KycProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "unassignKycProfiles", childObj}
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
func RemoveKycProfilesFromCustomer( customerId uuid.UUID, kycProfilesIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KycProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KycProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove KycProfileObj from the KycProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KycProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "removeKycProfiles", childObj}
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
func AddConsentsToCustomer ( customerId uuid.UUID, consentsIds uuid.UUID )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
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

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", childObj}
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
func RemoveConsentsFromCustomer( customerId uuid.UUID, consentsIds uuid.UUID )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
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

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", childObj}
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

