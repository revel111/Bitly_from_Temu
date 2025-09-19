package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func LoadEnvVariables() *ConfigData {
	return &ConfigData{
		DbHost:    getEnv("DB_HOST", "localhost"),
		DbPort:    getEnv("DB_PORT", "5432"),
		DbUser:    getEnv("DB_USER", "admin"),
		DbPass:    getEnv("DB_PASS", "admin"),
		DbName:    getEnv("DB_NAME", "postgres"),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
		RedisPass: getEnv("REDIS_PASS", "admin"),
	}
}

type ConfigData struct {
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	RedisHost string
	RedisPort string
	RedisPass string
}
