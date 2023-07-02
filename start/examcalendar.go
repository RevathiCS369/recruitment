package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
)

func CreateExamCalendar(client *ent.Client, newExamCalendar *ent.ExamCalendar) (*ent.ExamCalendar, error) {
	ctx := context.Background()
	u, err := client.ExamCalendar.Create().
		SetExamYear(newExamCalendar.ExamYear).
		SetExamName(newExamCalendar.ExamName).
		SetNotificationDate(newExamCalendar.NotificationDate).
		SetModelNotificationDate(newExamCalendar.ModelNotificationDate).
		SetApplicationEndDate(newExamCalendar.ApplicationEndDate).
		SetTentativeResultDate(newExamCalendar.TentativeResultDate).
		SetApprovedOrderDate(newExamCalendar.ApprovedOrderDate).
		SetApprovedOrderNumber(newExamCalendar.ApprovedOrderNumber).
		SetCreatedDate(newExamCalendar.CreatedDate).
		SetVacancyYears(newExamCalendar.VacancyYears).
		SetExamPapers(newExamCalendar.ExamPapers).
		SetExamCode(newExamCalendar.ExamCode).
		//SetPaperCode(newExamCalendar.PaperCode).
		//SetVacancyYearCode(newExamCalendar.VacancyYearCode).
		Save(ctx)

	if err != nil {
		log.Println("error at Creating Exam Calendar: ", newExamCalendar)
		return nil, fmt.Errorf("failed creating Exam Calendar: %w", err)
	}
	log.Println("Exam Calendar was created: ", u)

	return u, nil
}

func QueryExamCalendarID(ctx context.Context, client *ent.Client, id int32) (*ent.ExamCalendar, error) {
	//Can use GetX as well

	ExamCalendar, err := client.ExamCalendar.Get(ctx, id)
	if err != nil {
		log.Println("error at getting ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying ExamCalendar: %w", err)
	}
	log.Println("ExamCalendar returned: ", ExamCalendar)
	return ExamCalendar, nil
}

func QueryExamCalendars(ctx context.Context, client *ent.Client) ([]*ent.ExamCalendar, error) {
	//Array of exams
	examcalendar, err := client.ExamCalendar.Query().
		All(ctx)
	if err != nil {
		log.Println("error at ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying ExamCalendar: %w", err)
	}
	log.Println("Exam Calendar data returned: ", examcalendar)
	return examcalendar, nil
}

/*func QueryExamCalendarIDWithDetails(ctx context.Context, client *ent.Client, id int32) (*ent.ExamCalendar, error) {
	//Can use GetX as well

	ExamCalendar, err := client.ExamCalendar.Get(ctx, id)
	if err != nil {
		log.Println("error at getting ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying ExamCalendar: %w", err)
	}
	log.Println("ExamCalendar returned: ", ExamCalendar)
	return ExamCalendar, nil
}*/

/*func DeleteExamcalendarID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.ExamCalendar.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateExamCalendarID(client *ent.Client, id int32, newExamcalendar *ent.ExamCalendar) (*ent.ExamCalendar, error) {

	ctx := context.Background()
	_, err := QueryExamCalendarID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedExamCalendardata, err := client.ExamCalendar.UpdateOneID(id).
		SetExamCode(newExamcalendar.ExamCode).
		SetExamYear(newExamcalendar.ExamYear).
		SetCalendarNotificationDate(newExamcalendar.CalendarNotificationDate).
		SetModelNotificationDate(newExamcalendar.ModelNotificationDate).
		SetApplicationEndDate(newExamcalendar.ApplicationEndDate).
		SetDateOfExam(newExamcalendar.DateOfExam).
		SetTentativeResultDate(newExamcalendar.TentativeResultDate).
		SetApprovedOrderDate(newExamcalendar.ApprovedOrderDate).
		SetOrderNumber(newExamcalendar.OrderNumber).
		SetCreatedDate(newExamcalendar.CreatedDate).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedExamCalendardata, nil
}*/

/*func QueryExamCalendarsWithVacancyAndPapers(ctx context.Context, client *ent.Client, examcode int32) ([]*ent.ExamCalendar, error) {
	//Array of exams
	examcalendar, err := client.ExamCalendar.Query().
		WithExams(func(epq *ent.ExamQuery) {
			epq.Select(exam.FieldExamName, exam.FieldCalendarCode, exam.FieldNumOfPapers)
		}).
		WithVacancyYears(func(epq *ent.VacancyYearQuery) {
			epq.Select(vacancyyear.FieldVacancyYear, vacancyyear.FieldFromDate, vacancyyear.FieldToDate).Aggregate()
		}).
		WithPapers(func(epq *ent.ExamPapersQuery) {
			epq.Select(exampapers.FieldID, exampapers.FieldPaperDescription)
		}).
		Where(examcalendar.ExamCodeEQ(examcode)).Aggregate().
		All(ctx)

	if err != nil {
		log.Println("error at ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying ExamCalendar: %w", err)
	}
	log.Println("Exam Calendar data returned: ", examcalendar)
	return examcalendar, nil
}*/
