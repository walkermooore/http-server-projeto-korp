package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Quantidade total de requisições",
})

func IniciarMetrics() {
	prometheus.MustRegister(RequestsTotal)
}
