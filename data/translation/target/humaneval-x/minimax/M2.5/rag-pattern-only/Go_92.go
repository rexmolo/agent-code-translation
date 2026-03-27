package main

import (
	"fmt"
)

func AnyInt(x, y, z interface{}) bool {
	// Convert all three values to int64 to check if they're integers
	ix, ok1 := toInt64(x)
	if !ok1 {
		return false
	}

	iy, ok2 := toInt64(y)
	if !ok2 {
		return false
	}

	iz, ok3 := toInt64(z)
	if !ok3 {
		return false
	}

	// Check if one equals the sum of the other two
	if ix+iy == iz || ix+iz == iy || iy+iz == ix {
		return true
	}

	return false
}

// Helper function to convert interface{} to int64 if it's an integer type
func toInt64(v interface{}) (int64, bool) {
	switch val := v.(type) {
	case int:
		return int64(val), true
	case int8:
		return int64(val), true
	case int16:
		return int64(val), true
	case int32:
		return int64(val), true
	case int64:
		return val, true
	case uint:
		return int64(val), true
	case uint8:
		return int64(val), true
	case uint16:
		return int64(val), true
	case uint32:
		return int64(val), true
	case uint64:
		return int64(val), true
	default:
		return 0, false
	}
}

func main() {
	// Test cases
	fmt.Println(AnyInt(5, 2, 7))   // true
	fmt.Println(AnyInt(3, 2, 2))  // false
	fmt.Println(AnyInt(3, -2, 1)) // true
	fmt.Println(AnyInt(3.6, -2.2, 2)) // false
}
