package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Fibfib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter n: ")
	if scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		result := Fibfib(n)
		fmt.Println(result)
	}
}
