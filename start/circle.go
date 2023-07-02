package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
)

func CreateCircleMaster(client *ent.Client, newCircle *ent.CircleMaster) (*ent.CircleMaster, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx := context.Background()
	u, err := client.CircleMaster.
		Create().
		SetCircleCode(newCircle.CircleCode).
		SetCircleOfficeId(newCircle.CircleOfficeId).
		SetCircleOfficeName(newCircle.CircleOfficeName).
		SetOfficeType(newCircle.OfficeType).
		SetEmailID(newCircle.EmailID).
		SetMobileNumber(newCircle.MobileNumber).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Circle Master: ", newCircle)
		return nil, fmt.Errorf("failed creating Circle Master: %w", err)
	}
	log.Println("Circle was created: ", u)

	return u, nil
}

func QueryCircleMaster(ctx context.Context, client *ent.Client) ([]*ent.CircleMaster, error) {
	//Array of exams
	circles, err := client.CircleMaster.Query().
		All(ctx)
	if err != nil {
		log.Println("error at Circle Master: ", err)
		return nil, fmt.Errorf("failed querying Circle Master: %w", err)
	}
	log.Println("Circles returned: ", circles)
	return circles, nil
}

func QueryCircleMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.CircleMaster, error) {
	//Can use GetX as well

	Circle_Master, err := client.CircleMaster.Get(ctx, id)
	if err != nil {

		log.Println("error at getting CircleMaster ID: ", err)
		return nil, fmt.Errorf("failed querying CircleMaster: %w", err)
	}
	log.Println("CircleMaster details returned: ", Circle_Master)
	return Circle_Master, nil
}
