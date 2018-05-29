package main

import (
	"regexp"
)

var (
	reTitle = regexp.MustCompile(`(?i)<title>([^<>]+)</title>`)
)

func getTitle(s string) string {
	if m := reTitle.FindAllStringSubmatch(s, 1); m != nil {
		return m[0][1]
	}
	return "Noo title"
}
