package application

import (
	"github.com/Netflix/go-env"
	"github.com/pkg/errors"
)

func LoadConfigFromEnv() (Config, error) {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	return config, errors.Wrap(err, "failed to get config from environment")
}

type Config struct {
}
