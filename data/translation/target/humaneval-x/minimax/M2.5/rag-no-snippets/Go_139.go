package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SpecialFactorial(n int) int {
	factI := 1
	specialFact := 1
	for i := 1; i <= n; i++ {
		factI *= i
		specialFact *= factI
	}
	return specialFact
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error: invalid input")
			return
		}
		result := SpecialFactorial(n)
		fmt.Println(result)
	}
}
