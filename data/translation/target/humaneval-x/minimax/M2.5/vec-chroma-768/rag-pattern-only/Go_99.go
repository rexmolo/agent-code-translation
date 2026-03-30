package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	if len(value) == 0 {
		return 0
	}

	// Check if original value ends with .5 before removing trailing zeros
	hasDotFive := strings.HasSuffix(value, ".5")

	// Remove trailing zeros
	for len(value) > 0 && value[len(value)-1] == '0' {
		value = value[:len(value)-1]
	}
	// Remove trailing decimal point if no digits after
	if len(value) > 0 && value[len(value)-1] == '.' {
		value = value[:len(value)-1]
	}

	// Special case: round .5 away from zero
	if hasDotFive {
		num, _ := strconv.ParseFloat(value, 64)
		if num >= 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Standard rounding (round half away from zero)
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return int(math.Round(num))
}

func main() {
	// Test cases for manual verification
	println(ClosestInteger("10"))    // 10
	println(ClosestInteger("15.3")) // 15
	println(ClosestInteger("14.5")) // 15
	println(ClosestInteger("-14.5")) // -15
}