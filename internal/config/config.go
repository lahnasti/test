package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API    string
	Header string
	Token  string
}

func SetupConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}
	api := os.Getenv("API_URL")
	header := os.Getenv("HEADER")
	token := os.Getenv("TOKEN")

	if api == "" || header == "" || token == "" {
		return nil, fmt.Errorf("missing required environment variables: API, HEADER, and TOKEN")
	}

	return &Config{
		API:    api,
		Header: header,
		Token:  token,
	}, nil
}
