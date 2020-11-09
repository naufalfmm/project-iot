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

	RedisHost         string        `envconfig:"REDIS_HOST" default:"127.0.0.1" required:"true"`
	RedisPort         string        `envconfig:"REDIS_PORT" default:"6379" required:"true"`
	RedisPassword     string        `envconfig:"REDIS_PASSWORD" required:"true"`
	RedisPoolSize     int           `envconfig:"REDIS_POOL_SIZE" default:"100" required:"true"`
	RedisPoolTimeout  time.Duration `envconfig:"REDIS_POOL_TIMEOUT" default:"10s" required:"true"`
	RedisWriteTimeout time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"3s" required:"true"`
	RedisReadTimeout  time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"1s" required:"true"`
	RedisDialTimeout  time.Duration `envconfig:"REDIS_DIAL_TIMEOUT" default:"1s" required:"true"`
	RedisMinIdleConns int           `envconfig:"REDIS_MIN_IDLE_CONNS" default:"10" required:"true"`
	RedisMaxConnAge   time.Duration `envconfig:"REDIS_MAX_CONN_AGE" default:"3m" required:"true"`

	CacheExpiration time.Duration `envconfig:"CACHE_EXPIRATION" default:"10m"`

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
