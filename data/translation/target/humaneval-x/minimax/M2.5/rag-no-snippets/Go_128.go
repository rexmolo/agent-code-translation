package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check if there's a 0 in the array
	hasZero := false
	for _, v := range arr {
		if v == 0 {
			hasZero = true
			break
		}
	}

	// Count negative numbers
	negCount := 0
	for _, v := range arr {
		if v < 0 {
			negCount++
		}
	}

	// Calculate prod: 0 if hasZero, otherwise (-1)^negCount
	var prod int
	if hasZero {
		prod = 0
	} else {
		if negCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	// Calculate sum of absolute values
	sumAbs := 0
	for _, v := range arr {
		if v < 0 {
			sumAbs += -v
		} else {
			sumAbs += v
		}
	}

	return prod * sumAbs
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter array elements separated by spaces: ")
	scanner.Scan()
	line := scanner.Text()

	if line == "" {
		// Empty array test
		result := ProdSigns([]int{})
		fmt.Printf("Result: %v\n", result)
		return
	}

	parts := strings.Split(line, " ")
	arr := make([]int, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid input %q\n", p)
			return
		}
		arr[i] = n
	}

	result := ProdSigns(arr)
	fmt.Printf("Result: %v\n", result)
}
