package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

var (
	RequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "bank_service_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"code", "method"},
	)

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
		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		duration := time.Since(start).Seconds()
		defer func() {
			RequestsCounter.WithLabelValues(r.Method, strconv.Itoa(sw.statusCode)).Inc()
			ResponseDuration.Observe(duration)
		}()
		next.ServeHTTP(sw, r)

	})
}
