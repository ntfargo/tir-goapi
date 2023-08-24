package others

import (
	"errors"
	"net/http"
)

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

func ErrorHandler(err error, w http.ResponseWriter) {
	switch {
	case IsNotFound(err):
		w.WriteHeader(http.StatusNotFound)
	case IsUnauthorized(err):
		w.WriteHeader(http.StatusUnauthorized)
	case IsInvalidInput(err):
		w.WriteHeader(http.StatusBadRequest)
	case IsInternal(err):
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(err.Error()))
}
