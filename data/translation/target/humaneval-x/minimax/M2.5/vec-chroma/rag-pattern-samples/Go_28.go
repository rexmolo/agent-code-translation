package main

import (
	"fmt"
	"strings"
)

func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

func main() {
	// Test cases
	fmt.Println(Concatenate([]string{}))
	fmt.Println(Concatenate([]string{"a", "b", "c"}))
}