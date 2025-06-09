package main

import (
	"encoding/json"
	"net/http"
	"reloop-backend/internal/dto"
	"reloop-backend/internal/views"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userContextKey).(uint)
	if !ok {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get user from context", "")
		return
	}

	var req dto.CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid request body", err.Error())
		return
	}

	item, err := app.store.Facades.Item.CreateItem(r.Context(), userID, req)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Failed to create item", err.Error())
		return
	}

	views.WriteCreatedResponse(w, "Item created successfully", item)
}

func (app *application) getItemHandler(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid item ID", "")
		return
	}

	item, err := app.store.Facades.Item.GetItem(r.Context(), uint(itemID))
	if err != nil {
		views.WriteErrorResponse(w, http.StatusNotFound, views.ErrCodeItemNotFound, "Item not found", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Item retrieved successfully", item)
}

func (app *application) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userContextKey).(uint)
	if !ok {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get user from context", "")
		return
	}

	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid item ID", "")
		return
	}

	var req dto.UpdateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid request body", err.Error())
		return
	}

	item, err := app.store.Facades.Item.UpdateItem(r.Context(), uint(itemID), userID, req)
	if err != nil {
		if err.Error() == "permission denied: not the owner of this item" {
			views.WriteErrorResponse(w, http.StatusForbidden, views.ErrCodePermissionDenied, "Permission denied", err.Error())
			return
		}
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Failed to update item", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Item updated successfully", item)
}

func (app *application) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userContextKey).(uint)
	if !ok {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get user from context", "")
		return
	}

	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid item ID", "")
		return
	}

	err = app.store.Facades.Item.DeleteItem(r.Context(), uint(itemID), userID)
	if err != nil {
		if err.Error() == "permission denied: not the owner of this item" {
			views.WriteErrorResponse(w, http.StatusForbidden, views.ErrCodePermissionDenied, "Permission denied", err.Error())
			return
		}
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Failed to delete item", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Item deleted successfully", nil)
}

func (app *application) getMyItemsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userContextKey).(uint)
	if !ok {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get user from context", "")
		return
	}

	items, err := app.store.Facades.Item.GetItemsBySeller(r.Context(), userID)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get items", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Items retrieved successfully", items)
}

func (app *application) browseItemsHandler(w http.ResponseWriter, r *http.Request) {
	req := dto.BrowseItemsRequest{}

	// Parse query parameters
	if categoryIDStr := r.URL.Query().Get("categoryId"); categoryIDStr != "" {
		if categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil {
			categoryIDUint := uint(categoryID)
			req.CategoryID = &categoryIDUint
		}
	}

	if minPriceStr := r.URL.Query().Get("minPrice"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			req.MinPrice = &minPrice
		}
	}

	if maxPriceStr := r.URL.Query().Get("maxPrice"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			req.MaxPrice = &maxPrice
		}
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if status := r.URL.Query().Get("status"); status != "" {
		req.Status = &status
	}

	items, err := app.store.Facades.Item.BrowseItems(r.Context(), req)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to browse items", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Items retrieved successfully", items)
}

func (app *application) updateItemStatusHandler(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid item ID", "")
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid request body", err.Error())
		return
	}

	err = app.store.Facades.Item.UpdateItemStatus(r.Context(), uint(itemID), req.Status)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Failed to update item status", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Item status updated successfully", nil)
}
