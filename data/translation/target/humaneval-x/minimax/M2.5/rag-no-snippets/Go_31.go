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
	for k := 2; k < n-1; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	
	n, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Error: invalid input")
		return
	}
	
	result := IsPrime(n)
	if result {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
