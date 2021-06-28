package app

import (
	"reflect"
	"testing"
)

func TestValidateEndpoint(t *testing.T) {

	tests := []struct {
		scenario string
		given    Endpoint
		want     Endpoint
	}{
		{
			scenario: "Should Validate and initialise paths for an endpoint",
			given: Endpoint{
				Name: "test",
				Host: "http://www.test.com",
				Request: Request{
					Timeout: 5,
					Period:  20,
				},
				Paths: []Path{},
			},
			want: Endpoint{
				Name: "test",
				Host: "http://www.test.com",
				Request: Request{
					Timeout: 5,
					Period:  20,
				},
				Paths: []Path{""},
			},
		},
		{
			scenario: "Should set up default values for the request",
			given: Endpoint{
				Name:    "test",
				Host:    "http://www.test.com",
				Request: Request{},
				Paths:   []Path{"/test"},
			},
			want: Endpoint{
				Name: "test",
				Host: "http://www.test.com",
				Request: Request{
					Timeout: 3,
					Period:  15,
				},
				Paths: []Path{"/test"},
			},
		},
	}

	for _, test := range tests {

		err := test.given.ValidateEndpoint()

		if err != nil {
			t.Errorf("scenario: %s:\nGot unexpected error: %s\n", test.scenario, err)
		}

		if !reflect.DeepEqual(test.given, test.want) {
			t.Errorf("scenario: %s:\nGot: %q\nWant: %q", test.scenario, test.given, test.want)
		}

	}

}

func TestNoValidEndpoint(t *testing.T) {

	tests := []struct {
		scenario string
		given    Endpoint
	}{
		{
			scenario: "Endpoint should not be valid if name is not present",
			given: Endpoint{
				Name:    "",
				Host:    "http://www.test.com",
				Request: Request{},
				Paths:   []Path{},
			},
		},
		{
			scenario: "Endpoint should not be valid if host is not present",
			given: Endpoint{
				Name:    "test",
				Host:    "",
				Request: Request{},
				Paths:   []Path{""},
			},
		},
	}

	for _, test := range tests {

		err := test.given.ValidateEndpoint()

		if err == nil {
			t.Errorf("Expecting error for scenario: %s:\n", test.scenario)
		}
	}

}
