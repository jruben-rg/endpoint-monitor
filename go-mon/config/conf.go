package config

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/jruben-rg/endpoint-monitor/go-mon/app"
	"gopkg.in/yaml.v2"
)

const metricsPortEnvVar = "METRICS_PORT"

// Config represents the configuration for Prometheus and the Sites to monitor
type AppConfig struct {
	Prometheus app.Prometheus
	Endpoints  []app.Endpoint `yaml:"endpoints"`
}

// NewConfig reads yml config file.
func New(filePath string) (*AppConfig, error) {

	// Read
	appConfig, err := readConfigFile(filePath)
	if err != nil {
		return &AppConfig{}, err
	}

	for _, endpoint := range appConfig.Endpoints {
		if err := endpoint.ValidateEndpoint(); err != nil {
			return &AppConfig{}, err
		}
	}

	promConfig, err := getPromConf()
	if err != nil {
		return &AppConfig{}, err
	}

	appConfig.Prometheus = *promConfig

	return appConfig, nil
}

func getPromConf() (*app.Prometheus, error) {

	promPort := os.Getenv(metricsPortEnvVar)

	if promPort == "" {
		return &app.Prometheus{}, fmt.Errorf("make environment variable %s has been set", metricsPortEnvVar)
	}

	prometheus := &app.Prometheus{}

	prometheus.SetPath(app.PromMetricsPath)
	prometheus.SetPort(promPort)
	prometheus.SetBuckets(app.PromBuckets)

	return prometheus, nil

}

func readConfigFile(fileName string) (*AppConfig, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return &AppConfig{}, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	return read(reader)
}

func read(reader io.Reader) (*AppConfig, error) {

	file, err := io.ReadAll(reader)
	if err != nil {
		return &AppConfig{}, err
	}

	conf := AppConfig{}

	err = yaml.Unmarshal([]byte(file), &conf)
	if err != nil {
		return &AppConfig{}, err
	}

	return &conf, nil
}
