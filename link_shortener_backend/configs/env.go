package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return &EnvKey{{
		DB_HOST: "DB_HOST",
	}}
}

// todo: why not config struct?
type EnvKey struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
}

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
