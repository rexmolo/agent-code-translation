package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert both values to float64 for comparison
	floatA, err := toFloat64(a)
	if err != nil {
		return nil
	}
	floatB, err := toFloat64(b)
	if err != nil {
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
		// Replace comma with dot for string representation
		s := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(s, 64)
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}

func main() {
	// Test examples
	fmt.Println(CompareOne(1, 2.5))      // Expected: 2.5
	fmt.Println(CompareOne(1, "2,3"))     // Expected: 2,3
	fmt.Println(CompareOne("5,1", "6")) // Expected: 6
	fmt.Println(CompareOne("1", 1))      // Expected: <nil>
}