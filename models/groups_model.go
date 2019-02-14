package models

import (
	"time"
)

type Groups struct {
	Id                                       int64      `json:"groups_id,omitempty"`
	CarrierId                                int64      `json:"groups_carrier_id,omitempty"`
	GroupNum                                 string     `json:"groups_group_num,omitempty"`
	GroupName                                string     `json:"groups_group_name,omitempty"`
	Address1                                 *string    `json:"groups_address1,omitempty"`
	Address2                                 *string    `json:"groups_address2,omitempty"`
	City                                     *string    `json:"groups_city,omitempty"`
	State                                    *string    `json:"groups_state,omitempty"`
	Zip                                      *string    `json:"groups_zip,omitempty"`
	Phone                                    *string    `json:"groups_phone,omitempty"`
	PhoneExt                                 *string    `json:"groups_phone_ext,omitempty"`
	Fax                                      *string    `json:"groups_fax,omitempty"`
	BenefitYearStart                         *time.Time `json:"groups_benefit_year_start,omitempty"`
	MemberDetermination                      string     `json:"groups_member_determination,omitempty"`
	PersoncodeDetermination                  *string    `json:"groups_personcode_determination,omitempty"`
	DobDetermination                         *string    `json:"groups_dob_determination,omitempty"`
	DobDays                                  *int32     `json:"groups_dob_days,omitempty"`
	CardType                                 *string    `json:"groups_card_type,omitempty"`
	CardName                                 *string    `json:"groups_card_name,omitempty"`
	DynamicEnrollmentDays                    *int32     `json:"groups_dynamic_enrollment_days,omitempty"`
	DynamicEnrollmentIdLogic                 *string    `json:"groups_dynamic_enrollment_id_logic,omitempty"`
	PrescriberLogic                          *string    `json:"groups_prescriber_logic,omitempty"`
	MaxRxPeriod                              *string    `json:"groups_max_rx_period,omitempty"`
	MaxRx                                    *int32     `json:"groups_max_rx,omitempty"`
	TaxLogic                                 *string    `json:"groups_tax_logic,omitempty"`
	ChildCutoffAge                           *int32     `json:"groups_child_cutoff_age,omitempty"`
	ChildCutoffAgeType                       *string    `json:"groups_child_cutoff_age_type,omitempty"`
	StudentCutoffAge                         *int32     `json:"groups_student_cutoff_age,omitempty"`
	StudentCutoffAgeType                     *string    `json:"groups_student_cutoff_age_type,omitempty"`
	OtherCutoffAge                           *int32     `json:"groups_other_cutoff_age,omitempty"`
	OtherCutoffAgeType                       *string    `json:"groups_other_cutoff_age_type,omitempty"`
	RefillTooSoonGpiDigits                   int32      `json:"groups_refill_too_soon_gpi_digits,omitempty"`
	ContactName                              *string    `json:"groups_contact_name,omitempty"`
	ContactEmailAddress                      *string    `json:"groups_contact_email_address,omitempty"`
	OptSecondaryPayer                        *bool      `json:"groups_opt_secondary_payer,omitempty"`
	OptIgnoreHistory                         *bool      `json:"groups_opt_ignore_history,omitempty"`
	OptTestOnly                              *bool      `json:"groups_opt_test_only,omitempty"`
	OptMaxDollarInclCopayAmt                 *bool      `json:"groups_opt_max_dollar_incl_copay_amt,omitempty"`
	StartDate                                time.Time  `json:"groups_start_date,omitempty"`
	EndDate                                  time.Time  `json:"groups_end_date,omitempty"`
	ModStartDate                             time.Time  `json:"groups_mod_start_date,omitempty"`
	ModUser                                  int64      `json:"groups_mod_user,omitempty"`
	Opt001PatientsOnly                       *bool      `json:"groups_opt_001_patients_only,omitempty"`
	InjuryDateOffset                         *int32     `json:"groups_injury_date_offset,omitempty"`
	TermPatientIfNoClaimsOffset              *int32     `json:"groups_term_patient_if_no_claims_offset,omitempty"`
	OptActivatePatientProgramCodes           *bool      `json:"groups_opt_activate_patient_program_codes,omitempty"`
	OptDifferentPrescriptionNumberIgnoresRts *bool      `json:"groups_opt_different_prescription_number_ignores_rts,omitempty"`
	GroupType                                string     `json:"groups_group_type,omitempty"`
	OptUseGroupStartAsEnrollmentDate         *bool      `json:"groups_opt_use_group_start_as_enrollment_date,omitempty"`
	OptAssignNextAvailablePersonCode         *bool      `json:"groups_opt_assign_next_available_person_code,omitempty"`
	DynamicEnrollmentId                      *int64     `json:"groups_dynamic_enrollment_id,omitempty"`
	CarrierCode                              *string    `json:"groups_carrier_carrier_code,omitempty"`
	CarrierName                              *string    `json:"groups_carrier_carrier_name,omitempty"`
	// List
	GroupsPatientRequirements         []*GroupsPatientRequirements         `json:"groups_groups_patient_requirements,omitempty"`
	GroupsDynamicEnrollmentPrefixList []*GroupsDynamicEnrollmentPrefixList `json:"groups_groups_dynamic_enrollment_prefix_list,omitempty"`
	GroupUserDefinedField             []*GroupUserDefinedField             `json:"groups_group_user_defined_field,omitempty"`
	// ModList
	GroupsLocation                           []*GroupsLocation                           `json:"groups_groups_location,omitempty"`
	GroupsPlanList                           []*GroupsPlanList                           `json:"groups_groups_plan_list,omitempty"`
	GroupsPriorAuth                          []*GroupsPriorAuth                          `json:"groups_groups_prior_auth,omitempty"`
	GroupsDedCapMgmt                         []*GroupsDedCapMgmt                         `json:"groups_groups_ded_cap_mgmt,omitempty"`
	GroupsClaimAdminList                     []*GroupsClaimAdminList                     `json:"groups_groups_claim_admin_list,omitempty"`
	GroupsClaimAdminFeeList                  []*GroupsClaimAdminFeeList                  `json:"groups_groups_claim_admin_fee_list,omitempty"`
	GroupsSubPlan                            []*GroupsSubPlan                            `json:"groups_groups_sub_plan,omitempty"`
	GroupsDynamicEnrollmentPharmacyDaysHours []*GroupsDynamicEnrollmentPharmacyDaysHours `json:"groups_groups_dynamic_enrollment_pharmacy_days_hours,omitempty"`
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

type GroupsClaimAdminList struct {
	ItemId              int64     `json:"groups_claim_admin_list_item_id,omitempty"`
	ParentId            int64     `json:"groups_claim_admin_list_parent_id,omitempty"`
	ClaimType           string    `json:"groups_claim_admin_list_claim_type,omitempty"`
	ClaimDaysOld        int32     `json:"groups_claim_admin_list_claim_days_old,omitempty"`
	ReversalDaysOld     int32     `json:"groups_claim_admin_list_reversal_days_old,omitempty"`
	MaxDollarsPerScript float64   `json:"groups_claim_admin_list_max_dollars_per_script,omitempty"`
	MaxDollarsMessage   *string   `json:"groups_claim_admin_list_max_dollars_message,omitempty"`
	UpdatedAt           time.Time `json:"groups_claim_admin_list_updated_at,omitempty"`
	UpdatedBy           int64     `json:"groups_claim_admin_list_updated_by,omitempty"`
}

type GroupsDynamicEnrollmentPharmacyDaysHours struct {
	ItemId    int64     `json:"groups_dynamic_enrollment_pharmacy_days_hours_item_id,omitempty"`
	ParentId  int64     `json:"groups_dynamic_enrollment_pharmacy_days_hours_parent_id,omitempty"`
	Nabp      *string   `json:"groups_dynamic_enrollment_pharmacy_days_hours_nabp,omitempty"`
	Chain     *string   `json:"groups_dynamic_enrollment_pharmacy_days_hours_chain,omitempty"`
	DayOfWeek string    `json:"groups_dynamic_enrollment_pharmacy_days_hours_day_of_week,omitempty"`
	StartTime string    `json:"groups_dynamic_enrollment_pharmacy_days_hours_start_time,omitempty"`
	EndTime   string    `json:"groups_dynamic_enrollment_pharmacy_days_hours_end_time,omitempty"`
	StartDate time.Time `json:"groups_dynamic_enrollment_pharmacy_days_hours_start_date,omitempty"`
	EndDate   time.Time `json:"groups_dynamic_enrollment_pharmacy_days_hours_end_date,omitempty"`
	UpdatedAt time.Time `json:"groups_dynamic_enrollment_pharmacy_days_hours_updated_at,omitempty"`
	UpdatedBy int64     `json:"groups_dynamic_enrollment_pharmacy_days_hours_updated_by,omitempty"`
}

type GroupsSubPlan struct {
	ItemId                     int64     `json:"groups_sub_plan_item_id,omitempty"`
	ParentId                   int64     `json:"groups_sub_plan_parent_id,omitempty"`
	SubPlanName                string    `json:"groups_sub_plan_sub_plan_name,omitempty"`
	Description                *string   `json:"groups_sub_plan_description,omitempty"`
	StartDate                  time.Time `json:"groups_sub_plan_start_date,omitempty"`
	EndDate                    time.Time `json:"groups_sub_plan_end_date,omitempty"`
	PlanReplacement            *int64    `json:"groups_sub_plan_plan_replacement,omitempty"`
	PlanRequiredForReplacement *int64    `json:"groups_sub_plan_plan_required_for_replacement,omitempty"`
	BenefitListId              *int64    `json:"groups_sub_plan_benefit_list_id,omitempty"`
	UpdatedAt                  time.Time `json:"groups_sub_plan_updated_at,omitempty"`
	UpdatedBy                  int64     `json:"groups_sub_plan_updated_by,omitempty"`
}

type GroupsLocation struct {
	ItemId       int64     `json:"groups_location_item_id,omitempty"`
	ParentId     int64     `json:"groups_location_parent_id,omitempty"`
	LocationCode string    `json:"groups_location_location_code,omitempty"`
	Description  *string   `json:"groups_location_description,omitempty"`
	StartDate    time.Time `json:"groups_location_start_date,omitempty"`
	EndDate      time.Time `json:"groups_location_end_date,omitempty"`
	UpdatedAt    time.Time `json:"groups_location_updated_at,omitempty"`
	UpdatedBy    int64     `json:"groups_location_updated_by,omitempty"`
}

type GroupsClaimAdminFeeList struct {
	ItemId      int64     `json:"groups_claim_admin_fee_list_item_id,omitempty"`
	ParentId    int64     `json:"groups_claim_admin_fee_list_parent_id,omitempty"`
	ClaimType   string    `json:"groups_claim_admin_fee_list_claim_type,omitempty"`
	FeeType     string    `json:"groups_claim_admin_fee_list_fee_type,omitempty"`
	ClaimFee    *float64  `json:"groups_claim_admin_fee_list_claim_fee,omitempty"`
	ReversalFee *float64  `json:"groups_claim_admin_fee_list_reversal_fee,omitempty"`
	StartDate   time.Time `json:"groups_claim_admin_fee_list_start_date,omitempty"`
	EndDate     time.Time `json:"groups_claim_admin_fee_list_end_date,omitempty"`
	UpdatedAt   time.Time `json:"groups_claim_admin_fee_list_updated_at,omitempty"`
	UpdatedBy   int64     `json:"groups_claim_admin_fee_list_updated_by,omitempty"`
}

type GroupsDedCapMgmt struct {
	ItemId                         int64     `json:"groups_ded_cap_mgmt_item_id,omitempty"`
	ParentId                       int64     `json:"groups_ded_cap_mgmt_parent_id,omitempty"`
	Tier                           string    `json:"groups_ded_cap_mgmt_tier,omitempty"`
	PeriodType                     string    `json:"groups_ded_cap_mgmt_period_type,omitempty"`
	PeriodX                        *int32    `json:"groups_ded_cap_mgmt_period_x,omitempty"`
	DedSatisfaction                *string   `json:"groups_ded_cap_mgmt_ded_satisfaction,omitempty"`
	CapSatisfaction                *string   `json:"groups_ded_cap_mgmt_cap_satisfaction,omitempty"`
	CapType                        *string   `json:"groups_ded_cap_mgmt_cap_type,omitempty"`
	AfterCapCopayModifier          *int64    `json:"groups_ded_cap_mgmt_after_cap_copay_modifier,omitempty"`
	MailorderAfterCapCopayModifier *int64    `json:"groups_ded_cap_mgmt_mailorder_after_cap_copay_modifier,omitempty"`
	IndividualAmtDed               *float64  `json:"groups_ded_cap_mgmt_individual_amt_ded,omitempty"`
	IndividualAmtCap               *float64  `json:"groups_ded_cap_mgmt_individual_amt_cap,omitempty"`
	FamilyAmtDed                   *float64  `json:"groups_ded_cap_mgmt_family_amt_ded,omitempty"`
	FamilyAmtCap                   *float64  `json:"groups_ded_cap_mgmt_family_amt_cap,omitempty"`
	OptGenericsNaToDeductible      *bool     `json:"groups_ded_cap_mgmt_opt_generics_na_to_deductible,omitempty"`
	OptGenericsNaToBenCap          *bool     `json:"groups_ded_cap_mgmt_opt_generics_na_to_ben_cap,omitempty"`
	OptDeductibleNaOopCap          *bool     `json:"groups_ded_cap_mgmt_opt_deductible_na_oop_cap,omitempty"`
	StartDate                      time.Time `json:"groups_ded_cap_mgmt_start_date,omitempty"`
	EndDate                        time.Time `json:"groups_ded_cap_mgmt_end_date,omitempty"`
	UpdatedAt                      time.Time `json:"groups_ded_cap_mgmt_updated_at,omitempty"`
	UpdatedBy                      int64     `json:"groups_ded_cap_mgmt_updated_by,omitempty"`
}

type GroupsPlanList struct {
	ItemId    int64     `json:"groups_plan_list_item_id,omitempty"`
	ParentId  int64     `json:"groups_plan_list_parent_id,omitempty"`
	PlanId    int64     `json:"groups_plan_list_plan_id,omitempty"`
	StartDate time.Time `json:"groups_plan_list_start_date,omitempty"`
	EndDate   time.Time `json:"groups_plan_list_end_date,omitempty"`
	UpdatedAt time.Time `json:"groups_plan_list_updated_at,omitempty"`
	UpdatedBy int64     `json:"groups_plan_list_updated_by,omitempty"`
}
type GroupsPriorAuth struct {
	ItemId                   int64     `json:"groups_prior_auth_item_id,omitempty"`
	ParentId                 int64     `json:"groups_prior_auth_parent_id,omitempty"`
	PaNumber                 *int64    `json:"groups_prior_auth_pa_number,omitempty"`
	StartDate                time.Time `json:"groups_prior_auth_start_date,omitempty"`
	EndDate                  time.Time `json:"groups_prior_auth_end_date,omitempty"`
	PriorAuthId              int64     `json:"groups_prior_auth_prior_auth_id,omitempty"`
	SubmissionClarification  *string   `json:"groups_prior_auth_submission_clarification,omitempty"`
	LevelOfService           *string   `json:"groups_prior_auth_level_of_service,omitempty"`
	UpdatedAt                time.Time `json:"groups_prior_auth_updated_at,omitempty"`
	UpdatedBy                int64     `json:"groups_prior_auth_updated_by,omitempty"`
	OptCompoundMedsApplyToPa *bool     `json:"groups_prior_auth_opt_compound_meds_apply_to_pa,omitempty"`
}
