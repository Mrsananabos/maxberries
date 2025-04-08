package configs

import (
	"fmt"
	"os"
)

type Config struct {
	Database   Database
	ServerPort string `envconfig:"SERVER_PORT" default:":8080"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func NewParsedConfig() Config {
	return Config{
		ServerPort: getEnv("PORT"),
		Database: Database{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
		},
	}
}

func getEnv(name string) string {
	value, exist := os.LookupEnv(name)
	if !exist {
		panic(fmt.Sprintf("Not exist %s env", name))
	}

	return value
}
