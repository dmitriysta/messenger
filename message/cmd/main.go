package main

import (
	"github.com/dmitriysta/messenger/message/internal/api"
	"github.com/dmitriysta/messenger/message/internal/pkg/cache"
	"github.com/dmitriysta/messenger/message/internal/pkg/tracer"
	"github.com/dmitriysta/messenger/message/internal/repository"
	"github.com/dmitriysta/messenger/message/internal/service"
	"github.com/joho/godotenv"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	err := godotenv.Load()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
			"error":  err.Error(),
		}).Fatalf("failed to load .env file: %v", err)
	}

	trace, closer, err := tracer.NewJaegerTracer("message", logger)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
			"error":  err.Error(),
		}).Fatalf("failed to create new tracer: %v", err)
	}
	defer closer.Close()

	db := repository.DatabaseConnect(logger)

	cache.InitRedis(logger)

	messageRepo := repository.NewMessageRepository(db, logger)
	messageService := service.NewMessageService(messageRepo, logger, cache.RedisClient)
	messageHandler := api.NewMessageHandler(messageService, logger, trace)

	router := api.SetupRouter(messageHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
		}).Fatalf("failed to run server: %v", err)
	}
}
