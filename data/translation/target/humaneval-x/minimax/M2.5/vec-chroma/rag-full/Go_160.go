package main

import (
	"math"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	if len(operand) < 2 || len(operator) < 1 {
		return 0
	}

	// Build expression tokens in postfix (RPN) order using Shunting-Yard algorithm
	rpn := []string{}
	stack := []string{}

	precedence := map[string]int{
		"+":  1,
		"-":  1,
		"*":  2,
		"//": 2,
		"**": 3,
	}

	for i, op := range operator {
		// Add left operand
		rpn = append(rpn, strconv.Itoa(operand[i]))

		// Pop operators with higher precedence (for left-associative) or greater (for **)
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			pTop := precedence[top]
			pCurr := precedence[op]
			if pTop > pCurr || (pTop == pCurr && top != "**") {
				rpn = append(rpn, top)
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, op)
	}

	// Add last operand
	rpn = append(rpn, strconv.Itoa(operand[len(operand)-1]))

	// Pop remaining operators
	for len(stack) > 0 {
		rpn = append(rpn, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	// Evaluate RPN expression
	evalStack := []int{}
	for _, token := range rpn {
		switch token {
		case "+", "-", "*", "//", "**":
			if len(evalStack) < 2 {
				continue
			}
			n1 := evalStack[len(evalStack)-1]
			n2 := evalStack[len(evalStack)-2]
			evalStack = evalStack[:len(evalStack)-2]
			var result int
			switch token {
			case "+":
				result = n2 + n1
			case "-":
				result = n2 - n1
			case "*":
				result = n2 * n1
			case "//":
				result = n2 / n1
			case "**":
				result = int(math.Pow(float64(n2), float64(n1)))
			}
			evalStack = append(evalStack, result)
		default:
			num, _ := strconv.Atoi(token)
			evalStack = append(evalStack, num)
		}
	}

	return evalStack[0]
}

func main() {
	// Test with example: 2 + 3 * 4 - 5 = 9
	operators := []string{"+", "*", "-"}
	operands := []int{2, 3, 4, 5}
	result := DoAlgebra(operators, operands)
	println("Result:", result) // Should be 9
}