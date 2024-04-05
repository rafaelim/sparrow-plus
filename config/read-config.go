package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Port:   getEnv("PORT", "3000"),
		DBName: fmt.Sprintf("./%v", getEnv("DB_NAME", "sparrow-plus.db")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
