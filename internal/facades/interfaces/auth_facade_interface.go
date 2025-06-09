package interfaces

import (
	"context"
	"reloop-backend/internal/dto"
)

type AuthFacadeInterface interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	GetProfile(ctx context.Context, userID uint) (*dto.UserResponse, error)
}
