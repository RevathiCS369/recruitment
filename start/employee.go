package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
)

func CreateEmployeeProfile(client *ent.Client, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {
	ctx := context.Background()
	u, err := client.Employees.Create().
		SetEmployeedID(newEmployeeprofile.EmployeedID).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemStatus(newEmployeeprofile.IDRemStatus).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetEmployeeName(newEmployeeprofile.EmployeeName).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemStatus(newEmployeeprofile.NameRemStatus).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetEmployeeFathersName(newEmployeeprofile.EmployeeFathersName).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemStatus(newEmployeeprofile.FathersNameRemStatus).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOB(newEmployeeprofile.DOB).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemStatus(newEmployeeprofile.DOBRemStatus).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGender(newEmployeeprofile.Gender).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemStatus(newEmployeeprofile.GenderRemStatus).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetCategoryid(newEmployeeprofile.Categoryid).
		SetEmployeeCategoryCode(newEmployeeprofile.EmployeeCategoryCode).
		SetEmployeeCategory(newEmployeeprofile.EmployeeCategory).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemStatus(newEmployeeprofile.EmployeeCategoryCodeRemStatus).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisability(newEmployeeprofile.WithDisability).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemStatus(newEmployeeprofile.WithDisabilityRemStatus).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityType(newEmployeeprofile.DisabilityType).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemStatus(newEmployeeprofile.DisabilityTypeRemStatus).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentage(newEmployeeprofile.DisabilityPercentage).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemStatus(newEmployeeprofile.DisabilityPercentageRemStatus).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		//SetEmployeeCadre(newEmployeeprofile.EmployeeCadre).
		SetSignature(newEmployeeprofile.Signature).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemStatus(newEmployeeprofile.SignatureRemStatus).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhoto(newEmployeeprofile.Photo).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemStatus(newEmployeeprofile.PhotoRemStatus).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).
		SetCadreid(newEmployeeprofile.Cadreid).
		SetEmployeeCadre(newEmployeeprofile.EmployeeCadre).
		SetEmployeeCadreVerified(newEmployeeprofile.EmployeeCadreVerified).
		SetEmployeeCadreRemStatus(newEmployeeprofile.EmployeeCadreRemStatus).
		SetEmployeeCadreRemarks(newEmployeeprofile.EmployeeCadreRemarks).
		SetDesignationID(newEmployeeprofile.DesignationID).
		SetEmployeeDesignation(newEmployeeprofile.EmployeeDesignation).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemStatus(newEmployeeprofile.EmployeeDesignationRemStatus).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleID(newEmployeeprofile.CircleID).
		SetCircleName(newEmployeeprofile.CircleName).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemStatus(newEmployeeprofile.CircleRemStatus).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionID(newEmployeeprofile.RegionID).
		SetRegionName(newEmployeeprofile.RegionName).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemStatus(newEmployeeprofile.RegionRemStatus).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionID(newEmployeeprofile.DivisionID).
		SetDivisionName(newEmployeeprofile.DivisionName).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemStatus(newEmployeeprofile.DivisionRemStatus).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeID(newEmployeeprofile.OfficeID).
		SetOfficeName(newEmployeeprofile.OfficeName).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemStatus(newEmployeeprofile.OfficeRemStatus).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRole(newEmployeeprofile.Role).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemStatus(newEmployeeprofile.RoleRemStatus).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCS(newEmployeeprofile.DCCS).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemStatus(newEmployeeprofile.DCCSRemStatus).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadre(newEmployeeprofile.DCInPresentCadre).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemStatus(newEmployeeprofile.DCInPresentCadreRemStatus).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorking(newEmployeeprofile.APSWorking).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemStatus(newEmployeeprofile.APSWorkingRemStatus).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(ctx)

	if err != nil {
		log.Println("error at Creating Employees: ", newEmployeeprofile)
		return nil, fmt.Errorf("failed creating Employees: %w", err)
	}
	log.Println("Employees profile is created: ", u)

	return u, nil
}

func QueryEmployeesWithID(ctx context.Context, client *ent.Client, id int32) (*ent.Employees, error) {
	//Can use GetX as well

	employees, err := client.Employees.Get(ctx, id)
	if err != nil {
		log.Println("error at getting EmployeeID: ", err)
		return nil, fmt.Errorf("failed querying Employees: %w", err)
	}
	log.Println("Employee returned: ", employees)
	return employees, nil
}

func QueryEmployees(ctx context.Context, client *ent.Client) ([]*ent.Employees, error) {
	//Array of exams
	newemployee, err := client.Employees.Query().
		All(ctx)
	if err != nil {
		log.Println("error at ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying ExamCalendar: %w", err)
	}
	log.Println("Exam Employees data returned: ", newemployee)
	return newemployee, nil
}

//Update fields by CA

func UpdateVerificationDetails(client *ent.Client, id int32, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {

	ctx := context.Background()
	_, err := QueryEmployeesWithID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedEmployee, err := client.Employees.UpdateOneID(id).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetMobileNumber(newEmployeeprofile.MobileNumber).
		SetEmailID(newEmployeeprofile.EmailID).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).
		SetEmployeeCadreVerified(newEmployeeprofile.EmployeeCadreVerified).
		SetEmployeeCadreRemarks(newEmployeeprofile.EmployeeCadreRemarks).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(context.Background())

	//.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedEmployee, nil
}

// ResubmitDetails by Candidate
func ResubmitCandidateDetails(client *ent.Client, id int32, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {

	ctx := context.Background()
	_, err := QueryEmployeesWithID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedEmployee, err := client.Employees.UpdateOneID(id).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).
		SetEmployeeCadreVerified(newEmployeeprofile.EmployeeCadreVerified).
		SetEmployeeCadreRemarks(newEmployeeprofile.EmployeeCadreRemarks).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(context.Background())

	//.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedEmployee, nil
}
