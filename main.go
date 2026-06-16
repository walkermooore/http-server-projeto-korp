package main

import (
	"http-server-projeto-korp/handlers"
	"http-server-projeto-korp/metrics"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Health(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("UP"))
	if err != nil {
		log.Println("erro ao responder health:", err)
	}
}

func main() {

	metrics.IniciarMetrics()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", Health)

	http.HandleFunc(
		"/projeto-korp",
		handlers.ProjetoKorp,
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
