package main

import (
	"fmt"
)

// OddCount translates the Python odd_count function.
// Given a slice of strings, where each string consists of only digits, it returns a new slice.
// Each element in the output corresponds to an input string and describes the count of odd digits
// found in that string, embedded into a template sentence.
func OddCount(lst []string) []string {
	res := make([]string, 0, len(lst))
	for _, s := range lst {
		n := 0
		// Iterate over each character (rune) of the string.
		for _, r := range s {
			// Convert rune '0'-'9' to int 0-9 and check if it's odd.
			// This works because rune is an alias for int32 and the operation
			// is performed on their integer values.
			if (r-'0')%2 == 1 {
				n++
			}
		}
		// Format the output string using the count n.
		formattedString := fmt.Sprintf("the number of odd elements %dn the str%dng %d of the %dnput.", n, n, n, n)
		res = append(res, formattedString)
	}
	return res
}

// main function to make the code runnable and demonstrate OddCount.
func main() {
	// Example 1 from the original docstring
	fmt.Println(OddCount([]string{"1234567"}))
	// Expected output: [the number of odd elements 4n the str4ng 4 of the 4nput.]

	// Example 2 from the original docstring
	fmt.Println(OddCount([]string{"3", "11111111"}))
	// Expected output: [the number of odd elements 1n the str1ng 1 of the 1nput. the number of odd elements 8n the str8ng 8 of the 8nput.]
}
