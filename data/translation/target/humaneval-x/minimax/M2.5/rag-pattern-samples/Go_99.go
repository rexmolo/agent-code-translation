package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Handle empty string case
	if len(value) == 0 {
		return 0
	}

	// Remove trailing zeros after decimal point
	if strings.Count(value, ".") == 1 {
		// Remove trailing zeros
		value = strings.TrimRight(value, "0")
		// Also remove trailing decimal point if all zeros were removed
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if the trimmed value ends with .5 (for rounding away from zero)
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Standard rounding (round half away from zero in Go)
	return int(math.Round(num))
}
