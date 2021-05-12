package application

import (
	"github.com/Netflix/go-env"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/infrastructure/postgres"
)

func LoadConfigFromEnv() (Config, error) {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	return config, errors.Wrap(err, "failed to get config from environment")
}

type Config struct {
	WebServerAddr string `env:"WEB_SERVER_ADDR"`
	SWApiBaseUrl  string `env:"SW_API_BASE_URL"`
	Postgres      postgres.Config
}
