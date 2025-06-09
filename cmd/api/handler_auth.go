package main

import (
	"encoding/json"
	"net/http"
	"reloop-backend/internal/dto"
	_ "reloop-backend/internal/models"
	"reloop-backend/internal/views"
)

// @Param user body models.User true "user data"

// @Summary Registrasi pengguna baru
// @Description Membuat akun pengguna baru dengan memberikan detail yang diperlukan.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   payload body dto.RegisterRequest true "Payload untuk registrasi pengguna"
// @Success 201 {object} views.APIResponse{data=models.User} "Pengguna berhasil diregistrasi"
// @Failure 400 {object} views.APIResponse "Request tidak valid atau validasi gagal"
// @Router /auth/register [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid request body", err.Error())
		return
	}

	user, err := app.store.Facades.Auth.Register(r.Context(), req)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Registration failed", err.Error())
		return
	}

	views.WriteCreatedResponse(w, "User registered successfully", user)
}

// @Summary Login pengguna
// @Description Mengautentikasi pengguna dengan email dan password, dan mengembalikan token JWT jika berhasil.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   credentials body dto.LoginRequest true "Kredensial untuk login"
// @Success 200 {object} views.APIResponse{data=dto.LoginResponse} "Login berhasil"
// @Failure 400 {object} views.APIResponse "Request tidak valid"
// @Failure 401 {object} views.APIResponse "Kredensial tidak valid"
// @Router /auth/login [post]
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.WriteErrorResponse(w, http.StatusBadRequest, views.ErrCodeValidationFailed, "Invalid request body", err.Error())
		return
	}

	response, err := app.store.Facades.Auth.Login(r.Context(), req)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusUnauthorized, views.ErrCodeInvalidCredentials, "Login failed", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Login successful", response)
}

// @Summary Mendapatkan profil pengguna
// @Description Mengambil detail profil untuk pengguna yang sedang login (memerlukan token otentikasi).
// @Tags Users
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} views.APIResponse{data=models.User} "Profil berhasil diambil"
// @Failure 401 {object} views.APIResponse "Unauthorized (token tidak valid atau tidak ada)"
// @Failure 404 {object} views.APIResponse "Pengguna tidak ditemukan"
// @Failure 500 {object} views.APIResponse "Gagal mengambil pengguna dari konteks"
// @Router /me [get]
func (app *application) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userContextKey).(uint)
	if !ok {
		views.WriteErrorResponse(w, http.StatusInternalServerError, views.ErrCodeInternalError, "Failed to get user from context", "")
		return
	}

	user, err := app.store.Facades.Auth.GetProfile(r.Context(), userID)
	if err != nil {
		views.WriteErrorResponse(w, http.StatusNotFound, views.ErrCodeUserNotFound, "User not found", err.Error())
		return
	}

	views.WriteSuccessResponse(w, "Profile retrieved successfully", user)
}
