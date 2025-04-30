package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port             string `envconfig:"PORT" default:":8080"`
	Database         Database
	Redis            Redis
	Kafka            Kafka
	Services         Services
	FixerAccessToken string `envconfig:"FIXER_TOKEN"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" default:":localhost"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Schema   string `envconfig:"DB_SCHEMA" required:"true" default:":background"`
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST"`
	Port string `envconfig:"REDIS_PORT"`
	DB   int    `envconfig:"REDIS_DB"`
	TTL  int    `envconfig:"REDIS_TTL"`
}

type Kafka struct {
	Host string `envconfig:"KAFKA_HOST" required:"true"`
	Port string `envconfig:"KAFKA_PORT" required:"true"`
}

type Services struct {
	OrderServiceAddress string `envconfig:"ORDER_SERVICE_ADDR" default:"http://localhost:8081"`
}

func NewParsedConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
