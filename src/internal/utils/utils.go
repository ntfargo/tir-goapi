package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type PasswordValidator interface {
	Validate(password string) (bool, string)
}
type EmailValidator interface {
	Validate(email string) (bool, string)
}
type ComplexValidator struct{}
type SimpleValidator struct{}
type HTTPClient interface {
	Post(url string, contentType string, body []byte) (*http.Response, error)
}
type DefaultHTTPClient struct{}
type BcryptHasher struct{}
type RandomAPIKeyGenerator struct{}

func (cv ComplexValidator) Validate(password string) (bool, string) {
	if len(password) < 10 {
		return false, "The password must be at least 10 characters long."
	}
	patterns := []string{
		`[A-Z]`,
		`[a-z]`,
		`[0-9]`,
		`[!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]`,
	}
	for _, pattern := range patterns {
		if !regexp.MustCompile(pattern).MatchString(password) {
			return false, "The password must meet complexity requirements."
		}
	}
	return true, ""
}

func (v SimpleValidator) Validate(email string) (bool, error) {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return false, errors.New("Invalid email address")
	}
	return true, nil
}

func (c DefaultHTTPClient) Post(url string, contentType string, body []byte) (*http.Response, error) {
	return http.Post(url, contentType, bytes.NewBuffer(body))
}

func SendPostRequest(httpClient HTTPClient, url string, body interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(url, "application/json", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		message, ok := result["message"].(string)
		if !ok {
			return nil, errors.New("unknown error occurred")
		}
		return nil, errors.New(message)
	}

	return result, nil
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func (b BcryptHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (r RandomAPIKeyGenerator) Generate() (string, error) {
	keyBytes := make([]byte, 32)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(keyBytes), nil
}

func GenerateJWTSecretKey() (string, error) {
	key := make([]byte, 32) // 256 bits
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
