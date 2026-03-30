package main

import (
	"strings"
)

// Concatenate concatenates a slice of strings into a single string.
func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}