package main

import (
	"encoding/json"
	"net/http"
	"reloop-backend/internal/dto"
	"reloop-backend/internal/views"
)

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
