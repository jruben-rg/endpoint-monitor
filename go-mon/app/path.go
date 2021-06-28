package app

import "strings"

type Path string

func (p *Path) Name() string {

	if *p == "" {
		return "/"
	}

	if !strings.HasPrefix(string(*p), "/") {
		return "/" + string(*p)
	}

	return string(*p)
}

func (p *Path) Value() string {

	return string(*p)
}
