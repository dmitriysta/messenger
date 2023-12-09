package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(userHandler *UserHandler) *gin.Engine {
	router := gin.Default()

	router.Use(PrometheusMiddleware())

	router.POST("/user", userHandler.CreateUserHandler)
	router.GET("/user", userHandler.GetUserHandler)
	router.PUT("/user/:id", userHandler.UpdateUserHandler)
	router.DELETE("/user/:id", userHandler.DeleteUserHandler)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
