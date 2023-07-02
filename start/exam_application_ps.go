package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exam_applications_ps"
	//"strconv"
)

// with json array column for div, cadre preferences in same table.
func CreatePSApplications(client *ent.Client, newAppln *ent.Exam_Applications_PS) (*ent.Exam_Applications_PS, error) {
	if newAppln == nil {
		return nil, fmt.Errorf("applicant details are missing")
	}
	ctx := context.Background()

	/*// Get the last generated application number
		lastApplication, err := client.Exam_Applications_IP.
			Query().
			Order(ent.Desc(exam_applications_ip.FieldID)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		if lastApplication != nil {
			lastYear, err := strconv.Atoi(lastApplication.ExamYear)
			if err != nil {
				return nil, fmt.Errorf("invalid ExamYear format")
			}
			year = lastYear
		}
	// Generate the new serial number
	id := int32(1)
	if lastApplication != nil {
		id = lastApplication.ID + 1
	}

	// Generate the new application number
	applicationNo := fmt.Sprintf("IP%d%08d", year, serialNo)*/

	createdAppln, err := client.Exam_Applications_PS.
		Create().
		//SetApplicationNumber(applicationNo).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetDOB(newAppln.DOB).
		SetGender(newAppln.Gender).
		SetMobileNumber(newAppln.MobileNumber).
		SetEmailID(newAppln.EmailID).
		SetEmployeeCategory(newAppln.EmployeeCategory).
		SetCadre(newAppln.Cadre).
		SetEmployeePost(newAppln.EmployeePost).
		SetFacilityID(newAppln.FacilityID).
		SetDCCS(newAppln.DCCS).
		SetDCInPresentCadre(newAppln.DCInPresentCadre).
		SetDeputationOfficeId(newAppln.DeputationOfficeId).
		SetDisabilityType(newAppln.DisabilityType).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetEducation(newAppln.Education).
		SetExamNameCode(newAppln.ExamNameCode).
		SetExamYear(newAppln.ExamYear).
		SetCentrePreference(newAppln.CentrePreference).
		SetSignature(newAppln.Signature).
		SetPhoto(newAppln.Photo).
		SetApplicationStatus("Submitted By Candidate").
		SetApplnSubmittedDate(newAppln.ApplnSubmittedDate).
		SetCadrePreferences(newAppln.CadrePreferences).
		SetDivisionPreferences(newAppln.DivisionPreferences).
		//SetUpdatedAt(newAppln.UpdatedAt).
		//SetUpdatedBy(newAppln.UpdatedBy).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return createdAppln, nil
}

// Query IPExam Application with Emp ID.
func QueryPSExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64) (*ent.Exam_Applications_PS, error) {
	//Can use GetX as well

	newAppln, err := client.Exam_Applications_PS.Query().
		Where(exam_applications_ps.EmployeeIDEQ(empid)).
		First(ctx)

	if err != nil {
		log.Println("error at gettting Emp ID Application Details: ", err)
		return nil, fmt.Errorf("failed querying PS Exam Application details: %w", err)
	}
	log.Println("details returned by PS Exam Applications for the Employee is  : ", newAppln)
	return newAppln, nil
}
