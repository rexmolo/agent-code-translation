package main

import (
	"fmt"
)

// DoAlgebra evaluates a mathematical expression given operators and operands.
// It follows standard operator precedence: ** > * and // > + and -
func DoAlgebra(operator []string, operand []int) int {
	// Create working copies to avoid modifying input
	values := make([]int, len(operand))
	copy(values, operand)
	ops := make([]string, len(operator))
	copy(ops, operator)

	// First pass: handle high precedence operators (**, *, //)
	for i := 0; i < len(ops); {
		switch ops[i] {
		case "**":
			values[i] = power(values[i], values[i+1])
			values = append(values[:i+1], values[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
		case "*":
			values[i] *= values[i+1]
			values = append(values[:i+1], values[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
		case "//":
			values[i] /= values[i+1]
			values = append(values[:i+1], values[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
		default:
			i++
		}
	}

	// Second pass: handle low precedence operators (+, -)
	result := values[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "+":
			result += values[i+1]
		case "-":
			result -= values[i+1]
		}
	}

	return result
}

// power computes base raised to the power of exp
func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	// Test case from the example
	operator1 := []string{"+", "*", "-"}
	operand1 := []int{2, 3, 4, 5}
	fmt.Printf("Result: %d (expected: 9)\n", DoAlgebra(operator1, operand1))

	// Additional test case with exponentiation
	operator2 := []string{"+", "**"}
	operand2 := []int{2, 3, 4}
	fmt.Printf("Result: %d (expected: 83, since 2+3^4=2+81)\n", DoAlgebra(operator2, operand2))
}