package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Eldius/k3s-dashboard-go/config"
	"github.com/Eldius/k3s-dashboard-go/logger"
	"github.com/Eldius/k3s-dashboard-go/metricsclient"
)

var (
	log = logger.Log()
)

func MetricsHandler(rw http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	data := make(map[string]interface{})

	cpu, err := metricsclient.GetNodesData()
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		rw.WriteHeader(http.StatusInternalServerError)
		result["error"] = err.Error()
		result["status"] = "error"
		_ = json.NewEncoder(rw).Encode(&result)
	}
	data["nodes"] = cpu.Data.Result[0].Value[1]

	result["data"] = data
	rw.WriteHeader(200)
	rw.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(rw).Encode(result)
}

func Start(port int) {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	mux.HandleFunc("/metrics", MetricsHandler)

	log.Infof("prometheus endpoint: %s", config.GetPrometheusEndpoint())
	log.Infof("DASHBOARD_PROMETHEUS_ENDPOINT: %s", os.Getenv("DASHBOARD_PROMETHEUS_ENDPOINT"))

	host := fmt.Sprintf(":%d", port)

	log.Infof("Starting admin server on port %d\n", port)
	log.WithError(http.ListenAndServe(host, mux)).Error("Failed to start HTTP server")
}
