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
	if strings.Contains(value, ".") {
		// Remove trailing zeros
		for strings.HasSuffix(value, "0") {
			value = value[:len(value)-1]
		}
		// If it ends with ".", remove it
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if it ends with .5 (round away from zero)
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Regular rounding using math.Round
	return int(math.Round(num))
}
