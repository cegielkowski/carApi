package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPORT     string
	DatabaseURL    string
	CacheURL       string
	LoggerLevel    string
	ContextTimeout int
}

// LoadConfig will load config from environment variable
func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	serverPORT := os.Getenv("SERVER_PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	cacheURL := os.Getenv("CACHE_URL")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))

	return &Config{
		ServerPORT:     serverPORT,
		DatabaseURL:    databaseURL,
		CacheURL:       cacheURL,
		LoggerLevel:    loggerLevel,
		ContextTimeout: contextTimeout,
	}
}
