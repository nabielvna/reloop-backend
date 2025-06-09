package store

import (
	"reloop-backend/internal/facades"
	"reloop-backend/internal/repositories/implementations"
	"reloop-backend/internal/repositories/interfaces"
)

type Storage struct {
	// Repositories
	Users        interfaces.UserRepositoryInterface
	Items        interfaces.ItemRepositoryInterface
	Categories   interfaces.CategoryRepositoryInterface
	Sellers      interfaces.SellerRepositoryInterface
	FraudReports interfaces.FraudReportRepositoryInterface

	// Facades
	Facades facades.FacadeStorage
}

func NewStorage(jwtSecret string) Storage {
	// Create repositories
	userRepo := implementations.NewUserRepository()
	itemRepo := implementations.NewItemRepository()
	categoryRepo := implementations.NewCategoryRepository()
	sellerRepo := implementations.NewSellerRepository()
	fraudRepo := implementations.NewFraudReportRepository()

	// Create facades
	facadeStorage := facades.NewFacadeStorage(
		userRepo,
		itemRepo,
		categoryRepo,
		sellerRepo,
		jwtSecret,
	)

	return Storage{
		Users:        userRepo,
		Items:        itemRepo,
		Categories:   categoryRepo,
		Sellers:      sellerRepo,
		FraudReports: fraudRepo,
		Facades:      facadeStorage,
	}
}
