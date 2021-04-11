import React, { useEffect, useState } from 'react';

import { getMetrics } from '../../services/metrics';

export default function Metrics() {
  const [metrics, setMetrics] = useState({});

  useEffect(() => {
    getMetrics()
      .then(d => {
        setMetrics(d)
      });
    const id = setInterval(() => {
      getMetrics()
      .then(d => {
        setMetrics(d)
      });
    }, 10000);
    return () => clearInterval(id);
  }, []);

  return(
    <div className="metricsBoard">
        <div>
          <header>
            <h2>Nodes</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.nodes}</p>
          </div>
        </div>
        <div>
          <header>
            <h2>Pods</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.pods}</p>
          </div>
        </div>
        <div>
          <header>
            <h2>CPU Usage (%)</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.cpu} %</p>
          </div>
        </div>
        <div>
          <header>
            <h2>Memory Usage (%)</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.memory} %</p>
          </div>
        </div>
    </div>
  )
}
