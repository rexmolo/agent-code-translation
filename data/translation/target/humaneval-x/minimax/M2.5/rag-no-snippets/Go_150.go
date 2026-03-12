package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func XOrY(n, x, y int) int {
	if n == 1 {
		return y
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return y
		}
	}
	return x
}

func main() {
	// Read input: n x y
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		return
	}

	n, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])

	result := XOrY(n, x, y)
	fmt.Println(result)
}
