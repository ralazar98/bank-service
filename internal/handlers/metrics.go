package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var (
	RequestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "bank_service_http_requests_total",
			Help: "Total number of HTTP requests",
		})

	// Создаем гистограмму для замера времени отклика
	ResponseDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "bank_service_http_response_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		})
)

func init() {
	prometheus.MustRegister(RequestsCounter)
	prometheus.MustRegister(ResponseDuration)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()
		RequestsCounter.Inc()
		ResponseDuration.Observe(duration)
	})
}
