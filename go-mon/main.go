package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jruben-rg/endpoint-monitor/go-mon/config"
	"github.com/jruben-rg/endpoint-monitor/go-mon/monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("provide a path for a yaml file so configuration can be loaded")
	}

	config, err := config.New(os.Args[1])
	if err != nil {
		log.Fatalf("application configuration couldn't be loaded: %s\n", err)
	}

	monitor.New(config.Prometheus.GetBuckets()).Watch(config.Endpoints)

	promPath := config.Prometheus.GetPath()
	http.Handle(promPath.Value(), promhttp.Handler())
	http.ListenAndServe(config.Prometheus.GetPort(), nil)
}
