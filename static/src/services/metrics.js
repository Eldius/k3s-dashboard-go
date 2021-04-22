
let apiEndpoint = (() => {
  console.log("variable: " + process.env.REACT_APP_SUMMARY_ENDPOINT);
  return (process.env.REACT_APP_SUMMARY_ENDPOINT !== "" ? process.env.REACT_APP_SUMMARY_ENDPOINT + "summary" : "./summary");
})();

export function getSummary() {

  console.log("endpoint: " + apiEndpoint);

  return fetch(apiEndpoint)
    .then(response => response.json())
    .then(response => response.data);
}

