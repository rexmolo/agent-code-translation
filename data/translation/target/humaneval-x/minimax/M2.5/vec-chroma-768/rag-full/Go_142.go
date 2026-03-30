package main

func SumSquares(lst []int) int {
	result := 0
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			result += lst[i] * lst[i]
		} else if i%4 == 0 && i%3 != 0 {
			result += lst[i] * lst[i] * lst[i]
		} else {
			result += lst[i]
		}
	}
	return result
}

func main() {
	// Test cases
	lst1 := []int{1, 2, 3}
	lst2 := []int{}
	lst3 := []int{-1, -5, 2, -1, -5}
	
	println(SumSquares(lst1)) // Expected: 6
	println(SumSquares(lst2)) // Expected: 0
	println(SumSquares(lst3)) // Expected: -126
}
