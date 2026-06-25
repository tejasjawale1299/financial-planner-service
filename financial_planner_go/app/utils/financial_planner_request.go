package utils

type FinancialPlannerRequest struct {
	FullName                string `json:"full_name" binding:"required"`
	Contact                 string `json:"contact" binding:"required"`
	Email                   string `json:"email" binding:"required,email"`
	DOB                     string `json:"dob" binding:"required"`
	City                    string `json:"city" binding:"required"`
	MaritalStatus           string `json:"marital_status"`
	SpouseAge               string `json:"spouse_age"`
	SpouseIncome            string `json:"spouse_income"`
	ChildrenCount           string `json:"children_count"`
	Dependents              string `json:"dependents"`
	MonthlyIncome           string `json:"monthly_income"`
	MonthlyExpenses         string `json:"monthly_expenses"`
	MonthlySavings          string `json:"monthly_savings"`
	Education1TodayCost     string `json:"education_1_today_cost"`
	Marriage1TodayCost      string `json:"marriage_1_today_cost"`
	RetirementAge           string `json:"retirement_age"`
	LifeExpectancy          string `json:"life_expectancy"`
	WealthTargetToday       string `json:"wealth_target_today"`
	WealthTargetAge         string `json:"wealth_target_age"`
	Education1Enabled       string `json:"education_1_enabled"`
	Marriage1Enabled        string `json:"marriage_1_enabled"`
	RetirementEnabled       string `json:"retirement_enabled"`
	WealthEnabled           string `json:"wealth_enabled"`
	EducationPriority       string `json:"education_priority"`
	MarriagePriority        string `json:"marriage_priority"`
	RetirementPriority      string `json:"retirement_priority"`
	WealthPriority          string `json:"wealth_priority"`
	Education1TargetAge     string `json:"education_1_target_age"`
	Education1Inflation     string `json:"education_1_inflation"`
	Education1SIPReturn     string `json:"education_1_sip_return"`
	Education1LumpsumReturn string `json:"education_1_lumpsum_return"`
	Marriage1TargetAge      string `json:"marriage_1_target_age"`
	Marriage1Inflation      string `json:"marriage_1_inflation"`
	Marriage1SIPReturn      string `json:"marriage_1_sip_return"`
	Marriage1LumpsumReturn  string `json:"marriage_1_lumpsum_return"`
	RetirementInflation     string `json:"retirement_inflation"`
	RetirementPreReturn     string `json:"retirement_pre_return"`
	RetirementPostReturn    string `json:"retirement_post_return"`
	WealthInflation         string `json:"wealth_inflation"`
	WealthSIPReturn         string `json:"wealth_sip_return"`
	WealthLumpsumReturn     string `json:"wealth_lumpsum_return"`
	ExistingLifeTerm        string `json:"existing_life_term"`
	ExistingHealth          string `json:"existing_health"`
	ExistingCritical        string `json:"existing_critical"`
	ExistingPersonal        string `json:"existing_personal"`
	RiskAppetite            string `json:"risk_appetite"`
	InvestmentExperience    string `json:"investment_experience"`
}
