package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Eldius/k3s-dashboard-go/config"
	"github.com/Eldius/k3s-dashboard-go/logger"
	"github.com/Eldius/k3s-dashboard-go/metricsclient"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.Log()
)

func MetricsHandler(rw http.ResponseWriter, r *http.Request) {
	metrics := metricsclient.GetMetrics()
	rw.WriteHeader(200)
	rw.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(rw).Encode(metrics)
}

func QueryHandler(rw http.ResponseWriter, r *http.Request) {

}

func Start(port int) {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	mux.HandleFunc("/sumary", MetricsHandler)
	mux.HandleFunc("/query", QueryHandler)
	host := fmt.Sprintf(":%d", port)

	log.WithFields(logrus.Fields{
		"prometheus_endpoint": config.GetPrometheusEndpoint(),
		"listeningAt":         host,
	}).Infof("Stating server")

	log.WithError(http.ListenAndServe(host, mux)).Error("Failed to start HTTP server")
}
