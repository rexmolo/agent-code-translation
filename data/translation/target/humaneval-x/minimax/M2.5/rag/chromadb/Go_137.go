package main

import (
	"strconv"
	"strings"
)

func CompareOne(a, b interface{}) interface{} {
	// Helper function to convert interface{} to float64
	toFloat := func(v interface{}) float64 {
		switch x := v.(type) {
		case string:
			s := strings.ReplaceAll(x, ",", ".")
			f, _ := strconv.ParseFloat(s, 64)
			return f
		case int:
			return float64(x)
		case int8:
			return float64(x)
		case int16:
			return float64(x)
		case int32:
			return float64(x)
		case int64:
			return float64(x)
		case float32:
			return float64(x)
		case float64:
			return x
		}
		return 0
	}

	floatA := toFloat(a)
	floatB := toFloat(b)

	if floatA == floatB {
		return nil
	}

	if floatA > floatB {
		return a
	}
	return b
}
