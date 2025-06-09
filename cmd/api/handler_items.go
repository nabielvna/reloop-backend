package main

import (
	"encoding/json"
	"net/http"
	"reloop-backend/internal/dto"
	_ "reloop-backend/internal/models"
	"reloop-backend/internal/views"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary Membuat item baru
// @Description Membuat produk atau item baru. Hanya dapat diakses oleh pengguna yang sudah login (penjual).
// @Tags Items
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   payload body dto.CreateItemRequest true "Data untuk membuat item baru"
// @Success 201 {object} views.APIResponse{data=models.Item} "Item berhasil dibuat"
// @Failure 400 {object} views.APIResponse "Request tidak valid atau validasi gagal"
// @Failure 401 {object} views.APIResponse "Unauthorized"
// @Failure 500 {object} views.APIResponse "Gagal mengambil pengguna dari konteks"
// @Router /items [post]
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

// @Summary Mendapatkan detail item
// @Description Endpoint publik untuk mengambil detail spesifik dari sebuah item berdasarkan ID-nya.
// @Tags Items
// @Produce  json
// @Param   itemID path int true "ID Item"
// @Success 200 {object} views.APIResponse{data=models.Item} "Item berhasil diambil"
// @Failure 400 {object} views.APIResponse "ID item tidak valid"
// @Failure 404 {object} views.APIResponse "Item tidak ditemukan"
// @Router /items/{itemID} [get]
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

// @Summary Memperbarui item
// @Description Memperbarui detail item. Hanya bisa dilakukan oleh pemilik item.
// @Tags Items
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   itemID path int true "ID Item"
// @Param   payload body dto.UpdateItemRequest true "Data item yang akan diperbarui"
// @Success 200 {object} views.APIResponse{data=models.Item} "Item berhasil diperbarui"
// @Failure 400 {object} views.APIResponse "Request atau ID item tidak valid"
// @Failure 401 {object} views.APIResponse "Unauthorized"
// @Failure 403 {object} views.APIResponse "Akses ditolak (bukan pemilik item)"
// @Failure 404 {object} views.APIResponse "Item tidak ditemukan"
// @Router /items/{itemID} [put]
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

// @Summary Menghapus item
// @Description Menghapus sebuah item. Hanya bisa dilakukan oleh pemilik item.
// @Tags Items
// @Produce  json
// @Security BearerAuth
// @Param   itemID path int true "ID Item"
// @Success 200 {object} views.APIResponse "Item berhasil dihapus"
// @Failure 400 {object} views.APIResponse "ID item tidak valid"
// @Failure 401 {object} views.APIResponse "Unauthorized"
// @Failure 403 {object} views.APIResponse "Akses ditolak (bukan pemilik item)"
// @Failure 404 {object} views.APIResponse "Item tidak ditemukan"
// @Router /items/{itemID} [delete]
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

// @Summary Mendapatkan item milik penjual
// @Description Mengambil daftar semua item yang dimiliki oleh penjual yang sedang login.
// @Tags Items
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} views.APIResponse{data=[]models.Item} "Item berhasil diambil"
// @Failure 401 {object} views.APIResponse "Unauthorized"
// @Failure 500 {object} views.APIResponse "Gagal mengambil item"
// @Router /my-items [get]
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

// @Summary Mencari dan memfilter item
// @Description Endpoint publik untuk mencari dan memfilter item berdasarkan kriteria.
// @Tags Items
// @Produce  json
// @Param   categoryId query int false "Filter berdasarkan ID kategori"
// @Param   minPrice query number false "Filter harga minimum"
// @Param   maxPrice query number false "Filter harga maksimum"
// @Param   search query string false "Kata kunci pencarian pada nama atau deskripsi"
// @Param   status query string false "Filter berdasarkan status (e.g., 'approved')"
// @Success 200 {object} views.APIResponse{data=[]models.Item} "Item berhasil diambil"
// @Failure 500 {object} views.APIResponse "Gagal mencari item"
// @Router /items [get]
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

// @Summary Memperbarui status item (Admin)
// @Description Mengubah status sebuah item. Hanya bisa diakses oleh Admin.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   itemID path int true "ID Item"
// @Param   payload body dto.UpdateItemStatusRequest true "Status baru untuk item"
// @Success 200 {object} views.APIResponse "Status item berhasil diperbarui"
// @Failure 400 {object} views.APIResponse "Request atau ID item tidak valid"
// @Failure 401 {object} views.APIResponse "Unauthorized"
// @Failure 403 {object} views.APIResponse "Akses ditolak (bukan admin)"
// @Router /admin/items/{itemID}/status [patch]
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
