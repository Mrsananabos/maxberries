package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Services   Services
	Mongo      MongoDB
	ServerPort string `envconfig:"PORT" default:":8080"`
}

type MongoDB struct {
	Host       string `envconfig:"MONGO_HOST" required:"true" default:"localhost"`
	Port       int    `envconfig:"MONGO_PORT" required:"true"`
	Name       string `envconfig:"MONGO_NAME" required:"true"`
	Collection string `envconfig:"MONGO_COLLECTION" required:"true"`
}

type Services struct {
	CatalogServiceAddress string `envconfig:"CATALOG_SERVICE_ADDR"  required:"true" default:"http://localhost:8080"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
