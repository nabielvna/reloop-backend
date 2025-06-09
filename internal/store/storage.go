package store

import (
	"reloop-backend/internal/facades"
	"reloop-backend/internal/repositories/implementations"
	"reloop-backend/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type Storage struct {
	Users        interfaces.UserRepositoryInterface
	Items        interfaces.ItemRepositoryInterface
	Categories   interfaces.CategoryRepositoryInterface
	Sellers      interfaces.SellerRepositoryInterface
	FraudReports interfaces.FraudReportRepositoryInterface

	Facades facades.FacadeStorage
}

func NewStorage(db *gorm.DB, jwtSecret string) *Storage {
	userRepo := implementations.NewUserRepository(db)
	itemRepo := implementations.NewItemRepository(db)
	categoryRepo := implementations.NewCategoryRepository(db)
	sellerRepo := implementations.NewSellerRepository(db)
	fraudRepo := implementations.NewFraudReportRepository(db)

	facadeStorage := facades.NewFacadeStorage(
		userRepo,
		itemRepo,
		categoryRepo,
		sellerRepo,
		jwtSecret,
	)

	return &Storage{
		Users:        userRepo,
		Items:        itemRepo,
		Categories:   categoryRepo,
		Sellers:      sellerRepo,
		FraudReports: fraudRepo,
		Facades:      facadeStorage,
	}
}