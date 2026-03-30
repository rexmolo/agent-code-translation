package main

import (
	"fmt"
)

func FilterIntegers(values []interface{}) []int {
	result := make([]int, 0, len(values))
	for _, v := range values {
		if i, ok := v.(int); ok {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// Test examples
	example1 := []interface{}{"a", 3.14, 5}
	example2 := []interface{}{1, 2, 3, "abc", map[string]interface{}{}, []interface{}{}}

	fmt.Println(FilterIntegers(example1))
	fmt.Println(FilterIntegers(example2))
}
