package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/divisionmaster"
)

func CreateDivisionMaster(client *ent.Client, newDivisionMaster *ent.DivisionMaster) (*ent.DivisionMaster, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx := context.Background()
	u, err := client.DivisionMaster.
		Create().
		SetRegionCode(newDivisionMaster.RegionCode).
		SetDivisionCode(newDivisionMaster.DivisionCode).
		SetDivisionOfficeID(newDivisionMaster.DivisionOfficeID).
		SetDivisionOfficeName(newDivisionMaster.DivisionOfficeName).
		SetOfficeType(newDivisionMaster.OfficeType).
		SetReportingOfficeType(newDivisionMaster.ReportingOfficeType).
		SetReportingOfficeCode(newDivisionMaster.ReportingOfficeCode).
		SetEmailID(newDivisionMaster.EmailID).
		SetMobileNumber(newDivisionMaster.MobileNumber).

		//.AddCircleRefIDs().
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Division Master : ", newDivisionMaster)
		return nil, fmt.Errorf("failed creating Division Master: %w", err)
	}
	log.Println("Division was created: ", u)

	return u, nil
}

func QueryDivisionMaster(ctx context.Context, client *ent.Client) ([]*ent.DivisionMaster, error) {
	//Array of exams
	newDivisionMaster, err := client.DivisionMaster.Query().
		All(ctx)
	if err != nil {
		log.Println("error at Division Master: ", err)
		return nil, fmt.Errorf("failed querying Division Master: %w", err)
	}
	log.Println("Division returned: ", newDivisionMaster)
	return newDivisionMaster, nil
}

func QueryDivisionMasterByRegionCode(ctx context.Context, client *ent.Client, Rcode int32) ([]*ent.DivisionMaster, error) {
	//Can use GetX as well

	Division_Master, err := client.DivisionMaster.Query().
		Where(divisionmaster.RegionCodeEQ(Rcode)).All(ctx)

	if err != nil {
		log.Println("error at gettting Division master: ", err)
		return nil, fmt.Errorf("failed Division master: %w", err)
	}
	log.Println("Division returned by RegionCode : ", Division_Master)
	return Division_Master, nil
}
func QueryDivisionMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.DivisionMaster, error) {
	//Can use GetX as well

	Division_Master, err := client.DivisionMaster.Get(ctx, id)
	if err != nil {

		log.Println("error at getting Division Master ID: ", err)
		return nil, fmt.Errorf("failed querying Division Master: %w", err)
	}
	log.Println("Division Master details returned: ", Division_Master)
	return Division_Master, nil
}
