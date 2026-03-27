package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func FruitDistribution(s string, n int) int {
	parts := strings.Split(s, " ")
	sum := 0
	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			sum += num
		}
	}
	return n - sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the fruit distribution string: ")
	scanner.Scan()
	s := scanner.Text()
	
	fmt.Print("Enter the total number of fruits: ")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	result := FruitDistribution(s, n)
	fmt.Printf("Number of mango fruits: %d\n", result)
}