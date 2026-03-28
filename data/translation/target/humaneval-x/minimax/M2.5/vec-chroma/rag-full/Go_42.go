package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IncrList(l []int) []int {
	result := make([]int, 0, len(l))
	for _, e := range l {
		result = append(result, e+1)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var nums []int
		for _, p := range parts {
			if n, err := strconv.Atoi(p); err == nil {
				nums = append(nums, n)
			}
		}
		result := IncrList(nums)
		var output []string
		for _, v := range result {
			output = append(output, strconv.Itoa(v))
		}
		fmt.Println(strings.Join(output, " "))
	}
}
