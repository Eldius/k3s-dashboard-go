package metricsclient

import (
	"testing"

	"github.com/Eldius/k3s-dashboard-go/config"
	"gopkg.in/h2non/gock.v1"
)

func TestGetNodesData(t *testing.T) {
	defer gock.Off()

	gock.New(config.GetPrometheusEndpoint()).
		Get("/api/v1/query").
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
