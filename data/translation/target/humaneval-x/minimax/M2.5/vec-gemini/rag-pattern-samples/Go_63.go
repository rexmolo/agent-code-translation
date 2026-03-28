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
	return Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter n: ")
	scanner.Scan()
	input := scanner.Text()

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	result := Fibfib(n)
	fmt.Println(result)
}
