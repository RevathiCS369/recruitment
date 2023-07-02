// Code generated by ent, DO NOT EDIT.

package ent

import (
	"recruit/ent/employees"
	"recruit/ent/examcalendar"
	"recruit/ent/exampapers"
	"recruit/ent/papertypes"
	"recruit/ent/schema"
	"recruit/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	employeesFields := schema.Employees{}.Fields()
	_ = employeesFields
	// employeesDescIDVerified is the schema descriptor for IDVerified field.
	employeesDescIDVerified := employeesFields[2].Descriptor()
	// employees.DefaultIDVerified holds the default value on creation for the IDVerified field.
	employees.DefaultIDVerified = employeesDescIDVerified.Default.(bool)
	// employeesDescIDRemStatus is the schema descriptor for IDRemStatus field.
	employeesDescIDRemStatus := employeesFields[3].Descriptor()
	// employees.DefaultIDRemStatus holds the default value on creation for the IDRemStatus field.
	employees.DefaultIDRemStatus = employeesDescIDRemStatus.Default.(bool)
	// employeesDescNameVerified is the schema descriptor for nameVerified field.
	employeesDescNameVerified := employeesFields[6].Descriptor()
	// employees.DefaultNameVerified holds the default value on creation for the nameVerified field.
	employees.DefaultNameVerified = employeesDescNameVerified.Default.(bool)
	// employeesDescNameRemStatus is the schema descriptor for nameRemStatus field.
	employeesDescNameRemStatus := employeesFields[7].Descriptor()
	// employees.DefaultNameRemStatus holds the default value on creation for the nameRemStatus field.
	employees.DefaultNameRemStatus = employeesDescNameRemStatus.Default.(bool)
	// employeesDescFathersNameVerified is the schema descriptor for FathersNameVerified field.
	employeesDescFathersNameVerified := employeesFields[10].Descriptor()
	// employees.DefaultFathersNameVerified holds the default value on creation for the FathersNameVerified field.
	employees.DefaultFathersNameVerified = employeesDescFathersNameVerified.Default.(bool)
	// employeesDescFathersNameRemStatus is the schema descriptor for FathersNameRemStatus field.
	employeesDescFathersNameRemStatus := employeesFields[11].Descriptor()
	// employees.DefaultFathersNameRemStatus holds the default value on creation for the FathersNameRemStatus field.
	employees.DefaultFathersNameRemStatus = employeesDescFathersNameRemStatus.Default.(bool)
	// employeesDescDOBVerified is the schema descriptor for DOBVerified field.
	employeesDescDOBVerified := employeesFields[14].Descriptor()
	// employees.DefaultDOBVerified holds the default value on creation for the DOBVerified field.
	employees.DefaultDOBVerified = employeesDescDOBVerified.Default.(bool)
	// employeesDescDOBRemStatus is the schema descriptor for DOBRemStatus field.
	employeesDescDOBRemStatus := employeesFields[15].Descriptor()
	// employees.DefaultDOBRemStatus holds the default value on creation for the DOBRemStatus field.
	employees.DefaultDOBRemStatus = employeesDescDOBRemStatus.Default.(bool)
	// employeesDescGenderVerified is the schema descriptor for genderVerified field.
	employeesDescGenderVerified := employeesFields[18].Descriptor()
	// employees.DefaultGenderVerified holds the default value on creation for the genderVerified field.
	employees.DefaultGenderVerified = employeesDescGenderVerified.Default.(bool)
	// employeesDescGenderRemStatus is the schema descriptor for genderRemStatus field.
	employeesDescGenderRemStatus := employeesFields[19].Descriptor()
	// employees.DefaultGenderRemStatus holds the default value on creation for the genderRemStatus field.
	employees.DefaultGenderRemStatus = employeesDescGenderRemStatus.Default.(bool)
	// employeesDescEmployeeCategoryCodeVerified is the schema descriptor for EmployeeCategoryCodeVerified field.
	employeesDescEmployeeCategoryCodeVerified := employeesFields[26].Descriptor()
	// employees.DefaultEmployeeCategoryCodeVerified holds the default value on creation for the EmployeeCategoryCodeVerified field.
	employees.DefaultEmployeeCategoryCodeVerified = employeesDescEmployeeCategoryCodeVerified.Default.(bool)
	// employeesDescEmployeeCategoryCodeRemStatus is the schema descriptor for EmployeeCategoryCodeRemStatus field.
	employeesDescEmployeeCategoryCodeRemStatus := employeesFields[27].Descriptor()
	// employees.DefaultEmployeeCategoryCodeRemStatus holds the default value on creation for the EmployeeCategoryCodeRemStatus field.
	employees.DefaultEmployeeCategoryCodeRemStatus = employeesDescEmployeeCategoryCodeRemStatus.Default.(bool)
	// employeesDescWithDisabilityVerified is the schema descriptor for WithDisabilityVerified field.
	employeesDescWithDisabilityVerified := employeesFields[30].Descriptor()
	// employees.DefaultWithDisabilityVerified holds the default value on creation for the WithDisabilityVerified field.
	employees.DefaultWithDisabilityVerified = employeesDescWithDisabilityVerified.Default.(bool)
	// employeesDescWithDisabilityRemStatus is the schema descriptor for WithDisabilityRemStatus field.
	employeesDescWithDisabilityRemStatus := employeesFields[31].Descriptor()
	// employees.DefaultWithDisabilityRemStatus holds the default value on creation for the WithDisabilityRemStatus field.
	employees.DefaultWithDisabilityRemStatus = employeesDescWithDisabilityRemStatus.Default.(bool)
	// employeesDescDisabilityTypeVerified is the schema descriptor for DisabilityTypeVerified field.
	employeesDescDisabilityTypeVerified := employeesFields[34].Descriptor()
	// employees.DefaultDisabilityTypeVerified holds the default value on creation for the DisabilityTypeVerified field.
	employees.DefaultDisabilityTypeVerified = employeesDescDisabilityTypeVerified.Default.(bool)
	// employeesDescDisabilityTypeRemStatus is the schema descriptor for DisabilityTypeRemStatus field.
	employeesDescDisabilityTypeRemStatus := employeesFields[35].Descriptor()
	// employees.DefaultDisabilityTypeRemStatus holds the default value on creation for the DisabilityTypeRemStatus field.
	employees.DefaultDisabilityTypeRemStatus = employeesDescDisabilityTypeRemStatus.Default.(bool)
	// employeesDescDisabilityPercentageVerified is the schema descriptor for DisabilityPercentageVerified field.
	employeesDescDisabilityPercentageVerified := employeesFields[38].Descriptor()
	// employees.DefaultDisabilityPercentageVerified holds the default value on creation for the DisabilityPercentageVerified field.
	employees.DefaultDisabilityPercentageVerified = employeesDescDisabilityPercentageVerified.Default.(bool)
	// employeesDescDisabilityPercentageRemStatus is the schema descriptor for DisabilityPercentageRemStatus field.
	employeesDescDisabilityPercentageRemStatus := employeesFields[39].Descriptor()
	// employees.DefaultDisabilityPercentageRemStatus holds the default value on creation for the DisabilityPercentageRemStatus field.
	employees.DefaultDisabilityPercentageRemStatus = employeesDescDisabilityPercentageRemStatus.Default.(bool)
	// employeesDescSignatureVerified is the schema descriptor for SignatureVerified field.
	employeesDescSignatureVerified := employeesFields[42].Descriptor()
	// employees.DefaultSignatureVerified holds the default value on creation for the SignatureVerified field.
	employees.DefaultSignatureVerified = employeesDescSignatureVerified.Default.(bool)
	// employeesDescSignatureRemStatus is the schema descriptor for SignatureRemStatus field.
	employeesDescSignatureRemStatus := employeesFields[43].Descriptor()
	// employees.DefaultSignatureRemStatus holds the default value on creation for the SignatureRemStatus field.
	employees.DefaultSignatureRemStatus = employeesDescSignatureRemStatus.Default.(bool)
	// employeesDescPhotoVerified is the schema descriptor for PhotoVerified field.
	employeesDescPhotoVerified := employeesFields[46].Descriptor()
	// employees.DefaultPhotoVerified holds the default value on creation for the PhotoVerified field.
	employees.DefaultPhotoVerified = employeesDescPhotoVerified.Default.(bool)
	// employeesDescPhotoRemStatus is the schema descriptor for PhotoRemStatus field.
	employeesDescPhotoRemStatus := employeesFields[47].Descriptor()
	// employees.DefaultPhotoRemStatus holds the default value on creation for the PhotoRemStatus field.
	employees.DefaultPhotoRemStatus = employeesDescPhotoRemStatus.Default.(bool)
	// employeesDescEmployeeCadreVerified is the schema descriptor for EmployeeCadreVerified field.
	employeesDescEmployeeCadreVerified := employeesFields[51].Descriptor()
	// employees.DefaultEmployeeCadreVerified holds the default value on creation for the EmployeeCadreVerified field.
	employees.DefaultEmployeeCadreVerified = employeesDescEmployeeCadreVerified.Default.(bool)
	// employeesDescEmployeeCadreRemStatus is the schema descriptor for EmployeeCadreRemStatus field.
	employeesDescEmployeeCadreRemStatus := employeesFields[52].Descriptor()
	// employees.DefaultEmployeeCadreRemStatus holds the default value on creation for the EmployeeCadreRemStatus field.
	employees.DefaultEmployeeCadreRemStatus = employeesDescEmployeeCadreRemStatus.Default.(bool)
	// employeesDescEmployeeDesignationVerified is the schema descriptor for EmployeeDesignationVerified field.
	employeesDescEmployeeDesignationVerified := employeesFields[56].Descriptor()
	// employees.DefaultEmployeeDesignationVerified holds the default value on creation for the EmployeeDesignationVerified field.
	employees.DefaultEmployeeDesignationVerified = employeesDescEmployeeDesignationVerified.Default.(bool)
	// employeesDescEmployeeDesignationRemStatus is the schema descriptor for EmployeeDesignationRemStatus field.
	employeesDescEmployeeDesignationRemStatus := employeesFields[57].Descriptor()
	// employees.DefaultEmployeeDesignationRemStatus holds the default value on creation for the EmployeeDesignationRemStatus field.
	employees.DefaultEmployeeDesignationRemStatus = employeesDescEmployeeDesignationRemStatus.Default.(bool)
	// employeesDescCircleVerified is the schema descriptor for CircleVerified field.
	employeesDescCircleVerified := employeesFields[61].Descriptor()
	// employees.DefaultCircleVerified holds the default value on creation for the CircleVerified field.
	employees.DefaultCircleVerified = employeesDescCircleVerified.Default.(bool)
	// employeesDescCircleRemStatus is the schema descriptor for CircleRemStatus field.
	employeesDescCircleRemStatus := employeesFields[62].Descriptor()
	// employees.DefaultCircleRemStatus holds the default value on creation for the CircleRemStatus field.
	employees.DefaultCircleRemStatus = employeesDescCircleRemStatus.Default.(bool)
	// employeesDescRegionVerified is the schema descriptor for RegionVerified field.
	employeesDescRegionVerified := employeesFields[66].Descriptor()
	// employees.DefaultRegionVerified holds the default value on creation for the RegionVerified field.
	employees.DefaultRegionVerified = employeesDescRegionVerified.Default.(bool)
	// employeesDescRegionRemStatus is the schema descriptor for RegionRemStatus field.
	employeesDescRegionRemStatus := employeesFields[67].Descriptor()
	// employees.DefaultRegionRemStatus holds the default value on creation for the RegionRemStatus field.
	employees.DefaultRegionRemStatus = employeesDescRegionRemStatus.Default.(bool)
	// employeesDescDivisionVerified is the schema descriptor for DivisionVerified field.
	employeesDescDivisionVerified := employeesFields[71].Descriptor()
	// employees.DefaultDivisionVerified holds the default value on creation for the DivisionVerified field.
	employees.DefaultDivisionVerified = employeesDescDivisionVerified.Default.(bool)
	// employeesDescDivisionRemStatus is the schema descriptor for DivisionRemStatus field.
	employeesDescDivisionRemStatus := employeesFields[72].Descriptor()
	// employees.DefaultDivisionRemStatus holds the default value on creation for the DivisionRemStatus field.
	employees.DefaultDivisionRemStatus = employeesDescDivisionRemStatus.Default.(bool)
	// employeesDescOfficeVerified is the schema descriptor for OfficeVerified field.
	employeesDescOfficeVerified := employeesFields[76].Descriptor()
	// employees.DefaultOfficeVerified holds the default value on creation for the OfficeVerified field.
	employees.DefaultOfficeVerified = employeesDescOfficeVerified.Default.(bool)
	// employeesDescOfficeRemStatus is the schema descriptor for OfficeRemStatus field.
	employeesDescOfficeRemStatus := employeesFields[77].Descriptor()
	// employees.DefaultOfficeRemStatus holds the default value on creation for the OfficeRemStatus field.
	employees.DefaultOfficeRemStatus = employeesDescOfficeRemStatus.Default.(bool)
	// employeesDescRoleVerified is the schema descriptor for RoleVerified field.
	employeesDescRoleVerified := employeesFields[80].Descriptor()
	// employees.DefaultRoleVerified holds the default value on creation for the RoleVerified field.
	employees.DefaultRoleVerified = employeesDescRoleVerified.Default.(bool)
	// employeesDescRoleRemStatus is the schema descriptor for RoleRemStatus field.
	employeesDescRoleRemStatus := employeesFields[81].Descriptor()
	// employees.DefaultRoleRemStatus holds the default value on creation for the RoleRemStatus field.
	employees.DefaultRoleRemStatus = employeesDescRoleRemStatus.Default.(bool)
	// employeesDescDCCSVerified is the schema descriptor for DCCSVerified field.
	employeesDescDCCSVerified := employeesFields[84].Descriptor()
	// employees.DefaultDCCSVerified holds the default value on creation for the DCCSVerified field.
	employees.DefaultDCCSVerified = employeesDescDCCSVerified.Default.(bool)
	// employeesDescDCCSRemStatus is the schema descriptor for DCCSRemStatus field.
	employeesDescDCCSRemStatus := employeesFields[85].Descriptor()
	// employees.DefaultDCCSRemStatus holds the default value on creation for the DCCSRemStatus field.
	employees.DefaultDCCSRemStatus = employeesDescDCCSRemStatus.Default.(bool)
	// employeesDescDCInPresentCadreVerified is the schema descriptor for DCInPresentCadreVerified field.
	employeesDescDCInPresentCadreVerified := employeesFields[88].Descriptor()
	// employees.DefaultDCInPresentCadreVerified holds the default value on creation for the DCInPresentCadreVerified field.
	employees.DefaultDCInPresentCadreVerified = employeesDescDCInPresentCadreVerified.Default.(bool)
	// employeesDescDCInPresentCadreRemStatus is the schema descriptor for DCInPresentCadreRemStatus field.
	employeesDescDCInPresentCadreRemStatus := employeesFields[89].Descriptor()
	// employees.DefaultDCInPresentCadreRemStatus holds the default value on creation for the DCInPresentCadreRemStatus field.
	employees.DefaultDCInPresentCadreRemStatus = employeesDescDCInPresentCadreRemStatus.Default.(bool)
	// employeesDescAPSWorkingVerified is the schema descriptor for APSWorkingVerified field.
	employeesDescAPSWorkingVerified := employeesFields[92].Descriptor()
	// employees.DefaultAPSWorkingVerified holds the default value on creation for the APSWorkingVerified field.
	employees.DefaultAPSWorkingVerified = employeesDescAPSWorkingVerified.Default.(bool)
	// employeesDescAPSWorkingRemStatus is the schema descriptor for APSWorkingRemStatus field.
	employeesDescAPSWorkingRemStatus := employeesFields[93].Descriptor()
	// employees.DefaultAPSWorkingRemStatus holds the default value on creation for the APSWorkingRemStatus field.
	employees.DefaultAPSWorkingRemStatus = employeesDescAPSWorkingRemStatus.Default.(bool)
	// employeesDescProfilestatus is the schema descriptor for profilestatus field.
	employeesDescProfilestatus := employeesFields[95].Descriptor()
	// employees.DefaultProfilestatus holds the default value on creation for the profilestatus field.
	employees.DefaultProfilestatus = employeesDescProfilestatus.Default.(bool)
	examcalendarFields := schema.ExamCalendar{}.Fields()
	_ = examcalendarFields
	// examcalendarDescExamName is the schema descriptor for ExamName field.
	examcalendarDescExamName := examcalendarFields[2].Descriptor()
	// examcalendar.ExamNameValidator is a validator for the "ExamName" field. It is called by the builders before save.
	examcalendar.ExamNameValidator = examcalendarDescExamName.Validators[0].(func(string) error)
	exampapersFields := schema.ExamPapers{}.Fields()
	_ = exampapersFields
	// exampapersDescPaperDescription is the schema descriptor for PaperDescription field.
	exampapersDescPaperDescription := exampapersFields[1].Descriptor()
	// exampapers.PaperDescriptionValidator is a validator for the "PaperDescription" field. It is called by the builders before save.
	exampapers.PaperDescriptionValidator = func() func(string) error {
		validators := exampapersDescPaperDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_PaperDescription string) error {
			for _, fn := range fns {
				if err := fn(_PaperDescription); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// exampapersDescCompetitiveQualifying is the schema descriptor for competitiveQualifying field.
	exampapersDescCompetitiveQualifying := exampapersFields[3].Descriptor()
	// exampapers.CompetitiveQualifyingValidator is a validator for the "competitiveQualifying" field. It is called by the builders before save.
	exampapers.CompetitiveQualifyingValidator = exampapersDescCompetitiveQualifying.Validators[0].(func(string) error)
	// exampapersDescExceptionForDisability is the schema descriptor for exceptionForDisability field.
	exampapersDescExceptionForDisability := exampapersFields[4].Descriptor()
	// exampapers.ExceptionForDisabilityValidator is a validator for the "exceptionForDisability" field. It is called by the builders before save.
	exampapers.ExceptionForDisabilityValidator = exampapersDescExceptionForDisability.Validators[0].(func(string) error)
	// exampapersDescMaximumMarks is the schema descriptor for MaximumMarks field.
	exampapersDescMaximumMarks := exampapersFields[5].Descriptor()
	// exampapers.MaximumMarksValidator is a validator for the "MaximumMarks" field. It is called by the builders before save.
	exampapers.MaximumMarksValidator = exampapersDescMaximumMarks.Validators[0].(func(int) error)
	// exampapersDescDuration is the schema descriptor for Duration field.
	exampapersDescDuration := exampapersFields[6].Descriptor()
	// exampapers.DurationValidator is a validator for the "Duration" field. It is called by the builders before save.
	exampapers.DurationValidator = exampapersDescDuration.Validators[0].(func(int) error)
	// exampapersDescLocalLanguageAllowedQuestionPaper is the schema descriptor for localLanguageAllowedQuestionPaper field.
	exampapersDescLocalLanguageAllowedQuestionPaper := exampapersFields[7].Descriptor()
	// exampapers.LocalLanguageAllowedQuestionPaperValidator is a validator for the "localLanguageAllowedQuestionPaper" field. It is called by the builders before save.
	exampapers.LocalLanguageAllowedQuestionPaperValidator = exampapersDescLocalLanguageAllowedQuestionPaper.Validators[0].(func(string) error)
	// exampapersDescLocalLanguageAllowedAnswerPaper is the schema descriptor for localLanguageAllowedAnswerPaper field.
	exampapersDescLocalLanguageAllowedAnswerPaper := exampapersFields[8].Descriptor()
	// exampapers.LocalLanguageAllowedAnswerPaperValidator is a validator for the "localLanguageAllowedAnswerPaper" field. It is called by the builders before save.
	exampapers.LocalLanguageAllowedAnswerPaperValidator = exampapersDescLocalLanguageAllowedAnswerPaper.Validators[0].(func(string) error)
	// exampapersDescOrderNumber is the schema descriptor for OrderNumber field.
	exampapersDescOrderNumber := exampapersFields[9].Descriptor()
	// exampapers.OrderNumberValidator is a validator for the "OrderNumber" field. It is called by the builders before save.
	exampapers.OrderNumberValidator = exampapersDescOrderNumber.Validators[0].(func(string) error)
	// exampapersDescPaperStatus is the schema descriptor for PaperStatus field.
	exampapersDescPaperStatus := exampapersFields[10].Descriptor()
	// exampapers.PaperStatusValidator is a validator for the "PaperStatus" field. It is called by the builders before save.
	exampapers.PaperStatusValidator = func() func(string) error {
		validators := exampapersDescPaperStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_PaperStatus string) error {
			for _, fn := range fns {
				if err := fn(_PaperStatus); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	papertypesFields := schema.PaperTypes{}.Fields()
	_ = papertypesFields
	// papertypesDescPaperTypeDescription is the schema descriptor for PaperTypeDescription field.
	papertypesDescPaperTypeDescription := papertypesFields[2].Descriptor()
	// papertypes.PaperTypeDescriptionValidator is a validator for the "PaperTypeDescription" field. It is called by the builders before save.
	papertypes.PaperTypeDescriptionValidator = func() func(string) error {
		validators := papertypesDescPaperTypeDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_PaperTypeDescription string) error {
			for _, fn := range fns {
				if err := fn(_PaperTypeDescription); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// papertypesDescOrderNumber is the schema descriptor for OrderNumber field.
	papertypesDescOrderNumber := papertypesFields[3].Descriptor()
	// papertypes.OrderNumberValidator is a validator for the "OrderNumber" field. It is called by the builders before save.
	papertypes.OrderNumberValidator = papertypesDescOrderNumber.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIDVerified is the schema descriptor for IDVerified field.
	userDescIDVerified := userFields[1].Descriptor()
	// user.DefaultIDVerified holds the default value on creation for the IDVerified field.
	user.DefaultIDVerified = userDescIDVerified.Default.(bool)
	// userDescIDRemStatus is the schema descriptor for IDRemStatus field.
	userDescIDRemStatus := userFields[2].Descriptor()
	// user.DefaultIDRemStatus holds the default value on creation for the IDRemStatus field.
	user.DefaultIDRemStatus = userDescIDRemStatus.Default.(bool)
	// userDescNameVerified is the schema descriptor for nameVerified field.
	userDescNameVerified := userFields[5].Descriptor()
	// user.DefaultNameVerified holds the default value on creation for the nameVerified field.
	user.DefaultNameVerified = userDescNameVerified.Default.(bool)
	// userDescNameRemStatus is the schema descriptor for nameRemStatus field.
	userDescNameRemStatus := userFields[6].Descriptor()
	// user.DefaultNameRemStatus holds the default value on creation for the nameRemStatus field.
	user.DefaultNameRemStatus = userDescNameRemStatus.Default.(bool)
	// userDescDOBVerified is the schema descriptor for DOBVerified field.
	userDescDOBVerified := userFields[9].Descriptor()
	// user.DefaultDOBVerified holds the default value on creation for the DOBVerified field.
	user.DefaultDOBVerified = userDescDOBVerified.Default.(bool)
	// userDescDOBRemStatus is the schema descriptor for DOBRemStatus field.
	userDescDOBRemStatus := userFields[10].Descriptor()
	// user.DefaultDOBRemStatus holds the default value on creation for the DOBRemStatus field.
	user.DefaultDOBRemStatus = userDescDOBRemStatus.Default.(bool)
	// userDescGenderVerified is the schema descriptor for genderVerified field.
	userDescGenderVerified := userFields[13].Descriptor()
	// user.DefaultGenderVerified holds the default value on creation for the genderVerified field.
	user.DefaultGenderVerified = userDescGenderVerified.Default.(bool)
	// userDescGenderRemStatus is the schema descriptor for genderRemStatus field.
	userDescGenderRemStatus := userFields[14].Descriptor()
	// user.DefaultGenderRemStatus holds the default value on creation for the genderRemStatus field.
	user.DefaultGenderRemStatus = userDescGenderRemStatus.Default.(bool)
	// userDescCadreidVerified is the schema descriptor for cadreidVerified field.
	userDescCadreidVerified := userFields[17].Descriptor()
	// user.DefaultCadreidVerified holds the default value on creation for the cadreidVerified field.
	user.DefaultCadreidVerified = userDescCadreidVerified.Default.(bool)
	// userDescCadreidRemStatus is the schema descriptor for cadreidRemStatus field.
	userDescCadreidRemStatus := userFields[18].Descriptor()
	// user.DefaultCadreidRemStatus holds the default value on creation for the cadreidRemStatus field.
	user.DefaultCadreidRemStatus = userDescCadreidRemStatus.Default.(bool)
	// userDescOfficeIDVerified is the schema descriptor for officeIDVerified field.
	userDescOfficeIDVerified := userFields[21].Descriptor()
	// user.DefaultOfficeIDVerified holds the default value on creation for the officeIDVerified field.
	user.DefaultOfficeIDVerified = userDescOfficeIDVerified.Default.(bool)
	// userDescOfficeIDRemStatus is the schema descriptor for officeIDRemStatus field.
	userDescOfficeIDRemStatus := userFields[22].Descriptor()
	// user.DefaultOfficeIDRemStatus holds the default value on creation for the officeIDRemStatus field.
	user.DefaultOfficeIDRemStatus = userDescOfficeIDRemStatus.Default.(bool)
	// userDescPHVerified is the schema descriptor for PHVerified field.
	userDescPHVerified := userFields[25].Descriptor()
	// user.DefaultPHVerified holds the default value on creation for the PHVerified field.
	user.DefaultPHVerified = userDescPHVerified.Default.(bool)
	// userDescPHRemStatus is the schema descriptor for PHRemStatus field.
	userDescPHRemStatus := userFields[26].Descriptor()
	// user.DefaultPHRemStatus holds the default value on creation for the PHRemStatus field.
	user.DefaultPHRemStatus = userDescPHRemStatus.Default.(bool)
	// userDescPHDetailsVerified is the schema descriptor for PHDetailsVerified field.
	userDescPHDetailsVerified := userFields[29].Descriptor()
	// user.DefaultPHDetailsVerified holds the default value on creation for the PHDetailsVerified field.
	user.DefaultPHDetailsVerified = userDescPHDetailsVerified.Default.(bool)
	// userDescPHDetailsRemStatus is the schema descriptor for PHDetailsRemStatus field.
	userDescPHDetailsRemStatus := userFields[30].Descriptor()
	// user.DefaultPHDetailsRemStatus holds the default value on creation for the PHDetailsRemStatus field.
	user.DefaultPHDetailsRemStatus = userDescPHDetailsRemStatus.Default.(bool)
	// userDescAPSWorkingVerified is the schema descriptor for APSWorkingVerified field.
	userDescAPSWorkingVerified := userFields[33].Descriptor()
	// user.DefaultAPSWorkingVerified holds the default value on creation for the APSWorkingVerified field.
	user.DefaultAPSWorkingVerified = userDescAPSWorkingVerified.Default.(bool)
	// userDescAPSWorkingRemStatus is the schema descriptor for APSWorkingRemStatus field.
	userDescAPSWorkingRemStatus := userFields[34].Descriptor()
	// user.DefaultAPSWorkingRemStatus holds the default value on creation for the APSWorkingRemStatus field.
	user.DefaultAPSWorkingRemStatus = userDescAPSWorkingRemStatus.Default.(bool)
	// userDescProfilestatus is the schema descriptor for profilestatus field.
	userDescProfilestatus := userFields[36].Descriptor()
	// user.DefaultProfilestatus holds the default value on creation for the profilestatus field.
	user.DefaultProfilestatus = userDescProfilestatus.Default.(bool)
}
