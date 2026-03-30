package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt(nums []int) int {
	result := nums[0]
	for _, n := range nums[1:] {
		if n > result {
			result = n
		}
	}
	return result
}

func Minsubarraysum(nums []int) int {
	maxSum := 0
	s := 0
	for _, num := range nums {
		s += -num
		if s < 0 {
			s = 0
		}
		maxSum = max(s, maxSum)
	}
	if maxSum == 0 {
		maxSum = maxInt(nums)
		maxSum = -maxSum
	}
	minSum := -maxSum
	return minSum
}

func main() {
	// Read input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// Parse input string to slice of integers
	input = strings.Trim(input, "[]")
	var nums []int
	if input != "" {
		parts := strings.Split(input, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				num, _ := strconv.Atoi(p)
				nums = append(nums, num)
			}
		}
	}
	
	result := Minsubarraysum(nums)
	fmt.Println(result)
}
