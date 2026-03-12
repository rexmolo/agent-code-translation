package main

import (
	"fmt"
	"strconv"
)

/*
Write a function that takes an array of numbers as input and returns
the number of elements in the array that are greater than 10 and both
first and last digits of a number are odd (1, 3, 5, 7, 9).
For example:
Specialfilter([15, -73, 14, -15]) => 1 
Specialfilter([33, -2, -3, 45, 21, 109]) => 2
*/
func Specialfilter(nums []int) int {
	count := 0
	// Use a map as a set for efficient lookup of odd digit characters.
	oddDigits := map[byte]bool{
		'1': true, '3': true, '5': true, '7': true, '9': true,
	}

	for _, num := range nums {
		if num > 10 {
			numberAsString := strconv.Itoa(num)
			firstDigit := numberAsString[0]
			lastDigit := numberAsString[len(numberAsString)-1]

			// The map lookup returns the zero value (false) for keys not present.
			if oddDigits[firstDigit] && oddDigits[lastDigit] {
				count++
			}
		}
	}

	return count
}

// main function to demonstrate the usage of Specialfilter
func main() {
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109}))
}
