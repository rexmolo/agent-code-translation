package main

import "fmt"

func WillItFly(q []int, w int) bool {
	// Check if sum of elements is less than or equal to max weight w
	sum := 0
	for _, v := range q {
		sum += v
	}
	if sum > w {
		return false
	}

	// Check if the list is balanced (palindromic)
	i, j := 0, len(q)-1
	for i < j {
		if q[i] != q[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(WillItFly([]int{1, 2}, 5))   // false - unbalanced
	fmt.Println(WillItFly([]int{3, 2, 3}, 1)) // false - too heavy
	fmt.Println(WillItFly([]int{3, 2, 3}, 9)) // true
	fmt.Println(WillItFly([]int{3}, 5))       // true
}
