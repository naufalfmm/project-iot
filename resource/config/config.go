package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	ServerName string `envconfig:"SERVER_NAME" required:"true"`
	ServerPort int    `envconfig:"SERVER_PORT" default:"7010" required:"true"`

	PostgresDbURI             string        `envconfig:"POSTGRES_DB_URI" required:"true"`
	PostgresMaxIdleConnection int           `envconfig:"POSTGRES_MAX_IDLE_CONNECTION" default:"10"`
	PostgresMaxOpenConnection int           `envconfig:"POSTGRES_MAX_OPEN_CONNECTION" default:"10"`
	PostgresConnMaxLifetime   time.Duration `envconfig:"POSTGRES_CONNECTION_MAX_LIFE_TIME" default:"60s"`
	PostgresLogMode           bool          `envconfig:"POSTGRES_LOG_MODE" default:"false"`
	PostgresLogColorful       bool          `envconfig:"POSTGRES_LOG_COLORFUL" default:"false"`

	JwtPrivateKey        string        `envconfig:"JWT_PRIVATE_KEY" required:"true"`
	JwtPublicKey         string        `envconfig:"JWT_PUBLIC_KEY" required:"true"`
	JwtExpiresInDuration time.Duration `envconfig:"JWT_EXPIRES_IN_DURATION" default:"5d" required:"true"`
}

func NewConfig() (*EnvConfig, error) {
	var config EnvConfig

	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load env file")
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process env")
	}

	return &config, nil
}
