package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Services   Services
	Database   Database
	Kafka      Kafka
	ServerPort string `envconfig:"PORT" default:":8080"`
}

type Services struct {
	CatalogServiceAddress    string `envconfig:"CATALOG_SERVICE_ADDR" default:"http://localhost:8080"`
	BackgroundServiceAddress string `envconfig:"BACKGROUND_SERVICE_ADDR" default:"http://localhost:8082"`
	AuthServiceAddress       string `envconfig:"AUTH_SERVICE_ADDR" default:"http://localhost:8084"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Schema   string `envconfig:"DB_SCHEMA" required:"true" default:"orders"`
}

type Kafka struct {
	Host string `envconfig:"KAFKA_HOST" required:"true"`
	Port string `envconfig:"KAFKA_PORT" required:"true"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
