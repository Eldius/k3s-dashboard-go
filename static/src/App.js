import React from 'react';

import './App.css';

import Metrics from './components/Metrics/Metrics';

function App() {
/*
  const [data, setData] = useState({metrics: {}, isFetching: false});
  fetch("http://192.168.100.195/dashboard/summary")
  .then(response => response.json())
  .then(result => {
      console.log(result);
      setData({"metrics": result.data});
  })
  .catch(e => {
      console.log(e);
  });
*/
  return (
    <div className="App">
      <Metrics />
    </div>
  );
}

export default App;
