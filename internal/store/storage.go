package store

import (
	"reloop-backend/internal/repositories/implementations"
	"reloop-backend/internal/repositories/interfaces"
)

type Storage struct {
	Users        interfaces.UserRepositoryInterface
	Items        interfaces.ItemRepositoryInterface
	Categories   interfaces.CategoryRepositoryInterface
	Sellers      interfaces.SellerRepositoryInterface
	FraudReports interfaces.FraudReportRepositoryInterface
}

func NewStorage() Storage {
	return Storage{
		Users:        implementations.NewUserRepository(),
		Items:        implementations.NewItemRepository(),
		Categories:   implementations.NewCategoryRepository(),
		Sellers:      implementations.NewSellerRepository(),
		FraudReports: implementations.NewFraudReportRepository(),
	}
}
