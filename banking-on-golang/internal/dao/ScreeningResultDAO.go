package dao

import (
    "bankingongolang/internal/model"
    "bankingongolang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ScreeningResultDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateScreeningResult - creates a new db entry
//----------------------------------------------------------------------------
func CreateScreeningResult(obj model.ScreeningResult)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ScreeningResult with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ScreeningResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateScreeningResult", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetScreeningResult - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetScreeningResult(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ScreeningResult

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ScreeningResult with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ScreeningResult using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ScreeningResult using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetScreeningResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllScreeningResult - returns all
//----------------------------------------------------------------------------
func GetAllScreeningResult()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ScreeningResult

	//----------------------------------------------------------------------------
	// Request the ORM to find all ScreeningResult
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ScreeningResult" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ScreeningResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllScreeningResult", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateScreeningResult - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateScreeningResult(obj model.ScreeningResult)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ScreeningResult using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ScreeningResult using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateScreeningResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteScreeningResult - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteScreeningResult(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ScreeningResult with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetScreeningResult(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScreeningResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ScreeningResult)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ScreeningResult using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ScreeningResult using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteScreeningResult", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a ScreeningResult
//----------------------------------------------------------------------------
func AssignKycProfileToScreeningResult( screeningResultId uint64, kycProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ScreeningResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreeningResult(screeningResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScreeningResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ScreeningResultObj,_ := parentRequestResult.Data. (model.ScreeningResult)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var KycProfileObj model.KycProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KycProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&KycProfileObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the ScreeningResult
			//----------------------------------------------------------------------------
			ScreeningResultObj.KycProfile = &KycProfileObj

			//----------------------------------------------------------------------------
			// save the ScreeningResult
			//----------------------------------------------------------------------------
			return UpdateScreeningResult(ScreeningResultObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )
			return utils.RequestResult{false, msg, "assignKycProfile", KycProfileObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a ScreeningResult
//----------------------------------------------------------------------------
func UnassignKycProfileFromScreeningResult(screeningResultId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ScreeningResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreeningResult(screeningResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScreeningResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		ScreeningResultObj,_ := parentRequestResult.Data. (model.ScreeningResult)

		//----------------------------------------------------------------------------
		// assign an empty KycProfile to the KycProfile
		//----------------------------------------------------------------------------
		ScreeningResultObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		ScreeningResultObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the ScreeningResult
		//----------------------------------------------------------------------------
		return UpdateScreeningResult(ScreeningResultObj)

	} else {
		return parentRequestResult
	}

}


