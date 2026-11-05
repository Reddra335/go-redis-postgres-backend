package config

import (
	"fmt"
	"newProjek/internal/shared/logger"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis(config Config) *redis.Client {
	log := logger.New()
	dns := fmt.Sprintf("redis://%s:%s@%s:%s", config.Redis.User, config.Redis.Password, config.DB.DbHost, config.Redis.Port)
	opt, err := redis.ParseURL(dns)
	if err != nil {

		log.WithFields(logrus.Fields{
			"operation": "Init Redis",
			"layer":     "Config",
		}).WithError(err).Fatal("ada kesalahan pada Redis")

	}
	//"redis://<user>:<pass>@localhost:6379"
	// User, Host, Port

	client := redis.NewClient(opt)

	return client
}
