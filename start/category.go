package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
)

func CreateEmployeeCategory(client *ent.Client, newcategory *ent.EmployeeCategory) (*ent.EmployeeCategory, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx := context.Background()
	u, err := client.EmployeeCategory.
		Create().
		SetCategrycode(newcategory.Categrycode).
		SetCategoryDescription(newcategory.CategoryDescription).
		SetMinimumMarks(newcategory.MinimumMarks).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Employee Category: ", newcategory)
		return nil, fmt.Errorf("failed creating employee category: %w", err)
	}
	log.Println("Employee category was created: ", u)

	return u, nil
}

func QueryEmployeeCategoryID(ctx context.Context, client *ent.Client, id int32) (*ent.EmployeeCategory, error) {
	//Can use GetX as well
	empcat, err := client.EmployeeCategory.Get(ctx, id)
	if err != nil {
		log.Println("error at getting examid: ", err)
		return nil, fmt.Errorf("failed querying employee category: %w", err)
	}
	log.Println("EmployeeCategory returned: ", empcat)
	return empcat, nil
}

func QueryEmployeeCategories(ctx context.Context, client *ent.Client) ([]*ent.EmployeeCategory, error) {
	//Array of exams
	employeecategory, err := client.EmployeeCategory.Query().
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, fmt.Errorf("failed querying Employee Category: %w", err)
	}
	log.Println("EmployeeCategory returned: ", employeecategory)
	return employeecategory, nil
}
