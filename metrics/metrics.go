package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTPRequests — счётчик HTTP-запросов.
	HTTPRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",            // Имя метрики
			Help: "Total number of HTTP requests.", // Описание метрики
		},
		[]string{"method", "endpoint"}, // Метки (labels) для метрики
	)
)

// Init регистрирует метрики.
func Init() {
	prometheus.MustRegister(HTTPRequests)
}

// Handler возвращает HTTP-обработчик для эндпоинта /metrics.
func Handler() http.Handler {
	return promhttp.Handler()
}
