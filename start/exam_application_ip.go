package start

import (
	"context"
	//"encoding/json"
	//"entgo.io/ent/dialect/sql"
	//"entgo.io/ent/entc/sql/migrate"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exam_applications_ip"

	//"recruit/ent/placeofpreferenceip"
	"recruit/ent/recommendationsipapplications"

	//"strings"
	"time"
	//"strconv"
)

// Create Applications with Circle Preferences ...
func CreateIPApplications(client *ent.Client, newAppln *ent.Exam_Applications_IP) (*ent.Exam_Applications_IP, error) {
	ctx := context.Background()
	if newAppln == nil {
		return nil, fmt.Errorf("newAppln is nil or its ID is nil")
	}

	if newAppln.Edges.CirclePrefRef == nil || len(newAppln.Edges.CirclePrefRef) == 0 {
		return nil, fmt.Errorf("Circle preference values are missing! , Please provide Circle preferences..")
	}

	if len(newAppln.Edges.CirclePrefRef) != 23 {
		return nil, fmt.Errorf("Invalid number of Circle preferences. Must provide preferences for all 23 circles.")
	}
	exists, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending", "PendingWithCandidate", "VerifiedByCA"),
		).
		Exist(ctx)

	if err != nil {
		log.Println("Failed to check existing application:", err)
		return nil, fmt.Errorf("failed to check existing application: %v", err)
	}

	if exists {
		//return nil, fmt.Errorf("The employee has already submitted the application and is pending for verification by CA / Pending with Candidate / Verified/ ")
		return nil, fmt.Errorf("The employee has already submitted the application.")

	}

	// Generate Application number
	// Generate application number in the format "IP2023XXXXXX"

	applicationNumber, err := generateApplicationNumber(client, newAppln.EmployeeID)
	if err != nil {
		log.Printf("Failed to generate application number: %v", err)
		return nil, err
	}

	// Use the generated application number
	log.Printf("Generated application number: %s", applicationNumber)

	// Save the Exam_Applications_IP record.
	createdAppln, err := client.Exam_Applications_IP.
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
		SetFacilityName(newAppln.FacilityName).
		SetCircleName(newAppln.CircleName).
		SetRegionName(newAppln.RegionName).
		SetDivisionName(newAppln.DivisionName).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetEntryCadre(newAppln.EntryCadre).
		SetPresentCadre(newAppln.PresentCadre).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetEligibleCadre(newAppln.EligibleCadre).
		SetEligibleCadreDate(newAppln.EligibleCadreDate).
		SetServiceLength(newAppln.ServiceLength).
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
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("CAVerificationPending").
		Save(ctx)

	if err != nil {
		log.Println("Failed to save Exam_Applications_IP:", err)
		return nil, fmt.Errorf("failed to save Exam_Applications_IP: %v", err)
	}

	// Save the PlaceOfPreferenceIP records.
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CirclePrefRef))
	for i, circlePrefRef := range newAppln.Edges.CirclePrefRef {
		if circlePrefRef == nil {
			return nil, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		//log.Printf("Processing CirclePrefRef[%d]: %+v\n", i, circlePrefRef)
		circlePrefRefEntity, err := client.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(circlePrefRef.PlacePrefNo).
			SetEmployeeID(newAppln.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)

		if err != nil {
			log.Printf("Failed to save PlaceOfPreferenceIP[%d]: %v\n", i, err)
			return nil, fmt.Errorf("failed to save PlaceOfPreferenceIP: %v", err)
		}
		//log.Printf("Saved PlaceOfPreferenceIP[%d]: %+v\n", i, circlePrefRefEntity)
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = createdAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		log.Println("Failed to add PlacePrefRef references:", err)
		return nil, fmt.Errorf("failed to add PlacePrefRef references: %v", err)
	}

	return createdAppln, nil
}

func generateApplicationNumber(client *ent.Client, employeeID int64) (string, error) {
	nextApplicationNumber, err := getNextApplicationNumberFromDatabase(client)
	if err != nil {
		return "", err
	}

	// Get the current year
	currentYear := time.Now().Year()

	// Format the application number as "IPYYYYXXXXXX"
	applicationNumber := fmt.Sprintf("IP%d%06d", currentYear, nextApplicationNumber)

	return applicationNumber, nil
}

