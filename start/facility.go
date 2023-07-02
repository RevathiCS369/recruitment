package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/facility"
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
