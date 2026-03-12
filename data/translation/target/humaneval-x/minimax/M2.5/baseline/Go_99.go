package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		value = strings.TrimRight(value, "0")
		// Remove trailing decimal point if no digits after
		if strings.HasSuffix(value, ".") {
			value = strings.TrimRight(value, ".")
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if the string ends with .5 (round away from zero)
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	} else if len(value) > 0 {
		// Standard rounding - Go's math.Round rounds half away from zero
		return int(math.Round(num))
	}

	return 0
}