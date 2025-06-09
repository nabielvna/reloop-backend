package implementations

import (
	"context"
	"reloop-backend/internal/models"
	"reloop-backend/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type FraudReportRepository struct {
	db *gorm.DB
}

func NewFraudReportRepository(db *gorm.DB) interfaces.FraudReportRepositoryInterface {
	return &FraudReportRepository{db: db}
}

func (r *FraudReportRepository) Create(ctx context.Context, report *models.FraudReport) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *FraudReportRepository) GetByID(ctx context.Context, id uint) (*models.FraudReport, error) {
	var report models.FraudReport
	err := r.db.WithContext(ctx).
		Preload("Reporter").
		Preload("ReportedItem").
		Preload("ReportedItem.Seller").
		Preload("ReportedItem.Seller.User").
		First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *FraudReportRepository) GetPendingReports(ctx context.Context) ([]models.FraudReport, error) {
	var reports []models.FraudReport
	err := r.db.WithContext(ctx).
		Preload("Reporter").
		Preload("ReportedItem").
		Preload("ReportedItem.Seller").
		Preload("ReportedItem.Seller.User").
		Where("status = ?", "pending").
		Find(&reports).Error
	return reports, err
}

func (r *FraudReportRepository) GetByReporter(ctx context.Context, reporterID uint) ([]models.FraudReport, error) {
	var reports []models.FraudReport
	err := r.db.WithContext(ctx).
		Preload("ReportedItem").
		Where("reporter_id = ?", reporterID).
		Find(&reports).Error
	return reports, err
}

func (r *FraudReportRepository) GetByItem(ctx context.Context, itemID uint) ([]models.FraudReport, error) {
	var reports []models.FraudReport
	err := r.db.WithContext(ctx).
		Preload("Reporter").
		Where("reported_item_id = ?", itemID).
		Find(&reports).Error
	return reports, err
}

func (r *FraudReportRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.FraudReport{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *FraudReportRepository) GetAll(ctx context.Context) ([]models.FraudReport, error) {
	var reports []models.FraudReport
	err := r.db.WithContext(ctx).
		Preload("Reporter").
		Preload("ReportedItem").
		Preload("ReportedItem.Seller").
		Preload("ReportedItem.Seller.User").
		Find(&reports).Error
	return reports, err
}