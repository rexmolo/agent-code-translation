package main

import "fmt"

// IncrList returns a new slice with each element of the input slice incremented by 1.
func IncrList(l []int) []int {
	// The Python list comprehension creates a new list, so we create a new slice in Go.
	// Pre-allocating the slice with make is more efficient than appending in a loop.
	result := make([]int, len(l))

	// Iterate over the input slice using a for...range loop.
	for i, e := range l {
		// Increment the element and assign it to the new slice.
		result[i] = e + 1
	}

	return result
}

// main function to demonstrate the usage of IncrList
func main() {
	// Example 1 from docstring: incr_list([1, 2, 3]) -> [2, 3, 4]
	list1 := []int{1, 2, 3}
	fmt.Println(IncrList(list1))

	// Example 2 from docstring: incr_list([5, 3, 5, 2, 3, 3, 9, 0, 123]) -> [6, 4, 6, 3, 4, 4, 10, 1, 124]
	list2 := []int{5, 3, 5, 2, 3, 3, 9, 0, 123}
	fmt.Println(IncrList(list2))
}
