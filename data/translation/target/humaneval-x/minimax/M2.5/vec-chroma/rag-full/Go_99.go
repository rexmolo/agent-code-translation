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

	// Create a copy to manipulate
	val := value

	// Remove trailing zeros after decimal point
	if strings.Contains(val, ".") {
		// Remove trailing zeros
		val = strings.TrimRight(val, "0")
		// Remove trailing dot if it's the last character
		if strings.HasSuffix(val, ".") {
			val = val[:len(val)-1]
		}
	}

	// Convert to float64
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}

	// Check if value ends with .5 (exact half)
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			// Round away from zero for positive numbers
			return int(math.Ceil(num))
		} else {
			// Round away from zero for negative numbers
			return int(math.Floor(num))
		}
	}

	// Regular rounding using Go's Round (round half away from zero)
	res := int(math.Round(num))
	return res
}

func main() {
	// Test cases can be added here for verification
	fmt.Println(ClosestInteger("10"))      // 10
	fmt.Println(ClosestInteger("15.3"))    // 15
	fmt.Println(ClosestInteger("14.5"))   // 15
	fmt.Println(ClosestInteger("-14.5"))  // -15
}
