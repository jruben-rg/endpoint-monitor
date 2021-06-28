package config

import (
	"bytes"
	"github.com/jruben-rg/endpoint-monitor/go-mon/app"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"testing"
)

// var emptyPromConfig = app.Prometheus{
// 	// path:    "",
// 	// port:    "",
// 	Buckets: []float64{},
// }

// var promConfig = &app.Prometheus{}

// var customPromConfig = app.Prometheus{
// 	// Path:    "/metrics",
// 	// Port:    "8080",
// 	Buckets: []float64{100, 200, 300},
// }

var customRequest = app.Request{
	Timeout: 2,
	Period:  3,
}

var endpointWithNoPaths = app.Endpoint{
	Name:    "No paths endpoint",
	Host:    "http://www.test.com",
	Request: customRequest,
	Paths:   []app.Path{""},
}

var endpointWithOnePath = app.Endpoint{
	Name:    "One path endpoint",
	Host:    "http://www.test.com",
	Request: customRequest,
	Paths:   []app.Path{"/one"},
}

var endpointWithTwoPaths = app.Endpoint{
	Name:    "Two paths endpoint",
	Host:    "http://www.test.com",
	Request: customRequest,
	Paths:   []app.Path{"/one", "/two"},
}

func TestPrometheusConfigError(t *testing.T) {

	os.Setenv(metricsPortEnvVar, "")

	_, err := getPromConf()

	if err == nil {
		t.Errorf("Expecting error if Environment Variable for Prometheus Port is not set")
	}

}

func TestPrometheusConfig(t *testing.T) {

	testPort := "8080"

	os.Setenv(metricsPortEnvVar, testPort)
	wantPromConf := &app.Prometheus{}
	wantPromConf.SetPath(app.PromMetricsPath)
	wantPromConf.SetBuckets(app.PromBuckets)
	wantPromConf.SetPort(testPort)

	gotPromConf, err := getPromConf()

	if !reflect.DeepEqual(wantPromConf, gotPromConf) || err != nil {
		t.Errorf("Scenario: Prometheus Configuration.\nWanted: %#v\nGot: %#v\n", wantPromConf, gotPromConf)
	}

}

func TestEndpointsConfig(t *testing.T) {

	tests := []struct {
		scenario   string
		yamlConfig AppConfig
		wantConfig AppConfig
	}{
		{
			scenario: "Empty prometheus configuration should result in default prometheus configuration",
			yamlConfig: AppConfig{
				Endpoints: []app.Endpoint{},
			},
			wantConfig: AppConfig{
				Endpoints: []app.Endpoint{},
			},
		},
		{
			scenario: "Empty prometheus configuration with one path endpoint",
			yamlConfig: AppConfig{
				// Prometheus: emptyPromConfig,
				Endpoints: []app.Endpoint{endpointWithOnePath},
			},
			wantConfig: AppConfig{
				// Prometheus: *defaultPromConfig,
				Endpoints: []app.Endpoint{endpointWithOnePath},
			},
		},
		{
			scenario: "Empty prometheus configuration with two paths endpoint",
			yamlConfig: AppConfig{
				// Prometheus: emptyPromConfig,
				Endpoints: []app.Endpoint{endpointWithTwoPaths},
			},
			wantConfig: AppConfig{
				// Prometheus: *defaultPromConfig,
				Endpoints: []app.Endpoint{endpointWithTwoPaths},
			},
		},
		{
			scenario: "Default Prometheus configuration with multiple endpoints",
			yamlConfig: AppConfig{
				// Prometheus: emptyPromConfig,
				Endpoints: []app.Endpoint{endpointWithNoPaths, endpointWithOnePath, endpointWithTwoPaths},
			},
			wantConfig: AppConfig{
				// Prometheus: *defaultPromConfig,
				Endpoints: []app.Endpoint{endpointWithNoPaths, endpointWithOnePath, endpointWithTwoPaths},
			},
		},
		{
			scenario: "Custom Prometheus configuration with multiple endpoints",
			yamlConfig: AppConfig{
				// Prometheus: *defaultPromConfig,
				Endpoints: []app.Endpoint{endpointWithNoPaths, endpointWithOnePath, endpointWithTwoPaths},
			},
			wantConfig: AppConfig{
				// Prometheus: *defaultPromConfig,
				Endpoints: []app.Endpoint{endpointWithNoPaths, endpointWithOnePath, endpointWithTwoPaths},
			},
		},
	}

	for _, test := range tests {

		wantSerialized, err := yaml.Marshal(&test.yamlConfig)
		if err != nil {
			t.Error("failed to serialize application configuration")
		}

		var buffer bytes.Buffer
		buffer.Write(wantSerialized)
		gotConfig, err := read(&buffer)

		if err != nil {
			t.Error("Failed to read application configuration.")
		}

		// if !reflect.DeepEqual(test.wantConfig.Prometheus, gotConfig.Prometheus) {
		// 	t.Errorf("Scenario: %s.\nWanted: %#v\nGot: %#v\n", test.scenario, test.wantConfig.Prometheus, gotConfig.Prometheus)
		// }

		if !reflect.DeepEqual(test.wantConfig.Endpoints, gotConfig.Endpoints) {
			t.Errorf("Scenario: %s.\nWanted: %#v\nGot: %#v\n", test.scenario, test.wantConfig.Endpoints, gotConfig.Endpoints)
		}

		// if test.wantConfig.Prometheus != gotConfig.Prometheus {
		// 	t.Errorf("Scenario: %s.\nWanted: %#v\nGot: %#v\n", test.scenario, test.wantConfig.Prometheus, gotConfig.Prometheus)
		// }
	}

}
