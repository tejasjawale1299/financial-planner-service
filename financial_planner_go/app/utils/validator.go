package utils

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var phoneRegex = regexp.MustCompile(`^[6-9][0-9]{9}$`)

func NormalizeFinancialPlannerRequest(req *FinancialPlannerRequest) {
	req.FullName = strings.TrimSpace(req.FullName)
	req.Contact = strings.TrimSpace(req.Contact)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.DOB = strings.TrimSpace(req.DOB)
	req.City = strings.TrimSpace(req.City)
	req.MaritalStatus = strings.ToLower(strings.TrimSpace(req.MaritalStatus))
	req.RiskAppetite = strings.ToLower(strings.TrimSpace(req.RiskAppetite))
	req.InvestmentExperience = strings.TrimSpace(req.InvestmentExperience)
}

func ValidateFinancialPlannerRequest(req FinancialPlannerRequest) error {
	if len(req.FullName) < 3 || len(req.FullName) > 100 {
		return errors.New("full_name must be between 3 and 100 characters")
	}

	if !phoneRegex.MatchString(req.Contact) {
		return errors.New("contact must be valid 10 digit Indian mobile number")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("email must be valid")
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return errors.New("dob must be in YYYY-MM-DD format")
	}

	currentYear := time.Now().Year()
	if dob.Year() < 1900 || dob.Year() > currentYear || dob.After(time.Now()) {
		return errors.New("dob year must be between 1900 and current year and cannot be future date")
	}

	if req.City == "" || len(req.City) > 100 {
		return errors.New("city is required and cannot be more than 100 characters")
	}

	if err := validateNonNegativeNumbers(req); err != nil {
		return err
	}

	income, _ := atoiDefault(req.MonthlyIncome)
	expenses, _ := atoiDefault(req.MonthlyExpenses)
	savings, _ := atoiDefault(req.MonthlySavings)

	if income > 0 && expenses > income {
		return errors.New("monthly expenses cannot be greater than monthly income")
	}
	if income > 0 && savings > income {
		return errors.New("monthly savings cannot be greater than monthly income")
	}

	return nil
}

func validateNonNegativeNumbers(req FinancialPlannerRequest) error {
	fields := map[string]string{
		"spouse_income":              req.SpouseIncome,
		"children_count":             req.ChildrenCount,
		"dependents":                 req.Dependents,
		"monthly_income":             req.MonthlyIncome,
		"monthly_expenses":           req.MonthlyExpenses,
		"monthly_savings":            req.MonthlySavings,
		"education_1_today_cost":     req.Education1TodayCost,
		"marriage_1_today_cost":      req.Marriage1TodayCost,
		"retirement_age":             req.RetirementAge,
		"life_expectancy":            req.LifeExpectancy,
		"wealth_target_today":        req.WealthTargetToday,
		"wealth_target_age":          req.WealthTargetAge,
		"education_priority":         req.EducationPriority,
		"marriage_priority":          req.MarriagePriority,
		"retirement_priority":        req.RetirementPriority,
		"wealth_priority":            req.WealthPriority,
		"education_1_target_age":     req.Education1TargetAge,
		"education_1_inflation":      req.Education1Inflation,
		"education_1_sip_return":     req.Education1SIPReturn,
		"education_1_lumpsum_return": req.Education1LumpsumReturn,
		"marriage_1_target_age":      req.Marriage1TargetAge,
		"marriage_1_inflation":       req.Marriage1Inflation,
		"marriage_1_sip_return":      req.Marriage1SIPReturn,
		"marriage_1_lumpsum_return":  req.Marriage1LumpsumReturn,
		"retirement_inflation":       req.RetirementInflation,
		"retirement_pre_return":      req.RetirementPreReturn,
		"retirement_post_return":     req.RetirementPostReturn,
		"wealth_inflation":           req.WealthInflation,
		"wealth_sip_return":          req.WealthSIPReturn,
		"wealth_lumpsum_return":      req.WealthLumpsumReturn,
	}

	for field, value := range fields {
		if strings.TrimSpace(value) == "" {
			continue
		}
		number, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%s must be numeric", field)
		}
		if number < 0 {
			return fmt.Errorf("%s cannot be negative", field)
		}
	}

	return nil
}

func atoiDefault(value string) (int, error) {
	if strings.TrimSpace(value) == "" {
		return 0, nil
	}
	return strconv.Atoi(value)
}
