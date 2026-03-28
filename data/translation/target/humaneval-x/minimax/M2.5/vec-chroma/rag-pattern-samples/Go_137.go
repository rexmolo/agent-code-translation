package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert a to float64
	var floatA float64
	switch val := a.(type) {
	case string:
		val = strings.ReplaceAll(val, ",", ".")
		f, _ := strconv.ParseFloat(val, 64)
		floatA = f
	case int:
		floatA = float64(val)
	case float64:
		floatA = val
	}

	// Convert b to float64
	var floatB float64
	switch val := b.(type) {
	case string:
		val = strings.ReplaceAll(val, ",", ".")
		f, _ := strconv.ParseFloat(val, 64)
		floatB = f
	case int:
		floatB = float64(val)
	case float64:
		floatB = val
	}

	// Compare and return result
	if floatA == floatB {
		return nil
	}
	if floatA > floatB {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(CompareOne(1, 2.5))
	fmt.Println(CompareOne(1, "2,3"))
	fmt.Println(CompareOne("5,1", "6"))
	fmt.Println(CompareOne("1", 1))
}