func getNextApplicationNumberFromDatabase(client *ent.Client) (int64, error) {
	ctx := context.TODO()
	lastApplication, err := client.Exam_Applications_IP.
		Query().
		//Order(ent.Desc(ent.Exam_Applications_IPColumnID)).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		//Order(ent.Asc(ent.Exam_Applications_IPColumnApplnSubmittedDate)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// No existing applications, start from 100001
			return 100001, nil
		}
		return 0, fmt.Errorf("failed to retrieve last application: %v", err)
	}

	return lastApplication.ID + 1, nil
}

// Query IPExam Application with Emp ID.
func QueryIPExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64) (*ent.Exam_Applications_IP, error) {
	newAppln, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.EmployeeIDEQ(empid)).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		//Order(ent.Asc(ent.Exam_Applications_IPColumnApplnSubmittedDate)).
		WithCirclePrefRef().
		First(ctx)

	if err != nil {
		log.Println("error getting Emp ID Application Details: ", err)
		return nil, fmt.Errorf("failed querying IP Exam Application details: %w", err)
	}

	// Extract only the desired fields from the CirclePrefRef edge
	var circlePrefs []*ent.PlaceOfPreferenceIP
	for _, edge := range newAppln.Edges.CirclePrefRef {
		circlePrefs = append(circlePrefs, &ent.PlaceOfPreferenceIP{
			PlacePrefNo:    edge.PlacePrefNo,
			PlacePrefValue: edge.PlacePrefValue,
		})
	}

	// Update the CirclePrefRef edge with the filtered values
	newAppln.Edges.CirclePrefRef = circlePrefs

	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)

	// log.Println("details returned by IP Exam Applications for the Employee: ", newAppln)
	return newAppln, nil
}

// Update / Verification of IP Exam Application By CA
// Update Resubmission By Candidate.

