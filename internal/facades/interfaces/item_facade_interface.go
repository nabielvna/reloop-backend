package interfaces

import (
	"context"
	"reloop-backend/internal/dto"
)

type ItemFacadeInterface interface {
	CreateItem(ctx context.Context, sellerID uint, req dto.CreateItemRequest) (*dto.ItemResponse, error)
	GetItem(ctx context.Context, itemID uint) (*dto.ItemResponse, error)
	UpdateItem(ctx context.Context, itemID uint, sellerID uint, req dto.UpdateItemRequest) (*dto.ItemResponse, error)
	DeleteItem(ctx context.Context, itemID uint, sellerID uint) error
	GetItemsBySeller(ctx context.Context, sellerID uint) ([]dto.ItemResponse, error)
	BrowseItems(ctx context.Context, req dto.BrowseItemsRequest) ([]dto.ItemResponse, error)
	UpdateItemStatus(ctx context.Context, itemID uint, status string) error
}
