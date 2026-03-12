package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Minsubarraysum(nums []int) int {
	maxSum := 0
	s := 0
	for _, num := range nums {
		s += -num
		if s < 0 {
			s = 0
		}
		if s > maxSum {
			maxSum = s
		}
	}
	if maxSum == 0 {
		// Find max of -i for each i in nums
		maxNeg := -nums[0]
		for i := 1; i < len(nums); i++ {
			negVal := -nums[i]
			if negVal > maxNeg {
				maxNeg = negVal
			}
		}
		maxSum = maxNeg
	}
	minSum := -maxSum
	return minSum
}

func main() {
	// Read input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// Parse input string like "[2, 3, 4, 1, 2, 4]" to slice
	input = strings.Trim(input, "[]")
	parts := strings.Split(input, ", ")
	
	nums := make([]int, len(parts))
	for i, p := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		nums[i] = num
	}
	
	result := Minsubarraysum(nums)
	fmt.Println(result)
}
