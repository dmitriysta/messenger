package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dmitriysta/messenger/message/internal/pkg/metrics"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func PrometheusMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next(rw, r)

		elapsed := time.Since(start)

		status := strconv.Itoa(rw.statusCode)
		endpoint := r.URL.Path
		method := r.Method

		metrics.RequestCount.WithLabelValues(method, endpoint, status).Inc()
		metrics.ResponseTime.WithLabelValues(method, endpoint).Observe(elapsed.Seconds())

		if rw.statusCode >= 400 {
			metrics.ErrorCount.WithLabelValues(method, endpoint, status).Inc()
		}
	}
}
