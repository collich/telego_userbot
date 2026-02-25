package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	TelegramAPIID      int
	TelegramAPIHash    string
	TelegramPhone      string
	Telegram2FAPassword string // Optional - only needed if 2FA is enabled

	TargetGroupName string
	DownloadDir     string
	DatabasePath    string
	SessionDir      string

	LogLevel        string
	MaxConcurrentDL int
	DownloadTimeout int // seconds
}

func Load() (*Config, error) {
	apiID, err := strconv.Atoi(getEnv("TELEGRAM_API_ID", ""))
	if err != nil {
		return nil, fmt.Errorf("invalid TELEGRAM_API_ID: %w", err)
	}

	apiHash := getEnv("TELEGRAM_API_HASH", "")
	if apiHash == "" {
		return nil, fmt.Errorf("TELEGRAM_API_HASH is required")
	}

	phone := getEnv("TELEGRAM_PHONE", "")
	if phone == "" {
		return nil, fmt.Errorf("TELEGRAM_PHONE is required")
	}

	config := &Config{
		TelegramAPIID:       apiID,
		TelegramAPIHash:     apiHash,
		TelegramPhone:       phone,
		Telegram2FAPassword: getEnv("TELEGRAM_2FA_PASSWORD", ""), // Optional, empty if no 2FA
		TargetGroupName:     getEnv("TARGET_GROUP_NAME", "dwel"),
		DownloadDir:         getEnv("DOWNLOAD_DIR", "./data/downloads"),
		DatabasePath:        getEnv("DATABASE_PATH", "./data/db.sqlite3"),
		SessionDir:          getEnv("SESSION_DIR", "./session"),
		LogLevel:            getEnv("LOG_LEVEL", "info"),
		MaxConcurrentDL:     getEnvInt("MAX_CONCURRENT_DOWNLOADS", 3),
		DownloadTimeout:     getEnvInt("DOWNLOAD_TIMEOUT", 300),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
