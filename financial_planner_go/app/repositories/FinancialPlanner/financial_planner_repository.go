package FinancialPlanner

import (
	dao "FinancialPlannerGo/app/domain/dao/financialPlanner"

	"gorm.io/gorm"
)

type FinancialPlannerRepository interface {
	Save(report *dao.FinancialPlannerReport) error
	Update(report *dao.FinancialPlannerReport) error
	GetByID(id uint) (*dao.FinancialPlannerReport, error)
	GetByIDForUser(id uint, userID uint) (*dao.FinancialPlannerReport, error)
}

type financialPlannerRepository struct {
	db *gorm.DB
}

func NewFinancialPlannerRepository(db *gorm.DB) FinancialPlannerRepository {
	return &financialPlannerRepository{db: db}
}

func (r *financialPlannerRepository) Save(report *dao.FinancialPlannerReport) error {
	return r.db.Create(report).Error
}

func (r *financialPlannerRepository) Update(report *dao.FinancialPlannerReport) error {
	return r.db.Save(report).Error
}

func (r *financialPlannerRepository) GetByID(id uint) (*dao.FinancialPlannerReport, error) {
	var report dao.FinancialPlannerReport
	if err := r.db.First(&report, id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *financialPlannerRepository) GetByIDForUser(id uint, userID uint) (*dao.FinancialPlannerReport, error) {
	var report dao.FinancialPlannerReport
	query := r.db.Where("id = ?", id)
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}
