package main

import (
	"fmt"
)

// MaxElement returns the maximum element in a slice of integers.
// It assumes the slice is not empty. If the slice is empty, the function
// will panic, which is equivalent to the IndexError in the original Python code.
func MaxElement(l []int) int {
	m := l[0]
	for _, e := range l {
		if e > m {
			m = e
		}
	}
	return m
}

// main function to demonstrate the MaxElement function with examples from the Python docstring.
func main() {
	list1 := []int{1, 2, 3}
	fmt.Println(MaxElement(list1))

	list2 := []int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10}
	fmt.Println(MaxElement(list2))
}
