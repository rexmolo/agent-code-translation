package main

import (
	"bufio"
	"fmt"
	"os"
)

func Search(lst []int) int {
	if len(lst) == 0 {
		return -1
	}

	// Find the maximum value in the list
	maxVal := lst[0]
	for _, v := range lst {
		if v > maxVal {
			maxVal = v
		}
	}

	// Create frequency array
	frq := make([]int, maxVal+1)
	for _, v := range lst {
		frq[v]++
	}

	// Find the greatest integer i such that frq[i] >= i
	ans := -1
	for i := 1; i < len(frq); i++ {
		if frq[i] >= i {
			ans = i
		}
	}

	return ans
}

func main() {
	// Example usage - read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// Parse input (simplified for demonstration)
	// This is a placeholder - actual usage would parse the list
	fmt.Println("Use Search function with actual input")
}