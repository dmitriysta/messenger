package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "user_response_time_seconds",
			Help: "Response time of user service",
		},
		[]string{"method", "endpoint"},
	)

	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_request_count",
			Help: "Request count of user service",
		},
		[]string{"method", "endpoint", "http_status"},
	)

	ErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_error_count",
			Help: "Error count of user service",
		},
		[]string{"method", "endpoint", "http_status"},
	)
)
