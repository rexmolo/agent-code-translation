package main

import (
	"fmt"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build expression by alternating operands and operators
	expr := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expr += oprt + strconv.Itoa(operand[i+1])
	}

	// Evaluate the expression respecting operator precedence
	// Order: ** (highest), then * //, then + - (lowest)
	// We process operators in precedence groups

	// Create mutable copies
	ops := make([]string, len(operator))
	nums := make([]int, len(operand))
	copy(ops, operator)
	copy(nums, operand)

	// Handle exponentiation (**) - right to left
	for i := 0; i < len(ops); i++ {
		if ops[i] == "**" {
			// Calculate result
			result := 1
			base := nums[i]
			exp := nums[i+1]
			for j := 0; j < exp; j++ {
				result *= base
			}
			nums[i] = result
			// Remove the used operator and operand
			copy(ops[i:], ops[i+1:])
			copy(nums[i+1:], nums[i+2:])
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			i-- // Stay at current index to check again
		}
	}

	// Handle multiplication (*) and floor division (//) - left to right
	for i := 0; i < len(ops); i++ {
		if ops[i] == "*" {
			nums[i] = nums[i] * nums[i+1]
			copy(ops[i:], ops[i+1:])
			copy(nums[i+1:], nums[i+2:])
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			i--
		} else if ops[i] == "//" {
			nums[i] = nums[i] / nums[i+1]
			copy(ops[i:], ops[i+1:])
			copy(nums[i+1:], nums[i+2:])
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			i--
		}
	}

	// Handle addition (+) and subtraction (-) - left to right
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" {
			nums[i] = nums[i] + nums[i+1]
			copy(ops[i:], ops[i+1:])
			copy(nums[i+1:], nums[i+2:])
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			i--
		} else if ops[i] == "-" {
			nums[i] = nums[i] - nums[i+1]
			copy(ops[i:], ops[i+1:])
			copy(nums[i+1:], nums[i+2:])
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			i--
		}
	}

	return nums[0]
}

func main() {
	// Test the function
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	fmt.Println(result) // Expected: 9

	// Additional test cases
	operator2 := []string{"**", "+"}
	operand2 := []int{2, 3, 4}
	result2 := DoAlgebra(operator2, operand2)
	fmt.Println(result2) // 2**3 + 4 = 8 + 4 = 12

	operator3 := []string{"//", "-"}
	operand3 := []int{10, 2, 3}
	result3 := DoAlgebra(operator3, operand3)
	fmt.Println(result3) // 10 // 2 - 3 = 5 - 3 = 2
}