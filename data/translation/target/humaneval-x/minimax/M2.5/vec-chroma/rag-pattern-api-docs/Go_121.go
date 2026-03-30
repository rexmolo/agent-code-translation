package main

func Solution(lst []int) int {
	sum := 0
	for idx, x := range lst {
		// idx%2 == 0: even position (0, 2, 4, ...)
		// x%2 == 1: odd value
		if idx%2 == 0 && x%2 == 1 {
			sum += x
		}
	}
	return sum
}
