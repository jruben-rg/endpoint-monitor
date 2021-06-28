package app

import (
	"errors"
	"fmt"
	"time"
)

const (
	defaultTimeout = 3
	defaultPeriod  = 15
)

type Request struct {
	Timeout time.Duration `yaml:"timeout"`
	Period  time.Duration `yaml:"period"`
}

// Site target of the monitoring.
type Endpoint struct {
	Name    Name    `yaml:"name"`
	Host    string  `yaml:"host"`
	Request Request `yaml:"request"`
	Paths   []Path  `yaml:"paths"`
}

func (e *Endpoint) ValidateEndpoint() error {

	if e.Name == "" {
		return errors.New("the endpoint must have a name")
	}

	if e.Host == "" {
		return fmt.Errorf("host for endpoint %s must have a value", e.Name)
	}

	if e.Request.Timeout == 0 {
		e.Request.Timeout = defaultTimeout
	}

	if e.Request.Period == 0 {
		e.Request.Period = defaultPeriod
	}

	if len(e.Paths) == 0 {
		e.Paths = []Path{""}
	}

	return nil
}
