package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	TimeToLive = time.Hour
)

var RedisClient *redis.Client

func InitRedis(logger *logrus.Logger) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "cache",
			"func":   "InitRedis",
			"error":  err.Error(),
		}).Fatalf("failed to ping redis: %v", err)
	}
}
