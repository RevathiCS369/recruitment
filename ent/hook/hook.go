// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"recruit/ent"
)

// The AdminLoginFunc type is an adapter to allow the use of ordinary
// function as AdminLogin mutator.
type AdminLoginFunc func(context.Context, *ent.AdminLoginMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminLoginFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AdminLoginMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminLoginMutation", m)
}

// The AgeEligibilityFunc type is an adapter to allow the use of ordinary
// function as AgeEligibility mutator.
type AgeEligibilityFunc func(context.Context, *ent.AgeEligibilityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgeEligibilityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AgeEligibilityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgeEligibilityMutation", m)
}

// The ApplicationFunc type is an adapter to allow the use of ordinary
// function as Application mutator.
type ApplicationFunc func(context.Context, *ent.ApplicationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ApplicationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ApplicationMutation", m)
}

// The Cadre_Choice_IPFunc type is an adapter to allow the use of ordinary
// function as Cadre_Choice_IP mutator.
type Cadre_Choice_IPFunc func(context.Context, *ent.CadreChoiceIPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Cadre_Choice_IPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CadreChoiceIPMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CadreChoiceIPMutation", m)
}

// The Cadre_Choice_PAFunc type is an adapter to allow the use of ordinary
// function as Cadre_Choice_PA mutator.
type Cadre_Choice_PAFunc func(context.Context, *ent.CadreChoicePAMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Cadre_Choice_PAFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CadreChoicePAMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CadreChoicePAMutation", m)
}

// The Cadre_Choice_PMFunc type is an adapter to allow the use of ordinary
// function as Cadre_Choice_PM mutator.
type Cadre_Choice_PMFunc func(context.Context, *ent.CadreChoicePMMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Cadre_Choice_PMFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CadreChoicePMMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CadreChoicePMMutation", m)
}

// The Cadre_Choice_PSFunc type is an adapter to allow the use of ordinary
// function as Cadre_Choice_PS mutator.
type Cadre_Choice_PSFunc func(context.Context, *ent.CadreChoicePSMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Cadre_Choice_PSFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CadreChoicePSMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CadreChoicePSMutation", m)
}

// The CenterFunc type is an adapter to allow the use of ordinary
// function as Center mutator.
type CenterFunc func(context.Context, *ent.CenterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CenterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CenterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CenterMutation", m)
}

// The CircleMasterFunc type is an adapter to allow the use of ordinary
// function as CircleMaster mutator.
type CircleMasterFunc func(context.Context, *ent.CircleMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CircleMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CircleMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CircleMasterMutation", m)
}

// The DirectorateUsersFunc type is an adapter to allow the use of ordinary
// function as DirectorateUsers mutator.
type DirectorateUsersFunc func(context.Context, *ent.DirectorateUsersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DirectorateUsersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DirectorateUsersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DirectorateUsersMutation", m)
}

// The DisabilityFunc type is an adapter to allow the use of ordinary
// function as Disability mutator.
type DisabilityFunc func(context.Context, *ent.DisabilityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DisabilityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DisabilityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DisabilityMutation", m)
}

// The DivisionMasterFunc type is an adapter to allow the use of ordinary
// function as DivisionMaster mutator.
type DivisionMasterFunc func(context.Context, *ent.DivisionMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DivisionMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DivisionMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DivisionMasterMutation", m)
}

// The Division_Choice_PAFunc type is an adapter to allow the use of ordinary
// function as Division_Choice_PA mutator.
type Division_Choice_PAFunc func(context.Context, *ent.DivisionChoicePAMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Division_Choice_PAFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DivisionChoicePAMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DivisionChoicePAMutation", m)
}

// The Division_Choice_PMFunc type is an adapter to allow the use of ordinary
// function as Division_Choice_PM mutator.
type Division_Choice_PMFunc func(context.Context, *ent.DivisionChoicePMMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Division_Choice_PMFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DivisionChoicePMMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DivisionChoicePMMutation", m)
}

// The Division_Choice_PSFunc type is an adapter to allow the use of ordinary
// function as Division_Choice_PS mutator.
type Division_Choice_PSFunc func(context.Context, *ent.DivisionChoicePSMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Division_Choice_PSFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DivisionChoicePSMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DivisionChoicePSMutation", m)
}

// The EligibilityMasterFunc type is an adapter to allow the use of ordinary
// function as EligibilityMaster mutator.
type EligibilityMasterFunc func(context.Context, *ent.EligibilityMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EligibilityMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EligibilityMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EligibilityMasterMutation", m)
}

// The EmployeeCadreFunc type is an adapter to allow the use of ordinary
// function as EmployeeCadre mutator.
type EmployeeCadreFunc func(context.Context, *ent.EmployeeCadreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeCadreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeeCadreMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeCadreMutation", m)
}

// The EmployeeCategoryFunc type is an adapter to allow the use of ordinary
// function as EmployeeCategory mutator.
type EmployeeCategoryFunc func(context.Context, *ent.EmployeeCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeeCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeCategoryMutation", m)
}

// The EmployeeDesignationFunc type is an adapter to allow the use of ordinary
// function as EmployeeDesignation mutator.
type EmployeeDesignationFunc func(context.Context, *ent.EmployeeDesignationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeDesignationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeeDesignationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeDesignationMutation", m)
}

// The EmployeeMasterFunc type is an adapter to allow the use of ordinary
// function as EmployeeMaster mutator.
type EmployeeMasterFunc func(context.Context, *ent.EmployeeMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeeMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeMasterMutation", m)
}

// The EmployeePostsFunc type is an adapter to allow the use of ordinary
// function as EmployeePosts mutator.
type EmployeePostsFunc func(context.Context, *ent.EmployeePostsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeePostsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeePostsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeePostsMutation", m)
}

