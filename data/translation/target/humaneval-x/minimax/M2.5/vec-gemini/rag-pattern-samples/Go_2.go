package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	
	var number float64
	if scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%f", &number)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	}

	result := TruncateNumber(number)
	fmt.Printf("%.1f\n", result)
}