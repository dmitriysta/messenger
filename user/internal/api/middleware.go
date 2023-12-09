package api

import (
	"strconv"
	"time"

	"github.com/dmitriysta/messenger/user/internal/pkg/metrics"

	"github.com/gin-gonic/gin"
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
