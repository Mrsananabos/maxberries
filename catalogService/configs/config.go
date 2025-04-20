package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Database   Database
	ServerPort string `envconfig:"PORT" default:":8080"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" default:":localhost"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Schema   string `envconfig:"DB_SCHEMA" required:"true"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
