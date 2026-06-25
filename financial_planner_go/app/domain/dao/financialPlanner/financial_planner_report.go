package financialPlanner

import (
	"time"

	"gorm.io/datatypes"
)

type FinancialPlannerReport struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `gorm:"index" json:"user_id"`
	FullName        string         `gorm:"size:255;not null" json:"full_name"`
	Contact         string         `gorm:"size:20;not null" json:"contact"`
	Email           string         `gorm:"size:255;not null;index" json:"email"`
	DOB             string         `gorm:"size:20;not null" json:"dob"`
	City            string         `gorm:"size:100;not null" json:"city"`
	ScoreValue      int            `json:"score_value"`
	ScoreLabel      string         `gorm:"size:100" json:"score_label"`
	Status          string         `gorm:"size:50;not null;default:report_requested;index" json:"status"`
	ReportPDFURL    string         `gorm:"type:text" json:"report_pdf_url"`
	ReportPDFPath   string         `gorm:"type:text" json:"report_pdf_path"`
	RequestPayload  datatypes.JSON `gorm:"type:jsonb" json:"-"`
	ResponsePayload datatypes.JSON `gorm:"type:jsonb" json:"-"`
	ErrorMessage    string         `gorm:"type:text" json:"-"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (FinancialPlannerReport) TableName() string {
	return "financial_planner_reports"
}
