package main

import "fmt"

// CountUpper counts the number of uppercase vowels in even indices of a string.
// For example:
// CountUpper("aBCdEf") returns 1
// CountUpper("abcdefg") returns 0
// CountUpper("dBBE") returns 0
func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		switch s[i] {
		case 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

// main function to demonstrate the usage of CountUpper
func main() {
	fmt.Println(CountUpper("aBCdEf"))
	fmt.Println(CountUpper("abcdefg"))
	fmt.Println(CountUpper("dBBE"))
}
