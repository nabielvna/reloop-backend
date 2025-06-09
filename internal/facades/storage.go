package facades

import (
	"reloop-backend/internal/facades/implementations"
	"reloop-backend/internal/facades/interfaces"
	repoInterfaces "reloop-backend/internal/repositories/interfaces"
)

type FacadeStorage struct {
	Auth interfaces.AuthFacadeInterface
	Item interfaces.ItemFacadeInterface
}

func NewFacadeStorage(
	userRepo repoInterfaces.UserRepositoryInterface,
	itemRepo repoInterfaces.ItemRepositoryInterface,
	categoryRepo repoInterfaces.CategoryRepositoryInterface,
	sellerRepo repoInterfaces.SellerRepositoryInterface,
	jwtSecret string,
) FacadeStorage {
	return FacadeStorage{
		Auth: implementations.NewAuthFacade(userRepo, jwtSecret),
		Item: implementations.NewItemFacade(itemRepo, categoryRepo, sellerRepo),
	}
}
