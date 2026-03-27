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
	// Convert both values to float64 for comparison
	floatA, errA := toFloat64(a)
	if errA != nil {
		return nil
	}

	floatB, errB := toFloat64(b)
	if errB != nil {
		return nil
	}

	// Compare the float values
	if floatA == floatB {
		return nil
	}

	// Return the original value that is larger
	if floatA > floatB {
		return a
	}
	return b
}

// toFloat64 converts an interface{} to float64
// Handles strings (with comma to period replacement), int, and float64 types
func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case string:
		// Replace comma with period for string representation
		val = strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(val, 64)
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}

func main() {
	// Test cases
	fmt.Println(CompareOne(1, 2.5))      // 2.5
	fmt.Println(CompareOne(1, "2,3"))  // 2,3
	fmt.Println(CompareOne("5,1", "6")) // 6
	fmt.Println(CompareOne("1", 1))    // <nil>
}