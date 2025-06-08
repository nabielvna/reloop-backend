package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Definisikan key untuk context
type contextKey string

const userContextKey = contextKey("user")

func (app *application) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Header Authorization tidak ada", http.StatusUnauthorized)
			return
		}

		// Header harus berformat "Bearer <token>"
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Format Authorization header tidak valid", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing adalah HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(app.config.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token tidak valid", http.StatusUnauthorized)
			return
		}
		
		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token tidak valid", http.StatusUnauthorized)
			return
		}

		// Simpan user ID ke dalam context request
		userID := uint(claims["sub"].(float64))
		ctx := context.WithValue(r.Context(), userContextKey, userID)

		// Lanjutkan ke handler selanjutnya dengan context yang baru
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}