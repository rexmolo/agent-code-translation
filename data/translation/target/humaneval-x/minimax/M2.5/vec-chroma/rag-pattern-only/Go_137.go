package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Helper function to convert a value (int, float64, or string) to float64
	toFloat := func(v interface{}) (float64, bool) {
		switch val := v.(type) {
		case int:
			return float64(val), true
		case float64:
			return val, true
		case string:
			// Replace comma with dot for string representation of numbers
			val = strings.ReplaceAll(val, ",", ".")
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0, false
			}
			return f, true
		default:
			return 0, false
		}
	}

	// Convert both values to float64 for comparison
	aFloat, aOk := toFloat(a)
	bFloat, bOk := toFloat(b)

	// If either conversion fails, we cannot compare
	if !aOk || !bOk {
		return nil
	}

	// Compare the numeric values
	if aFloat == bFloat {
		return nil
	}

	// Return the larger value in its original type
	if aFloat > bFloat {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(CompareOne(1, 2.5))      // Expected: 2.5
	fmt.Println(CompareOne(1, "2,3"))    // Expected: 2,3
	fmt.Println(CompareOne("5,1", "6")) // Expected: 6
	fmt.Println(CompareOne("1", 1))      // Expected: <nil>
}