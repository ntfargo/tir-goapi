package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	others "github.com/ntfargo/tir-goapi/src/internal/others"
)

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		err      error
		expected int
	}{
		{err: others.ErrNotFound, expected: http.StatusNotFound},
		{err: others.ErrUnauthorized, expected: http.StatusUnauthorized},
		{err: others.ErrInvalidInput, expected: http.StatusBadRequest},
		{err: others.ErrInternal, expected: http.StatusInternalServerError},
		{err: errors.New("random error"), expected: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		others.ErrorHandler(tt.err, w)
		resp := w.Result()
		if resp.StatusCode != tt.expected {
			t.Errorf("for error %v, expected status %v but got %v", tt.err, tt.expected, resp.StatusCode)
		}

		body := strings.TrimSpace(w.Body.String())
		if body != tt.err.Error() {
			t.Errorf("for error %v, expected response body %v but got %v", tt.err, tt.err.Error(), body)
		}
	}
}
