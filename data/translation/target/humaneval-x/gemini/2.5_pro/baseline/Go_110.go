package main

import "fmt"

// Exchange determines if it is possible to perform an exchange of elements
// between lst1 and lst2 to make lst1 a list of only even numbers.
// There is no limit on the number of exchanged elements.
// It returns "YES" if possible, otherwise "NO".
func Exchange(lst1, lst2 []int) string {
	// Count the number of odd numbers in lst1.
	// These must be swapped out for even numbers.
	oddCount := 0
	for _, num := range lst1 {
		if num%2 != 0 {
			oddCount++
		}
	}

	// Count the number of even numbers in lst2.
	// These are the available numbers to swap into lst1.
	evenCount := 0
	for _, num := range lst2 {
		if num%2 == 0 {
			evenCount++
		}
	}

	// If there are enough even numbers in lst2 to replace all
	// the odd numbers in lst1, the exchange is possible.
	if evenCount >= oddCount {
		return "YES"
	}

	return "NO"
}

// main function to run example cases
func main() {
	// Example 1: exchange([1, 2, 3, 4], [1, 2, 3, 4]) => "YES"
	lst1a := []int{1, 2, 3, 4}
	lst2a := []int{1, 2, 3, 4}
	fmt.Println(Exchange(lst1a, lst2a))

	// Example 2: exchange([1, 2, 3, 4], [1, 5, 3, 4]) => "NO"
	lst1b := []int{1, 2, 3, 4}
	lst2b := []int{1, 5, 3, 4}
	fmt.Println(Exchange(lst1b, lst2b))
}