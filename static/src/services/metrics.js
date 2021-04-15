
var apiEndpoint = "./summary";

export function getSummary() {
  return fetch(apiEndpoint)
    .then(response => response.json())
    .then(response => response.data);
}
