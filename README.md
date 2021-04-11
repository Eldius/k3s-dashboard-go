# k3s-dashboard-go #

## links ##

- [Prometheus - HTTP API](https://prometheus.io/docs/prometheus/latest/querying/api/)
- [Tutorial: Build a Prometheus Dashboard for K3s with Wio Terminal](https://thenewstack.io/tutorial-build-a-prometheus-dashboard-for-k3s-with-wio-terminal/)
- [Kubernetes PromQL (Prometheus Query Language) CPU aggregation walkthrough](https://medium.com/@amimahloof/kubernetes-promql-prometheus-cpu-aggregation-walkthrough-2c6fd2f941eb)
- [Kubernetes in Production: The Ultimate Guide to Monitoring Resource Metrics with Prometheus](https://www.replex.io/blog/kubernetes-in-production-the-ultimate-guide-to-monitoring-resource-metrics)

## commands to test requests ##

```shell
kubectl run -i --tty alpine --image=alpine --restart=Never -- sh
```

```shell
# nodes:           count(kube_node_info)
# cpu:             (1-avg(rate(node_cpu_seconds_total{mode="idle", cluster=""}[5m])))*100
# memory:          (1 - sum(:node_memory_MemAvailable_bytes:sum{cluster=""}) / sum(kube_node_status_allocatable_memory_bytes{cluster=""}))*100
# pod count:       sum(kubelet_running_pods{cluster="", job="kubelet", metrics_path="/metrics"})
# container count: sum(kubelet_running_containers{cluster="", job="kubelet", metrics_path="/metrics"})
# build info:      kubernetes_build_info

curl 'http://prometheus-kube-prometheus-prometheus.prometheus.svc.cluster.local:9090/api/v1/query?query=<metric-query-urlencoded>'
```