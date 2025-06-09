package implementations

import (
	"context"
	"errors"
	"reloop-backend/internal/dto"
	"reloop-backend/internal/facades/interfaces"
	"reloop-backend/internal/models"
	repoInterfaces "reloop-backend/internal/repositories/interfaces"
	"reloop-backend/internal/validation"

	"gorm.io/gorm"
)

type ItemFacade struct {
	itemRepo     repoInterfaces.ItemRepositoryInterface
	categoryRepo repoInterfaces.CategoryRepositoryInterface
	sellerRepo   repoInterfaces.SellerRepositoryInterface
}

func NewItemFacade(
	itemRepo repoInterfaces.ItemRepositoryInterface,
	categoryRepo repoInterfaces.CategoryRepositoryInterface,
	sellerRepo repoInterfaces.SellerRepositoryInterface,
) interfaces.ItemFacadeInterface {
	return &ItemFacade{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
		sellerRepo:   sellerRepo,
	}
}

func (f *ItemFacade) CreateItem(ctx context.Context, sellerID uint, req dto.CreateItemRequest) (*dto.ItemResponse, error) {
	// Validate input
	if err := validation.ValidateCreateItemRequest(req); err != nil {
		return nil, err
	}

	// Verify seller exists
	seller, err := f.sellerRepo.GetByUserID(ctx, sellerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seller not found")
		}
		return nil, errors.New("failed to verify seller")
	}

	// Verify category exists
	_, err = f.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, errors.New("failed to verify category")
	}

	// Create item
	item := &models.Item{
		SellerID:    seller.UserID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Status:      "pending",
	}

	if err := f.itemRepo.Create(ctx, item); err != nil {
		return nil, errors.New("failed to create item")
	}

	// Load relationships for response
	createdItem, err := f.itemRepo.GetByID(ctx, item.ID)
	if err != nil {
		return nil, errors.New("failed to load created item")
	}

	return f.mapItemToResponse(createdItem), nil
}

func (f *ItemFacade) GetItem(ctx context.Context, itemID uint) (*dto.ItemResponse, error) {
	item, err := f.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, errors.New("failed to get item")
	}

	return f.mapItemToResponse(item), nil
}

func (f *ItemFacade) UpdateItem(ctx context.Context, itemID uint, sellerID uint, req dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	// Validate input
	if err := validation.ValidateUpdateItemRequest(req); err != nil {
		return nil, err
	}

	// Get item and verify ownership
	item, err := f.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, errors.New("failed to get item")
	}

	if item.SellerID != sellerID {
		return nil, errors.New("permission denied: not the owner of this item")
	}

	// Update fields if provided
	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Description != nil {
		item.Description = *req.Description
	}
	if req.Price != nil {
		item.Price = *req.Price
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	if req.CategoryID != nil {
		// Verify new category exists
		_, err := f.categoryRepo.GetByID(ctx, *req.CategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("category not found")
			}
			return nil, errors.New("failed to verify category")
		}
		item.CategoryID = *req.CategoryID
	}

	// Save updates
	if err := f.itemRepo.Update(ctx, item); err != nil {
		return nil, errors.New("failed to update item")
	}

	// Reload with relationships
	updatedItem, err := f.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, errors.New("failed to load updated item")
	}

	return f.mapItemToResponse(updatedItem), nil
}

func (f *ItemFacade) DeleteItem(ctx context.Context, itemID uint, sellerID uint) error {
	// Get item and verify ownership
	item, err := f.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return errors.New("failed to get item")
	}

	if item.SellerID != sellerID {
		return errors.New("permission denied: not the owner of this item")
	}

	// Delete item
	if err := f.itemRepo.Delete(ctx, itemID); err != nil {
		return errors.New("failed to delete item")
	}

	return nil
}

func (f *ItemFacade) GetItemsBySeller(ctx context.Context, sellerID uint) ([]dto.ItemResponse, error) {
	items, err := f.itemRepo.GetBySeller(ctx, sellerID)
	if err != nil {
		return nil, errors.New("failed to get items")
	}

	var response []dto.ItemResponse
	for _, item := range items {
		response = append(response, *f.mapItemToResponse(&item))
	}

	return response, nil
}

func (f *ItemFacade) BrowseItems(ctx context.Context, req dto.BrowseItemsRequest) ([]dto.ItemResponse, error) {
	filters := repoInterfaces.ItemFilters{
		CategoryID: req.CategoryID,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		Search:     req.Search,
		Status:     req.Status,
	}

	items, err := f.itemRepo.Browse(ctx, filters)
	if err != nil {
		return nil, errors.New("failed to browse items")
	}

	var response []dto.ItemResponse
	for _, item := range items {
		response = append(response, *f.mapItemToResponse(&item))
	}

	return response, nil
}

func (f *ItemFacade) UpdateItemStatus(ctx context.Context, itemID uint, status string) error {
	// Verify item exists
	_, err := f.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return errors.New("failed to get item")
	}

	// Update status
	if err := f.itemRepo.UpdateStatus(ctx, itemID, status); err != nil {
		return errors.New("failed to update item status")
	}

	return nil
}

func (f *ItemFacade) mapItemToResponse(item *models.Item) *dto.ItemResponse {
	response := &dto.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
		Status:      item.Status,
	}

	if item.Category != nil {
		response.Category = dto.CategoryResponse{
			ID:   item.Category.ID,
			Name: item.Category.Name,
		}
	}

	if item.Seller != nil {
		response.Seller = dto.SellerResponse{
			ID:           item.Seller.UserID,
			BusinessName: item.Seller.BusinessName,
		}

		if item.Seller.User != nil {
			response.Seller.User = dto.UserResponse{
				ID:       item.Seller.User.ID,
				Email:    item.Seller.User.Email,
				UserName: item.Seller.User.UserName,
				Role:     item.Seller.User.Role,
			}
		}
	}

	return response
}
