package main

import (
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Convert to int64 to handle large numbers from exponentiation
	values := make([]int64, len(operand))
	for i, v := range operand {
		values[i] = int64(v)
	}
	ops := operator

	// Process by operator precedence:
	// 1. ** (exponentiation) - right to left associativity
	// 2. * and // (multiplication and floor division) - left to right
	// 3. + and - (addition and subtraction) - left to right

	// ** - right to left (e.g., 2**3**2 = 2**(3**2) = 2**9 = 512)
	for i := len(ops) - 1; i >= 0; i-- {
		if ops[i] == "**" {
			result := powerInt64(values[i], values[i+1])
			values = append(values[:i], append([]int64{result}, values[i+2:]...)...)
			ops = append(ops[:i], ops[i+1:]...)
		}
	}

	// * and // - left to right
	for i := 0; i < len(ops); i++ {
		if ops[i] == "*" || ops[i] == "//" {
			result := applyOp(values[i], values[i+1], ops[i])
			values = append(values[:i], append([]int64{result}, values[i+2:]...)...)
			ops = append(ops[:i], ops[i+1:]...)
			i-- // Stay at same position after removal
		}
	}

	// + and - - left to right
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" || ops[i] == "-" {
			result := applyOp(values[i], values[i+1], ops[i])
			values = append(values[:i], append([]int64{result}, values[i+2:]...)...)
			ops = append(ops[:i], ops[i+1:]...)
			i-- // Stay at same position after removal
		}
	}

	return int(values[0])
}

func applyOp(a, b int64, op string) int64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "//":
		return a / b
	}
	return 0
}

func powerInt64(base, exp int64) int64 {
	result := int64(1)
	for i := int64(0); i < exp; i++ {
		result *= base
	}
	return result
}

// Helper to test the function (not part of the required signature)
func main() {
	// Example test: operator=['+', '*', '-'], operand=[2, 3, 4, 5]
	// Expression: 2 + 3 * 4 - 5 = 2 + 12 - 5 = 9
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	println("Result:", result)
	
	// Test with exponentiation: operator=['**', '+'], operand=[2, 3, 1]
	// Expression: 2 ** 3 + 1 = 8 + 1 = 9
	operator2 := []string{"**", "+"}
	operand2 := []int{2, 3, 1}
	result2 := DoAlgebra(operator2, operand2)
	println("Result2:", result2)
	
	// Test with floor division: operator=['//'], operand=[10, 3]
	// Expression: 10 // 3 = 3
	operator3 := []string{"//"}
	operand3 := []int{10, 3}
	result3 := DoAlgebra(operator3, operand3)
	println("Result3:", result3)
	
	// Test: operator=['*', '+'], operand=[2, 3, 4]
	// Expression: 2 * 3 + 4 = 6 + 4 = 10
	operator4 := []string{"*", "+"}
	operand4 := []int{2, 3, 4}
	result4 := DoAlgebra(operator4, operand4)
	println("Result4:", result4)
}