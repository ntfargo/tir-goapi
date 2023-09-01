package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	middleware "github.com/ntfargo/tir-goapi/src/internal/middlewares"
)

func TestDeserializeUserValidToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/example", nil)
	req.Header.Set("x-access-token", "Bearer valid-token-here")
	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middleware.DeserializeUser(testHandler).ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
	}
}

func TestDeserializeUserMissingToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/example", nil)
	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Handler should not be called when token is missing")
	})
	middleware.DeserializeUser(testHandler).ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status %v, but got %v", http.StatusUnauthorized, status)
	}
}
