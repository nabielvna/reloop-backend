package validation

import (
	"errors"
	"regexp"
	"reloop-backend/internal/dto"
	"strings"
)

func ValidateRegisterRequest(req dto.RegisterRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}

	if strings.TrimSpace(req.UserName) == "" {
		return errors.New("userName is required")
	}

	if len(req.UserName) < 3 {
		return errors.New("userName must be at least 3 characters")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func ValidateLoginRequest(req dto.LoginRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
