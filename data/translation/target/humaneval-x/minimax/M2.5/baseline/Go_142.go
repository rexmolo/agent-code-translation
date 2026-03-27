package main

import "fmt"

func SumSquares(lst []int) int {
	result := make([]int, 0, len(lst))
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			result = append(result, lst[i]*lst[i])
		} else if i%4 == 0 && i%3 != 0 {
			result = append(result, lst[i]*lst[i]*lst[i])
		} else {
			result = append(result, lst[i])
		}
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(SumSquares([]int{1, 2, 3}))           // 6
	fmt.Println(SumSquares([]int{}))                 // 0
	fmt.Println(SumSquares([]int{-1, -5, 2, -1, -5})) // -126
}