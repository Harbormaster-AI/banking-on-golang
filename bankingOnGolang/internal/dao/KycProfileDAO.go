
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
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
		createMsg = fmt.Sprintf( "Failed trying to create a KycProfile. Result: %s", result )
		success = false
	}

	return utils.RequestResult{success, createMsg, "CreateKycProfile", obj}
}


//----------------------------------------------------------------------------
// GetKycProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetKycProfile(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
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

	return utils.RequestResult{success, getMsg, "GetKycProfile", obj}

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
	    getAllMsg = "Retrieved all KycProfile"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all KycProfile. Result: %s", result )
		success = false
	}

	return utils.RequestResult{success, getAllMsg, "GetAllKycProfile", objs}
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

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateKycProfile",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteKycProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteKycProfile(id uuid.UUID)(requestResult utils.RequestResult){
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
		obj,_ := requestResult.Data.(model.KycProfile)

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

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteKycProfile",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a KycProfile
//----------------------------------------------------------------------------
func AssignCustomerToKycProfile( kycProfileId uuid.UUID, customerId uuid.UUID )(utils.RequestResult){

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

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignCustomer",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a KycProfile
//----------------------------------------------------------------------------
func UnassignCustomerFromKycProfile(kycProfileId uuid.UUID)(utils.RequestResult) {

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
func AddIdentityDocumentsToKycProfile ( kycProfileId uuid.UUID, identityDocumentsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, identityDocumentsId:= range identityDocumentsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "unassignIdentityDocuments",
                    Data:       childObj,
                }
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
func RemoveIdentityDocumentsFromKycProfile( kycProfileId uuid.UUID, identityDocumentsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, identityDocumentsId:= range identityDocumentsIds {
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
                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeIdentityDocuments",
                    Data:       childObj,
                }
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
func AddRiskAssessmentsToKycProfile ( kycProfileId uuid.UUID, riskAssessmentsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, riskAssessmentsId:= range riskAssessmentsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "unassignRiskAssessments",
                    Data:       childObj,
                }
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
func RemoveRiskAssessmentsFromKycProfile( kycProfileId uuid.UUID, riskAssessmentsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, riskAssessmentsId:= range riskAssessmentsIds {
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
                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeRiskAssessments",
                    Data:       childObj,
                }
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
func AddScreeningsToKycProfile ( kycProfileId uuid.UUID, screeningsIds []uuid.UUID )(utils.RequestResult) {

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

		for _, screeningsId:= range screeningsIds {
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

                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "unassignScreenings",
                    Data:       childObj,
                }
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
func RemoveScreeningsFromKycProfile( kycProfileId uuid.UUID, screeningsIds []uuid.UUID )(utils.RequestResult) {
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

		for _, screeningsId:= range screeningsIds {
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
                return utils.RequestResult{
                    Success:    false,
                    Msg:        msg,
                    Call:       "removeScreenings",
                    Data:       childObj,
                }
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

