package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

func main() {
	// Read input line
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// Parse the float64 number
	var number float64
	fmt.Sscan(input, &number)
	
	// Get the decimal part
	dec := TruncateNumber(number)
	fmt.Println(dec)
}
