package configs

import (
	"fmt"
	"os"
)

type Config struct {
	Database   Database
	ServerPort string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewParsedConfig() Config {
	return Config{
		ServerPort: getEnv("CATALOG_SERVICE_PORT"),
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
