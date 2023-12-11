package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"os"
)

func SetupRouter(userHandler *UserHandler, logger *logrus.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(PrometheusMiddleware())

	secret := getJWTSecret(logger)

	authGroup := router.Group("/").Use(JWTMiddleware(secret))
	{
		authGroup.POST("/user", userHandler.CreateUserHandler)
		authGroup.GET("/user", userHandler.GetUserHandler)
		authGroup.PUT("/user/:id", userHandler.UpdateUserHandler)
		authGroup.DELETE("/user/:id", userHandler.DeleteUserHandler)
	}

	router.POST("/auth/login", userHandler.AuthenticateUserHandler)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}

func getJWTSecret(logger *logrus.Logger) string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Warn("JWT_SECRET is not set, using default value")
		return "default_secret"
	}

	return secret
}
