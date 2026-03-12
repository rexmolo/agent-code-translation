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

func main() {
	// Example usage
	lst1 := []int{5, 8, 7, 1}
	lst2 := []int{3, 3, 3, 3, 3}
	lst3 := []int{30, 13, 24, 321}
	println(Solution(lst1)) // 12
	println(Solution(lst2)) // 9
	println(Solution(lst3)) // 0
}