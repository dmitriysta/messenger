package api

import (
	"github.com/gin-gonic/gin"
	"message/internal/pkg/metrics"
	"strconv"
	"time"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		elapsed := time.Since(start)

		status := strconv.Itoa(c.Writer.Status())
		endpoint := c.Request.URL.Path
		method := c.Request.Method

		metrics.RequestCount.WithLabelValues(method, endpoint, status).Inc()
		metrics.ResponseTime.WithLabelValues(method, endpoint).Observe(elapsed.Seconds())
		metrics.ErrorCount.WithLabelValues(method, endpoint, status).Inc()
	}
}
