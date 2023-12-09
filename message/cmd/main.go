package main

import (
	"github.com/dmitriysta/messenger/message/internal/api"
	"github.com/dmitriysta/messenger/message/internal/pkg/tracer"
	"github.com/dmitriysta/messenger/message/internal/repository"
	"github.com/dmitriysta/messenger/message/internal/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

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

	messageRepo := repository.NewMessageRepository(db, logger)
	messageService := service.NewMessageService(messageRepo, logger)
	messageHandler := api.NewMessageHandler(messageService, logger, trace)

	router := api.SetupRouter(messageHandler)

	if err := router.Run(":8080"); err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
		}).Fatalf("failed to run server: %v", err)
	}
}
