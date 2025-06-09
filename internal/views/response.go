package views

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func WriteSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	WriteJSONResponse(w, http.StatusOK, response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, code, message, details string) {
	response := APIResponse{
		Success: false,
		Message: "Request failed",
		Error: &APIError{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
	WriteJSONResponse(w, statusCode, response)
}

func WriteCreatedResponse(w http.ResponseWriter, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	WriteJSONResponse(w, http.StatusCreated, response)
}
