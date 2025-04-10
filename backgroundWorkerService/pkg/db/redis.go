package db

import (
	"backgroundWorkerService/configs"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Connect(cnf configs.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cnf.Host, cnf.Port),
		DB:   cnf.DB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return redisClient, fmt.Errorf("error connection to redis: %v", err)
	}

	return redisClient, nil
}
