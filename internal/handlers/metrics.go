package handlers

import "github.com/prometheus/client_golang/prometheus"

var (
	// Создаем счетчик для подсчета количества запросов
	RequestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "bank_service_http_requests_total", // Имя метрики
			Help: "Total number of HTTP requests",    // Описание метрики
		})

	// Создаем гистограмму для замера времени отклика
	ResponseDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "bank_service_http_response_duration_seconds", // Имя метрики
			Help:    "Duration of HTTP requests in seconds",        // Описание метрики
			Buckets: prometheus.DefBuckets,                         // Диапазоны значений (по умолчанию)
		})
)

func RegisMetrics() {
	prometheus.MustRegister(RequestsCounter)
	prometheus.MustRegister(ResponseDuration)
}
