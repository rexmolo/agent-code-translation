package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMaxTriples(n int) int {
	// Create array A where A[i] = i*i - i + 1 for i from 1 to n
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i*i - i + 1
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (A[i]+A[j]+A[k])%3 == 0 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)
	result := GetMaxTriples(n)
	fmt.Println(result)
}
