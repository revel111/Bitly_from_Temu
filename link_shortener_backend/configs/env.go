package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Printf("No .env file %s found\n", path)
	}
}

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	DbHost    EnvKey = "DB_HOST"
	DbPort    EnvKey = "DB_PORT"
	DbUser    EnvKey = "DB_USER"
	DbPass    EnvKey = "DB_PASS"
	DbName    EnvKey = "DB_NAME"
	RedisHost EnvKey = "REDIS_HOST"
	RedisPort EnvKey = "REDIS_PORT"
	RedisPass EnvKey = "REDIS_PASS"
	BaseUrl   EnvKey = "BASE_URL"
)
