package implementations

import (
	"context"
	"errors"
	"reloop-backend/internal/dto"
	"reloop-backend/internal/facades/interfaces"
	"reloop-backend/internal/models"
	repoInterfaces "reloop-backend/internal/repositories/interfaces"
	"reloop-backend/internal/validation"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthFacade struct {
	userRepo  repoInterfaces.UserRepositoryInterface
	jwtSecret string
}

func NewAuthFacade(userRepo repoInterfaces.UserRepositoryInterface, jwtSecret string) interfaces.AuthFacadeInterface {
	return &AuthFacade{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (f *AuthFacade) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	// Validate input
	if err := validation.ValidateRegisterRequest(req); err != nil {
		return nil, err
	}

	// Check if user already exists
	existingUser, _ := f.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// Create user
	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		UserName:     req.UserName,
		Role:         "user",
	}

	if err := f.userRepo.Create(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Return user response
	return &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.UserName,
		Role:     user.Role,
	}, nil
}

func (f *AuthFacade) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// Validate input
	if err := validation.ValidateLoginRequest(req); err != nil {
		return nil, err
	}

	// Get user by email
	user, err := f.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, errors.New("server error")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := f.generateJWTToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Return login response
	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			UserName: user.UserName,
			Role:     user.Role,
		},
	}, nil
}

func (f *AuthFacade) GetProfile(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	user, err := f.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("server error")
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.UserName,
		Role:     user.Role,
	}, nil
}

func (f *AuthFacade) generateJWTToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"rol": user.Role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(f.jwtSecret))
}
