package main

import "fmt"

// WillItFly returns true if the object q will fly, and false otherwise.
// The object q will fly if it's balanced (it is a palindromic slice)
// and the sum of its elements is less than or equal to the maximum possible weight w.
func WillItFly(q []int, w int) bool {
	// Calculate the sum of elements in q
	sum := 0
	for _, val := range q {
		sum += val
	}

	// Check if the sum is greater than the maximum weight
	if sum > w {
		return false
	}

	// Check if the slice is a palindrome using a two-pointer approach
	i, j := 0, len(q)-1
	for i < j {
		if q[i] != q[j] {
			return false
		}
		i++
		j--
	}

	// If both checks pass, it will fly
	return true
}

func main() {
	// Example 1: unbalanced
	fmt.Printf("will_it_fly([1, 2], 5) -> %v\n", WillItFly([]int{1, 2}, 5)) // Expected: false

	// Example 2: overweight
	fmt.Printf("will_it_fly([3, 2, 3], 1) -> %v\n", WillItFly([]int{3, 2, 3}, 1)) // Expected: false

	// Example 3: balanced and within weight
	fmt.Printf("will_it_fly([3, 2, 3], 9) -> %v\n", WillItFly([]int{3, 2, 3}, 9)) // Expected: true

	// Example 4: single element, balanced and within weight
	fmt.Printf("will_it_fly([3], 5) -> %v\n", WillItFly([]int{3}, 5)) // Expected: true
}
