package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddElements(arr []int, k int) int {
	sum := 0
	// Ensure k doesn't exceed array length
	if k > len(arr) {
		k = len(arr)
	}
	for i := 0; i < k; i++ {
		s := strconv.Itoa(arr[i])
		if len(s) <= 2 {
			sum += arr[i]
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Read the array (first line)
	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	arr := make([]int, len(arrStr))
	for i, s := range arrStr {
		arr[i], _ = strconv.Atoi(s)
	}
	
	// Read k (second line)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	
	// Calculate and output result
	result := AddElements(arr, k)
	fmt.Println(result)
}
