
package dao

import (
    "bankingOnGolang/internal/model"
    "bankingOnGolang/internal/utils"
    "fmt"
    "strings"
    "github.com/google/uuid"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RiskAssessmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRiskAssessment - creates a new db entry
//----------------------------------------------------------------------------
func CreateRiskAssessment(obj model.RiskAssessment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RiskAssessment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RiskAssessment. Result: %s", result )
		success = false
	}

	return utils.RequestResult{success, createMsg, "CreateRiskAssessment", obj}
}


//----------------------------------------------------------------------------
// GetRiskAssessment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRiskAssessment(id uuid.UUID)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RiskAssessment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RiskAssessment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RiskAssessment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RiskAssessment using ID=%v", id )
		success = false
	}

	return utils.RequestResult{success, getMsg, "GetRiskAssessment", obj}

}

//----------------------------------------------------------------------------
// GetAllRiskAssessment - returns all
//----------------------------------------------------------------------------
func GetAllRiskAssessment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RiskAssessment

	//----------------------------------------------------------------------------
	// Request the ORM to find all RiskAssessment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = "Retrieved all RiskAssessment"
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RiskAssessment. Result: %s", result )
		success = false
	}

	return utils.RequestResult{success, getAllMsg, "GetAllRiskAssessment", objs}
}

//----------------------------------------------------------------------------
// UpdateRiskAssessment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRiskAssessment(obj model.RiskAssessment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RiskAssessment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RiskAssessment using ID=%v", obj.ID )
		success = false
	}

	return utils.RequestResult{
        Success:    success,
        Msg:        updateMsg,
        Call:       "UpdateRiskAssessment",
        Data:       obj,
    }

}

//----------------------------------------------------------------------------
// DeleteRiskAssessment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRiskAssessment(id uuid.UUID)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RiskAssessment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRiskAssessment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RiskAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data.(model.RiskAssessment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RiskAssessment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RiskAssessment using ID=%v", id )
			success = false
		}

        requestResult = utils.RequestResult{
            Success:    success,
            Msg:        deleteMsg,
            Call:       "DeleteRiskAssessment",
            Data:       requestResult.Data,
        }

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a RiskAssessment
//----------------------------------------------------------------------------
func AssignKycProfileToRiskAssessment( riskAssessmentId uuid.UUID, kycProfileId uuid.UUID )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RiskAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRiskAssessment(riskAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RiskAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RiskAssessment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.KycProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KycProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the RiskAssessment
			//----------------------------------------------------------------------------
			parentObj.KycProfile = &childObj

			//----------------------------------------------------------------------------
			// save the RiskAssessment
			//----------------------------------------------------------------------------
			return UpdateRiskAssessment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )

            return utils.RequestResult{
                        Success:    false,
                        Msg:        msg,
                        Call:       "assignKycProfile",
                        Data:       childObj,
            }
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a RiskAssessment
//----------------------------------------------------------------------------
func UnassignKycProfileFromRiskAssessment(riskAssessmentId uuid.UUID)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RiskAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRiskAssessment(riskAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RiskAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RiskAssessment)

		//----------------------------------------------------------------------------
		// assign an empty KycProfile to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the RiskAssessment
		//----------------------------------------------------------------------------
		return UpdateRiskAssessment(parentObj)

	} else {
		return parentRequestResult
	}

}


