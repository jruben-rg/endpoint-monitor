package app

import (
	"regexp"
	"strings"
)

type Name string

// ToSnakeCase transforms a string into snake case.
func (n *Name) ToSnakeCase() string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	replacer := strings.NewReplacer(" ", "_")

	snake := matchFirstCap.ReplaceAllString(string(*n), "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return replacer.Replace(strings.ToLower(snake))
}
