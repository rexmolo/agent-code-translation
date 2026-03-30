package main

import (
	"math"
	"strconv"
	"strings"
)

// ClosestInteger takes a string representing a number and returns
// the closest integer. If the number is equidistant from two integers,
// it rounds away from zero.
func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		value = strings.TrimRight(value, "0")
		// Remove trailing decimal point if no digits after it
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	// Handle empty string case
	if len(value) == 0 {
		return 0
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if the number ends with ".5" - round away from zero
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		}
		return int(math.Floor(num))
	}

	// Otherwise, use standard rounding (math.Round uses round-to-nearest)
	return int(math.Round(num))
}

func main() {
	// Example usage
	println(ClosestInteger("10"))
	println(ClosestInteger("15.3"))
	println(ClosestInteger("14.5"))
	println(ClosestInteger("-14.5"))
}