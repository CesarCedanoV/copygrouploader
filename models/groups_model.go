package models

import (
	"time"
)

type Groups struct {
	Id                                       int64                                `json:"groups_id,omitempty"`
	CarrierId                                int64                                `json:"groups_carrier_id,omitempty"`
	GroupNum                                 string                               `json:"groups_group_num,omitempty"`
	GroupName                                *string                              `json:"groups_group_name,omitempty"`
	Address1                                 *string                              `json:"groups_address1,omitempty"`
	Address2                                 *string                              `json:"groups_address2,omitempty"`
	City                                     *string                              `json:"groups_city,omitempty"`
	State                                    *string                              `json:"groups_state,omitempty"`
	Zip                                      *string                              `json:"groups_zip,omitempty"`
	Phone                                    *string                              `json:"groups_phone,omitempty"`
	PhoneExt                                 *string                              `json:"groups_phone_ext,omitempty"`
	Fax                                      *string                              `json:"groups_fax,omitempty"`
	BenefitYearStart                         *time.Time                           `json:"groups_benefit_year_start,omitempty"`
	MemberDetermination                      string                               `json:"groups_member_determination,omitempty"`
	PersoncodeDetermination                  *string                              `json:"groups_personcode_determination,omitempty"`
	DobDetermination                         *string                              `json:"groups_dob_determination,omitempty"`
	DobDays                                  *int32                               `json:"groups_dob_days,omitempty"`
	CardType                                 *string                              `json:"groups_card_type,omitempty"`
	CardName                                 *string                              `json:"groups_card_name,omitempty"`
	DynamicEnrollmentDays                    *int32                               `json:"groups_dynamic_enrollment_days,omitempty"`
	DynamicEnrollmentIdLogic                 *string                              `json:"groups_dynamic_enrollment_id_logic,omitempty"`
	PrescriberLogic                          *string                              `json:"groups_prescriber_logic,omitempty"`
	MaxRxPeriod                              *string                              `json:"groups_max_rx_period,omitempty"`
	MaxRx                                    *int32                               `json:"groups_max_rx,omitempty"`
	TaxLogic                                 *string                              `json:"groups_tax_logic,omitempty"`
	ChildCutoffAge                           *int32                               `json:"groups_child_cutoff_age,omitempty"`
	ChildCutoffAgeType                       *string                              `json:"groups_child_cutoff_age_type,omitempty"`
	StudentCutoffAge                         *int32                               `json:"groups_student_cutoff_age,omitempty"`
	StudentCutoffAgeType                     *string                              `json:"groups_student_cutoff_age_type,omitempty"`
	OtherCutoffAge                           *int32                               `json:"groups_other_cutoff_age,omitempty"`
	OtherCutoffAgeType                       *string                              `json:"groups_other_cutoff_age_type,omitempty"`
	RefillTooSoonGpiDigits                   int32                                `json:"groups_refill_too_soon_gpi_digits,omitempty"`
	ContactName                              *string                              `json:"groups_contact_name,omitempty"`
	ContactEmailAddress                      *string                              `json:"groups_contact_email_address,omitempty"`
	OptSecondaryPayer                        *bool                                `json:"groups_opt_secondary_payer,omitempty"`
	OptIgnoreHistory                         *bool                                `json:"groups_opt_ignore_history,omitempty"`
	OptTestOnly                              *bool                                `json:"groups_opt_test_only,omitempty"`
	OptMaxDollarInclCopayAmt                 *bool                                `json:"groups_opt_max_dollar_incl_copay_amt,omitempty"`
	StartDate                                time.Time                            `json:"groups_start_date,omitempty"`
	EndDate                                  time.Time                            `json:"groups_end_date,omitempty"`
	ModStartDate                             time.Time                            `json:"groups_mod_start_date,omitempty"`
	ModUser                                  int64                                `json:"groups_mod_user,omitempty"`
	Opt001PatientsOnly                       *bool                                `json:"groups_opt_001_patients_only,omitempty"`
	InjuryDateOffset                         *int32                               `json:"groups_injury_date_offset,omitempty"`
	TermPatientIfNoClaimsOffset              *int32                               `json:"groups_term_patient_if_no_claims_offset,omitempty"`
	OptActivatePatientProgramCodes           *bool                                `json:"groups_opt_activate_patient_program_codes,omitempty"`
	OptDifferentPrescriptionNumberIgnoresRts *bool                                `json:"groups_opt_different_prescription_number_ignores_rts,omitempty"`
	GroupType                                string                               `json:"groups_group_type,omitempty"`
	OptUseGroupStartAsEnrollmentDate         *bool                                `json:"groups_opt_use_group_start_as_enrollment_date,omitempty"`
	OptAssignNextAvailablePersonCode         *bool                                `json:"groups_opt_assign_next_available_person_code,omitempty"`
	DynamicEnrollmentId                      *int64                               `json:"groups_dynamic_enrollment_id,omitempty"`
	CarrierCode                              *string                              `json:"groups_carrier_carrier_code,omitempty"`
	CarrierName                              *string                              `json:"groups_carrier_carrier_name,omitempty"`
	GroupsPatientRequirements                []*GroupsPatientRequirements         `json:"groups_groups_patient_requirements,omitempty"`
	GroupsDynamicEnrollmentPrefixList        []*GroupsDynamicEnrollmentPrefixList `json:"groups_groups_dynamic_enrollment_prefix_list,omitempty"`
	GroupUserDefinedField                    []*GroupUserDefinedField             `json:"groups_group_user_defined_field,omitempty"`
	Username                                 *string                              `json:"groups_mod_user_user_name,omitempty"`
	DynamicEnrollmentName                    *string                              `json:"groups_dynamic_enrollment_name,omitempty"`
	DynamicEnrollmentDescription             *string                              `json:"groups_dynamic_enrollment_description,omitempty"`
}

type GroupsPatientRequirements struct {
	GroupId          int64   `json:"groups_patient_requirements_group_id,omitempty"`
	PatientFieldName string  `json:"groups_patient_requirements_patient_field_name,omitempty"`
	RequireValue     *bool   `json:"groups_patient_requirements_require_value,omitempty"`
	DefaultValue     *string `json:"groups_patient_requirements_default_value,omitempty"`
}

type GroupsDynamicEnrollmentPrefixList struct {
	GroupId int64  `json:"groups_dynamic_enrollment_prefix_list_group_id,omitempty"`
	Prefix  string `json:"groups_dynamic_enrollment_prefix_list_prefix,omitempty"`
}

type GroupUserDefinedField struct {
	GroupId    int64   `json:"group_user_defined_field_group_id,omitempty"`
	FieldCode  string  `json:"group_user_defined_field_field_code,omitempty"`
	FieldValue *string `json:"group_user_defined_field_field_value,omitempty"`
}
