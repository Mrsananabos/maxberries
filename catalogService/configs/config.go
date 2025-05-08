package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Database   Database
	ServerPort string `envconfig:"PORT" default:":8080"`
	Redis      Redis
	Services   Services
}

type Database struct {
	Host     string `envconfig:"DB_HOST" default:":localhost"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Schema   string `envconfig:"DB_SCHEMA" required:"true"`
}

type Services struct {
	AuthServiceAddress string `envconfig:"AUTH_SERVICE_ADDR" default:"http://localhost:8084"`
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST"  required:"true"`
	Port string `envconfig:"REDIS_PORT"  required:"true"`
	DB   int    `envconfig:"REDIS_DB"  required:"true"`
	TTL  int    `envconfig:"REDIS_TTL"  required:"true"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
