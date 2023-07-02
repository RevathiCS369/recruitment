// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgeEligibilityColumns holds the columns for the "AgeEligibility" table.
	AgeEligibilityColumns = []*schema.Column{
		{Name: "AgeElibilityCode", Type: field.TypeInt32, Increment: true},
		{Name: "age", Type: field.TypeInt32, Nullable: true},
		{Name: "category_id", Type: field.TypeInt32, Nullable: true},
		{Name: "eligibility_code", Type: field.TypeInt32, Nullable: true},
	}
	// AgeEligibilityTable holds the schema information for the "AgeEligibility" table.
	AgeEligibilityTable = &schema.Table{
		Name:       "AgeEligibility",
		Columns:    AgeEligibilityColumns,
		PrimaryKey: []*schema.Column{AgeEligibilityColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "AgeEligibility_ExamEligibility_age_eligibilities",
				Columns:    []*schema.Column{AgeEligibilityColumns[3]},
				RefColumns: []*schema.Column{ExamEligibilityColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ApplicationColumns holds the columns for the "Application" table.
	ApplicationColumns = []*schema.Column{
		{Name: "ApplicationCode", Type: field.TypeInt32, Increment: true},
		{Name: "employee_id", Type: field.TypeInt32},
		{Name: "hall_ticket_number", Type: field.TypeString, Nullable: true},
		{Name: "applied_stamp", Type: field.TypeTime},
		{Name: "center_code", Type: field.TypeInt32, Nullable: true},
		{Name: "notify_code", Type: field.TypeInt32, Nullable: true},
	}
	// ApplicationTable holds the schema information for the "Application" table.
	ApplicationTable = &schema.Table{
		Name:       "Application",
		Columns:    ApplicationColumns,
		PrimaryKey: []*schema.Column{ApplicationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Application_Center_applications",
				Columns:    []*schema.Column{ApplicationColumns[4]},
				RefColumns: []*schema.Column{CenterColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Application_Notification_applications",
				Columns:    []*schema.Column{ApplicationColumns[5]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CenterColumns holds the columns for the "Center" table.
	CenterColumns = []*schema.Column{
		{Name: "CenterCode", Type: field.TypeInt32, Increment: true},
		{Name: "center_name", Type: field.TypeString},
		{Name: "exam_papers_centers", Type: field.TypeInt32, Nullable: true},
		{Name: "nodal_officer_code", Type: field.TypeInt32, Nullable: true},
		{Name: "notify_code", Type: field.TypeInt32, Nullable: true},
	}
	// CenterTable holds the schema information for the "Center" table.
	CenterTable = &schema.Table{
		Name:       "Center",
		Columns:    CenterColumns,
		PrimaryKey: []*schema.Column{CenterColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Center_ExamPapers_centers",
				Columns:    []*schema.Column{CenterColumns[2]},
				RefColumns: []*schema.Column{ExamPapersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Center_NodalOfficers_centers",
				Columns:    []*schema.Column{CenterColumns[3]},
				RefColumns: []*schema.Column{NodalOfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Center_Notification_centers",
				Columns:    []*schema.Column{CenterColumns[4]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CircleMasterColumns holds the columns for the "CircleMaster" table.
	CircleMasterColumns = []*schema.Column{
		{Name: "CircleID", Type: field.TypeInt32, Increment: true},
		{Name: "circle_code", Type: field.TypeInt32},
		{Name: "circle_office_id", Type: field.TypeString},
		{Name: "circle_office_name", Type: field.TypeString},
		{Name: "office_type", Type: field.TypeString},
		{Name: "email_id", Type: field.TypeString, Nullable: true},
		{Name: "mobile_number", Type: field.TypeInt32, Nullable: true},
		{Name: "facility_circle_ref", Type: field.TypeInt32, Nullable: true},
		{Name: "region_master_circle_ref", Type: field.TypeInt32, Nullable: true},
	}
	// CircleMasterTable holds the schema information for the "CircleMaster" table.
	CircleMasterTable = &schema.Table{
		Name:       "CircleMaster",
		Columns:    CircleMasterColumns,
		PrimaryKey: []*schema.Column{CircleMasterColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "CircleMaster_Facility_circle_ref",
				Columns:    []*schema.Column{CircleMasterColumns[7]},
				RefColumns: []*schema.Column{FacilityColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "CircleMaster_RegionMaster_circle_ref",
				Columns:    []*schema.Column{CircleMasterColumns[8]},
				RefColumns: []*schema.Column{RegionMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DisabilityColumns holds the columns for the "Disability" table.
	DisabilityColumns = []*schema.Column{
		{Name: "DisabilityTypeID", Type: field.TypeInt32, Increment: true},
		{Name: "disability_type_code", Type: field.TypeString},
		{Name: "disability_type_description", Type: field.TypeString},
		{Name: "disability_percentage", Type: field.TypeInt32},
		{Name: "disability_flag", Type: field.TypeEnum, Enums: []string{"Temporary", "Permanent"}},
	}
	// DisabilityTable holds the schema information for the "Disability" table.
	DisabilityTable = &schema.Table{
		Name:       "Disability",
		Columns:    DisabilityColumns,
		PrimaryKey: []*schema.Column{DisabilityColumns[0]},
	}
	// DivisionMasterColumns holds the columns for the "DivisionMaster" table.
	DivisionMasterColumns = []*schema.Column{
		{Name: "DivisionID", Type: field.TypeInt32, Increment: true},
		{Name: "division_code", Type: field.TypeInt32},
		{Name: "office_type", Type: field.TypeString},
		{Name: "division_office_id", Type: field.TypeString},
		{Name: "division_office_name", Type: field.TypeString},
		{Name: "reporting_office_type", Type: field.TypeString, Nullable: true},
		{Name: "reporting_office_code", Type: field.TypeString, Nullable: true},
		{Name: "email_id", Type: field.TypeString, Nullable: true},
		{Name: "mobile_number", Type: field.TypeInt32, Nullable: true},
		{Name: "region_code", Type: field.TypeInt32, Nullable: true},
		{Name: "region_master_regions", Type: field.TypeInt32, Nullable: true},
	}
	// DivisionMasterTable holds the schema information for the "DivisionMaster" table.
	DivisionMasterTable = &schema.Table{
		Name:       "DivisionMaster",
		Columns:    DivisionMasterColumns,
		PrimaryKey: []*schema.Column{DivisionMasterColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "DivisionMaster_RegionMaster_regions",
				Columns:    []*schema.Column{DivisionMasterColumns[10]},
				RefColumns: []*schema.Column{RegionMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmployeeCadreColumns holds the columns for the "EmployeeCadre" table.
	EmployeeCadreColumns = []*schema.Column{
		{Name: "cadreid", Type: field.TypeInt32, Increment: true},
		{Name: "cadrecode", Type: field.TypeString},
		{Name: "cadredescription", Type: field.TypeString},
		{Name: "pay_level", Type: field.TypeString},
		{Name: "scale", Type: field.TypeString},
	}
	// EmployeeCadreTable holds the schema information for the "EmployeeCadre" table.
	EmployeeCadreTable = &schema.Table{
		Name:       "EmployeeCadre",
		Columns:    EmployeeCadreColumns,
		PrimaryKey: []*schema.Column{EmployeeCadreColumns[0]},
	}
	// EmployeeCategoryColumns holds the columns for the "EmployeeCategory" table.
	EmployeeCategoryColumns = []*schema.Column{
		{Name: "categoryidid", Type: field.TypeInt32, Increment: true},
		{Name: "categrycode", Type: field.TypeString},
		{Name: "category_description", Type: field.TypeString},
		{Name: "minimum_marks", Type: field.TypeInt32, Nullable: true},
	}
	// EmployeeCategoryTable holds the schema information for the "EmployeeCategory" table.
	EmployeeCategoryTable = &schema.Table{
		Name:       "EmployeeCategory",
		Columns:    EmployeeCategoryColumns,
		PrimaryKey: []*schema.Column{EmployeeCategoryColumns[0]},
	}
	// EmployeeDesignationColumns holds the columns for the "EmployeeDesignation" table.
	EmployeeDesignationColumns = []*schema.Column{
		{Name: "DesignationID", Type: field.TypeInt32, Increment: true},
		{Name: "designation_code", Type: field.TypeString},
		{Name: "designation_description", Type: field.TypeString},
	}
	// EmployeeDesignationTable holds the schema information for the "EmployeeDesignation" table.
	EmployeeDesignationTable = &schema.Table{
		Name:       "EmployeeDesignation",
		Columns:    EmployeeDesignationColumns,
		PrimaryKey: []*schema.Column{EmployeeDesignationColumns[0]},
	}
	// EmployeePostsColumns holds the columns for the "EmployeePosts" table.
	EmployeePostsColumns = []*schema.Column{
		{Name: "PostID", Type: field.TypeInt32, Increment: true},
		{Name: "post_code", Type: field.TypeString},
		{Name: "post_description", Type: field.TypeString},
		{Name: "group", Type: field.TypeString},
		{Name: "pay_level", Type: field.TypeString},
		{Name: "scale", Type: field.TypeString},
		{Name: "base_cadre_flag", Type: field.TypeBool},
	}
	// EmployeePostsTable holds the schema information for the "EmployeePosts" table.
	EmployeePostsTable = &schema.Table{
		Name:       "EmployeePosts",
		Columns:    EmployeePostsColumns,
		PrimaryKey: []*schema.Column{EmployeePostsColumns[0]},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "RegistrationID", Type: field.TypeInt32, Increment: true},
		{Name: "employeed_id", Type: field.TypeInt32},
		{Name: "id_verified", Type: field.TypeBool, Default: false},
		{Name: "id_rem_status", Type: field.TypeBool, Default: false},
		{Name: "id_remarks", Type: field.TypeString, Nullable: true},
		{Name: "employee_name", Type: field.TypeString},
		{Name: "name_verified", Type: field.TypeBool, Default: false},
		{Name: "name_rem_status", Type: field.TypeBool, Default: false},
		{Name: "name_remarks", Type: field.TypeString, Nullable: true},
		{Name: "employee_fathers_name", Type: field.TypeString},
		{Name: "fathers_name_verified", Type: field.TypeBool, Default: false},
		{Name: "fathers_name_rem_status", Type: field.TypeBool, Default: false},
		{Name: "fathers_name_remarks", Type: field.TypeString, Nullable: true},
		{Name: "dob", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "dob_verified", Type: field.TypeBool, Default: false},
		{Name: "dob_rem_status", Type: field.TypeBool, Default: false},
		{Name: "dob_remarks", Type: field.TypeString, Nullable: true},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"Male", "Female"}},
		{Name: "gender_verified", Type: field.TypeBool, Default: false},
		{Name: "gender_rem_status", Type: field.TypeBool, Default: false},
		{Name: "gender_remarks", Type: field.TypeString, Nullable: true},
		{Name: "mobile_number", Type: field.TypeInt32, Nullable: true},
		{Name: "email_id", Type: field.TypeString, Nullable: true},
		{Name: "categoryid", Type: field.TypeInt32, Nullable: true},
		{Name: "employee_category_code", Type: field.TypeString, Nullable: true},
		{Name: "employee_category", Type: field.TypeString},
		{Name: "employee_category_code_verified", Type: field.TypeBool, Default: false},
		{Name: "employee_category_code_rem_status", Type: field.TypeBool, Default: false},
		{Name: "employee_category_code_remarks", Type: field.TypeString, Nullable: true},
		{Name: "with_disability", Type: field.TypeString},
		{Name: "with_disability_verified", Type: field.TypeBool, Default: false},
		{Name: "with_disability_rem_status", Type: field.TypeBool, Default: false},
		{Name: "with_disability_remarks", Type: field.TypeBool, Nullable: true},
		{Name: "disability_type", Type: field.TypeString, Nullable: true},
		{Name: "disability_type_verified", Type: field.TypeBool, Default: false},
		{Name: "disability_type_rem_status", Type: field.TypeBool, Default: false},
		{Name: "disability_type_remarks", Type: field.TypeString, Nullable: true},
		{Name: "disability_percentage", Type: field.TypeInt32, Nullable: true},
		{Name: "disability_percentage_verified", Type: field.TypeBool, Default: false},
		{Name: "disability_percentage_rem_status", Type: field.TypeBool, Default: false},
		{Name: "disability_percentage_remarks", Type: field.TypeString, Nullable: true},
		{Name: "signature", Type: field.TypeString},
		{Name: "signature_verified", Type: field.TypeBool, Default: false},
		{Name: "signature_rem_status", Type: field.TypeBool, Default: false},
		{Name: "signature_remarks", Type: field.TypeString, Nullable: true},
		{Name: "photo", Type: field.TypeString},
		{Name: "photo_verified", Type: field.TypeBool, Default: false},
		{Name: "photo_rem_status", Type: field.TypeBool, Default: false},
		{Name: "photo_remarks", Type: field.TypeString, Nullable: true},
		{Name: "cadreid", Type: field.TypeInt32, Nullable: true},
		{Name: "employee_cadre", Type: field.TypeString},
		{Name: "employee_cadre_verified", Type: field.TypeBool, Default: false},
		{Name: "employee_cadre_rem_status", Type: field.TypeBool, Default: false},
		{Name: "employee_cadre_remarks", Type: field.TypeString, Nullable: true},
		{Name: "designation_id", Type: field.TypeInt32, Nullable: true},
		{Name: "employee_designation", Type: field.TypeString},
		{Name: "employee_designation_verified", Type: field.TypeBool, Default: false},
		{Name: "employee_designation_rem_status", Type: field.TypeBool, Default: false},
		{Name: "employee_designation_remarks", Type: field.TypeString, Nullable: true},
		{Name: "circle_id", Type: field.TypeInt32, Nullable: true},
		{Name: "circle_name", Type: field.TypeString},
		{Name: "circle_verified", Type: field.TypeBool, Default: false},
		{Name: "circle_rem_status", Type: field.TypeBool, Default: false},
		{Name: "circle_remarks", Type: field.TypeString, Nullable: true},
		{Name: "region_id", Type: field.TypeInt32, Nullable: true},
		{Name: "region_name", Type: field.TypeString},
		{Name: "region_verified", Type: field.TypeBool, Default: false},
		{Name: "region_rem_status", Type: field.TypeBool, Default: false},
		{Name: "region_remarks", Type: field.TypeString, Nullable: true},
		{Name: "division_id", Type: field.TypeInt32, Nullable: true},
		{Name: "division_name", Type: field.TypeString},
		{Name: "division_verified", Type: field.TypeBool, Default: false},
		{Name: "division_rem_status", Type: field.TypeBool, Default: false},
		{Name: "division_remarks", Type: field.TypeString},
		{Name: "office_id", Type: field.TypeInt32, Nullable: true},
		{Name: "office_name", Type: field.TypeString},
		{Name: "office_verified", Type: field.TypeBool, Default: false},
		{Name: "office_rem_status", Type: field.TypeBool, Default: false},
		{Name: "office_remarks", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeString},
		{Name: "role_verified", Type: field.TypeBool, Default: false},
		{Name: "role_rem_status", Type: field.TypeBool, Default: false},
		{Name: "role_remarks", Type: field.TypeString},
		{Name: "dccs", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "dccs_verified", Type: field.TypeBool, Default: false},
		{Name: "dccs_rem_status", Type: field.TypeBool, Default: false},
		{Name: "dccs_remarks", Type: field.TypeString, Nullable: true},
		{Name: "dc_in_present_cadre", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "dc_in_present_cadre_verified", Type: field.TypeBool, Default: false},
		{Name: "dc_in_present_cadre_rem_status", Type: field.TypeBool, Default: false},
		{Name: "dc_in_present_cadre_remarks", Type: field.TypeString, Nullable: true},
		{Name: "aps_working", Type: field.TypeBool, Nullable: true},
		{Name: "aps_working_verified", Type: field.TypeBool, Default: false},
		{Name: "aps_working_rem_status", Type: field.TypeBool, Default: false},
		{Name: "aps_working_remarks", Type: field.TypeString, Nullable: true},
		{Name: "profilestatus", Type: field.TypeBool, Default: false},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
	}
	// ExamColumns holds the columns for the "Exam" table.
	ExamColumns = []*schema.Column{
		{Name: "ExamCode", Type: field.TypeInt32, Increment: true},
		{Name: "exam_name", Type: field.TypeString},
		{Name: "num_of_papers", Type: field.TypeInt32},
		{Name: "notification_by", Type: field.TypeString},
		{Name: "conducted_by", Type: field.TypeString},
		{Name: "nodal_officer_level", Type: field.TypeInt32, Nullable: true},
		{Name: "calendar_code", Type: field.TypeInt32, Nullable: true},
		{Name: "paper_code", Type: field.TypeInt32, Nullable: true},
		{Name: "vacancy_year_exams", Type: field.TypeInt32, Nullable: true},
	}
	// ExamTable holds the schema information for the "Exam" table.
	ExamTable = &schema.Table{
		Name:       "Exam",
		Columns:    ExamColumns,
		PrimaryKey: []*schema.Column{ExamColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Exam_VacancyYears_exams",
				Columns:    []*schema.Column{ExamColumns[8]},
				RefColumns: []*schema.Column{VacancyYearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExamCalendarColumns holds the columns for the "ExamCalendar" table.
	ExamCalendarColumns = []*schema.Column{
		{Name: "CalendarCode", Type: field.TypeInt32, Increment: true},
		{Name: "exam_year", Type: field.TypeInt32},
		{Name: "exam_name", Type: field.TypeString},
		{Name: "notification_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "model_notification_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "application_end_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "approved_order_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "tentative_result_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "created_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "approved_order_number", Type: field.TypeString},
		{Name: "vacancy_years", Type: field.TypeJSON, Nullable: true},
		{Name: "exam_papers", Type: field.TypeJSON, Nullable: true},
		{Name: "exam_code", Type: field.TypeInt32, Nullable: true},
		{Name: "paper_code", Type: field.TypeInt32, Nullable: true},
		{Name: "vacancy_year_code", Type: field.TypeInt32, Nullable: true},
	}
	// ExamCalendarTable holds the schema information for the "ExamCalendar" table.
	ExamCalendarTable = &schema.Table{
		Name:       "ExamCalendar",
		Columns:    ExamCalendarColumns,
		PrimaryKey: []*schema.Column{ExamCalendarColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ExamCalendar_Exam_exams_ref",
				Columns:    []*schema.Column{ExamCalendarColumns[12]},
				RefColumns: []*schema.Column{ExamColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ExamCalendar_ExamPapers_papers_ref",
				Columns:    []*schema.Column{ExamCalendarColumns[13]},
				RefColumns: []*schema.Column{ExamPapersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ExamCalendar_VacancyYears_vacancy_ref",
				Columns:    []*schema.Column{ExamCalendarColumns[14]},
				RefColumns: []*schema.Column{VacancyYearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExamEligibilityColumns holds the columns for the "ExamEligibility" table.
	ExamEligibilityColumns = []*schema.Column{
		{Name: "EligibilityCode", Type: field.TypeInt32, Increment: true},
		{Name: "examcode", Type: field.TypeInt32},
		{Name: "age_criteria", Type: field.TypeString, Nullable: true},
		{Name: "service_criteria", Type: field.TypeString, Nullable: true},
		{Name: "driving_license_criteria", Type: field.TypeString, Nullable: true},
		{Name: "notify_code", Type: field.TypeInt32, Nullable: true},
		{Name: "employee_cadre_id", Type: field.TypeInt32, Nullable: true},
		{Name: "category_id", Type: field.TypeInt32, Nullable: true},
	}
	// ExamEligibilityTable holds the schema information for the "ExamEligibility" table.
	ExamEligibilityTable = &schema.Table{
		Name:       "ExamEligibility",
		Columns:    ExamEligibilityColumns,
		PrimaryKey: []*schema.Column{ExamEligibilityColumns[0]},
	}
	// ExamPapersColumns holds the columns for the "ExamPapers" table.
	ExamPapersColumns = []*schema.Column{
		{Name: "PaperCode", Type: field.TypeInt32, Increment: true},
		{Name: "paper_description", Type: field.TypeString, Size: 100},
		{Name: "competitive_qualifying", Type: field.TypeString, Size: 10},
		{Name: "exception_for_disability", Type: field.TypeString, Size: 50},
		{Name: "maximum_marks", Type: field.TypeInt},
		{Name: "duration", Type: field.TypeInt},
		{Name: "local_language_allowed_question_paper", Type: field.TypeString, Size: 10},
		{Name: "local_language_allowed_answer_paper", Type: field.TypeString, Size: 10},
		{Name: "order_number", Type: field.TypeString, Size: 20},
		{Name: "paper_status", Type: field.TypeString, Size: 10},
		{Name: "calendar_code", Type: field.TypeInt32, Nullable: true},
		{Name: "created_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "exam_code", Type: field.TypeInt32, Nullable: true},
	}
	// ExamPapersTable holds the schema information for the "ExamPapers" table.
	ExamPapersTable = &schema.Table{
		Name:       "ExamPapers",
		Columns:    ExamPapersColumns,
		PrimaryKey: []*schema.Column{ExamPapersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ExamPapers_Exam_papers",
				Columns:    []*schema.Column{ExamPapersColumns[12]},
				RefColumns: []*schema.Column{ExamColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FacilityColumns holds the columns for the "Facility" table.
	FacilityColumns = []*schema.Column{
		{Name: "FacilityID", Type: field.TypeInt32, Increment: true},
		{Name: "facility_code", Type: field.TypeString},
		{Name: "office_type", Type: field.TypeString},
		{Name: "facility_name", Type: field.TypeString},
		{Name: "reporting_office_type", Type: field.TypeString, Nullable: true},
		{Name: "reporting_office_code", Type: field.TypeString, Nullable: true},
		{Name: "email_id", Type: field.TypeString, Nullable: true},
		{Name: "mobile_number", Type: field.TypeInt32, Nullable: true},
		{Name: "division_code", Type: field.TypeInt32, Nullable: true},
		{Name: "region_code", Type: field.TypeInt32, Nullable: true},
		{Name: "circle_code", Type: field.TypeInt32, Nullable: true},
		{Name: "circle_master_circle_ref", Type: field.TypeInt32, Nullable: true},
		{Name: "region_master_region_ref_ref", Type: field.TypeInt32, Nullable: true},
	}
	// FacilityTable holds the schema information for the "Facility" table.
	FacilityTable = &schema.Table{
		Name:       "Facility",
		Columns:    FacilityColumns,
		PrimaryKey: []*schema.Column{FacilityColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Facility_CircleMaster_circle_ref",
				Columns:    []*schema.Column{FacilityColumns[11]},
				RefColumns: []*schema.Column{CircleMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Facility_RegionMaster_region_ref_ref",
				Columns:    []*schema.Column{FacilityColumns[12]},
				RefColumns: []*schema.Column{RegionMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NodalOfficersColumns holds the columns for the "NodalOfficers" table.
	NodalOfficersColumns = []*schema.Column{
		{Name: "NodalOfficerCode", Type: field.TypeInt32, Increment: true},
		{Name: "nodal_officer_name", Type: field.TypeString},
		{Name: "designation_id", Type: field.TypeInt32},
		{Name: "email_id", Type: field.TypeString},
		{Name: "mobile_number", Type: field.TypeString},
		{Name: "hall_ticket_approved", Type: field.TypeString, Nullable: true},
		{Name: "exam_code", Type: field.TypeInt32, Nullable: true},
		{Name: "notify_code", Type: field.TypeInt32, Nullable: true},
	}
	// NodalOfficersTable holds the schema information for the "NodalOfficers" table.
	NodalOfficersTable = &schema.Table{
		Name:       "NodalOfficers",
		Columns:    NodalOfficersColumns,
		PrimaryKey: []*schema.Column{NodalOfficersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "NodalOfficers_Exam_nodal_officers",
				Columns:    []*schema.Column{NodalOfficersColumns[6]},
				RefColumns: []*schema.Column{ExamColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "NodalOfficers_Notification_nodal_officers",
				Columns:    []*schema.Column{NodalOfficersColumns[7]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationColumns holds the columns for the "Notification" table.
	NotificationColumns = []*schema.Column{
		{Name: "NotifyCode", Type: field.TypeInt32, Increment: true},
		{Name: "exam_year", Type: field.TypeInt32},
		{Name: "application_start_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "application_end_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "verification_date_by_controller", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "correction_date_by_candidate", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "correction_veriy_date_by_controller", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "hall_ticket_allotment_date_by_nodal_officer", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "hall_ticket_download_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "notify_file", Type: field.TypeString, Nullable: true},
		{Name: "syllabus_file", Type: field.TypeString, Nullable: true},
		{Name: "vacancies_file", Type: field.TypeString, Nullable: true},
		{Name: "exam_code", Type: field.TypeInt32, Nullable: true},
		{Name: "exam_calendar_notify_ref", Type: field.TypeInt32, Nullable: true},
	}
	// NotificationTable holds the schema information for the "Notification" table.
	NotificationTable = &schema.Table{
		Name:       "Notification",
		Columns:    NotificationColumns,
		PrimaryKey: []*schema.Column{NotificationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Notification_Exam_notifications",
				Columns:    []*schema.Column{NotificationColumns[12]},
				RefColumns: []*schema.Column{ExamColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Notification_ExamCalendar_Notify_ref",
				Columns:    []*schema.Column{NotificationColumns[13]},
				RefColumns: []*schema.Column{ExamCalendarColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaperTypesColumns holds the columns for the "paper_types" table.
	PaperTypesColumns = []*schema.Column{
		{Name: "PaperTypeCode", Type: field.TypeInt32, Increment: true},
		{Name: "paper_type_description", Type: field.TypeString, Size: 100},
		{Name: "order_number", Type: field.TypeString, Size: 100},
		{Name: "sequence_number", Type: field.TypeInt32},
		{Name: "created_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "paper_code", Type: field.TypeInt32, Nullable: true},
	}
	// PaperTypesTable holds the schema information for the "paper_types" table.
	PaperTypesTable = &schema.Table{
		Name:       "paper_types",
		Columns:    PaperTypesColumns,
		PrimaryKey: []*schema.Column{PaperTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "paper_types_ExamPapers_exampapers_types",
				Columns:    []*schema.Column{PaperTypesColumns[5]},
				RefColumns: []*schema.Column{ExamPapersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RegionMasterColumns holds the columns for the "RegionMaster" table.
	RegionMasterColumns = []*schema.Column{
		{Name: "RegionID", Type: field.TypeInt32, Increment: true},
		{Name: "region_code", Type: field.TypeInt32},
		{Name: "region_office_id", Type: field.TypeString},
		{Name: "office_type", Type: field.TypeString},
		{Name: "region_office_name", Type: field.TypeString},
		{Name: "reporting_office_type", Type: field.TypeString, Nullable: true},
		{Name: "reporting_office_code", Type: field.TypeString, Nullable: true},
		{Name: "email_id", Type: field.TypeString, Nullable: true},
		{Name: "mobile_number", Type: field.TypeInt32, Nullable: true},
		{Name: "circle_code", Type: field.TypeInt32, Nullable: true},
		{Name: "circle_master_region_ref", Type: field.TypeInt32, Nullable: true},
		{Name: "division_master_regions", Type: field.TypeInt32, Nullable: true},
		{Name: "facility_region_ref", Type: field.TypeInt32, Nullable: true},
	}
	// RegionMasterTable holds the schema information for the "RegionMaster" table.
	RegionMasterTable = &schema.Table{
		Name:       "RegionMaster",
		Columns:    RegionMasterColumns,
		PrimaryKey: []*schema.Column{RegionMasterColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "RegionMaster_CircleMaster_region_ref",
				Columns:    []*schema.Column{RegionMasterColumns[10]},
				RefColumns: []*schema.Column{CircleMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "RegionMaster_DivisionMaster_regions",
				Columns:    []*schema.Column{RegionMasterColumns[11]},
				RefColumns: []*schema.Column{DivisionMasterColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "RegionMaster_Facility_region_ref",
				Columns:    []*schema.Column{RegionMasterColumns[12]},
				RefColumns: []*schema.Column{FacilityColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "employeed_id", Type: field.TypeString, Unique: true},
		{Name: "id_verified", Type: field.TypeBool, Default: false},
		{Name: "id_rem_status", Type: field.TypeBool, Default: false},
		{Name: "id_remarks", Type: field.TypeString},
		{Name: "employeed_name", Type: field.TypeString},
		{Name: "name_verified", Type: field.TypeBool, Default: false},
		{Name: "name_rem_status", Type: field.TypeBool, Default: false},
		{Name: "name_remarks", Type: field.TypeString},
		{Name: "dob", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "dob_verified", Type: field.TypeBool, Default: false},
		{Name: "dob_rem_status", Type: field.TypeBool, Default: false},
		{Name: "dob_remarks", Type: field.TypeString},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"Male", "Female"}},
		{Name: "gender_verified", Type: field.TypeBool, Default: false},
		{Name: "gender_rem_status", Type: field.TypeBool, Default: false},
		{Name: "gender_remarks", Type: field.TypeString},
		{Name: "cadreid", Type: field.TypeInt32},
		{Name: "cadreid_verified", Type: field.TypeBool, Default: false},
		{Name: "cadreid_rem_status", Type: field.TypeBool, Default: false},
		{Name: "cadreid_remarks", Type: field.TypeString},
		{Name: "office_id", Type: field.TypeInt32},
		{Name: "office_id_verified", Type: field.TypeBool, Default: false},
		{Name: "office_id_rem_status", Type: field.TypeBool, Default: false},
		{Name: "office_id_remarks", Type: field.TypeString},
		{Name: "ph", Type: field.TypeBool},
		{Name: "ph_verified", Type: field.TypeBool, Default: false},
		{Name: "ph_rem_status", Type: field.TypeBool, Default: false},
		{Name: "ph_remarks", Type: field.TypeString},
		{Name: "ph_details", Type: field.TypeString, Nullable: true},
		{Name: "ph_details_verified", Type: field.TypeBool, Default: false},
		{Name: "ph_details_rem_status", Type: field.TypeBool, Default: false},
		{Name: "ph_details_remarks", Type: field.TypeString},
		{Name: "aps_working", Type: field.TypeBool},
		{Name: "aps_working_verified", Type: field.TypeBool, Default: false},
		{Name: "aps_working_rem_status", Type: field.TypeBool, Default: false},
		{Name: "aps_working_remarks", Type: field.TypeString},
		{Name: "profilestatus", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VacancyYearsColumns holds the columns for the "VacancyYears" table.
	VacancyYearsColumns = []*schema.Column{
		{Name: "VacancyYearCode", Type: field.TypeInt32, Increment: true},
		{Name: "from_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "to_date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "notify_code", Type: field.TypeInt32, Nullable: true},
		{Name: "vacancy_year", Type: field.TypeString},
		{Name: "calendar_code", Type: field.TypeInt32, Nullable: true},
		{Name: "notification_vacancy_years", Type: field.TypeInt32, Nullable: true},
	}
	// VacancyYearsTable holds the schema information for the "VacancyYears" table.
	VacancyYearsTable = &schema.Table{
		Name:       "VacancyYears",
		Columns:    VacancyYearsColumns,
		PrimaryKey: []*schema.Column{VacancyYearsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "VacancyYears_Notification_vacancy_years",
				Columns:    []*schema.Column{VacancyYearsColumns[6]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationNotifyRefColumns holds the columns for the "notification_notify_ref" table.
	NotificationNotifyRefColumns = []*schema.Column{
		{Name: "notification_id", Type: field.TypeInt32},
		{Name: "notify_ref_id", Type: field.TypeInt32},
	}
	// NotificationNotifyRefTable holds the schema information for the "notification_notify_ref" table.
	NotificationNotifyRefTable = &schema.Table{
		Name:       "notification_notify_ref",
		Columns:    NotificationNotifyRefColumns,
		PrimaryKey: []*schema.Column{NotificationNotifyRefColumns[0], NotificationNotifyRefColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_notify_ref_notification_id",
				Columns:    []*schema.Column{NotificationNotifyRefColumns[0]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notification_notify_ref_notify_ref_id",
				Columns:    []*schema.Column{NotificationNotifyRefColumns[1]},
				RefColumns: []*schema.Column{NotificationColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgeEligibilityTable,
		ApplicationTable,
		CenterTable,
		CircleMasterTable,
		DisabilityTable,
		DivisionMasterTable,
		EmployeeCadreTable,
		EmployeeCategoryTable,
		EmployeeDesignationTable,
		EmployeePostsTable,
		EmployeesTable,
		ExamTable,
		ExamCalendarTable,
		ExamEligibilityTable,
		ExamPapersTable,
		FacilityTable,
		NodalOfficersTable,
		NotificationTable,
		PaperTypesTable,
		RegionMasterTable,
		UsersTable,
		VacancyYearsTable,
		NotificationNotifyRefTable,
	}
)

func init() {
	AgeEligibilityTable.ForeignKeys[0].RefTable = ExamEligibilityTable
	AgeEligibilityTable.Annotation = &entsql.Annotation{
		Table: "AgeEligibility",
	}
	ApplicationTable.ForeignKeys[0].RefTable = CenterTable
	ApplicationTable.ForeignKeys[1].RefTable = NotificationTable
	ApplicationTable.Annotation = &entsql.Annotation{
		Table: "Application",
	}
	CenterTable.ForeignKeys[0].RefTable = ExamPapersTable
	CenterTable.ForeignKeys[1].RefTable = NodalOfficersTable
	CenterTable.ForeignKeys[2].RefTable = NotificationTable
	CenterTable.Annotation = &entsql.Annotation{
		Table: "Center",
	}
	CircleMasterTable.ForeignKeys[0].RefTable = FacilityTable
	CircleMasterTable.ForeignKeys[1].RefTable = RegionMasterTable
	CircleMasterTable.Annotation = &entsql.Annotation{
		Table: "CircleMaster",
	}
	DisabilityTable.Annotation = &entsql.Annotation{
		Table: "Disability",
	}
	DivisionMasterTable.ForeignKeys[0].RefTable = RegionMasterTable
	DivisionMasterTable.Annotation = &entsql.Annotation{
		Table: "DivisionMaster",
	}
	EmployeeCadreTable.Annotation = &entsql.Annotation{
		Table: "EmployeeCadre",
	}
	EmployeeCategoryTable.Annotation = &entsql.Annotation{
		Table: "EmployeeCategory",
	}
	EmployeeDesignationTable.Annotation = &entsql.Annotation{
		Table: "EmployeeDesignation",
	}
	EmployeePostsTable.Annotation = &entsql.Annotation{
		Table: "EmployeePosts",
	}
	ExamTable.ForeignKeys[0].RefTable = VacancyYearsTable
	ExamTable.Annotation = &entsql.Annotation{
		Table: "Exam",
	}
	ExamCalendarTable.ForeignKeys[0].RefTable = ExamTable
	ExamCalendarTable.ForeignKeys[1].RefTable = ExamPapersTable
	ExamCalendarTable.ForeignKeys[2].RefTable = VacancyYearsTable
	ExamCalendarTable.Annotation = &entsql.Annotation{
		Table: "ExamCalendar",
	}
	ExamEligibilityTable.Annotation = &entsql.Annotation{
		Table: "ExamEligibility",
	}
	ExamPapersTable.ForeignKeys[0].RefTable = ExamTable
	ExamPapersTable.Annotation = &entsql.Annotation{
		Table: "ExamPapers",
	}
	FacilityTable.ForeignKeys[0].RefTable = CircleMasterTable
	FacilityTable.ForeignKeys[1].RefTable = RegionMasterTable
	FacilityTable.Annotation = &entsql.Annotation{
		Table: "Facility",
	}
	NodalOfficersTable.ForeignKeys[0].RefTable = ExamTable
	NodalOfficersTable.ForeignKeys[1].RefTable = NotificationTable
	NodalOfficersTable.Annotation = &entsql.Annotation{
		Table: "NodalOfficers",
	}
	NotificationTable.ForeignKeys[0].RefTable = ExamTable
	NotificationTable.ForeignKeys[1].RefTable = ExamCalendarTable
	NotificationTable.Annotation = &entsql.Annotation{
		Table: "Notification",
	}
	PaperTypesTable.ForeignKeys[0].RefTable = ExamPapersTable
	RegionMasterTable.ForeignKeys[0].RefTable = CircleMasterTable
	RegionMasterTable.ForeignKeys[1].RefTable = DivisionMasterTable
	RegionMasterTable.ForeignKeys[2].RefTable = FacilityTable
	RegionMasterTable.Annotation = &entsql.Annotation{
		Table: "RegionMaster",
	}
	VacancyYearsTable.ForeignKeys[0].RefTable = NotificationTable
	VacancyYearsTable.Annotation = &entsql.Annotation{
		Table: "VacancyYears",
	}
	NotificationNotifyRefTable.ForeignKeys[0].RefTable = NotificationTable
	NotificationNotifyRefTable.ForeignKeys[1].RefTable = NotificationTable
}
