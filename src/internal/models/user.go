package models

/*
import (
	"errors"

	utils "github.com/ntfargo/tir-goapi/src/internal/utils"
)
*/
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	FullName string `json:"fullName"`
	APIKey   string `json:"apiKey"`
	Role     string `json:"role"`
}
