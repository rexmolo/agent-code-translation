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
		// Also remove trailing dot if all zeros were removed
		if strings.HasSuffix(value, ".") {
			value = strings.TrimSuffix(value, ".")
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Check if ends with .5 - round away from zero
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// For other cases, use math.Round
	return int(math.Round(num))
}

func main() {
	// Test examples
	fmt.Println(ClosestInteger("10"))   // 10
	fmt.Println(ClosestInteger("15.3")) // 15
	fmt.Println(ClosestInteger("14.5")) // 15
	fmt.Println(ClosestInteger("-14.5")) // -15
}