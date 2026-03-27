package main

func SumSquares(lst []int) int {
	result := 0
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			result += lst[i] * lst[i]
		} else if i%4 == 0 {
			result += lst[i] * lst[i] * lst[i]
		} else {
			result += lst[i]
		}
	}
	return result
}

func main() {
	// Example test cases
	// fmt.Println(SumSquares([]int{1, 2, 3}))           // Output: 6
	// fmt.Println(SumSquares([]int{}))                  // Output: 0
	// fmt.Println(SumSquares([]int{-1, -5, 2, -1, -5})) // Output: -126
}
