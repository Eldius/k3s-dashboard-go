package metricsclient

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetNodesData() (cpu *QueryResponse, err error) {
	req, err := http.NewRequest("GET", getQueryEndpoint(), nil)
	if err != nil {
		log.WithError(err).
			Error("Failed to create metrics request")
		return
	}

	req.URL.Query().Add("query", "count(kube_node_info)")
	res, err := c.Do(req)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&cpu)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		return
	}

	return
}

func getQueryEndpoint() string {
	return fmt.Sprintf("%s/api/v1/query", config.GetPrometheusEndpoint())
}
