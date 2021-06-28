package app

import (
	"strings"
)

const (
	PromMetricsPath = "/metrics"
	PromMetricsPort = ":2112"
)

var PromBuckets = []float64{0, 50, 100, 200, 300, 500, 1000, 1500, 2000, 2500, 3000}

// Prometheus configuration
type Prometheus struct {
	port    string
	path    Path
	buckets []float64
}

func (p *Prometheus) SetPort(port string) {

	if strings.HasPrefix(port, ":") {
		p.port = port
	} else {
		p.port = ":" + port
	}

}

func (p *Prometheus) GetPort() string {

	return p.port
}

func (p *Prometheus) SetPath(path string) {
	p.path = Path(path)
}

func (p *Prometheus) GetPath() Path {
	return p.path
}

func (p *Prometheus) SetBuckets(buckets []float64) {
	p.buckets = buckets
}

func (p *Prometheus) GetBuckets() []float64 {
	return p.buckets
}
