package main

import (
	"github.com/dmitriysta/messenger/user/internal/api"
	"github.com/dmitriysta/messenger/user/internal/pkg/tracer"
	"github.com/dmitriysta/messenger/user/internal/repository"
	"github.com/dmitriysta/messenger/user/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	trace, closer, err := tracer.NewJaegerTracer("user", logger)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
			"error":  err.Error(),
		}).Fatalf("failed to create new tracer: %v", err)
	}
	defer closer.Close()

	db := repository.DatabaseConnect(logger)

	userRepo := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepo, logger)
	userHandler := api.NewUserHandler(userService, logger, trace)

	router := api.SetupRouter(userHandler)

	if err := router.Run(":8081"); err != nil {
		logger.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
		}).Fatalf("failed to run server: %v", err)
	}
}
