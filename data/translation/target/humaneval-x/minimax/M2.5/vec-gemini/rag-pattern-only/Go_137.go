package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Convert inputs to float64 for comparison
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
		// Replace comma with period for European decimal format
		str := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(str, 64)
	default:
		return 0, nil
	}
}