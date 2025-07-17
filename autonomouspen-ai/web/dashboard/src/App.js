import React, { useState, useEffect } from 'react';

function App() {
  const [metrics, setMetrics] = useState({});
  const [health, setHealth] = useState({});

  useEffect(() => {
    // Fetch data from API
    // setInterval(() => {
    //   fetch('/api/metrics').then(res => res.json()).then(data => setMetrics(data));
    //   fetch('/api/health').then(res => res.json()).then(data => setHealth(data));
    // }, 1000);
  }, []);

  return (
    <div>
      <h1>AUTONOMOUSPEN AI Dashboard</h1>
      <h2>Performance Metrics</h2>
      <p>Targets scanned: {metrics.targets_scanned}/hour</p>
      <p>Vulnerabilities found: {metrics.vulnerabilities_found} today</p>
      <p>Success rate: {metrics.success_rate}% accepted</p>
      <p>Total earnings: ${metrics.total_earnings} this month</p>
      <h2>System Health</h2>
      <p>Worker status: {health.workers_active}/{health.workers_total} active</p>
      <p>Queue depth: {health.queue_depth} pending</p>
      <p>API rate limits: HackerOne {health.hackerone_rate_limit}/60</p>
      <p>Error rate: {health.error_rate}%</p>
    </div>
  );
}

export default App;
