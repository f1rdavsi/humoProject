package repository

import (
	"time"

	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{
		db: db,
	}
}

func (r *ReportRepository) GetReport(rep models.Report) ([]models.Transaction, error) {
	var tr []models.Transaction
	query := r.db.Model(&models.Transaction{})

	if rep.Type != "" {
		query = query.Where("type = ?", rep.Type)
	}

	if rep.ToType != "" {
		query = query.Where("to_type = ?", rep.ToType)
	}

	if rep.From != (time.Time{}) {
		query = query.Where("created >= ?", rep.From)
	}

	if rep.To != (time.Time{}) {
		query = query.Where("created <= ?", rep.To)
	}

	page := rep.Page
	limit := rep.Limit

	if page > 0 {
		query = query.Limit(limit).Offset((page - 1) * limit)
	}

	err := query.Find(&tr).Error
	if err != nil {
		return nil, err
	}

	return tr, nil
}
