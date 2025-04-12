package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port             string
	Redis            Redis
	FixerAccessToken string
}

type Redis struct {
	Host string
	Port string
	DB   int
	TTL  int
}

func NewParsedConfig() Config {
	return Config{
		Redis: Redis{
			Host: getEnv("REDIS_HOST"),
			Port: getEnv("REDIS_PORT"),
			DB:   getIntEnv("REDIS_DB"),
			TTL:  getIntEnv("REDIS_TTL"),
		},
		Port:             getEnv("BACKGROUND_SERVICE_PORT"),
		FixerAccessToken: getEnv("FIXER_TOKEN"),
	}
}

func getEnv(name string) string {
	value, exist := os.LookupEnv(name)
	if !exist {
		panic(fmt.Sprintf("Not exist %s env", name))
	}

	return value
}

func getIntEnv(name string) int {
	redisDB, err := strconv.Atoi(getEnv(name))
	if err != nil {
		panic(fmt.Sprintf("%s is not integer", name))
	}

	return redisDB
}
