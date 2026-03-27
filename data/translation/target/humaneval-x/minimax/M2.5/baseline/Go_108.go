package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountNums(arr []int) int {
	// Helper function to calculate digit sum with signed first digit
	digitSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert to string and process each digit
		s := strconv.Itoa(n)
		sum := 0
		for i, c := range s {
			digit := int(c - '0')
			if i == 0 {
				digit *= neg
			}
			sum += digit
		}
		return sum
	}

	// Count elements where digit sum > 0
	count := 0
	for _, n := range arr {
		if digitSum(n) > 0 {
			count++
		}
	}
	return count
}

func main() {
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Parse input array (format: [1, 2, -3])
	input = strings.Trim(input, "[]")
	if input == "" {
		fmt.Println(CountNums([]int{}))
		return
	}

	parts := strings.Split(input, ", ")
	arr := make([]int, len(parts))
	for i, p := range parts {
		arr[i], _ = strconv.Atoi(strings.TrimSpace(p))
	}

	fmt.Println(CountNums(arr))
}
