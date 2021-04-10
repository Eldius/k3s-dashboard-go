package metricsclient

import (
	"net/http"
	"testing"

	"github.com/Eldius/k3s-dashboard-go/config"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	config.Setup("")
}

func TestGetNodesData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == nodesQuery
		}).
		Reply(200).
		File("./sample_results/nodes.json")

	d, err := GetNodesData()
	if err != nil {
		t.Errorf("Returned an error querying nodes data: '%s'", err.Error())
	}
	if d.Data.Result[0].Value[1] != "1" {
		t.Errorf("Must return result '1': '%s'", d.Data.Result[0].Value[1])
	}
}

func TestGetCpuData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == cpuQuery
		}).
		Reply(200).
		File("./sample_results/cpu.json")

	d, err := GetCpuData()
	if err != nil {
		t.Errorf("Returned an error querying cpu data: '%s'", err.Error())
	}
	if d.Data.Result[0].Value[1] != "13.332407407407466" {
		t.Errorf("Must return result '13.332407407407466': '%s'", d.Data.Result[0].Value[1])
	}
}

func TestGetMemoryData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == memoryQuery
		}).
		Reply(200).
		File("./sample_results/memory.json")

	d, err := GetMemoryData()
	if err != nil {
		t.Errorf("Returned an error querying nodes data: '%s'", err.Error())
	}
	if d.Data.Result[0].Value[1] != "17.731588183240767" {
		t.Errorf("Must return result '17.731588183240767': '%s'", d.Data.Result[0].Value[1])
	}
}

func TestGetPodCountData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == podQuery
		}).
		Reply(200).
		File("./sample_results/pods.json")

	d, err := GetPodCountData()
	if err != nil {
		t.Errorf("Returned an error querying nodes data: '%s'", err.Error())
	}
	if d.Data.Result[0].Value[1] != "16" {
		t.Errorf("Must return result '16': '%s'", d.Data.Result[0].Value[1])
	}
}

func TestGetContainerCountData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == containerQuery
		}).
		Reply(200).
		File("./sample_results/containers.json")

	d, err := GetContainerCountData()
	if err != nil {
		t.Errorf("Returned an error querying nodes data: '%s'", err.Error())
	}
	if d.Data.Result[0].Value[1] != "41" {
		t.Errorf("Must return result '41': '%s'", d.Data.Result[0].Value[1])
	}
}
