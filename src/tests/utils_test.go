package tests

import (
	"testing"

	utils "github.com/ntfargo/tir-goapi/src/internal/utils"
)

func TestComplexValidator_Validate(t *testing.T) {
	cv := utils.ComplexValidator{}

	tests := []struct {
		password string
		expected bool
	}{
		{"P@ssw0rd123", true},
		{"Weak", false},
		{"1234590", false},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			valid, _ := cv.Validate(test.password)
			if valid != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, valid)
			}
		})
	}
}

func TestSimpleValidator_Validate(t *testing.T) {
	sv := utils.SimpleValidator{}

	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"another@.com", false},
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			valid, _ := sv.Validate(test.email)
			if valid != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, valid)
			}
		})
	}
}
