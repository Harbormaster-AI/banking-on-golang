package test

import ( 
	"testing"
    dao "bankingOnGolang/internal/dao"
	"bankingOnGolang/internal/model"
	"bankingOnGolang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestBankCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Bank
	//----------------------------------------------------------------------------
	BankObj := model.Bank{
        Name:"test value for Name",
        LegalName:"test value for LegalName",
        SwiftBic:model.BIC{},
        HeadquartersCountry:"test value for HeadquartersCountry",
        Website:"test value for Website",
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBankRequestResult := dao.CreateBank( BankObj )
	
	if createBankRequestResult.Success == false {
		t.Error(createBankRequestResult.Msg)
	} else {
		fmt.Println("Check Create Bank success...")
	}
	
	createBankObj, ok := createBankRequestResult.Data.(model.Bank)

    if !ok {
        t.Fatalf("Expected Bank, got %T", createBankRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Bank Obj ID
	// --------------------------------------------------------------	
	if createBankObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Bank" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBankRequestResult := dao.GetBank( uint64(createBankObj.ID) )
	
	if getBankRequestResult.Success == false {
		t.Error(getBankRequestResult.Msg)
	} else {
		fmt.Println("Check Get Bank success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBankObj,_ := getBankRequestResult.Data (model.Bank)
	compareBank := cmp.Equal(createBankObj.ID, getBankObj.ID)
	
	if  compareBank == false	{
		t.Error( "Created Bank object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBankRequestResult := dao.GetAllBank()

	if getAllBankRequestResult.Success == false {
			t.Error(getAllBankRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Bank success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBankObj []model.Bank = getAllBankRequestResult.Data. ([]model.Bank)
		
	equalBank := cmp.Equal(createBankObj.ID, getAllBankObj[len(getAllBankObj)-1].ID)
		
	if equalBank == false {
		t.Error( "Created object is not equal to the last entry in Bank[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Bank
	// --------------------------------------------------------------	
	deleteBankRequestResult := dao.DeleteBank(uint64(createBankObj.ID))

	if deleteBankRequestResult.Success == false {
			t.Error(deleteBankRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Bank success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBankRequestResult = dao.GetBank( uint64(createBankObj.ID) )
	
	if getBankRequestResult.Success == true {
		t.Error(getBankRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBranchCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Branch
	//----------------------------------------------------------------------------
	BranchObj := model.Branch{
        Name:"test value for Name",
        BranchCode:"test value for BranchCode",
        Address:model.Address{},
        Phone:"test value for Phone",
        OpeningHours:"test value for OpeningHours",
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBranchRequestResult := dao.CreateBranch( BranchObj )
	
	if createBranchRequestResult.Success == false {
		t.Error(createBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Create Branch success...")
	}
	
	createBranchObj, ok := createBranchRequestResult.Data.(model.Branch)

    if !ok {
        t.Fatalf("Expected Branch, got %T", createBranchRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Branch Obj ID
	// --------------------------------------------------------------	
	if createBranchObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Branch" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBranchRequestResult := dao.GetBranch( uint64(createBranchObj.ID) )
	
	if getBranchRequestResult.Success == false {
		t.Error(getBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Get Branch success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBranchObj,_ := getBranchRequestResult.Data (model.Branch)
	compareBranch := cmp.Equal(createBranchObj.ID, getBranchObj.ID)
	
	if  compareBranch == false	{
		t.Error( "Created Branch object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBranchRequestResult := dao.GetAllBranch()

	if getAllBranchRequestResult.Success == false {
			t.Error(getAllBranchRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Branch success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBranchObj []model.Branch = getAllBranchRequestResult.Data. ([]model.Branch)
		
	equalBranch := cmp.Equal(createBranchObj.ID, getAllBranchObj[len(getAllBranchObj)-1].ID)
		
	if equalBranch == false {
		t.Error( "Created object is not equal to the last entry in Branch[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Branch
	// --------------------------------------------------------------	
	deleteBranchRequestResult := dao.DeleteBranch(uint64(createBranchObj.ID))

	if deleteBranchRequestResult.Success == false {
			t.Error(deleteBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Branch success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBranchRequestResult = dao.GetBranch( uint64(createBranchObj.ID) )
	
	if getBranchRequestResult.Success == true {
		t.Error(getBranchRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestATMCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ATM
	//----------------------------------------------------------------------------
	ATMObj := model.ATM{
        TerminalId:"test value for TerminalId",
        Location:model.Address{},
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createATMRequestResult := dao.CreateATM( ATMObj )
	
	if createATMRequestResult.Success == false {
		t.Error(createATMRequestResult.Msg)
	} else {
		fmt.Println("Check Create ATM success...")
	}
	
	createATMObj, ok := createATMRequestResult.Data.(model.ATM)

    if !ok {
        t.Fatalf("Expected ATM, got %T", createATMRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check ATM Obj ID
	// --------------------------------------------------------------	
	if createATMObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for ATM" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getATMRequestResult := dao.GetATM( uint64(createATMObj.ID) )
	
	if getATMRequestResult.Success == false {
		t.Error(getATMRequestResult.Msg)
	} else {
		fmt.Println("Check Get ATM success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getATMObj,_ := getATMRequestResult.Data (model.ATM)
	compareATM := cmp.Equal(createATMObj.ID, getATMObj.ID)
	
	if  compareATM == false	{
		t.Error( "Created ATM object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllATMRequestResult := dao.GetAllATM()

	if getAllATMRequestResult.Success == false {
			t.Error(getAllATMRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ATM success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllATMObj []model.ATM = getAllATMRequestResult.Data. ([]model.ATM)
		
	equalATM := cmp.Equal(createATMObj.ID, getAllATMObj[len(getAllATMObj)-1].ID)
		
	if equalATM == false {
		t.Error( "Created object is not equal to the last entry in ATM[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ATM
	// --------------------------------------------------------------	
	deleteATMRequestResult := dao.DeleteATM(uint64(createATMObj.ID))

	if deleteATMRequestResult.Success == false {
			t.Error(deleteATMRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ATM success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getATMRequestResult = dao.GetATM( uint64(createATMObj.ID) )
	
	if getATMRequestResult.Success == true {
		t.Error(getATMRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCustomerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Customer
	//----------------------------------------------------------------------------
	CustomerObj := model.Customer{
        FirstName:"test value for FirstName",
        LastName:"test value for LastName",
        LegalName:"test value for LegalName",
        DateOfBirth:time.Now(),
        TaxId:"test value for TaxId",
        Email:"test value for Email",
        Phone:"test value for Phone",
        Address:model.Address{},
        CustomerType:0,
        RiskRating:0,
        KycStatus:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCustomerRequestResult := dao.CreateCustomer( CustomerObj )
	
	if createCustomerRequestResult.Success == false {
		t.Error(createCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Create Customer success...")
	}
	
	createCustomerObj, ok := createCustomerRequestResult.Data.(model.Customer)

    if !ok {
        t.Fatalf("Expected Customer, got %T", createCustomerRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Customer Obj ID
	// --------------------------------------------------------------	
	if createCustomerObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Customer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCustomerRequestResult := dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == false {
		t.Error(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Get Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCustomerObj,_ := getCustomerRequestResult.Data (model.Customer)
	compareCustomer := cmp.Equal(createCustomerObj.ID, getCustomerObj.ID)
	
	if  compareCustomer == false	{
		t.Error( "Created Customer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCustomerRequestResult := dao.GetAllCustomer()

	if getAllCustomerRequestResult.Success == false {
			t.Error(getAllCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCustomerObj []model.Customer = getAllCustomerRequestResult.Data. ([]model.Customer)
		
	equalCustomer := cmp.Equal(createCustomerObj.ID, getAllCustomerObj[len(getAllCustomerObj)-1].ID)
		
	if equalCustomer == false {
		t.Error( "Created object is not equal to the last entry in Customer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Customer
	// --------------------------------------------------------------	
	deleteCustomerRequestResult := dao.DeleteCustomer(uint64(createCustomerObj.ID))

	if deleteCustomerRequestResult.Success == false {
			t.Error(deleteCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Customer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCustomerRequestResult = dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == true {
		t.Error(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestKycProfileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for KycProfile
	//----------------------------------------------------------------------------
	KycProfileObj := model.KycProfile{
        ProfileId:"test value for ProfileId",
        LastReviewedOn:time.Now(),
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createKycProfileRequestResult := dao.CreateKycProfile( KycProfileObj )
	
	if createKycProfileRequestResult.Success == false {
		t.Error(createKycProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Create KycProfile success...")
	}
	
	createKycProfileObj, ok := createKycProfileRequestResult.Data.(model.KycProfile)

    if !ok {
        t.Fatalf("Expected KycProfile, got %T", createKycProfileRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check KycProfile Obj ID
	// --------------------------------------------------------------	
	if createKycProfileObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for KycProfile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getKycProfileRequestResult := dao.GetKycProfile( uint64(createKycProfileObj.ID) )
	
	if getKycProfileRequestResult.Success == false {
		t.Error(getKycProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Get KycProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getKycProfileObj,_ := getKycProfileRequestResult.Data (model.KycProfile)
	compareKycProfile := cmp.Equal(createKycProfileObj.ID, getKycProfileObj.ID)
	
	if  compareKycProfile == false	{
		t.Error( "Created KycProfile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllKycProfileRequestResult := dao.GetAllKycProfile()

	if getAllKycProfileRequestResult.Success == false {
			t.Error(getAllKycProfileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll KycProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllKycProfileObj []model.KycProfile = getAllKycProfileRequestResult.Data. ([]model.KycProfile)
		
	equalKycProfile := cmp.Equal(createKycProfileObj.ID, getAllKycProfileObj[len(getAllKycProfileObj)-1].ID)
		
	if equalKycProfile == false {
		t.Error( "Created object is not equal to the last entry in KycProfile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for KycProfile
	// --------------------------------------------------------------	
	deleteKycProfileRequestResult := dao.DeleteKycProfile(uint64(createKycProfileObj.ID))

	if deleteKycProfileRequestResult.Success == false {
			t.Error(deleteKycProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion KycProfile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getKycProfileRequestResult = dao.GetKycProfile( uint64(createKycProfileObj.ID) )
	
	if getKycProfileRequestResult.Success == true {
		t.Error(getKycProfileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestIdentityDocumentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for IdentityDocument
	//----------------------------------------------------------------------------
	IdentityDocumentObj := model.IdentityDocument{
        DocumentNumber:"test value for DocumentNumber",
        IssuingCountry:"test value for IssuingCountry",
        ExpirationDate:time.Now(),
        DocumentType:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createIdentityDocumentRequestResult := dao.CreateIdentityDocument( IdentityDocumentObj )
	
	if createIdentityDocumentRequestResult.Success == false {
		t.Error(createIdentityDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Create IdentityDocument success...")
	}
	
	createIdentityDocumentObj, ok := createIdentityDocumentRequestResult.Data.(model.IdentityDocument)

    if !ok {
        t.Fatalf("Expected IdentityDocument, got %T", createIdentityDocumentRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check IdentityDocument Obj ID
	// --------------------------------------------------------------	
	if createIdentityDocumentObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for IdentityDocument" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getIdentityDocumentRequestResult := dao.GetIdentityDocument( uint64(createIdentityDocumentObj.ID) )
	
	if getIdentityDocumentRequestResult.Success == false {
		t.Error(getIdentityDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Get IdentityDocument success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getIdentityDocumentObj,_ := getIdentityDocumentRequestResult.Data (model.IdentityDocument)
	compareIdentityDocument := cmp.Equal(createIdentityDocumentObj.ID, getIdentityDocumentObj.ID)
	
	if  compareIdentityDocument == false	{
		t.Error( "Created IdentityDocument object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllIdentityDocumentRequestResult := dao.GetAllIdentityDocument()

	if getAllIdentityDocumentRequestResult.Success == false {
			t.Error(getAllIdentityDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll IdentityDocument success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllIdentityDocumentObj []model.IdentityDocument = getAllIdentityDocumentRequestResult.Data. ([]model.IdentityDocument)
		
	equalIdentityDocument := cmp.Equal(createIdentityDocumentObj.ID, getAllIdentityDocumentObj[len(getAllIdentityDocumentObj)-1].ID)
		
	if equalIdentityDocument == false {
		t.Error( "Created object is not equal to the last entry in IdentityDocument[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for IdentityDocument
	// --------------------------------------------------------------	
	deleteIdentityDocumentRequestResult := dao.DeleteIdentityDocument(uint64(createIdentityDocumentObj.ID))

	if deleteIdentityDocumentRequestResult.Success == false {
			t.Error(deleteIdentityDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion IdentityDocument success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getIdentityDocumentRequestResult = dao.GetIdentityDocument( uint64(createIdentityDocumentObj.ID) )
	
	if getIdentityDocumentRequestResult.Success == true {
		t.Error(getIdentityDocumentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRiskAssessmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RiskAssessment
	//----------------------------------------------------------------------------
	RiskAssessmentObj := model.RiskAssessment{
        Score:100,
        AssessedOn:time.Now(),
        Rating:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRiskAssessmentRequestResult := dao.CreateRiskAssessment( RiskAssessmentObj )
	
	if createRiskAssessmentRequestResult.Success == false {
		t.Error(createRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create RiskAssessment success...")
	}
	
	createRiskAssessmentObj, ok := createRiskAssessmentRequestResult.Data.(model.RiskAssessment)

    if !ok {
        t.Fatalf("Expected RiskAssessment, got %T", createRiskAssessmentRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check RiskAssessment Obj ID
	// --------------------------------------------------------------	
	if createRiskAssessmentObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for RiskAssessment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRiskAssessmentRequestResult := dao.GetRiskAssessment( uint64(createRiskAssessmentObj.ID) )
	
	if getRiskAssessmentRequestResult.Success == false {
		t.Error(getRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get RiskAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRiskAssessmentObj,_ := getRiskAssessmentRequestResult.Data (model.RiskAssessment)
	compareRiskAssessment := cmp.Equal(createRiskAssessmentObj.ID, getRiskAssessmentObj.ID)
	
	if  compareRiskAssessment == false	{
		t.Error( "Created RiskAssessment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRiskAssessmentRequestResult := dao.GetAllRiskAssessment()

	if getAllRiskAssessmentRequestResult.Success == false {
			t.Error(getAllRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RiskAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRiskAssessmentObj []model.RiskAssessment = getAllRiskAssessmentRequestResult.Data. ([]model.RiskAssessment)
		
	equalRiskAssessment := cmp.Equal(createRiskAssessmentObj.ID, getAllRiskAssessmentObj[len(getAllRiskAssessmentObj)-1].ID)
		
	if equalRiskAssessment == false {
		t.Error( "Created object is not equal to the last entry in RiskAssessment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RiskAssessment
	// --------------------------------------------------------------	
	deleteRiskAssessmentRequestResult := dao.DeleteRiskAssessment(uint64(createRiskAssessmentObj.ID))

	if deleteRiskAssessmentRequestResult.Success == false {
			t.Error(deleteRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RiskAssessment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRiskAssessmentRequestResult = dao.GetRiskAssessment( uint64(createRiskAssessmentObj.ID) )
	
	if getRiskAssessmentRequestResult.Success == true {
		t.Error(getRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestScreeningResultCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ScreeningResult
	//----------------------------------------------------------------------------
	ScreeningResultObj := model.ScreeningResult{
        ScreeningDate:time.Now(),
        Provider:"test value for Provider",
        Outcome:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createScreeningResultRequestResult := dao.CreateScreeningResult( ScreeningResultObj )
	
	if createScreeningResultRequestResult.Success == false {
		t.Error(createScreeningResultRequestResult.Msg)
	} else {
		fmt.Println("Check Create ScreeningResult success...")
	}
	
	createScreeningResultObj, ok := createScreeningResultRequestResult.Data.(model.ScreeningResult)

    if !ok {
        t.Fatalf("Expected ScreeningResult, got %T", createScreeningResultRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check ScreeningResult Obj ID
	// --------------------------------------------------------------	
	if createScreeningResultObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for ScreeningResult" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getScreeningResultRequestResult := dao.GetScreeningResult( uint64(createScreeningResultObj.ID) )
	
	if getScreeningResultRequestResult.Success == false {
		t.Error(getScreeningResultRequestResult.Msg)
	} else {
		fmt.Println("Check Get ScreeningResult success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getScreeningResultObj,_ := getScreeningResultRequestResult.Data (model.ScreeningResult)
	compareScreeningResult := cmp.Equal(createScreeningResultObj.ID, getScreeningResultObj.ID)
	
	if  compareScreeningResult == false	{
		t.Error( "Created ScreeningResult object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllScreeningResultRequestResult := dao.GetAllScreeningResult()

	if getAllScreeningResultRequestResult.Success == false {
			t.Error(getAllScreeningResultRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ScreeningResult success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllScreeningResultObj []model.ScreeningResult = getAllScreeningResultRequestResult.Data. ([]model.ScreeningResult)
		
	equalScreeningResult := cmp.Equal(createScreeningResultObj.ID, getAllScreeningResultObj[len(getAllScreeningResultObj)-1].ID)
		
	if equalScreeningResult == false {
		t.Error( "Created object is not equal to the last entry in ScreeningResult[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ScreeningResult
	// --------------------------------------------------------------	
	deleteScreeningResultRequestResult := dao.DeleteScreeningResult(uint64(createScreeningResultObj.ID))

	if deleteScreeningResultRequestResult.Success == false {
			t.Error(deleteScreeningResultRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ScreeningResult success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getScreeningResultRequestResult = dao.GetScreeningResult( uint64(createScreeningResultObj.ID) )
	
	if getScreeningResultRequestResult.Success == true {
		t.Error(getScreeningResultRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBankingProductCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BankingProduct
	//----------------------------------------------------------------------------
	BankingProductObj := model.BankingProduct{
        ProductCode:"test value for ProductCode",
        Name:"test value for Name",
        Description:"test value for Description",
        ProductCategory:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBankingProductRequestResult := dao.CreateBankingProduct( BankingProductObj )
	
	if createBankingProductRequestResult.Success == false {
		t.Error(createBankingProductRequestResult.Msg)
	} else {
		fmt.Println("Check Create BankingProduct success...")
	}
	
	createBankingProductObj, ok := createBankingProductRequestResult.Data.(model.BankingProduct)

    if !ok {
        t.Fatalf("Expected BankingProduct, got %T", createBankingProductRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check BankingProduct Obj ID
	// --------------------------------------------------------------	
	if createBankingProductObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for BankingProduct" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBankingProductRequestResult := dao.GetBankingProduct( uint64(createBankingProductObj.ID) )
	
	if getBankingProductRequestResult.Success == false {
		t.Error(getBankingProductRequestResult.Msg)
	} else {
		fmt.Println("Check Get BankingProduct success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBankingProductObj,_ := getBankingProductRequestResult.Data (model.BankingProduct)
	compareBankingProduct := cmp.Equal(createBankingProductObj.ID, getBankingProductObj.ID)
	
	if  compareBankingProduct == false	{
		t.Error( "Created BankingProduct object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBankingProductRequestResult := dao.GetAllBankingProduct()

	if getAllBankingProductRequestResult.Success == false {
			t.Error(getAllBankingProductRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BankingProduct success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBankingProductObj []model.BankingProduct = getAllBankingProductRequestResult.Data. ([]model.BankingProduct)
		
	equalBankingProduct := cmp.Equal(createBankingProductObj.ID, getAllBankingProductObj[len(getAllBankingProductObj)-1].ID)
		
	if equalBankingProduct == false {
		t.Error( "Created object is not equal to the last entry in BankingProduct[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BankingProduct
	// --------------------------------------------------------------	
	deleteBankingProductRequestResult := dao.DeleteBankingProduct(uint64(createBankingProductObj.ID))

	if deleteBankingProductRequestResult.Success == false {
			t.Error(deleteBankingProductRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BankingProduct success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBankingProductRequestResult = dao.GetBankingProduct( uint64(createBankingProductObj.ID) )
	
	if getBankingProductRequestResult.Success == true {
		t.Error(getBankingProductRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Account
	//----------------------------------------------------------------------------
	AccountObj := model.Account{
        AccountNumber:model.AccountNumber{},
        Iban:model.IBAN{},
        AccountName:"test value for AccountName",
        Currency:"test value for Currency",
        OpenedOn:time.Now(),
        ClosedOn:time.Now(),
        AccountType:0,
        OwnershipType:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAccountRequestResult := dao.CreateAccount( AccountObj )
	
	if createAccountRequestResult.Success == false {
		t.Error(createAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create Account success...")
	}
	
	createAccountObj, ok := createAccountRequestResult.Data.(model.Account)

    if !ok {
        t.Fatalf("Expected Account, got %T", createAccountRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Account Obj ID
	// --------------------------------------------------------------	
	if createAccountObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Account" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAccountRequestResult := dao.GetAccount( uint64(createAccountObj.ID) )
	
	if getAccountRequestResult.Success == false {
		t.Error(getAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get Account success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAccountObj,_ := getAccountRequestResult.Data (model.Account)
	compareAccount := cmp.Equal(createAccountObj.ID, getAccountObj.ID)
	
	if  compareAccount == false	{
		t.Error( "Created Account object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAccountRequestResult := dao.GetAllAccount()

	if getAllAccountRequestResult.Success == false {
			t.Error(getAllAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Account success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAccountObj []model.Account = getAllAccountRequestResult.Data. ([]model.Account)
		
	equalAccount := cmp.Equal(createAccountObj.ID, getAllAccountObj[len(getAllAccountObj)-1].ID)
		
	if equalAccount == false {
		t.Error( "Created object is not equal to the last entry in Account[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Account
	// --------------------------------------------------------------	
	deleteAccountRequestResult := dao.DeleteAccount(uint64(createAccountObj.ID))

	if deleteAccountRequestResult.Success == false {
			t.Error(deleteAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Account success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAccountRequestResult = dao.GetAccount( uint64(createAccountObj.ID) )
	
	if getAccountRequestResult.Success == true {
		t.Error(getAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAccountStatementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AccountStatement
	//----------------------------------------------------------------------------
	AccountStatementObj := model.AccountStatement{
        StatementNumber:"test value for StatementNumber",
        PeriodStart:time.Now(),
        PeriodEnd:time.Now(),
        OpeningBalance:model.Money{},
        ClosingBalance:model.Money{},
        DeliveryMethod:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAccountStatementRequestResult := dao.CreateAccountStatement( AccountStatementObj )
	
	if createAccountStatementRequestResult.Success == false {
		t.Error(createAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Create AccountStatement success...")
	}
	
	createAccountStatementObj, ok := createAccountStatementRequestResult.Data.(model.AccountStatement)

    if !ok {
        t.Fatalf("Expected AccountStatement, got %T", createAccountStatementRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check AccountStatement Obj ID
	// --------------------------------------------------------------	
	if createAccountStatementObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for AccountStatement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAccountStatementRequestResult := dao.GetAccountStatement( uint64(createAccountStatementObj.ID) )
	
	if getAccountStatementRequestResult.Success == false {
		t.Error(getAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Get AccountStatement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAccountStatementObj,_ := getAccountStatementRequestResult.Data (model.AccountStatement)
	compareAccountStatement := cmp.Equal(createAccountStatementObj.ID, getAccountStatementObj.ID)
	
	if  compareAccountStatement == false	{
		t.Error( "Created AccountStatement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAccountStatementRequestResult := dao.GetAllAccountStatement()

	if getAllAccountStatementRequestResult.Success == false {
			t.Error(getAllAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AccountStatement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAccountStatementObj []model.AccountStatement = getAllAccountStatementRequestResult.Data. ([]model.AccountStatement)
		
	equalAccountStatement := cmp.Equal(createAccountStatementObj.ID, getAllAccountStatementObj[len(getAllAccountStatementObj)-1].ID)
		
	if equalAccountStatement == false {
		t.Error( "Created object is not equal to the last entry in AccountStatement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AccountStatement
	// --------------------------------------------------------------	
	deleteAccountStatementRequestResult := dao.DeleteAccountStatement(uint64(createAccountStatementObj.ID))

	if deleteAccountStatementRequestResult.Success == false {
			t.Error(deleteAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AccountStatement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAccountStatementRequestResult = dao.GetAccountStatement( uint64(createAccountStatementObj.ID) )
	
	if getAccountStatementRequestResult.Success == true {
		t.Error(getAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTransactionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Transaction
	//----------------------------------------------------------------------------
	TransactionObj := model.Transaction{
        BookingDate:time.Now(),
        ValueDate:time.Now(),
        Amount:model.Money{},
        Description:"test value for Description",
        Direction:0,
        TransactionType:0,
        Status:0,
        Channel:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTransactionRequestResult := dao.CreateTransaction( TransactionObj )
	
	if createTransactionRequestResult.Success == false {
		t.Error(createTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Transaction success...")
	}
	
	createTransactionObj, ok := createTransactionRequestResult.Data.(model.Transaction)

    if !ok {
        t.Fatalf("Expected Transaction, got %T", createTransactionRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Transaction Obj ID
	// --------------------------------------------------------------	
	if createTransactionObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Transaction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTransactionRequestResult := dao.GetTransaction( uint64(createTransactionObj.ID) )
	
	if getTransactionRequestResult.Success == false {
		t.Error(getTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Transaction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTransactionObj,_ := getTransactionRequestResult.Data (model.Transaction)
	compareTransaction := cmp.Equal(createTransactionObj.ID, getTransactionObj.ID)
	
	if  compareTransaction == false	{
		t.Error( "Created Transaction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTransactionRequestResult := dao.GetAllTransaction()

	if getAllTransactionRequestResult.Success == false {
			t.Error(getAllTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Transaction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTransactionObj []model.Transaction = getAllTransactionRequestResult.Data. ([]model.Transaction)
		
	equalTransaction := cmp.Equal(createTransactionObj.ID, getAllTransactionObj[len(getAllTransactionObj)-1].ID)
		
	if equalTransaction == false {
		t.Error( "Created object is not equal to the last entry in Transaction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Transaction
	// --------------------------------------------------------------	
	deleteTransactionRequestResult := dao.DeleteTransaction(uint64(createTransactionObj.ID))

	if deleteTransactionRequestResult.Success == false {
			t.Error(deleteTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Transaction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTransactionRequestResult = dao.GetTransaction( uint64(createTransactionObj.ID) )
	
	if getTransactionRequestResult.Success == true {
		t.Error(getTransactionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExternalAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ExternalAccount
	//----------------------------------------------------------------------------
	ExternalAccountObj := model.ExternalAccount{
        Name:"test value for Name",
        Iban:model.IBAN{},
        AccountNumber:model.AccountNumber{},
        Bic:model.BIC{},
        BankName:"test value for BankName",
        Country:"test value for Country",
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExternalAccountRequestResult := dao.CreateExternalAccount( ExternalAccountObj )
	
	if createExternalAccountRequestResult.Success == false {
		t.Error(createExternalAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create ExternalAccount success...")
	}
	
	createExternalAccountObj, ok := createExternalAccountRequestResult.Data.(model.ExternalAccount)

    if !ok {
        t.Fatalf("Expected ExternalAccount, got %T", createExternalAccountRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check ExternalAccount Obj ID
	// --------------------------------------------------------------	
	if createExternalAccountObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for ExternalAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExternalAccountRequestResult := dao.GetExternalAccount( uint64(createExternalAccountObj.ID) )
	
	if getExternalAccountRequestResult.Success == false {
		t.Error(getExternalAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get ExternalAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExternalAccountObj,_ := getExternalAccountRequestResult.Data (model.ExternalAccount)
	compareExternalAccount := cmp.Equal(createExternalAccountObj.ID, getExternalAccountObj.ID)
	
	if  compareExternalAccount == false	{
		t.Error( "Created ExternalAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExternalAccountRequestResult := dao.GetAllExternalAccount()

	if getAllExternalAccountRequestResult.Success == false {
			t.Error(getAllExternalAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ExternalAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExternalAccountObj []model.ExternalAccount = getAllExternalAccountRequestResult.Data. ([]model.ExternalAccount)
		
	equalExternalAccount := cmp.Equal(createExternalAccountObj.ID, getAllExternalAccountObj[len(getAllExternalAccountObj)-1].ID)
		
	if equalExternalAccount == false {
		t.Error( "Created object is not equal to the last entry in ExternalAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ExternalAccount
	// --------------------------------------------------------------	
	deleteExternalAccountRequestResult := dao.DeleteExternalAccount(uint64(createExternalAccountObj.ID))

	if deleteExternalAccountRequestResult.Success == false {
			t.Error(deleteExternalAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ExternalAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExternalAccountRequestResult = dao.GetExternalAccount( uint64(createExternalAccountObj.ID) )
	
	if getExternalAccountRequestResult.Success == true {
		t.Error(getExternalAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFundsTransferCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FundsTransfer
	//----------------------------------------------------------------------------
	FundsTransferObj := model.FundsTransfer{
        TransferReference:"test value for TransferReference",
        Amount:model.Money{},
        RequestedDate:time.Now(),
        ExecutionDate:time.Now(),
        Purpose:"test value for Purpose",
        FeeAmount:model.Money{},
        Method:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFundsTransferRequestResult := dao.CreateFundsTransfer( FundsTransferObj )
	
	if createFundsTransferRequestResult.Success == false {
		t.Error(createFundsTransferRequestResult.Msg)
	} else {
		fmt.Println("Check Create FundsTransfer success...")
	}
	
	createFundsTransferObj, ok := createFundsTransferRequestResult.Data.(model.FundsTransfer)

    if !ok {
        t.Fatalf("Expected FundsTransfer, got %T", createFundsTransferRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check FundsTransfer Obj ID
	// --------------------------------------------------------------	
	if createFundsTransferObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for FundsTransfer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFundsTransferRequestResult := dao.GetFundsTransfer( uint64(createFundsTransferObj.ID) )
	
	if getFundsTransferRequestResult.Success == false {
		t.Error(getFundsTransferRequestResult.Msg)
	} else {
		fmt.Println("Check Get FundsTransfer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFundsTransferObj,_ := getFundsTransferRequestResult.Data (model.FundsTransfer)
	compareFundsTransfer := cmp.Equal(createFundsTransferObj.ID, getFundsTransferObj.ID)
	
	if  compareFundsTransfer == false	{
		t.Error( "Created FundsTransfer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFundsTransferRequestResult := dao.GetAllFundsTransfer()

	if getAllFundsTransferRequestResult.Success == false {
			t.Error(getAllFundsTransferRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FundsTransfer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFundsTransferObj []model.FundsTransfer = getAllFundsTransferRequestResult.Data. ([]model.FundsTransfer)
		
	equalFundsTransfer := cmp.Equal(createFundsTransferObj.ID, getAllFundsTransferObj[len(getAllFundsTransferObj)-1].ID)
		
	if equalFundsTransfer == false {
		t.Error( "Created object is not equal to the last entry in FundsTransfer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FundsTransfer
	// --------------------------------------------------------------	
	deleteFundsTransferRequestResult := dao.DeleteFundsTransfer(uint64(createFundsTransferObj.ID))

	if deleteFundsTransferRequestResult.Success == false {
			t.Error(deleteFundsTransferRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FundsTransfer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFundsTransferRequestResult = dao.GetFundsTransfer( uint64(createFundsTransferObj.ID) )
	
	if getFundsTransferRequestResult.Success == true {
		t.Error(getFundsTransferRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestStandingInstructionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for StandingInstruction
	//----------------------------------------------------------------------------
	StandingInstructionObj := model.StandingInstruction{
        InstructionId:"test value for InstructionId",
        Amount:model.Money{},
        NextExecutionDate:time.Now(),
        Frequency:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createStandingInstructionRequestResult := dao.CreateStandingInstruction( StandingInstructionObj )
	
	if createStandingInstructionRequestResult.Success == false {
		t.Error(createStandingInstructionRequestResult.Msg)
	} else {
		fmt.Println("Check Create StandingInstruction success...")
	}
	
	createStandingInstructionObj, ok := createStandingInstructionRequestResult.Data.(model.StandingInstruction)

    if !ok {
        t.Fatalf("Expected StandingInstruction, got %T", createStandingInstructionRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check StandingInstruction Obj ID
	// --------------------------------------------------------------	
	if createStandingInstructionObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for StandingInstruction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getStandingInstructionRequestResult := dao.GetStandingInstruction( uint64(createStandingInstructionObj.ID) )
	
	if getStandingInstructionRequestResult.Success == false {
		t.Error(getStandingInstructionRequestResult.Msg)
	} else {
		fmt.Println("Check Get StandingInstruction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getStandingInstructionObj,_ := getStandingInstructionRequestResult.Data (model.StandingInstruction)
	compareStandingInstruction := cmp.Equal(createStandingInstructionObj.ID, getStandingInstructionObj.ID)
	
	if  compareStandingInstruction == false	{
		t.Error( "Created StandingInstruction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllStandingInstructionRequestResult := dao.GetAllStandingInstruction()

	if getAllStandingInstructionRequestResult.Success == false {
			t.Error(getAllStandingInstructionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll StandingInstruction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllStandingInstructionObj []model.StandingInstruction = getAllStandingInstructionRequestResult.Data. ([]model.StandingInstruction)
		
	equalStandingInstruction := cmp.Equal(createStandingInstructionObj.ID, getAllStandingInstructionObj[len(getAllStandingInstructionObj)-1].ID)
		
	if equalStandingInstruction == false {
		t.Error( "Created object is not equal to the last entry in StandingInstruction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for StandingInstruction
	// --------------------------------------------------------------	
	deleteStandingInstructionRequestResult := dao.DeleteStandingInstruction(uint64(createStandingInstructionObj.ID))

	if deleteStandingInstructionRequestResult.Success == false {
			t.Error(deleteStandingInstructionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion StandingInstruction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getStandingInstructionRequestResult = dao.GetStandingInstruction( uint64(createStandingInstructionObj.ID) )
	
	if getStandingInstructionRequestResult.Success == true {
		t.Error(getStandingInstructionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentCardCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentCard
	//----------------------------------------------------------------------------
	PaymentCardObj := model.PaymentCard{
        CardNumber:model.CardPAN{},
        EmbossedName:"test value for EmbossedName",
        ExpiryMonth:100,
        ExpiryYear:100,
        CardType:0,
        CardStatus:0,
        Network:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentCardRequestResult := dao.CreatePaymentCard( PaymentCardObj )
	
	if createPaymentCardRequestResult.Success == false {
		t.Error(createPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentCard success...")
	}
	
	createPaymentCardObj, ok := createPaymentCardRequestResult.Data.(model.PaymentCard)

    if !ok {
        t.Fatalf("Expected PaymentCard, got %T", createPaymentCardRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check PaymentCard Obj ID
	// --------------------------------------------------------------	
	if createPaymentCardObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for PaymentCard" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentCardRequestResult := dao.GetPaymentCard( uint64(createPaymentCardObj.ID) )
	
	if getPaymentCardRequestResult.Success == false {
		t.Error(getPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentCard success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentCardObj,_ := getPaymentCardRequestResult.Data (model.PaymentCard)
	comparePaymentCard := cmp.Equal(createPaymentCardObj.ID, getPaymentCardObj.ID)
	
	if  comparePaymentCard == false	{
		t.Error( "Created PaymentCard object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentCardRequestResult := dao.GetAllPaymentCard()

	if getAllPaymentCardRequestResult.Success == false {
			t.Error(getAllPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentCard success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentCardObj []model.PaymentCard = getAllPaymentCardRequestResult.Data. ([]model.PaymentCard)
		
	equalPaymentCard := cmp.Equal(createPaymentCardObj.ID, getAllPaymentCardObj[len(getAllPaymentCardObj)-1].ID)
		
	if equalPaymentCard == false {
		t.Error( "Created object is not equal to the last entry in PaymentCard[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentCard
	// --------------------------------------------------------------	
	deletePaymentCardRequestResult := dao.DeletePaymentCard(uint64(createPaymentCardObj.ID))

	if deletePaymentCardRequestResult.Success == false {
			t.Error(deletePaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentCard success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentCardRequestResult = dao.GetPaymentCard( uint64(createPaymentCardObj.ID) )
	
	if getPaymentCardRequestResult.Success == true {
		t.Error(getPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLoanAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LoanAccount
	//----------------------------------------------------------------------------
	LoanAccountObj := model.LoanAccount{
        LoanNumber:"test value for LoanNumber",
        PrincipalAmount:model.Money{},
        OutstandingPrincipal:model.Money{},
        InterestRate:model.Percentage{},
        OriginationDate:time.Now(),
        MaturityDate:time.Now(),
        PaymentDayOfMonth:100,
        Currency:"test value for Currency",
        LoanType:0,
        RateType:0,
        Compounding:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLoanAccountRequestResult := dao.CreateLoanAccount( LoanAccountObj )
	
	if createLoanAccountRequestResult.Success == false {
		t.Error(createLoanAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create LoanAccount success...")
	}
	
	createLoanAccountObj, ok := createLoanAccountRequestResult.Data.(model.LoanAccount)

    if !ok {
        t.Fatalf("Expected LoanAccount, got %T", createLoanAccountRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check LoanAccount Obj ID
	// --------------------------------------------------------------	
	if createLoanAccountObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for LoanAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLoanAccountRequestResult := dao.GetLoanAccount( uint64(createLoanAccountObj.ID) )
	
	if getLoanAccountRequestResult.Success == false {
		t.Error(getLoanAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get LoanAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLoanAccountObj,_ := getLoanAccountRequestResult.Data (model.LoanAccount)
	compareLoanAccount := cmp.Equal(createLoanAccountObj.ID, getLoanAccountObj.ID)
	
	if  compareLoanAccount == false	{
		t.Error( "Created LoanAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLoanAccountRequestResult := dao.GetAllLoanAccount()

	if getAllLoanAccountRequestResult.Success == false {
			t.Error(getAllLoanAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LoanAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLoanAccountObj []model.LoanAccount = getAllLoanAccountRequestResult.Data. ([]model.LoanAccount)
		
	equalLoanAccount := cmp.Equal(createLoanAccountObj.ID, getAllLoanAccountObj[len(getAllLoanAccountObj)-1].ID)
		
	if equalLoanAccount == false {
		t.Error( "Created object is not equal to the last entry in LoanAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LoanAccount
	// --------------------------------------------------------------	
	deleteLoanAccountRequestResult := dao.DeleteLoanAccount(uint64(createLoanAccountObj.ID))

	if deleteLoanAccountRequestResult.Success == false {
			t.Error(deleteLoanAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LoanAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLoanAccountRequestResult = dao.GetLoanAccount( uint64(createLoanAccountObj.ID) )
	
	if getLoanAccountRequestResult.Success == true {
		t.Error(getLoanAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRepaymentScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RepaymentSchedule
	//----------------------------------------------------------------------------
	RepaymentScheduleObj := model.RepaymentSchedule{
        InstallmentNumber:100,
        DueDate:time.Now(),
        PrincipalDue:model.Money{},
        InterestDue:model.Money{},
        TotalDue:model.Money{},
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRepaymentScheduleRequestResult := dao.CreateRepaymentSchedule( RepaymentScheduleObj )
	
	if createRepaymentScheduleRequestResult.Success == false {
		t.Error(createRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create RepaymentSchedule success...")
	}
	
	createRepaymentScheduleObj, ok := createRepaymentScheduleRequestResult.Data.(model.RepaymentSchedule)

    if !ok {
        t.Fatalf("Expected RepaymentSchedule, got %T", createRepaymentScheduleRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check RepaymentSchedule Obj ID
	// --------------------------------------------------------------	
	if createRepaymentScheduleObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for RepaymentSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRepaymentScheduleRequestResult := dao.GetRepaymentSchedule( uint64(createRepaymentScheduleObj.ID) )
	
	if getRepaymentScheduleRequestResult.Success == false {
		t.Error(getRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get RepaymentSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRepaymentScheduleObj,_ := getRepaymentScheduleRequestResult.Data (model.RepaymentSchedule)
	compareRepaymentSchedule := cmp.Equal(createRepaymentScheduleObj.ID, getRepaymentScheduleObj.ID)
	
	if  compareRepaymentSchedule == false	{
		t.Error( "Created RepaymentSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRepaymentScheduleRequestResult := dao.GetAllRepaymentSchedule()

	if getAllRepaymentScheduleRequestResult.Success == false {
			t.Error(getAllRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RepaymentSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRepaymentScheduleObj []model.RepaymentSchedule = getAllRepaymentScheduleRequestResult.Data. ([]model.RepaymentSchedule)
		
	equalRepaymentSchedule := cmp.Equal(createRepaymentScheduleObj.ID, getAllRepaymentScheduleObj[len(getAllRepaymentScheduleObj)-1].ID)
		
	if equalRepaymentSchedule == false {
		t.Error( "Created object is not equal to the last entry in RepaymentSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RepaymentSchedule
	// --------------------------------------------------------------	
	deleteRepaymentScheduleRequestResult := dao.DeleteRepaymentSchedule(uint64(createRepaymentScheduleObj.ID))

	if deleteRepaymentScheduleRequestResult.Success == false {
			t.Error(deleteRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RepaymentSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRepaymentScheduleRequestResult = dao.GetRepaymentSchedule( uint64(createRepaymentScheduleObj.ID) )
	
	if getRepaymentScheduleRequestResult.Success == true {
		t.Error(getRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLoanPaymentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LoanPayment
	//----------------------------------------------------------------------------
	LoanPaymentObj := model.LoanPayment{
        PaymentReference:"test value for PaymentReference",
        Amount:model.Money{},
        PaymentDate:time.Now(),
        Method:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLoanPaymentRequestResult := dao.CreateLoanPayment( LoanPaymentObj )
	
	if createLoanPaymentRequestResult.Success == false {
		t.Error(createLoanPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Create LoanPayment success...")
	}
	
	createLoanPaymentObj, ok := createLoanPaymentRequestResult.Data.(model.LoanPayment)

    if !ok {
        t.Fatalf("Expected LoanPayment, got %T", createLoanPaymentRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check LoanPayment Obj ID
	// --------------------------------------------------------------	
	if createLoanPaymentObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for LoanPayment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLoanPaymentRequestResult := dao.GetLoanPayment( uint64(createLoanPaymentObj.ID) )
	
	if getLoanPaymentRequestResult.Success == false {
		t.Error(getLoanPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Get LoanPayment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLoanPaymentObj,_ := getLoanPaymentRequestResult.Data (model.LoanPayment)
	compareLoanPayment := cmp.Equal(createLoanPaymentObj.ID, getLoanPaymentObj.ID)
	
	if  compareLoanPayment == false	{
		t.Error( "Created LoanPayment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLoanPaymentRequestResult := dao.GetAllLoanPayment()

	if getAllLoanPaymentRequestResult.Success == false {
			t.Error(getAllLoanPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LoanPayment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLoanPaymentObj []model.LoanPayment = getAllLoanPaymentRequestResult.Data. ([]model.LoanPayment)
		
	equalLoanPayment := cmp.Equal(createLoanPaymentObj.ID, getAllLoanPaymentObj[len(getAllLoanPaymentObj)-1].ID)
		
	if equalLoanPayment == false {
		t.Error( "Created object is not equal to the last entry in LoanPayment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LoanPayment
	// --------------------------------------------------------------	
	deleteLoanPaymentRequestResult := dao.DeleteLoanPayment(uint64(createLoanPaymentObj.ID))

	if deleteLoanPaymentRequestResult.Success == false {
			t.Error(deleteLoanPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LoanPayment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLoanPaymentRequestResult = dao.GetLoanPayment( uint64(createLoanPaymentObj.ID) )
	
	if getLoanPaymentRequestResult.Success == true {
		t.Error(getLoanPaymentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCollateralCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Collateral
	//----------------------------------------------------------------------------
	CollateralObj := model.Collateral{
        CollateralIdentifier:"test value for CollateralIdentifier",
        AppraisedValue:model.Money{},
        Description:"test value for Description",
        Location:model.Address{},
        CollateralType:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCollateralRequestResult := dao.CreateCollateral( CollateralObj )
	
	if createCollateralRequestResult.Success == false {
		t.Error(createCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Create Collateral success...")
	}
	
	createCollateralObj, ok := createCollateralRequestResult.Data.(model.Collateral)

    if !ok {
        t.Fatalf("Expected Collateral, got %T", createCollateralRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Collateral Obj ID
	// --------------------------------------------------------------	
	if createCollateralObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Collateral" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCollateralRequestResult := dao.GetCollateral( uint64(createCollateralObj.ID) )
	
	if getCollateralRequestResult.Success == false {
		t.Error(getCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Get Collateral success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCollateralObj,_ := getCollateralRequestResult.Data (model.Collateral)
	compareCollateral := cmp.Equal(createCollateralObj.ID, getCollateralObj.ID)
	
	if  compareCollateral == false	{
		t.Error( "Created Collateral object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCollateralRequestResult := dao.GetAllCollateral()

	if getAllCollateralRequestResult.Success == false {
			t.Error(getAllCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Collateral success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCollateralObj []model.Collateral = getAllCollateralRequestResult.Data. ([]model.Collateral)
		
	equalCollateral := cmp.Equal(createCollateralObj.ID, getAllCollateralObj[len(getAllCollateralObj)-1].ID)
		
	if equalCollateral == false {
		t.Error( "Created object is not equal to the last entry in Collateral[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Collateral
	// --------------------------------------------------------------	
	deleteCollateralRequestResult := dao.DeleteCollateral(uint64(createCollateralObj.ID))

	if deleteCollateralRequestResult.Success == false {
			t.Error(deleteCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Collateral success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCollateralRequestResult = dao.GetCollateral( uint64(createCollateralObj.ID) )
	
	if getCollateralRequestResult.Success == true {
		t.Error(getCollateralRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFeeChargeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FeeCharge
	//----------------------------------------------------------------------------
	FeeChargeObj := model.FeeCharge{
        FeeCode:"test value for FeeCode",
        Amount:model.Money{},
        AppliedOn:time.Now(),
        FeeType:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFeeChargeRequestResult := dao.CreateFeeCharge( FeeChargeObj )
	
	if createFeeChargeRequestResult.Success == false {
		t.Error(createFeeChargeRequestResult.Msg)
	} else {
		fmt.Println("Check Create FeeCharge success...")
	}
	
	createFeeChargeObj, ok := createFeeChargeRequestResult.Data.(model.FeeCharge)

    if !ok {
        t.Fatalf("Expected FeeCharge, got %T", createFeeChargeRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check FeeCharge Obj ID
	// --------------------------------------------------------------	
	if createFeeChargeObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for FeeCharge" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFeeChargeRequestResult := dao.GetFeeCharge( uint64(createFeeChargeObj.ID) )
	
	if getFeeChargeRequestResult.Success == false {
		t.Error(getFeeChargeRequestResult.Msg)
	} else {
		fmt.Println("Check Get FeeCharge success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFeeChargeObj,_ := getFeeChargeRequestResult.Data (model.FeeCharge)
	compareFeeCharge := cmp.Equal(createFeeChargeObj.ID, getFeeChargeObj.ID)
	
	if  compareFeeCharge == false	{
		t.Error( "Created FeeCharge object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFeeChargeRequestResult := dao.GetAllFeeCharge()

	if getAllFeeChargeRequestResult.Success == false {
			t.Error(getAllFeeChargeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FeeCharge success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFeeChargeObj []model.FeeCharge = getAllFeeChargeRequestResult.Data. ([]model.FeeCharge)
		
	equalFeeCharge := cmp.Equal(createFeeChargeObj.ID, getAllFeeChargeObj[len(getAllFeeChargeObj)-1].ID)
		
	if equalFeeCharge == false {
		t.Error( "Created object is not equal to the last entry in FeeCharge[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FeeCharge
	// --------------------------------------------------------------	
	deleteFeeChargeRequestResult := dao.DeleteFeeCharge(uint64(createFeeChargeObj.ID))

	if deleteFeeChargeRequestResult.Success == false {
			t.Error(deleteFeeChargeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FeeCharge success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFeeChargeRequestResult = dao.GetFeeCharge( uint64(createFeeChargeObj.ID) )
	
	if getFeeChargeRequestResult.Success == true {
		t.Error(getFeeChargeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExchangeRateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ExchangeRate
	//----------------------------------------------------------------------------
	ExchangeRateObj := model.ExchangeRate{
        BaseCurrency:"test value for BaseCurrency",
        CounterCurrency:"test value for CounterCurrency",
        Rate:"test value",
        AsOf:time.Now(),
        Source:"test value for Source",
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExchangeRateRequestResult := dao.CreateExchangeRate( ExchangeRateObj )
	
	if createExchangeRateRequestResult.Success == false {
		t.Error(createExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Create ExchangeRate success...")
	}
	
	createExchangeRateObj, ok := createExchangeRateRequestResult.Data.(model.ExchangeRate)

    if !ok {
        t.Fatalf("Expected ExchangeRate, got %T", createExchangeRateRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check ExchangeRate Obj ID
	// --------------------------------------------------------------	
	if createExchangeRateObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for ExchangeRate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExchangeRateRequestResult := dao.GetExchangeRate( uint64(createExchangeRateObj.ID) )
	
	if getExchangeRateRequestResult.Success == false {
		t.Error(getExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Get ExchangeRate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExchangeRateObj,_ := getExchangeRateRequestResult.Data (model.ExchangeRate)
	compareExchangeRate := cmp.Equal(createExchangeRateObj.ID, getExchangeRateObj.ID)
	
	if  compareExchangeRate == false	{
		t.Error( "Created ExchangeRate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExchangeRateRequestResult := dao.GetAllExchangeRate()

	if getAllExchangeRateRequestResult.Success == false {
			t.Error(getAllExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ExchangeRate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExchangeRateObj []model.ExchangeRate = getAllExchangeRateRequestResult.Data. ([]model.ExchangeRate)
		
	equalExchangeRate := cmp.Equal(createExchangeRateObj.ID, getAllExchangeRateObj[len(getAllExchangeRateObj)-1].ID)
		
	if equalExchangeRate == false {
		t.Error( "Created object is not equal to the last entry in ExchangeRate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ExchangeRate
	// --------------------------------------------------------------	
	deleteExchangeRateRequestResult := dao.DeleteExchangeRate(uint64(createExchangeRateObj.ID))

	if deleteExchangeRateRequestResult.Success == false {
			t.Error(deleteExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ExchangeRate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExchangeRateRequestResult = dao.GetExchangeRate( uint64(createExchangeRateObj.ID) )
	
	if getExchangeRateRequestResult.Success == true {
		t.Error(getExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFXTradeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FXTrade
	//----------------------------------------------------------------------------
	FXTradeObj := model.FXTrade{
        TradeReference:"test value for TradeReference",
        TradeDate:time.Now(),
        SettlementDate:time.Now(),
        AmountSold:model.Money{},
        AmountBought:model.Money{},
        Rate:"test value",
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFXTradeRequestResult := dao.CreateFXTrade( FXTradeObj )
	
	if createFXTradeRequestResult.Success == false {
		t.Error(createFXTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Create FXTrade success...")
	}
	
	createFXTradeObj, ok := createFXTradeRequestResult.Data.(model.FXTrade)

    if !ok {
        t.Fatalf("Expected FXTrade, got %T", createFXTradeRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check FXTrade Obj ID
	// --------------------------------------------------------------	
	if createFXTradeObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for FXTrade" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFXTradeRequestResult := dao.GetFXTrade( uint64(createFXTradeObj.ID) )
	
	if getFXTradeRequestResult.Success == false {
		t.Error(getFXTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Get FXTrade success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFXTradeObj,_ := getFXTradeRequestResult.Data (model.FXTrade)
	compareFXTrade := cmp.Equal(createFXTradeObj.ID, getFXTradeObj.ID)
	
	if  compareFXTrade == false	{
		t.Error( "Created FXTrade object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFXTradeRequestResult := dao.GetAllFXTrade()

	if getAllFXTradeRequestResult.Success == false {
			t.Error(getAllFXTradeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FXTrade success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFXTradeObj []model.FXTrade = getAllFXTradeRequestResult.Data. ([]model.FXTrade)
		
	equalFXTrade := cmp.Equal(createFXTradeObj.ID, getAllFXTradeObj[len(getAllFXTradeObj)-1].ID)
		
	if equalFXTrade == false {
		t.Error( "Created object is not equal to the last entry in FXTrade[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FXTrade
	// --------------------------------------------------------------	
	deleteFXTradeRequestResult := dao.DeleteFXTrade(uint64(createFXTradeObj.ID))

	if deleteFXTradeRequestResult.Success == false {
			t.Error(deleteFXTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FXTrade success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFXTradeRequestResult = dao.GetFXTrade( uint64(createFXTradeObj.ID) )
	
	if getFXTradeRequestResult.Success == true {
		t.Error(getFXTradeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDisputeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Dispute
	//----------------------------------------------------------------------------
	DisputeObj := model.Dispute{
        DisputeReference:"test value for DisputeReference",
        RaisedOn:time.Now(),
        Reason:"test value for Reason",
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDisputeRequestResult := dao.CreateDispute( DisputeObj )
	
	if createDisputeRequestResult.Success == false {
		t.Error(createDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Dispute success...")
	}
	
	createDisputeObj, ok := createDisputeRequestResult.Data.(model.Dispute)

    if !ok {
        t.Fatalf("Expected Dispute, got %T", createDisputeRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Dispute Obj ID
	// --------------------------------------------------------------	
	if createDisputeObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Dispute" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDisputeRequestResult := dao.GetDispute( uint64(createDisputeObj.ID) )
	
	if getDisputeRequestResult.Success == false {
		t.Error(getDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Dispute success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDisputeObj,_ := getDisputeRequestResult.Data (model.Dispute)
	compareDispute := cmp.Equal(createDisputeObj.ID, getDisputeObj.ID)
	
	if  compareDispute == false	{
		t.Error( "Created Dispute object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDisputeRequestResult := dao.GetAllDispute()

	if getAllDisputeRequestResult.Success == false {
			t.Error(getAllDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Dispute success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDisputeObj []model.Dispute = getAllDisputeRequestResult.Data. ([]model.Dispute)
		
	equalDispute := cmp.Equal(createDisputeObj.ID, getAllDisputeObj[len(getAllDisputeObj)-1].ID)
		
	if equalDispute == false {
		t.Error( "Created object is not equal to the last entry in Dispute[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Dispute
	// --------------------------------------------------------------	
	deleteDisputeRequestResult := dao.DeleteDispute(uint64(createDisputeObj.ID))

	if deleteDisputeRequestResult.Success == false {
			t.Error(deleteDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Dispute success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDisputeRequestResult = dao.GetDispute( uint64(createDisputeObj.ID) )
	
	if getDisputeRequestResult.Success == true {
		t.Error(getDisputeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConsentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Consent
	//----------------------------------------------------------------------------
	ConsentObj := model.Consent{
        GrantedOn:time.Now(),
        ExpiresOn:time.Now(),
        ConsentType:0,
        Status:0,
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createConsentRequestResult := dao.CreateConsent( ConsentObj )
	
	if createConsentRequestResult.Success == false {
		t.Error(createConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Consent success...")
	}
	
	createConsentObj, ok := createConsentRequestResult.Data.(model.Consent)

    if !ok {
        t.Fatalf("Expected Consent, got %T", createConsentRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check Consent Obj ID
	// --------------------------------------------------------------	
	if createConsentObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for Consent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getConsentRequestResult := dao.GetConsent( uint64(createConsentObj.ID) )
	
	if getConsentRequestResult.Success == false {
		t.Error(getConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Consent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getConsentObj,_ := getConsentRequestResult.Data (model.Consent)
	compareConsent := cmp.Equal(createConsentObj.ID, getConsentObj.ID)
	
	if  compareConsent == false	{
		t.Error( "Created Consent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllConsentRequestResult := dao.GetAllConsent()

	if getAllConsentRequestResult.Success == false {
			t.Error(getAllConsentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Consent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllConsentObj []model.Consent = getAllConsentRequestResult.Data. ([]model.Consent)
		
	equalConsent := cmp.Equal(createConsentObj.ID, getAllConsentObj[len(getAllConsentObj)-1].ID)
		
	if equalConsent == false {
		t.Error( "Created object is not equal to the last entry in Consent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Consent
	// --------------------------------------------------------------	
	deleteConsentRequestResult := dao.DeleteConsent(uint64(createConsentObj.ID))

	if deleteConsentRequestResult.Success == false {
			t.Error(deleteConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Consent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getConsentRequestResult = dao.GetConsent( uint64(createConsentObj.ID) )
	
	if getConsentRequestResult.Success == true {
		t.Error(getConsentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestThirdPartyProviderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ThirdPartyProvider
	//----------------------------------------------------------------------------
	ThirdPartyProviderObj := model.ThirdPartyProvider{
        Name:"test value for Name",
        RegistrationId:"test value for RegistrationId",
        Website:"test value for Website",
	}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createThirdPartyProviderRequestResult := dao.CreateThirdPartyProvider( ThirdPartyProviderObj )
	
	if createThirdPartyProviderRequestResult.Success == false {
		t.Error(createThirdPartyProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ThirdPartyProvider success...")
	}
	
	createThirdPartyProviderObj, ok := createThirdPartyProviderRequestResult.Data.(model.ThirdPartyProvider)

    if !ok {
        t.Fatalf("Expected ThirdPartyProvider, got %T", createThirdPartyProviderRequestResult.Data)
    }

	// --------------------------------------------------------------
	// Check ThirdPartyProvider Obj ID
	// --------------------------------------------------------------	
	if createThirdPartyProviderObj.ID == 0 {
	    t.Error( "The ORM failed to assign and ID for ThirdPartyProvider" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getThirdPartyProviderRequestResult := dao.GetThirdPartyProvider( uint64(createThirdPartyProviderObj.ID) )
	
	if getThirdPartyProviderRequestResult.Success == false {
		t.Error(getThirdPartyProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ThirdPartyProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getThirdPartyProviderObj,_ := getThirdPartyProviderRequestResult.Data (model.ThirdPartyProvider)
	compareThirdPartyProvider := cmp.Equal(createThirdPartyProviderObj.ID, getThirdPartyProviderObj.ID)
	
	if  compareThirdPartyProvider == false	{
		t.Error( "Created ThirdPartyProvider object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllThirdPartyProviderRequestResult := dao.GetAllThirdPartyProvider()

	if getAllThirdPartyProviderRequestResult.Success == false {
			t.Error(getAllThirdPartyProviderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ThirdPartyProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllThirdPartyProviderObj []model.ThirdPartyProvider = getAllThirdPartyProviderRequestResult.Data. ([]model.ThirdPartyProvider)
		
	equalThirdPartyProvider := cmp.Equal(createThirdPartyProviderObj.ID, getAllThirdPartyProviderObj[len(getAllThirdPartyProviderObj)-1].ID)
		
	if equalThirdPartyProvider == false {
		t.Error( "Created object is not equal to the last entry in ThirdPartyProvider[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ThirdPartyProvider
	// --------------------------------------------------------------	
	deleteThirdPartyProviderRequestResult := dao.DeleteThirdPartyProvider(uint64(createThirdPartyProviderObj.ID))

	if deleteThirdPartyProviderRequestResult.Success == false {
			t.Error(deleteThirdPartyProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ThirdPartyProvider success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getThirdPartyProviderRequestResult = dao.GetThirdPartyProvider( uint64(createThirdPartyProviderObj.ID) )
	
	if getThirdPartyProviderRequestResult.Success == true {
		t.Error(getThirdPartyProviderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