// The EmployeesFunc type is an adapter to allow the use of ordinary
// function as Employees mutator.
type EmployeesFunc func(context.Context, *ent.EmployeesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeesMutation", m)
}

// The ExamFunc type is an adapter to allow the use of ordinary
// function as Exam mutator.
type ExamFunc func(context.Context, *ent.ExamMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExamFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamMutation", m)
}

// The ExamCalendarFunc type is an adapter to allow the use of ordinary
// function as ExamCalendar mutator.
type ExamCalendarFunc func(context.Context, *ent.ExamCalendarMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExamCalendarFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamCalendarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamCalendarMutation", m)
}

// The ExamPapersFunc type is an adapter to allow the use of ordinary
// function as ExamPapers mutator.
type ExamPapersFunc func(context.Context, *ent.ExamPapersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExamPapersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamPapersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamPapersMutation", m)
}

// The ExamTypeFunc type is an adapter to allow the use of ordinary
// function as ExamType mutator.
type ExamTypeFunc func(context.Context, *ent.ExamTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExamTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamTypeMutation", m)
}

// The Exam_Applications_IPFunc type is an adapter to allow the use of ordinary
// function as Exam_Applications_IP mutator.
type Exam_Applications_IPFunc func(context.Context, *ent.ExamApplicationsIPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_Applications_IPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamApplicationsIPMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamApplicationsIPMutation", m)
}

// The Exam_Applications_PSFunc type is an adapter to allow the use of ordinary
// function as Exam_Applications_PS mutator.
type Exam_Applications_PSFunc func(context.Context, *ent.ExamApplicationsPSMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_Applications_PSFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamApplicationsPSMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamApplicationsPSMutation", m)
}

// The Exam_IPFunc type is an adapter to allow the use of ordinary
// function as Exam_IP mutator.
type Exam_IPFunc func(context.Context, *ent.ExamIPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_IPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamIPMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamIPMutation", m)
}

// The Exam_PAFunc type is an adapter to allow the use of ordinary
// function as Exam_PA mutator.
type Exam_PAFunc func(context.Context, *ent.ExamPAMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_PAFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamPAMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamPAMutation", m)
}

// The Exam_PMFunc type is an adapter to allow the use of ordinary
// function as Exam_PM mutator.
type Exam_PMFunc func(context.Context, *ent.ExamPMMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_PMFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamPMMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamPMMutation", m)
}

// The Exam_PSFunc type is an adapter to allow the use of ordinary
// function as Exam_PS mutator.
type Exam_PSFunc func(context.Context, *ent.ExamPSMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Exam_PSFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ExamPSMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExamPSMutation", m)
}

// The FacilityFunc type is an adapter to allow the use of ordinary
// function as Facility mutator.
type FacilityFunc func(context.Context, *ent.FacilityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FacilityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FacilityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FacilityMutation", m)
}

// The LoginFunc type is an adapter to allow the use of ordinary
// function as Login mutator.
type LoginFunc func(context.Context, *ent.LoginMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LoginFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LoginMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LoginMutation", m)
}

// The NodalOfficerFunc type is an adapter to allow the use of ordinary
// function as NodalOfficer mutator.
type NodalOfficerFunc func(context.Context, *ent.NodalOfficerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NodalOfficerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NodalOfficerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NodalOfficerMutation", m)
}

// The NotificationFunc type is an adapter to allow the use of ordinary
// function as Notification mutator.
type NotificationFunc func(context.Context, *ent.NotificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotificationMutation", m)
}

// The PaperTypesFunc type is an adapter to allow the use of ordinary
// function as PaperTypes mutator.
type PaperTypesFunc func(context.Context, *ent.PaperTypesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaperTypesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PaperTypesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaperTypesMutation", m)
}

// The PlaceOfPreferenceIPFunc type is an adapter to allow the use of ordinary
// function as PlaceOfPreferenceIP mutator.
type PlaceOfPreferenceIPFunc func(context.Context, *ent.PlaceOfPreferenceIPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceOfPreferenceIPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlaceOfPreferenceIPMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceOfPreferenceIPMutation", m)
}

// The RecommendationsIPApplicationsFunc type is an adapter to allow the use of ordinary
// function as RecommendationsIPApplications mutator.
type RecommendationsIPApplicationsFunc func(context.Context, *ent.RecommendationsIPApplicationsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecommendationsIPApplicationsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RecommendationsIPApplicationsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecommendationsIPApplicationsMutation", m)
}

// The RegionMasterFunc type is an adapter to allow the use of ordinary
// function as RegionMaster mutator.
type RegionMasterFunc func(context.Context, *ent.RegionMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegionMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RegionMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegionMasterMutation", m)
}

// The Reversal_Application_IPFunc type is an adapter to allow the use of ordinary
// function as Reversal_Application_IP mutator.
type Reversal_Application_IPFunc func(context.Context, *ent.ReversalApplicationIPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f Reversal_Application_IPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReversalApplicationIPMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReversalApplicationIPMutation", m)
}

// The RoleMasterFunc type is an adapter to allow the use of ordinary
// function as RoleMaster mutator.
type RoleMasterFunc func(context.Context, *ent.RoleMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoleMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMasterMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserMasterFunc type is an adapter to allow the use of ordinary
// function as UserMaster mutator.
type UserMasterFunc func(context.Context, *ent.UserMasterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserMasterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMasterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMasterMutation", m)
}

// The VacancyYearFunc type is an adapter to allow the use of ordinary
// function as VacancyYear mutator.
type VacancyYearFunc func(context.Context, *ent.VacancyYearMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VacancyYearFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VacancyYearMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VacancyYearMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
