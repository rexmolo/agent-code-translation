package main

import "bufio"
import "fmt"
import "os"
import "strconv"

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
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	n, _ := strconv.Atoi(input)
	result := Fibfib(n)
	fmt.Println(result)
}