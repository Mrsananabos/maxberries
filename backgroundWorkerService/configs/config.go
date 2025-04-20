package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port             string `envconfig:"PORT" default:":8080"`
	Redis            Redis
	FixerAccessToken string `envconfig:"FIXER_TOKEN"`
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST"`
	Port string `envconfig:"REDIS_PORT"`
	DB   int    `envconfig:"REDIS_DB"`
	TTL  int    `envconfig:"REDIS_TTL"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
