package others

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

var (
	ErrNotFound                = errors.New("not found")
	ErrUnauthorized            = errors.New("unauthorized")
	ErrInvalidInput            = errors.New("invalid input")
	ErrInternal                = errors.New("internal error")
	ErrInvalidToken            = errors.New("invalid token")
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
	ErrMissingAuthHeader       = errors.New("missing authorization header")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, ErrInvalidInput)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}

func WriteErrorResponse(w http.ResponseWriter, code int, errType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := ErrorResponse{
		Code:    code,
		Type:    errType,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func ErrorHandler(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case IsNotFound(err):
		WriteErrorResponse(w, http.StatusNotFound, "NOT_FOUND", "Service not found")
	case IsUnauthorized(err):
		WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Missing x-access-token header variable")
	case IsInvalidInput(err):
		WriteErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", err.Error())
	case IsInternal(err):
		WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An internal error occurred")
	default:
		WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}
}
