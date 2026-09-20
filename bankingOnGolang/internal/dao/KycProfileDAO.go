
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing KycProfileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateKycProfile - creates a new db entry
//----------------------------------------------------------------------------
func CreateKycProfile(obj model.KycProfile)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a KycProfile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a KycProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateKycProfile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetKycProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetKycProfile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.KycProfile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a KycProfile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a KycProfile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a KycProfile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetKycProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllKycProfile - returns all
//----------------------------------------------------------------------------
func GetAllKycProfile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.KycProfile

	//----------------------------------------------------------------------------
	// Request the ORM to find all KycProfile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all KycProfile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all KycProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllKycProfile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateKycProfile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateKycProfile(obj model.KycProfile)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a KycProfile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a KycProfile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateKycProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteKycProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteKycProfile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetKycProfile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.KycProfile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a KycProfile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a KycProfile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteKycProfile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a KycProfile
//----------------------------------------------------------------------------
func AssignCustomerToKycProfile( kycProfileId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the KycProfile
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the KycProfile
			//----------------------------------------------------------------------------
			return UpdateKycProfile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a KycProfile
//----------------------------------------------------------------------------
func UnassignCustomerFromKycProfile(kycProfileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the KycProfile
		//----------------------------------------------------------------------------
		return UpdateKycProfile(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more identityDocumentsIds as a IdentityDocuments to a KycProfile
//----------------------------------------------------------------------------
func AddIdentityDocumentsToKycProfile ( kycProfileId uint64, identityDocumentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( identityDocumentsIds, ",")

		for _, identityDocumentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.IdentityDocument

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a IdentityDocument
			// with a matching identityDocumentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , identityDocumentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the IdentityDocuments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("IdentityDocuments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "IdentityDocuments", identityDocumentsId )
				return utils.RequestResult{false, msg, "unassignIdentityDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more identityDocumentsIds as a IdentityDocuments from a KycProfile
//----------------------------------------------------------------------------
func RemoveIdentityDocumentsFromKycProfile( kycProfileId uint64, identityDocumentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( identityDocumentsIds, ",")

		for _, identityDocumentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.IdentityDocument

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a IdentityDocument
			// with a matching identityDocumentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , identityDocumentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove IdentityDocumentObj from the IdentityDocuments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("IdentityDocuments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "IdentityDocuments", identityDocumentsId )
				return utils.RequestResult{false, msg, "removeIdentityDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more riskAssessmentsIds as a RiskAssessments to a KycProfile
//----------------------------------------------------------------------------
func AddRiskAssessmentsToKycProfile ( kycProfileId uint64, riskAssessmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( riskAssessmentsIds, ",")

		for _, riskAssessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RiskAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RiskAssessment
			// with a matching riskAssessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , riskAssessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RiskAssessments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RiskAssessments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RiskAssessments", riskAssessmentsId )
				return utils.RequestResult{false, msg, "unassignRiskAssessments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more riskAssessmentsIds as a RiskAssessments from a KycProfile
//----------------------------------------------------------------------------
func RemoveRiskAssessmentsFromKycProfile( kycProfileId uint64, riskAssessmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( riskAssessmentsIds, ",")

		for _, riskAssessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RiskAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RiskAssessment
			// with a matching riskAssessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , riskAssessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RiskAssessmentObj from the RiskAssessments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RiskAssessments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RiskAssessments", riskAssessmentsId )
				return utils.RequestResult{false, msg, "removeRiskAssessments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more screeningsIds as a Screenings to a KycProfile
//----------------------------------------------------------------------------
func AddScreeningsToKycProfile ( kycProfileId uint64, screeningsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ScreeningResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ScreeningResult
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Screenings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "unassignScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more screeningsIds as a Screenings from a KycProfile
//----------------------------------------------------------------------------
func RemoveScreeningsFromKycProfile( kycProfileId uint64, screeningsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KycProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKycProfile(kycProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KycProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KycProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ScreeningResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ScreeningResult
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ScreeningResultObj from the Screenings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "removeScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KycProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKycProfile(kycProfileId)

	} else {
		return parentRequestResult
	}
}

