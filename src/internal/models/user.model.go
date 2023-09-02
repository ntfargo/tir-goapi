package models

import (
	"errors"

	"github.com/ntfargo/tir-goapi/src/internal/config"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	FullName string `json:"fullName"`
	APIKey   string `json:"apiKey"`
	Role     string `json:"role"`
}

type AuthCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrDatabaseInstance = errors.New("error getting database instance")
	ErrCreateUser       = errors.New("error creating user")
)

func CreateUser(user User) (int64, error) {
	db, err := config.GetInstance()
	if err != nil {
		return 0, ErrDatabaseInstance
	}
	defer db.Close()
	conn := db.GetConnection()

	query := `
		INSERT INTO users (email, password, bio, fullName, apiKey, role)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var userID int64
	err = conn.QueryRow(query, user.Email, user.Password, user.Bio, user.FullName, user.APIKey, user.Role).Scan(&userID)
	if err != nil {
		return 0, ErrCreateUser
	}

	return userID, nil
}
