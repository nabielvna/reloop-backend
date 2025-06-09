package interfaces

import (
	"context"
	"reloop-backend/internal/models"
)

type FraudReportRepositoryInterface interface {
	Create(ctx context.Context, report *models.FraudReport) error
	GetByID(ctx context.Context, id uint) (*models.FraudReport, error)
	GetPendingReports(ctx context.Context) ([]models.FraudReport, error)
	GetByReporter(ctx context.Context, reporterID uint) ([]models.FraudReport, error)
	GetByItem(ctx context.Context, itemID uint) ([]models.FraudReport, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
	GetAll(ctx context.Context) ([]models.FraudReport, error)
}
