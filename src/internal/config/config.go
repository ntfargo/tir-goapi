package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ENV_FILE_PATH = "./.env"
	PORT          = "5211"
)

func LoadEnvVariables() (map[string]string, error) {
	err := godotenv.Load(ENV_FILE_PATH)
	if err != nil {
		return nil, err
	}

	envVars := make(map[string]string)

	for _, envVar := range os.Environ() {
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			envVars[key] = value
		}
	}

	return envVars, nil
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return PORT
	}
	return port
}
