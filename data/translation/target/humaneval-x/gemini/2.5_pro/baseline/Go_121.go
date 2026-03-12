package main

// Solution calculates the sum of all odd elements that are in even positions.
func Solution(lst []int) int {
	var sum int
	for index, value := range lst {
		if index%2 == 0 && value%2 != 0 {
			sum += value
		}
	}
	return sum
}
