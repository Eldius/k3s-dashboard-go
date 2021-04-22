import React, { useEffect, useState } from 'react';

import { getSummary } from '../../services/metrics';

import NumberFormat from 'react-number-format';

import "./Metrics.css"

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
        <div className="metricCard">
          <header>
            <h2>Nodes</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.nodes}</p>
          </div>
        </div>
        <div className="metricCard">
          <header>
            <h2>Pods</h2>
          </header>
          <div className="metricValue">
            <p>{metrics.pods}</p>
          </div>
        </div>
        <div className="metricCard">
          <header>
            <h2>CPU Usage (%)</h2>
          </header>
          <div className="metricValue">
            <p>
              <NumberFormat 
                value={metrics.cpu}
                className="metricValueFloat"
                displayType={'text'} 
                thousandSeparator={true}
                decimalScale={4}
                fixedDecimalScale={true}
                suffix={'%'} />
            </p>
          </div>
        </div>
        <div className="metricCard">
          <header>
            <h2>Memory Usage (%)</h2>
          </header>
          <div className="metricValue">
            <p>
              <NumberFormat 
                value={metrics.memory}
                className="metricValueFloat"
                displayType={'text'} 
                thousandSeparator={true}
                decimalScale={4}
                fixedDecimalScale={true}
                suffix={'%'} />
            </p>
          </div>
        </div>
    </div>
  )
}
