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
	// Example usage
	println(SumSquares([]int{1, 2, 3}))        // Output: 6
	println(SumSquares([]int{}))                // Output: 0
	println(SumSquares([]int{-1, -5, 2, -1, -5})) // Output: -126
}
