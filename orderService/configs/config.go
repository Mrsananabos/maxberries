package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Services   Services
	Database   Database
	ServerPort string `envconfig:"ORDER_SERVICE_PORT" default:":8080"`
}

type Services struct {
	CatalogServiceAddress    string `envconfig:"CATALOG_SERVICE_ADDR" default:"http://localhost:8080"`
	BackgroundServiceAddress string `envconfig:"BACKGROUND_SERVICE_ADDR" default:"http://localhost:8082"`
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
