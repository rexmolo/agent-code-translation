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
	// Test cases
	fmt.Println(FilterIntegers([]interface{}{"a", 3.14, 5}))       // Output: [5]
	fmt.Println(FilterIntegers([]interface{}{1, 2, 3, "abc", map[string]int{}, []int{}})) // Output: [1 2 3]
}