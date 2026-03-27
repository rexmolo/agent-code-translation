package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert both values to float64 for comparison
	floatA, errA := toFloat64(a)
	floatB, errB := toFloat64(b)

	// If either conversion fails, return nil (shouldn't happen with valid input)
	if errA != nil || errB != nil {
		return nil
	}

	// If values are equal, return nil
	if floatA == floatB {
		return nil
	}

	// Return the larger value in its original type
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
		// Replace commas with dots (as per Python code)
		s := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(s, 64)
	default:
		return 0, nil
	}
}
