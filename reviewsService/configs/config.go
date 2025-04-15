package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   MongoDB
	ServerPort string `envconfig:"PORT" default:":8082"`
}

type MongoDB struct {
	Host string `envconfig:"DB_HOST" required:"true" default:"localhost"`
	Port int    `envconfig:"DB_PORT" required:"true"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
