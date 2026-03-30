package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	floatA, okA := toFloat64(a)
	floatB, okB := toFloat64(b)

	if !okA || !okB {
		return nil
	}

	if floatA == floatB {
		return nil
	}

	if floatA > floatB {
		return a
	}
	return b
}

func toFloat64(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case int:
		return float64(val), true
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case string:
		cleaned := strings.ReplaceAll(val, ",", ".")
		f, err := strconv.ParseFloat(cleaned, 64)
		if err != nil {
			return 0, false
		}
		return f, true
	default:
		return 0, false
	}
}