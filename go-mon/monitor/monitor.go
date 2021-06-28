package monitor

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jruben-rg/endpoint-monitor/go-mon/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	EndpointName  = "endpoint_name"
	HttpCode      = "http_code"
	Path          = "path"
	Job           = "job"
	HistogramName = "endpoint_monitoring"
	JobName       = "endpoint_monitor"
)

var once sync.Once
var singleton *monitor

type monitor struct {
	histogram *prometheus.HistogramVec
	jobName   string
}

func New(promBuckets []float64) *monitor {
	once.Do(func() {
		singleton = &monitor{
			histogram: promauto.NewHistogramVec(prometheus.HistogramOpts{
				Name:    HistogramName,
				Help:    "Histogram for " + HistogramName,
				Buckets: promBuckets,
			},
				[]string{Job, EndpointName, Path, HttpCode},
			),
			jobName: JobName,
		}
	})
	return singleton
}

func (m *monitor) Watch(endpoints []app.Endpoint) {

	for i := range endpoints {
		go m.watchEndpoint(endpoints[i])
	}
}

// Monitor - watches for a monitor configuration
func (m *monitor) watchEndpoint(endpoint app.Endpoint) {

	client := http.Client{
		Timeout: endpoint.Request.Timeout * time.Second,
	}

	// TO DO
	endpointPaths := endpoint.Paths

	endpointName := endpoint.Name.ToSnakeCase()

	for {

		for _, path := range endpointPaths {
			go func(path app.Path) {
				start := time.Now()
				statusCode := request(&client, endpoint.Host+path.Value())
				duration := float64(time.Since(start).Milliseconds())
				m.histogram.WithLabelValues(m.jobName, endpointName, path.Name(), fmt.Sprint(statusCode)).Observe(duration)
			}(path)
		}

		time.Sleep(endpoint.Request.Period * time.Second)

	}
}

// PerformRequest to given url
func request(client *http.Client, url string) int {

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return http.StatusRequestTimeout
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
