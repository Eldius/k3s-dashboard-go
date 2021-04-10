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

func GetNodesData() (*QueryResponse, error) {
	return queryMetric("count(kube_node_info)")
}

func GetCpuData() (nodes *QueryResponse, err error) {
	return queryMetric(`(1-avg(rate(node_cpu_seconds_total{mode="idle", cluster=""}[5m])))*100`)
}

func GetMemoryData() (*QueryResponse, error) {
	return queryMetric(`(1 - sum(:node_memory_MemAvailable_bytes:sum{cluster=""}) / sum(kube_node_status_allocatable_memory_bytes{cluster=""}))*100`)
}

func GetPodCountData() (*QueryResponse, error) {
	return queryMetric(`sum(kubelet_running_pods{cluster="", job="kubelet", metrics_path="/metrics", instance=~"(10.0.0.70:10250|10.0.0.71:10250|10.0.0.72:10250)"})`)
}

func GetContainerCountData() (*QueryResponse, error) {
	return queryMetric(`sum(kubelet_running_containers{cluster="", job="kubelet", metrics_path="/metrics", instance=~"(10.0.0.70:10250|10.0.0.71:10250|10.0.0.72:10250)"})`)
}

func GetBuildinfoData() (*QueryResponse, error) {
	return queryMetric(`kubernetes_build_info`)
}

func GetMetrics() map[string]interface{} {
	metrics := map[string]interface{}{
		"status": "success",
	}
	nodes, err := GetNodesData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["nodes"] = nodes.Data.Result[0].Value[0]

	cpu, err := GetCpuData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["cpu"] = cpu.Data.Result[0].Value[0]

	memory, err := GetMemoryData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["memory"] = memory.Data.Result[0].Value[0]

	podCount, err := GetMemoryData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["pod_count"] = podCount.Data.Result[0].Value[0]

	containerCount, err := GetMemoryData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["container_count"] = containerCount.Data.Result[0].Value[0]

	buildInfo, err := GetMemoryData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	metrics["build_info"] = buildInfo.Data.Result[0].Value[0]

	return metrics
}

func getQueryEndpoint() string {
	return fmt.Sprintf("%s/api/v1/query", config.GetPrometheusEndpoint())
}

func queryMetric(query string) (nodes *QueryResponse, err error) {
	req, err := http.NewRequest("GET", getQueryEndpoint(), nil)
	if err != nil {
		log.WithError(err).
			Error("Failed to create metrics request")
		return
	}

	req.URL.Query().Add("query", query)
	res, err := c.Do(req)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&nodes)
	if err != nil {
		log.WithError(err).
			Error("Failed to fetch metrics")
		return
	}

	return
}
