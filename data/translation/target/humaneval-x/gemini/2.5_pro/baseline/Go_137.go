package main

import (
	"fmt"
	"strconv"
	"strings"
)

// toFloat64 is a helper function that converts a value of type int, float, or
// string into a float64 for comparison. It returns an error for unsupported types.
func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		// Replace comma decimal separator with a period.
		s := strings.ReplaceAll(val, ",", ".")
		return strconv.ParseFloat(s, 64)
	default:
		return 0, fmt.Errorf("unsupported type for comparison: %T", v)
	}
}

// CompareOne takes integers, floats, or strings representing real numbers,
// and returns the larger variable in its given variable type.
// It returns nil if the values are equal.
func CompareOne(a, b interface{}) interface{} {
	valA, errA := toFloat64(a)
	if errA != nil {
		// The original Python code would raise an exception on invalid conversion.
		// Panic is the Go equivalent for an unrecoverable error.
		panic(fmt.Sprintf("could not convert 'a' to float: %v", errA))
	}

	valB, errB := toFloat64(b)
	if errB != nil {
		panic(fmt.Sprintf("could not convert 'b' to float: %v", errB))
	}

	if valA == valB {
		return nil
	}

	if valA > valB {
		return a
	}

	return b
}

// main function to demonstrate the CompareOne function with examples.
func main() {
	fmt.Println("--- Demonstrating CompareOne function ---")

	// Example 1: compare_one(1, 2.5) -> 2.5
	res1 := CompareOne(1, 2.5)
	fmt.Printf("CompareOne(1, 2.5) -> %v (type: %T)\n", res1, res1)

	// Example 2: compare_one(1, "2,3") -> "2,3"
	res2 := CompareOne(1, "2,3")
	fmt.Printf("CompareOne(1, \"2,3\") -> %q (type: %T)\n", res2, res2)

	// Example 3: compare_one("5,1", "6") -> "6"
	res3 := CompareOne("5,1", "6")
	fmt.Printf("CompareOne(\"5,1\", \"6\") -> %q (type: %T)\n", res3, res3)

	// Example 4: compare_one("1", 1) -> None
	res4 := CompareOne("1", 1)
	fmt.Printf("CompareOne(\"1\", 1) -> %v (type: %T)\n", res4, res4)
}