package validation

import (
	"errors"
	"reloop-backend/internal/dto"
	"strings"
)

func ValidateCreateItemRequest(req dto.CreateItemRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if len(req.Name) > 255 {
		return errors.New("name must not exceed 255 characters")
	}

	if req.CategoryID == 0 {
		return errors.New("categoryId is required")
	}

	if req.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if req.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	return nil
}

func ValidateUpdateItemRequest(req dto.UpdateItemRequest) error {
	if req.Name != nil && strings.TrimSpace(*req.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if req.Name != nil && len(*req.Name) > 255 {
		return errors.New("name must not exceed 255 characters")
	}

	if req.Price != nil && *req.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if req.Quantity != nil && *req.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	return nil
}
