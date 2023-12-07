package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "message_response_time_seconds",
			Help: "Response time of message service",
		},
		[]string{"method", "endpoint"},
	)

	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "message_request_count",
			Help: "Request count of message service",
		},
		[]string{"method", "endpoint", "http_status"},
	)

	ErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "message_error_count",
			Help: "Error count of message service",
		},
		[]string{"method", "endpoint", "http_status"},
	)
)