func UpdateApplicationRemarks(client *ent.Client, newAppln *ent.Exam_Applications_IP) (*ent.Exam_Applications_IP, error) {
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

	// Insert a new record with the specified conditions.
	caRemarks := newAppln.CARemarks
	now := time.Now()

	if oldAppln != nil {
		// Update the existing record.
		if oldAppln.ApplicationStatus == "CAVerificationPending" {
			if caRemarks == "InCorrect" {

				log.Println("Hi ! Into CAVerification Pending , Incorrect remarks")
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
				log.Println("the updated application is:", updatedAppln)
				return updatedAppln, nil
			} else if caRemarks == "Correct" {
				log.Println("Hi ! Into CAVerification Pending , Correct remarks")
				updatedAppln, err := oldAppln.
					Update().
					SetApplicationStatus("VerifiedByCA").
					SetCARemarks(newAppln.CARemarks).
					SetCAUserName(newAppln.CAUserName).
					SetCADate(now).
					SetAppliactionRemarks(newAppln.AppliactionRemarks).
					//SetIPApplicationsRef(recommendationsRef...).
					Save(ctx)

				if err != nil {
					log.Println("Failed to update application:", err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}

				// Insert into recommendations.
				// Save the Recommendation records.
				log.Println("Into insert recommendations")
				recommendationsRef := make([]*ent.RecommendationsIPApplications, len(newAppln.Edges.IPApplicationsRef))
				for i, recommendation := range newAppln.Edges.IPApplicationsRef {
					if recommendation == nil {
						return nil, fmt.Errorf("Recommendations value at index %d is nil", i)
					}

					RecommendationsRefEntity, err := client.RecommendationsIPApplications.
						Create().
						SetApplicationID(updatedAppln.ID).
						SetEmployeeID(updatedAppln.EmployeeID).
						SetExamNameCode(updatedAppln.ExamNameCode).
						SetExamYear(updatedAppln.ExamYear).
						SetVacancyYear(recommendation.VacancyYear).
						SetCARecommendations(recommendation.CARecommendations).
						SetCAUserName(newAppln.CAUserName).     // Use newAppln.CAUserName instead of updatedAppln.CAUserName
						SetCARemarks(recommendation.CARemarks). // Use newAppln.CARemarks instead of updatedAppln.CARemarks
						SetApplicationStatus("VerifiedRecommendationsByCA").
						Save(ctx)
					log.Println("Into insert recommendations: ")
					if err != nil {
						log.Printf("Failed to save Recommendation[%d]: %v\n", i, err)
						return nil, fmt.Errorf("failed to save Recommendation: %v", err)
					}

					recommendationsRef[i] = RecommendationsRefEntity
				}

				updatedAppln.Update().
					ClearIPApplicationsRef().
					AddIPApplicationsRef(recommendationsRef...).
					Save(ctx)
				//log.Println("Updated application:", recommendationsRef)
				return updatedAppln, nil
			}
		}
		if oldAppln.ApplicationStatus == "PendingWithCandidate" {
			if newAppln.Edges.CirclePrefRef == nil || len(newAppln.Edges.CirclePrefRef) == 0 {
				return nil, fmt.Errorf("Circle preference values are missing! Please provide Circle preferences")
			}

			if len(newAppln.Edges.CirclePrefRef) != 23 {
				return nil, fmt.Errorf("Invalid number of Circle preferences. Must provide preferences for all 23 circles.")
			}

			// Insert a new record.
			log.Println("Hi! I am into Candidate Resubmission")
			log.Printf("Generated application number: %s", oldAppln.ApplicationNumber)
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
				SetApplicationNumber(oldAppln.ApplicationNumber).
				SetApplicationStatus("ResubmitCAVerificationPending").
				SetAppliactionRemarks(newAppln.AppliactionRemarks).
				Save(ctx)

			if err != nil {
				log.Println("Failed to save Exam_Applications_IP:", err)
				return nil, fmt.Errorf("save Exam_Applications_IP: %v", err)
			}

			// Save the PlaceOfPreferenceIP records.
			circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CirclePrefRef))
			for i, circlePrefRef := range newAppln.Edges.CirclePrefRef {
				if circlePrefRef == nil {
					return nil, fmt.Errorf("circle preference value at index %d is nil", i)
				}

				//log.Printf("Processing CirclePrefRef[%d]: %+v\n", i, circlePrefRef)
				circlePrefRefEntity, err := client.PlaceOfPreferenceIP.
					Create().
					SetPlacePrefNo(circlePrefRef.PlacePrefNo).
					SetEmployeeID(newAppln.EmployeeID).
					SetPlacePrefValue(circlePrefRef.PlacePrefValue).
					SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
					Save(ctx)

				if err != nil {
					log.Printf("Failed to save PlaceOfPreferenceIP[%d]: %v\n", i, err)
					return nil, fmt.Errorf("failed to save PlaceOfPreferenceIP: %v", err)
				}
				//log.Printf("Saved PlaceOfPreferenceIP[%d]: %+v\n", i, circlePrefRefEntity)
				circlePrefRefs[i] = circlePrefRefEntity
			}

			// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
			_, err = updatedAppln.
				Update().
				AddCirclePrefRef(circlePrefRefs...).
				Save(ctx)

			if err != nil {
				//log.Println("Failed to add PlacePrefRef references:", err)
				return nil, fmt.Errorf("failed to add PlacePrefRef references: %v", err)
			}

			// For Resubmission
			return updatedAppln, nil
		}

		// Resubmit with CA Pending

		if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
			log.Println("Hi , I am now into Resubmit CA Verification..")
			if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
				//previousRemarks, err = GetApplicationRemarksByEmployeeID(ctx, client, oldAppln.EmployeeID)
				if err != nil {
					log.Println("Failed to retrieve previous remarks:", err)
					return nil, fmt.Errorf("failed to retrieve previous remarks: %v", err)
				}
			}
			if caRemarks == "InCorrect" {

				updatedAppln, err := oldAppln.
					Update().
					SetApplicationStatus("PendingWithCandidate").
					SetCARemarks(newAppln.CARemarks).
					SetCAUserName(newAppln.CAUserName).
					SetCADate(now).
					SetAppliactionRemarks(newAppln.AppliactionRemarks).
					//SetCAPreviousRemarks(previousRemarks).
					Save(ctx)

				if err != nil {
					log.Println("Failed to update application:", err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}

				//log.Println("The previous remarks given was :", previousRemarks)
				return updatedAppln, nil

			} else if caRemarks == "Correct" {
				if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
					//previousRemarks, err = GetApplicationRemarksByEmployeeID(ctx, client, oldAppln.EmployeeID)
					if err != nil {
						log.Println("Failed to retrieve previous remarks:", err)
						return nil, fmt.Errorf("failed to retrieve previous remarks: %v", err)
					}
				}
				updatedAppln, err := oldAppln.
					Update().
					SetApplicationStatus("VerifiedByCA").
					SetCARemarks(newAppln.CARemarks).
					SetCAUserName(newAppln.CAUserName).
					SetCADate(now).
					SetAppliactionRemarks(newAppln.AppliactionRemarks).
					//SetCAPreviousRemarks(previousRemarks).
					Save(ctx)

				if err != nil {
					log.Println("Failed to update application:", err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}

				// Insert into recommendations.
				// Save the Recommendation records.
				log.Println("Into insert recommendations")
				recommendationsRef := make([]*ent.RecommendationsIPApplications, len(newAppln.Edges.IPApplicationsRef))
				for i, recommendation := range newAppln.Edges.IPApplicationsRef {
					if recommendation == nil {
						return nil, fmt.Errorf("Recommendations value at index %d is nil", i)
					}

					RecommendationsRefEntity, err := client.RecommendationsIPApplications.
						Create().
						SetApplicationID(updatedAppln.ID).
						SetEmployeeID(updatedAppln.EmployeeID).
						SetExamNameCode(updatedAppln.ExamNameCode).
						SetExamYear(updatedAppln.ExamYear).
						SetVacancyYear(recommendation.VacancyYear).
						SetCARecommendations(recommendation.CARecommendations).
						SetCAUserName(newAppln.CAUserName). //
						SetCARemarks(recommendation.CARemarks).
						SetApplicationStatus("VerifiedRecommendationsByCA"). //

						Save(ctx)

					if err != nil {
						log.Printf("Failed to save Recommendation[%d]: %v\n", i, err)
						return nil, fmt.Errorf("failed to save Recommendation: %v", err)
					}

					recommendationsRef[i] = RecommendationsRefEntity
				}

				updatedAppln.Update().
					ClearIPApplicationsRef().
					AddIPApplicationsRef(recommendationsRef...).
					Save(ctx)
				//log.Println("Updated application:", recommendationsRef)
				return updatedAppln, nil
			}

		}

	}

	log.Println("No updates or inserts performed.")
	return oldAppln, nil
}
func GetApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (string, error) {
	application, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ApplicationStatusEQ("PendingWithCandidate"),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		First(ctx)

	if err != nil {
		return "", fmt.Errorf("failed to retrieve application: %v", err)
	}

	return application.AppliactionRemarks, nil
}

