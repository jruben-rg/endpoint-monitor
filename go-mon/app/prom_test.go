package app

import (
	"testing"
)

func TestPromPort(t *testing.T) {

	prometheus := &Prometheus{}

	tests := []struct {
		inputPort string
		wantPort  string
	}{
		{
			inputPort: ":8080",
			wantPort:  ":8080",
		},
		{
			inputPort: "8080",
			wantPort:  ":8080",
		},
		{
			inputPort: "9090",
			wantPort:  ":9090",
		},
	}

	for _, test := range tests {
		prometheus.SetPort(test.inputPort)
		if prometheus.GetPort() != test.wantPort {
			t.Errorf("Expected: %s. Got: %s.\n", test.wantPort, prometheus.GetPort())
		}
	}

}
