package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for k := 2; k < n; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)
	fmt.Printf("IsPrime(%d) = %v\n", n, IsPrime(n))
}
