package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Eldius/k3s-dashboard-go/config"
	"github.com/Eldius/k3s-dashboard-go/logger"
)

var (
	log = logger.Log()
	c   = http.Client{
		Timeout: 2 * time.Second,
	}
)

type CPUResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type Metric struct {
}
type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}
type Data struct {
	Resulttype string   `json:"resultType"`
	Result     []Result `json:"result"`
}

func MetricsHandler(rw http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", config.GetPrometheusEndpoint()+"/api/v1/query", nil)
	if err != nil {
		log.WithError(err).
			Error("Failed to create metrics request")
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(&map[string]string{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	req.URL.Query().Add("query", "count(kube_node_info)")
	res, err := c.Do(req)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(&map[string]string{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(&map[string]string{
			"error":  err.Error(),
			"status": "error",
		})
	}

	rw.WriteHeader(200)
	rw.Header().Set("content-type", "application/json")
	_, _ = rw.Write(body)
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
