package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Eldius/k3s-dashboard-go/logger"
	"github.com/spf13/viper"
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

func getBaseEndpoint() string {
	return viper.GetString("metrics.server.endpoint")
}

func MetricsHandler(rw http.ResponseWriter, r *http.Request) {
	res, err := c.Get(getBaseEndpoint() + "/api/v1/query" + "'count(kube_node_info)'")
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(err)
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

	host := fmt.Sprintf(":%d", port)

	log.Infof("Starting admin server on port %d\n", port)
	log.WithError(http.ListenAndServe(host, mux)).Error("Failed to start HTTP server")
}
