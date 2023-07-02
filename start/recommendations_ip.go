package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exam_applications_ip"

	//"strings"
	"time"
	//"strconv"
)

func UpdateApplicationStatus(client *ent.Client, newAppln *ent.Exam_Applications_IP) (*ent.Exam_Applications_IP, error) {
	ctx := context.Background()

	// Check if newAppln is not nil.
	if newAppln == nil {
		return nil, fmt.Errorf("newAppln is nil")
	}

	// Check if the EmployeeID exists.
	oldAppln, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(newAppln.EmployeeID),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		First(ctx)
	//log.Println("The current Application status is ", oldAppln.ApplicationStatus)
	if err != nil {
		log.Println("Failed to retrieve Employee:", err)
		return nil, fmt.Errorf("failed to retrieve Employee: %v", err)
	}
	log.Println("The selected record is:", oldAppln.ApplicationStatus)
	// Insert a new record with the specified conditions.
	caRemarks := newAppln.CARemarks
	now := time.Now()

	if oldAppln != nil {
		// Update the existing record.
		if oldAppln.ApplicationStatus == "CAVerificationPending" {
			if caRemarks == "InCorrect" {
				updatedAppln, err := oldAppln.
					Update().
					SetApplicationStatus("PendingWithCandidate").
					SetCARemarks(newAppln.CARemarks).
					SetCAUserName(newAppln.CAUserName).
					SetCADate(now).
					SetAppliactionRemarks(newAppln.AppliactionRemarks).
					Save(ctx)

				if err != nil {
					log.Println("Failed to update application:", err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}

				log.Println("Updated application:", updatedAppln.ApplicationStatus)
				return updatedAppln, nil
			} else if caRemarks == "Correct" {
				updatedAppln, err := oldAppln.
					Update().
					SetApplicationStatus("VerifiedByCA").
					SetCARemarks(newAppln.CARemarks).
					SetCAUserName(newAppln.CAUserName).
					SetCADate(now).
					SetAppliactionRemarks(newAppln.AppliactionRemarks).
					Save(ctx)

				if err != nil {
					log.Println("Failed to update application:", err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}

				log.Println("Updated application:", updatedAppln.ApplicationStatus)
				return updatedAppln, nil
			}
		}

		log.Println("The current Application status is ", oldAppln.CARemarks)
		log.Println("The current Application status is ", oldAppln.ApplicationStatus)
	}
	if oldAppln.ApplicationStatus == "PendingWithCandidate" {
		// Insert a new record.
		log.Println("Hi! I am into PendingWithCandidate")
		updatedAppln, err := client.Exam_Applications_IP.
			Create().
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
			SetApplicationStatus("CAVerificationPending").
			SetAppliactionRemarks(newAppln.AppliactionRemarks).
			Save(ctx)

		if err != nil {
			log.Println("Failed to insert application:", err)
			return nil, fmt.Errorf("failed to insert application: %v", err)
		}

		log.Println("Inserted new application:", updatedAppln)
		return updatedAppln, nil
	}

	log.Println("No updates or inserts performed.")
	return oldAppln, nil
}
