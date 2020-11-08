package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/naufalfmm/project-iot/resource/config"
	"github.com/pkg/errors"
)

func New(config *config.EnvConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisPassword,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.Wrap(err, "failed to open redis connection")
	}

	return client, nil
}
