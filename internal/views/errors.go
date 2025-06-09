package views

const (
	// Authentication errors
	ErrCodeInvalidCredentials = "INVALID_CREDENTIALS"
	ErrCodeTokenExpired       = "TOKEN_EXPIRED"
	ErrCodeTokenInvalid       = "TOKEN_INVALID"
	ErrCodeUnauthorized       = "UNAUTHORIZED"

	// Validation errors
	ErrCodeValidationFailed = "VALIDATION_FAILED"
	ErrCodeRequiredField    = "REQUIRED_FIELD"
	ErrCodeInvalidFormat    = "INVALID_FORMAT"

	// Business logic errors
	ErrCodeUserExists       = "USER_EXISTS"
	ErrCodeUserNotFound     = "USER_NOT_FOUND"
	ErrCodeItemNotFound     = "ITEM_NOT_FOUND"
	ErrCodePermissionDenied = "PERMISSION_DENIED"

	// System errors
	ErrCodeInternalError = "INTERNAL_ERROR"
	ErrCodeDatabaseError = "DATABASE_ERROR"
)
