package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Helper function to convert interface{} to float64 for comparison
	toFloat := func(v interface{}) (float64, bool) {
		switch val := v.(type) {
		case int:
			return float64(val), true
		case float64:
			return val, true
		case string:
			s := strings.ReplaceAll(val, ",", ".")
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, false
			}
			return f, true
		default:
			return 0, false
		}
	}

	fa, okA := toFloat(a)
	if !okA {
		return nil
	}

	fb, okB := toFloat(b)
	if !okB {
		return nil
	}

	if fa == fb {
		return nil
	}

	if fa > fb {
		return a
	}
	return b
}