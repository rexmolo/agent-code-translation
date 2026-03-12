package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RightAngleTriangle(a, b, c int) bool {
	return a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter three sides of a triangle separated by spaces: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a, b, c int
	fmt.Sscanf(input, "%d %d %d", &a, &b, &c)

	result := RightAngleTriangle(a, b, c)
	fmt.Println(result)
}
