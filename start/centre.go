package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
)

func CreateCenter(client *ent.Client, newCenter *ent.Center) (*ent.Center, error) {
	ctx := context.Background()
	u, err := client.Center.Create().
		SetNotifyCode(newCenter.NotifyCode).
		SetNodalOfficerCode(newCenter.NodalOfficerCode).
		SetCenterName(newCenter.CenterName).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating newCenter: ", newCenter)
		return nil, fmt.Errorf("failed creating newCenter: %w", err)
	}
	log.Println("newCenter was created: ", u)

	return u, nil
}

/*func QueryExamByName(ctx context.Context, client *ent.Client, examname string) (*ent.Exam, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}*/

func QueryCenterID(ctx context.Context, client *ent.Client, id int32) (*ent.Center, error) {
	//Can use GetX as well

	u, err := client.Center.Get(ctx, id)
	if err != nil {
		log.Println("error at getting CenterID: ", err)
		return nil, fmt.Errorf("failed querying Center: %w", err)
	}
	log.Println("Center returned: ", u)
	return u, nil
}

func QueryCenter(ctx context.Context, client *ent.Client) ([]*ent.Center, error) {
	//Array of exams
	u, err := client.Center.Query().All(ctx)
	if err != nil {
		log.Println("error at Center: ", err)
		return nil, fmt.Errorf("failed querying Center: %w", err)
	}
	log.Println("Center returned: ", u)
	return u, nil
}

func DeleteCenterID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.Center.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateCenter(client *ent.Client, id int32, newCenter *ent.Center) (*ent.Center, error) {

	ctx := context.Background()
	_, err := QueryCenterID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.Center.UpdateOneID(id).
		SetNotifyCode(newCenter.NotifyCode).
		SetNodalOfficerCode(newCenter.NodalOfficerCode).
		SetCenterName(newCenter.CenterName).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}
