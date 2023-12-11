package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(userHandler *UserHandler) *gin.Engine {
	router := gin.Default()

	router.Use(PrometheusMiddleware())

	authGroup := router.Group("/").Use(JWTMiddleware("secret"))
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
