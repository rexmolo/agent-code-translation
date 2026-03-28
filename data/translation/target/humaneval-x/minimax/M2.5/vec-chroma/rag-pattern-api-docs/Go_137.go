package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert both values to float64 for comparison
	floatA, errA := toFloat64(a)
	floatB, errB := toFloat64(b)
	
	// If either conversion fails, we cannot compare
	if errA != nil || errB != nil {
		return nil
	}
	
	// Compare the float values
	if floatA == floatB {
		return nil
	}
	
	if floatA > floatB {
		return a
	}
	return b
}

// Helper function to convert various types to float64
func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		// Replace comma with period for string representation of numbers
		cleaned := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(cleaned, 64)
	default:
		return 0, nil
	}
}

func main() {
	// Test cases
	println(CompareOne(1, 2.5))       // Expected: 2.5
	println(CompareOne(1, "2,3"))    // Expected: 2,3
	println(CompareOne("5,1", "6")) // Expected: 6
	println(CompareOne("1", 1))      // Expected: <nil>
}