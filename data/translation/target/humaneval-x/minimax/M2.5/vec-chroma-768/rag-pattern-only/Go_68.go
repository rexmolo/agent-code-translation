package main

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	minEven := 0
	minIndex := -1
	found := false

	for i, v := range arr {
		if v%2 == 0 {
			if !found || v < minEven {
				minEven = v
				minIndex = i
				found = true
			}
		}
	}

	if !found {
		return []int{}
	}

	return []int{minEven, minIndex}
}

func main() {
	// Test cases
	// Example 1: [4,2,3] -> [2, 1]
	// Example 2: [1,2,3] -> [2, 1]
	// Example 3: [] -> []
	// Example 4: [5, 0, 3, 0, 4, 2] -> [0, 1]
	_ = [][]int{
		Pluck([]int{4, 2, 3}),
		Pluck([]int{1, 2, 3}),
		Pluck([]int{}),
		Pluck([]int{5, 0, 3, 0, 4, 2}),
	}
}
