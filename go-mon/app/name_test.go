package app

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"already_snake", "already_snake"},
		{"A", "a"},
		{"AA", "aa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"iPhone", "i_phone"},
		{"Id0Value", "id0_value"},
		{"ID0Value", "id0_value"},
		{"something else", "something_else"},
		{"SOMETHING ELSE", "something_else"},
	}
	for _, test := range tests {
		name := Name(test.input)
		got := name.ToSnakeCase()
		if got != test.want {
			t.Errorf("input=%q:\nGot: %q\nWant: %q", test.input, got, test.want)
		}
	}
}