func getInputRecordByVacancyYear(inputRecords []*ent.RecommendationsIPApplications, vacancyYear int32) *ent.RecommendationsIPApplications {
	// Find the corresponding input record based on vacancy year
	for _, record := range inputRecords {
		if record.VacancyYear == vacancyYear {
			return record
		}
	}
	return nil
}

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
func UpdateNodalRecommendationsByEmpID(client *ent.Client, empID int64, newRecommendations []*ent.RecommendationsIPApplications) ([]*ent.RecommendationsIPApplications, error) {
	ctx := context.Background()

	// Check if empID exists
	exists, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if employee with ID %d exists: %v", empID, err)
	}
	if !exists {
		return nil, fmt.Errorf("employee with ID %d does not exist", empID)
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID %d: %v", empID, err)
	}

	if len(newRecommendations) == 0 {
		return nil, fmt.Errorf("input recommendations are empty")
	}

	// Update the records for each vacancy year
	for _, record := range records {
		vacancyYear := record.VacancyYear
		inputRecord := getInputRecordByVacancyYear(newRecommendations, int32(vacancyYear))

		if inputRecord != nil && inputRecord.NORemarks != "" {
			// Update the NO_Recommendations field and set the ApplicationStatus to "RecommendedByNO"
			record.Update().
				SetNORecommendations(inputRecord.NORecommendations).
				SetNOUserName(inputRecord.NOUserName).
				SetNORemarks(inputRecord.NORemarks).
				SetApplicationStatus("VerifiedRecommendationsByNO").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to save updated record for vacancy year %d: %v", vacancyYear, err)
			}

			// Log the input record
			//log.Printf("Input record for Vacancy Year %d: %+v", vacancyYear, inputRecord)
		} else {
			// Log if there is no matching input record or NORemarks is null
			//log.Printf("No update for Vacancy Year %d", vacancyYear)
		}
	}

	// Query the RecommendationsIPApplications table for the specific employee
	record, err := client.RecommendationsIPApplications.Query().
		Where(
			recommendationsipapplications.EmployeeIDEQ(empID),
			recommendationsipapplications.ApplicationStatusEQ("VerifiedRecommendationsByNO"),
		).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve RecommendedByNO record: %v", err)
	}

	if record != nil {
		// Retrieve the corresponding Exam_Applications_IP record using edges
		applicationRecord, err := client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.EmployeeIDEQ(empID),
				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Exam_Applications_IP record: %v", err)
		}

		// Update the Exam_Applications_IP record
		_, err = applicationRecord.Update().
			SetApplicationStatus("VerifiedByNA").
			//SetNARemarks(record.NORemarks).
			SetNAUserName(record.NOUserName).
			SetNADate(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update application status: %v", err)
		}
	}

	// Retrieve all records for the employee ID
	recordsupdated, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID after updation %d: %v", empID, err)
	}

	return recordsupdated, nil
}

