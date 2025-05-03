package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	Redis      Redis
	JWTConfig  JWTConfig
	ServerPort string `envconfig:"PORT" default:":8080"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Schema   string `envconfig:"DB_SCHEMA" required:"true"`
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST" required:"true"`
	Port string `envconfig:"REDIS_PORT" required:"true"`
	DB   int    `envconfig:"REDIS_DB" required:"true"`
}

type JWTConfig struct {
	Secret          string `envconfig:"JWT_SECRET" required:"true"`
	AccessTokenTTL  int    `envconfig:"ACCESS_TOKEN_TTL" required:"true"`
	RefreshTokenTTL int    `envconfig:"REFRESH_TOKEN_TTL" required:"true"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
