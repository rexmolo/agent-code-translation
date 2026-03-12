package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(N int) string {
	// Convert N to string
	strN := strconv.Itoa(N)

	// Sum all digits
	sum := 0
	for _, c := range strN {
		digit := int(c - '0')
		sum += digit
	}

	// Convert sum to binary string
	return fmt.Sprintf("%b", sum)
}

func main() {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		N, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		result := Solve(N)
		fmt.Println(result)
	}
}
