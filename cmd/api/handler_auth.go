package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"reloop-backend/internal/models"
	"reloop-backend/internal/store"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input RegisterUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Request Body tidak valid", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		http.Error(w, "Gagal memproses password", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		UserName:     input.UserName,
	}

	if err := app.store.Users.Create(r.Context(), user); err != nil {
		http.Error(w, "Gagal membuat user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User berhasil dibuat"})
}

func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Request Body tidak valid", http.StatusBadRequest)
		return
	}

	user, err := app.store.Users.(*store.UsersStore).GetByEmail(r.Context(), input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Email atau password salah", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Terjadi kesalahan server", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		http.Error(w, "Email atau password salah", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"rol": user.Role,    
		"iat": time.Now().Unix(), 
		"exp": time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(app.config.jwtSecret))
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}