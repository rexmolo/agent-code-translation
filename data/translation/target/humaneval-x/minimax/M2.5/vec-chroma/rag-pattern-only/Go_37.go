package main

import "fmt"

func SortEven(l []int) []int {
	// Extract elements at even indices (0, 2, 4, ...)
	evens := make([]int, 0, (len(l)+1)/2)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Extract elements at odd indices (1, 3, 5, ...)
	odds := make([]int, 0, len(l)/2)
	for i := 1; i < len(l); i += 2 {
		odds = append(odds, l[i])
	}

	// Sort the even-indexed elements
	// Using insertion sort since we need to sort in place
	for i := 1; i < len(evens); i++ {
		key := evens[i]
		j := i - 1
		for j >= 0 && evens[j] > key {
			evens[j+1] = evens[j]
			j--
		}
		evens[j+1] = key
	}

	// Interleave sorted evens with odds
	ans := make([]int, 0, len(l))
	for i := 0; i < len(evens); i++ {
		ans = append(ans, evens[i])
		if i < len(odds) {
			ans = append(ans, odds[i])
		}
	}

	// If there's one more even element than odd elements, append the last even
	if len(evens) > len(odds) {
		ans = append(ans, evens[len(evens)-1])
	}

	return ans
}

func main() {
	// Test cases from docstring
	fmt.Println(SortEven([]int{1, 2, 3}))       // [1 2 3]
	fmt.Println(SortEven([]int{5, 6, 3, 4}))   // [3 6 5 4]
}
