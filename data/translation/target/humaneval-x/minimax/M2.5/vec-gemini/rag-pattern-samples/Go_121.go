package main

func Solution(lst []int) int {
	sum := 0
	for idx, x := range lst {
		if idx%2 == 0 && x%2 == 1 {
			sum += x
		}
	}
	return sum
}
