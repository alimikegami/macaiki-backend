package utils

import (
	"errors"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
	// ErrBadParamInput will throw if the email already used
	ErrEmailAlreadyUsed = errors.New("Email already used")

	ErrUsernameAlreadyUsed = errors.New("Username already used")
	// ErrForbidden will throw if the authorization is fail
	ErrForbidden = errors.New("Forbidden")
	// ErrLoginFailed will throw if the email & password invalid
	ErrLoginFailed = errors.New("Invalid Email or Password")

	ErrEmailRequired = errors.New("Email is Required")

	ErrReportCategoryNameRequired = errors.New("Report Category Name is Required")

	ErrPasswordRequired = errors.New("Password is Required")

	ErrPasswordDontMatch = errors.New("Password don't match")

	ErrUnauthorizedAccess = errors.New("unauthorized access")

	ErrDuplicateEntry = errors.New("Duplicate entry")
)

// unfinished
func GetStatusCode(err error) int {
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadParamInput:
		return http.StatusBadRequest
	case ErrEmailAlreadyUsed:
		return http.StatusBadRequest
	case ErrEmailRequired:
		return http.StatusBadRequest
	case ErrDuplicateEntry:
		return http.StatusConflict
	case ErrReportCategoryNameRequired:
		return http.StatusBadRequest
	case ErrPasswordDontMatch:
		return http.StatusBadRequest
	case ErrForbidden:
		return http.StatusForbidden
	case ErrLoginFailed:
		return http.StatusUnauthorized
	case ErrUnauthorizedAccess:
		return http.StatusUnauthorized
	default:
		return http.StatusOK
	}
}
