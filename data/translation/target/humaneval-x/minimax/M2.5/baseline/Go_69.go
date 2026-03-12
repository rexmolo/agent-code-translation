package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Search(lst []int) int {
	// Count frequency of each integer using a map
	freq := make(map[int]int)
	maxVal := 0
	for _, v := range lst {
		freq[v]++
		if v > maxVal {
			maxVal = v
		}
	}

	// Find the greatest integer i where frequency >= i
	ans := -1
	for i := 1; i <= maxVal; i++ {
		if freq[i] >= i {
			ans = i
		}
	}

	return ans
}

func main() {
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// Parse the input list from space-separated values
	strs := strings.Fields(line)
	lst := make([]int, len(strs))
	for i, s := range strs {
		lst[i], _ = strconv.Atoi(s)
	}

	result := Search(lst)
	fmt.Println(result)
}