// Get All CA Pending records ...
func QueryIPApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client) ([]*ent.Exam_Applications_IP, error) {
	// Array of exams
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.Or(
				exam_applications_ip.ApplicationStatus("CAVerificationPending"),
				exam_applications_ip.ApplicationStatus("ResubmitCAVerificationPending"),
			),
		).
		WithCirclePrefRef(). // Add the Where clause with multiple statuses using Or
		All(ctx)
	if err != nil {
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}
	//log.Println("CA verifications pending returned: ", records)
	return records, nil
}

// Get All CA verified records
func QueryIPApplicationsByCAVerified(ctx context.Context, client *ent.Client) ([]*ent.Exam_Applications_IP, error) {
	// Array of exams
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.ApplicationStatus("VerifiedByCA"), // Add the Where clause
		).
		WithIPApplicationsRef().
		All(ctx)
	if err != nil {
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}
	//log.Println("CA verified records returned: ", records)
	return records, nil
}

// Get CA Verified with Emp ID
func QueryIPApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64) (*ent.Exam_Applications_IP, error) {
	employeeExists, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.EmployeeIDEQ(employeeID)).
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_applications_ip.EmployeeIDEQ(employeeID),
		).
		WithIPApplicationsRef().
		WithCirclePrefRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("record not found for employee ID: %d with 'CAVerified' status", employeeID)
		}
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}

	//log.Println("CA verified record returned: ", record)
	return record, nil
}

// Get CA Pending with EmpID
func QueryIPApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64) ([]*ent.Exam_Applications_IP, error) {
	// Check if employee ID exists
	employeeExists, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.Or(
				exam_applications_ip.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_ip.ApplicationStatusEQ("ResubmitCAVerificationPending"),
			),
		).
		//WithCirclePrefRef().
		//WithIPApplicationsRef().
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: or the verification is not pending with CA %d", empID)
	}

	// Retrieve the record
	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.Or(
				exam_applications_ip.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_ip.ApplicationStatusEQ("ResubmitCAVerificationPending"),
			),
		).
		WithIPApplicationsRef().
		WithCirclePrefRef().
		All(ctx)
	if err != nil {
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}

	//log.Println("CA pending records returned: ", record)
	return record, nil
}

// Get latest old Application Remarks given to Candidate for CA Verification
func GetOldCAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (*ent.Exam_Applications_IP, error) {
	employeeExists, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.EmployeeIDEQ(employeeID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check employee existence: %v", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	application, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ApplicationStatusEQ("PendingWithCandidate"),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("application not found for employee ID: %d with 'PendingWithCandidate' status", employeeID)
		}
		return nil, fmt.Errorf("failed to retrieve application: %v", err)
	}

	return application, nil
}

// Get Recommendations/ Remarks with Emp ID
func GetRecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsIPApplications, error) {
	ctx := context.Background()
	// Check if empID is null
	if empID == 0 {
		return nil, fmt.Errorf("no employee ID provided to process")
	}
	// Check if empID exists
	exists, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if employee with ID %d exists: %v", empID, err)
	}
	if !exists {
		return nil, fmt.Errorf("employee with ID %d does not exist", empID)
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID %d: %v", empID, err)
	}

	return records, nil
}

// Get All NA Verified Records
func QueryIPApplicationsByNAVerified(ctx context.Context, client *ent.Client) ([]*ent.Exam_Applications_IP, error) {
	// Array of exams
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.ApplicationStatus("VerifiedByNA"), // Add the Where clause
		).
		//WithIPApplicationsRef().
		All(ctx)
	if err != nil {
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}
	//log.Println("CA verified records returned: ", records)
	return records, nil
}

// Get All NA Verified Records with Emp ID
func QueryIPApplicationsByNAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64) (*ent.Exam_Applications_IP, error) {
	employeeExists, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.EmployeeIDEQ(employeeID)).
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ApplicationStatusEQ("VerifiedByNA"), // Check for "CAVerified" status
			exam_applications_ip.EmployeeIDEQ(employeeID),
		).
		WithIPApplicationsRef().
		WithCirclePrefRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("record not found for employee ID: %d with 'CAVerified' status", employeeID)
		}
		log.Println("error at IP Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying IP exams Applications: %w", err)
	}

	//log.Println("CA verified record returned: ", record)
	return record, nil
}
