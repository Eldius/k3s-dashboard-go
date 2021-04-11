
var apiEndpoint = "http://dashboard-svc.default.svc.cluster.local:8080";

export function getMetrics() {
  return fetch(apiEndpoint)
    .then(response => response.json())
    .then(response => response.data);
}