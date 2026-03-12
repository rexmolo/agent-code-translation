package main

import (
	"fmt"
	"math"
)

// power is a helper function for integer exponentiation, replicating Python's ** operator.
func power(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

// DoAlgebra evaluates an algebraic expression defined by a list of operators and operands.
func DoAlgebra(operator []string, operand []int) int {
	// Create copies to avoid modifying the input slices.
	operands := make([]int, len(operand))
	copy(operands, operand)
	operators := make([]string, len(operator))
	copy(operators, operator)

	// process is a helper closure that computes operations of a specific precedence level.
	process := func(targetOps map[string]bool) {
		i := 0
		for i < len(operators) {
			op := operators[i]
			if targetOps[op] {
				left := operands[i]
				right := operands[i+1]
				var result int
				switch op {
				case "**":
					result = power(left, right)
				case "*":
					result = left * right
				case "//":
					// Go's '/' on integers performs floor division for non-negative numbers,
					// matching Python's '//'. The problem states operands are non-negative.
					result = left / right
				}

				// Replace the left operand with the result.
				operands[i] = result
				// Remove the consumed right operand and the operator from their slices.
				operands = append(operands[:i+1], operands[i+2:]...)
				operators = append(operators[:i], operators[i+1:]...)
				// Do not increment 'i', so we can process the next operation at the current position
				// in case of consecutive high-precedence operators (e.g., 2*3*4).
			} else {
				i++
			}
		}
	}

	// 1st pass: Exponentiation (highest precedence).
	process(map[string]bool{"**": true})
	// 2nd pass: Multiplication and Division.
	process(map[string]bool{"*": true, "//": true})

	// 3rd pass: Addition and Subtraction (evaluated left to right).
	// After the first two passes, only these low-precedence operators remain.
	result := operands[0]
	for i, op := range operators {
		right := operands[i+1]
		switch op {
		case "+":
			result += right
		case "-":
			result -= right
		}
	}

	return result
}

func main() {
	// Example from the Python function's docstring.
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	
	// The expression is: 2 + 3 * 4 - 5
	// Evaluation: 2 + 12 - 5 = 9
	result := DoAlgebra(operator, operand)
	fmt.Printf("Operators: %v\n", operator)
	fmt.Printf("Operands: %v\n", operand)
	fmt.Printf("Result: %d\n", result)

	// Another test case: 2 ** 3 * 4 -> 8 * 4 = 32
	operator2 := []string{"**", "*"}
	operand2 := []int{2, 3, 4}
	result2 := DoAlgebra(operator2, operand2)
	fmt.Printf("\nOperators: %v\n", operator2)
	fmt.Printf("Operands: %v\n", operand2)
	fmt.Printf("Result: %d\n", result2)
}