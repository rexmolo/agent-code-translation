package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareOne takes integers, floats, or strings representing real numbers,
// and returns the larger variable in its given variable type.
// Returns nil if the values are equal.
// Note: If a real number is represented as a string, the floating point might be . or ,
func CompareOne(a, b interface{}) interface{} {
	// Convert a to float64 if possible
	floatA, errA := toFloat64(a)
	floatB, errB := toFloat64(b)

	// If either conversion fails, return nil
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

// toFloat64 converts various types to float64 for comparison
func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		// Replace comma with period for string representation of numbers
		s := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(s, 64)
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// Test function to verify the implementation
func main() {
	// Test cases
	fmt.Println(CompareOne(1, 2.5))       // Expected: 2.5
	fmt.Println(CompareOne(1, "2,3"))     // Expected: 2,3 (as string)
	fmt.Println(CompareOne("5,1", "6")) // Expected: 6
	fmt.Println(CompareOne("1", 1))      // Expected: nil
}