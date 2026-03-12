package main

import "fmt"

// FilterIntegers filters a given slice of any Go values only for integers.
func FilterIntegers(values []interface{}) []int {
	var integers []int
	for _, value := range values {
		// Use a type assertion to check if the value is an int.
		if intValue, ok := value.(int); ok {
			integers = append(integers, intValue)
		}
	}
	return integers
}

// main is a driver function to demonstrate FilterIntegers.
func main() {
	// Mimics: filter_integers(['a', 3.14, 5])
	slice1 := []interface{}{"a", 3.14, 5}
	fmt.Println(FilterIntegers(slice1))

	// Mimics: filter_integers([1, 2, 3, 'abc', {}, []])
	slice2 := []interface{}{1, 2, 3, "abc", struct{}{}, []string{}}
	fmt.Println(FilterIntegers(slice2))
}
