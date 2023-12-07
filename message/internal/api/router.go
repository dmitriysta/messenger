package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(messageHandler *MessageHandler) *gin.Engine {
	router := gin.Default()

	router.Use(PrometheusMiddleware())

	router.POST("/messages", messageHandler.CreateMessageHandler)
	router.GET("/messages", messageHandler.GetMessagesByChannelIdHandler)
	router.PUT("/messages/:id", messageHandler.UpdateMessageHandler)
	router.DELETE("/messages/:id", messageHandler.DeleteMessageHandler)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
