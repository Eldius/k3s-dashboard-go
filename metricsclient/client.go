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
	return QueryMetric(nodesQuery)
}

func GetCpuData() (nodes *QueryResponse, err error) {
	return QueryMetric(cpuQuery)
}

func GetMemoryData() (*QueryResponse, error) {
	return QueryMetric(memoryQuery)
}

func GetPodCountData() (*QueryResponse, error) {
	return QueryMetric(podQuery)
}

func GetContainerCountData() (*QueryResponse, error) {
	return QueryMetric(containerQuery)
}

/*
func GetBuildinfoData() (*QueryResponse, error) {
	return QueryMetric(buildQuery)
}
*/

func GetSummary() *SummaryResponse {

	metrics := &SummaryResponse{
		Status: StatusSuccess,
		Data:   SummaryData{},
	}
	nodes, err := GetNodesData()
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	log.WithField("nodes", nodes).Info("nodes")
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	metrics.Data.Nodes = nodes.GetValueAsInt()

	cpu, err := GetCpuData()
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	log.WithField("cpu", cpu).Info("cpu")
	metrics.Data.CPU = cpu.GetValueAsFloat()

	memory, err := GetMemoryData()
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	log.WithField("memory", memory).Info("memory")
	metrics.Data.Memory = memory.GetValueAsFloat()

	podCount, err := GetPodCountData()
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	log.WithField("podCount", podCount).Info("podCount")
	metrics.Data.Pods = podCount.GetValueAsInt()

	containerCount, err := GetContainerCountData()
	if err != nil {
		return &SummaryResponse{
			Error:  err.Error(),
			Status: StatusError,
		}
	}
	log.WithField("containerCount", containerCount).Info("containerCount")
	metrics.Data.Containers = containerCount.GetValueAsInt()

	return metrics
}

func getQueryEndpoint() string {
	return fmt.Sprintf("%s/api/v1/query", config.GetPrometheusEndpoint())
}

func QueryMetric(query string) (nodes *QueryResponse, err error) {
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
