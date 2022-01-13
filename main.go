package main

import (
	"net/http"

	"github.com/alexgtn/go-middleware-metrics/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := mux.NewRouter()

	metricsMiddleware := middleware.NewMetricsMiddleware()

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/bye", byeHandler).Methods(http.MethodPost)

	r.Use(metricsMiddleware.Metrics)

	http.ListenAndServe(":8080", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Lemon"))
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Potato"))
}
