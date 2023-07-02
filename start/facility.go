package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/circlemaster"
	"recruit/ent/divisionmaster"
	"recruit/ent/facility"
	"recruit/ent/regionmaster"
)

func CreateFacility(client *ent.Client, newfacility *ent.Facility) (*ent.Facility, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx := context.Background()
	u, err := client.Facility.
		Create().
		SetFacilityCode(newfacility.FacilityCode).
		SetFacilityName(newfacility.FacilityName).
		SetOfficeType(newfacility.OfficeType).
		SetReportingOfficeType(newfacility.ReportingOfficeCode).
		SetReportingOfficeType(newfacility.ReportingOfficeType).
		SetFacilityOfficeID(newfacility.FacilityOfficeID).
		SetMobileNumber(newfacility.MobileNumber).
		SetEmailID(newfacility.EmailID).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Employee designation: ", newfacility)
		return nil, fmt.Errorf("failed creating Facility: %w", err)
	}
	log.Println("Facility was created: ", u)

	return u, nil
}

func QueryFacilitiesMaster(ctx context.Context, client *ent.Client) ([]*ent.Facility, error) {
	//Array of exams
	Facilities, err := client.Facility.Query().
		All(ctx)
	if err != nil {
		log.Println("error at Facility Master: ", err)
		return nil, fmt.Errorf("failed querying Facility Master: %w", err)
	}
	log.Println("Facility returned: ", Facilities)
	return Facilities, nil
}

func QueryFacilityMasterByDivisionCode(ctx context.Context, client *ent.Client, Dcode int32) ([]*ent.Facility, error) {
	//Can use GetX as well

	Facility_Master, err := client.Facility.Query().
		Where(facility.DivisionCodeEQ(Dcode)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting Facility master: ", err)
		return nil, fmt.Errorf("failed Facility master: %w", err)
	}
	log.Println("Facility returned by Division code : ", Facility_Master)
	return Facility_Master, nil
}

func QueryFacilityMasterByCircleCode(ctx context.Context, client *ent.Client, Ccode int32) ([]*ent.Facility, error) {
	//Can use GetX as well

	Facility_Master, err := client.Facility.Query().
		Where(facility.CircleCodeEQ(Ccode)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting Facility master: ", err)
		return nil, fmt.Errorf("failed Facility master: %w", err)
	}
	log.Println("Facility returned by Circle code : ", Facility_Master)
	return Facility_Master, nil
}

func QueryFacilityMasterByRegionCode(ctx context.Context, client *ent.Client, Rcode int32) ([]*ent.Facility, error) {
	//Can use GetX as well

	Facility_Master, err := client.Facility.Query().
		Where(facility.RegionCodeEQ(Rcode)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting Facility master: ", err)
		return nil, fmt.Errorf("failed Facility master: %w", err)
	}
	log.Println("Facility returned by Region code : ", Facility_Master)
	return Facility_Master, nil
}

func QueryFacilityMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.Facility, error) {
	//Can use GetX as well

	Facility_Master, err := client.Facility.Get(ctx, id)
	if err != nil {

		log.Println("error at getting Facility ID: ", err)
		return nil, fmt.Errorf("failed querying Facility Master: %w", err)
	}
	log.Println("Facility Master details returned: ", Facility_Master)
	return Facility_Master, nil
}

// Query Facility by Facility Office ID

func QueryFacilityByOfficeID(ctx context.Context, client *ent.Client, officeID string) (*ent.Facility, error) {
	if officeID == "" {
		return nil, fmt.Errorf("officeID cannot be empty")
	}

	facility, err := client.Facility.
		Query().
		Where(facility.FacilityOfficeID(officeID)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying Facility by OfficeID:", err)
		return nil, fmt.Errorf("failed querying Facility: %w", err)
	}

	// Check if the facility exists
	if facility == nil {
		log.Println("Facility not found with OfficeID:", officeID)
		return nil, fmt.Errorf("facility not found with OfficeID: %s", officeID)
	}

	//log.Println("Facility retrieved:", facility)

	// Check if the facility has a division ID
	if facility.DivisionID != 0 {
		// Retrieve the division by its ID
		division, err := client.DivisionMaster.
			Query().
			Where(divisionmaster.ID(facility.DivisionID)).
			Only(ctx)

		if err != nil {
			log.Println("Error querying DivisionMaster by ID:", err)
			return nil, fmt.Errorf("failed querying DivisionMaster: %w", err)
		}

		// Update the DivisionName field in the facility
		facility, err = facility.Update().
			SetDivisionName(division.DivisionOfficeName).
			SetReportingOfficeID(division.DivisionOfficeID).
			SetReportingOfficeName(division.DivisionOfficeName).
			Save(ctx)

		if err != nil {
			log.Println("Error updating Facility with DivisionName:", err)
			return nil, fmt.Errorf("failed updating Facility: %w", err)
		}

		//log.Println("Facility updated with DivisionName:", facility.DivisionName)
	}

	// Check if the facility has a region ID
	if facility.RegionID != 0 {
		// Retrieve the region by its ID
		region, err := client.RegionMaster.
			Query().
			Where(regionmaster.ID(facility.RegionID)).
			Only(ctx)

		if err != nil {
			log.Println("Error querying RegionMaster by ID:", err)
			return nil, fmt.Errorf("failed querying RegionMaster: %w", err)
		}

		// Update the RegionName field in the facility
		facility, err = facility.Update().
			SetRegionName(region.RegionOfficeName).
			Save(ctx)

		if err != nil {
			log.Println("Error updating Facility with RegionName:", err)
			return nil, fmt.Errorf("failed updating Facility: %w", err)
		}

		//log.Println("Facility updated with RegionName:", facility.RegionName)
	}

	// Check if the facility has a circle ID
	if facility.CircleID != 0 {
		// Retrieve the circle by its ID
		circle, err := client.CircleMaster.
			Query().
			Where(circlemaster.ID(facility.CircleID)).
			Only(ctx)

		if err != nil {
			log.Println("Error querying CircleMaster by ID:", err)
			return nil, fmt.Errorf("failed querying CircleMaster: %w", err)
		}

		// Update the CircleName field in the facility
		facility, err = facility.Update().
			SetCircleName(circle.CircleOfficeName).
			Save(ctx)

		if err != nil {
			log.Println("Error updating Facility with CircleName:", err)
			return nil, fmt.Errorf("failed updating Facility: %w", err)
		}

		//log.Println("Facility updated with CircleName:", facility.CircleName)
	}
	return facility, nil
}
