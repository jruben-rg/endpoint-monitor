# Introduction

endpoint-monitor is a basic tool to monitoring `GET` Endpoints written in Golang. It uses Prometheus for monitoring, and Grafana for data visualization. All services are orchestrated by Docker.

endpoint-Monitor reads a monitoring configuration file and performs GET requests to the endpoints specified for each of the sites, e.g:

GET http://www.a-site.com/service-a
GET http://www.a-site.com/service-b

GET http://www.b-site.com/


# Building endpoint-monitor

Prerequisites:

<!-- - [Go Development Kit](https://golang.org/doc/devel/release.html#policy) -->
- [Docker & Docker Compose](https://docs.docker.com/compose/install/) 

Download Endpoint-Monitor application:

`git clone https://github.com/jruben-rg/endpoint-monitor.git`

## Basic Usage

endpoint-monitor reads a `.yaml` configuration file, with the following configuration:

```
endpoints:                  # List of Endpoints object of the monitoring. Multiple Endpoints and paths can be provided.
  - name: My App            # Endpoint name.  
    host: www.web.com
    timeout: 2              # Amount of seconds to wait before timeout.
    period: 3               # Perform a request each period of seconds.
    paths:                  # Paths is an optional configuation. If not specified, only the host is targeted.
      - /service-a/healhtz  
      - /service-b/healthz
  - name: My Other App
  ...    
```

# Run endpoint-monitor

Before running the application, a `conf.yml` file like the provided above must be set under `/config` directory.

`cd` into endpoint-monitor and execute `mon run`

This will start docker compose with three containers:
 - Monitoring app, based in Golang. Will expose Prometheus Scrape endpoint at :2112/metrics by default.
 - Prometheus container. Starts at port :9090 and scrapes metrics from Monitoring app.
 - Grafana container. Datasource for Prometheus

If you wish to customise the container ports, run `mon run <GRAFANA_PORT> <PROMETHEUS_PORT> <METRICS_PORT>`

The Grafana container already has provisioned the prometheus datasource and some dashboards to visualize the monitored endpoints and the duration of the requests.
