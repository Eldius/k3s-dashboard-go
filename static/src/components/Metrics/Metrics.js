import React, { useEffect, useState } from 'react';

import { getSummary } from '../../services/metrics';

export default function Metrics() {
  const [metrics, setMetrics] = useState({});

  useEffect(() => {
    getSummary()
      .then(d => {
        setMetrics(d)
      });
    const id = setInterval(() => {
      getSummary()
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
            <p>{metrics.cpu.toFixed(4)} %</p>
          </div>
        </div>
        <div>
          <header>
            <h2>Memory Usage (%)</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.memory.toFixed(4)} %</p>
          </div>
        </div>
    </div>
  )
}
