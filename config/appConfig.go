package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	JWTSecret  string
	DSN        string
}

func SetupEnv() (AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found or error loading .env file:", err)
	}

	config := AppConfig{}

	if err := loadEnv(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}

func loadEnv(config *AppConfig) error {
	if port := os.Getenv("HTTP_PORT"); port != "" {
		config.ServerPort = ":" + port
	}

	config.JWTSecret = os.Getenv("JWT_SECRET")
	if config.JWTSecret == "" {
		return errors.New("JWT_SECRET environment variable not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return errors.New("DB_HOST environment variable not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return errors.New("DB_USER environment variable not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return errors.New("DB_PASSWORD environment variable not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return errors.New("DB_NAME environment variable not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return errors.New("DB_PORT environment variable not set")
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		return errors.New("DB_SSLMODE environment variable not set")
	}

	config.DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSSLMode,
	)

	return nil
}
