package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dmitriysta/messenger/user/internal/pkg/metrics"

	"github.com/gin-gonic/gin"
)

const (
	BearerSchema = "Bearer "
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

		if c.Writer.Status() >= 400 {
			metrics.ErrorCount.WithLabelValues(method, endpoint, status).Inc()
		}
	}
}

func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BearerSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			return
		}

		tokenString := authHeader[len(BearerSchema):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
	}
}
