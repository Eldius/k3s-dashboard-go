package metricsclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Eldius/k3s-dashboard-go/config"
	"github.com/Eldius/k3s-dashboard-go/logger"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.Log()
	c   = http.Client{
		Timeout: 2 * time.Second,
	}
)

const (
	nodesQuery     = "count(kube_node_info)"
	cpuQuery       = `(1-avg(rate(node_cpu_seconds_total{mode="idle", cluster=""}[5m])))*100`
	memoryQuery    = `(1 - sum(:node_memory_MemAvailable_bytes:sum{cluster=""}) / sum(kube_node_status_allocatable_memory_bytes{cluster=""}))*100`
	podQuery       = `sum(kubelet_running_pods{cluster="", job="kubelet", metrics_path="/metrics"})`
	containerQuery = `sum(kubelet_running_containers{cluster="", job="kubelet", metrics_path="/metrics"})`
	buildQuery     = `kubernetes_build_info`
)

func GetNodesData() (*QueryResponse, error) {
	return queryMetric(nodesQuery)
}

func GetCpuData() (nodes *QueryResponse, err error) {
	return queryMetric(cpuQuery)
}

func GetMemoryData() (*QueryResponse, error) {
	return queryMetric(memoryQuery)
}

func GetPodCountData() (*QueryResponse, error) {
	return queryMetric(podQuery)
}

func GetContainerCountData() (*QueryResponse, error) {
	return queryMetric(containerQuery)
}

/*
func GetBuildinfoData() (*QueryResponse, error) {
	return queryMetric(buildQuery)
}
*/

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
	log.WithField("nodes", nodes).Info("nodes")
	metrics["nodes"] = nodes.Data.Result[0].Value[0]

	cpu, err := GetCpuData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	log.WithField("cpu", cpu).Info("cpu")
	metrics["cpu"] = cpu.Data.Result[0].Value[0]

	memory, err := GetMemoryData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	log.WithField("memory", memory).Info("memory")
	metrics["memory"] = memory.Data.Result[0].Value[0]

	podCount, err := GetPodCountData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	log.WithField("podCount", podCount).Info("podCount")
	metrics["pod_count"] = podCount.Data.Result[0].Value[0]

	containerCount, err := GetContainerCountData()
	if err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": "error",
		}
	}
	log.WithField("containerCount", containerCount).Info("containerCount")
	metrics["container_count"] = containerCount.Data.Result[0].Value[0]

	//buildInfo, err := GetBuildinfoData()
	//if err != nil {
	//	return map[string]interface{}{
	//		"error":  err.Error(),
	//		"status": "error",
	//	}
	//}
	//log.WithField("buildInfo", buildInfo).Info("buildInfo")
	//metrics["build_info"] = buildInfo.Data.Result[0].Value[0]

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
	q := req.URL.Query()
	q.Set("query", query)
	req.URL.RawQuery = q.Encode()

	log.WithFields(logrus.Fields{
		"query": req.URL.Query().Encode(),
		"url":   req.URL.String(),
	}).Info("test")

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
