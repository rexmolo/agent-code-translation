package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Handle empty string
	if len(value) == 0 {
		return 0
	}

	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		value = strings.TrimRight(value, "0")
		// Remove trailing decimal point if it becomes just "."
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if value ends with ".5" - round away from zero
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Standard rounding (round half away from zero)
	return int(math.Round(num))
}