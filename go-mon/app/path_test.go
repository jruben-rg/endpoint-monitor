package app

import (
	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct {
		scenario  string
		inputPath Path
		wantName  string
		wantValue string
	}{
		{
			"Empty path",
			Path(""),
			"/",
			"",
		},
		{
			"Path does not start with slash",
			Path("test"),
			"/test",
			"test",
		},
		{
			"Path starts with slash",
			Path("/test"),
			"/test",
			"/test",
		},
	}

	for _, test := range tests {
		gotPath := test.inputPath
		if gotPath.Name() != test.wantName {
			t.Errorf("Scenario: %s.\nGot %s. Expected Path Name: %s.", test.scenario, gotPath.Name(), test.wantName)
		}

		if gotPath.Value() != test.wantValue {
			t.Errorf("Scenario: %s.\nGot %s. Expected Path Value: %s.", test.scenario, gotPath.Name(), test.wantValue)
		}

	}

}
