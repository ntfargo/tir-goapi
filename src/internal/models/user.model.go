package models

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
