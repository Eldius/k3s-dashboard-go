
export function getMetrics() {
  return fetch("http://192.168.100.195/dashboard/summary")
    .then(response => response.json())
    .then(response => response.data);
}