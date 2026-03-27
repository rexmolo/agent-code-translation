package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		for strings.HasSuffix(value, "0") {
			value = strings.TrimSuffix(value, "0")
		}
	}

	// Handle empty string
	if len(value) == 0 {
		return 0
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// If ends with .5, round away from zero
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Otherwise use standard rounding
	res := int(math.Round(num))
	return res
}

func main() {
	// Test cases can be added here for verification
}