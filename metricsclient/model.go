package metricsclient

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type QueryResponse struct {
	Status    string  `json:"status"`
	Data      *Data   `json:"data"`
	ErrorType *string `json:"errorType"`
	Error     *string `json:"error"`
}

type Metric map[string]string

type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}
type Data struct {
	Resulttype string   `json:"resultType"`
	Result     []Result `json:"result"`
}

func (q *QueryResponse) IsSuccess() bool {
	return StatusSuccess == q.Status
}

type SummaryData struct {
	Nodes      int     `json:"nodes"`
	CPU        float64 `json:"cpu"`
	Memory     float64 `json:"memory"`
	Pods       int     `json:"pods"`
	Containers int     `json:"containers"`
}

type SummaryResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   SummaryData `json:"data"`
}
