package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Env is the must environment config for database
type Env struct {
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort string `envconfig:"DB_PORT" default:"3306"`
	DBUser string `envconfig:"DB_USER" required:"true"`
	DBPass string `envconfig:"DB_PASS" required:"true"`
	DBName string `envconfig:"DB_NAME" required:"true"`
}

// ReadFromEnv reads environmental variables of above config
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.New("environment variable is not correctly set")
	}
	return &env, nil
}
