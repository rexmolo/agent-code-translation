package main

import (
    "strconv"
    "strings"
)

func DoAlgebra(operator []string, operand []int) int {
    // Build expression string from operator and operand lists
    var exprBuilder strings.Builder
    exprBuilder.WriteString(strconv.Itoa(operand[0]))
    
    for i, oprt := range operator {
        exprBuilder.WriteString(oprt)
        exprBuilder.WriteString(strconv.Itoa(operand[i+1]))
    }
    
    expression := exprBuilder.String()
    
    // Parse and evaluate the expression
    tokens := parseTokens(expression)
    return evaluateExpression(tokens)
}

// Tokenize the expression into numbers and operators
func parseTokens(expr string) []string {
    var tokens []string
    var currentNum strings.Builder
    
    for i := 0; i < len(expr); i++ {
        ch := expr[i]
        if ch >= '0' && ch <= '9' {
            currentNum.WriteByte(ch)
        } else if ch == '-' && (i == 0 || expr[i-1] == '+' || expr[i-1] == '-' || expr[i-1] == '*' || expr[i-1] == '/' || expr[i-1] == '^') {
            // Negative number
            currentNum.WriteByte(ch)
        } else {
            if currentNum.Len() > 0 {
                tokens = append(tokens, currentNum.String())
                currentNum.Reset()
            }
            if string(ch) == "*" && i+1 < len(expr) && expr[i+1] == '*' {
                // Handle ** as a single token
                tokens = append(tokens, "**")
                i++ // skip next '*'
            } else if string(ch) == "/" && i+1 < len(expr) && expr[i+1] == '/' {
                // Handle // as a single token
                tokens = append(tokens, "//")
                i++ // skip next '/'
            } else {
                tokens = append(tokens, string(ch))
            }
        }
    }
    
    if currentNum.Len() > 0 {
        tokens = append(tokens, currentNum.String())
    }
    
    return tokens
}

// Evaluate expression with proper operator precedence:
// Highest: ** (right-associative)
// Middle: * and //
// Lowest: + and -
func evaluateExpression(tokens []string) int {
    pos := 0
    
    // Parse expression: handles + and -
    result := parseTerm(&tokens, &pos)
    
    for pos < len(tokens) {
        op := tokens[pos]
        if op != "+" && op != "-" {
            break
        }
        pos++
        right := parseTerm(&tokens, &pos)
        if op == "+" {
            result += right
        } else {
            result -= right
        }
    }
    
    return result
}

// Parse term: handles * and //
func parseTerm(tokens *[]string, pos *int) int {
    result := parseFactor(tokens, pos)
    
    for *pos < len(*tokens) {
        op := (*tokens)[*pos]
        if op != "*" && op != "//" {
            break
        }
        (*pos)++
        right := parseFactor(tokens, pos)
        if op == "*" {
            result *= right
        } else {
            result /= right // Floor division
        }
    }
    
    return result
}

// Parse factor: handles ** (right-associative)
func parseFactor(tokens *[]string, pos *int) int {
    result := parsePrimary(tokens, pos)
    
    if *pos < len(*tokens) && (*tokens)[*pos] == "**" {
        (*pos)++
        // ** is right-associative: parseFactor for right operand
        right := parseFactor(tokens, pos)
        result = powInt(result, right)
    }
    
    return result
}

// Parse primary: handles numbers (including negative)
func parsePrimary(tokens *[]string, pos *int) int {
    num, _ := strconv.Atoi((*tokens)[*pos])
    (*pos)++
    return num
}

// Integer exponentiation
func powInt(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}

// For testing
func main() {
    // Test example from docstring
    operators := []string{"+", "*", "-"}
    operands := []int{2, 3, 4, 5}
    result := DoAlgebra(operators, operands)
    println("Result:", result) // Should print 9
    
    // Additional tests
    // Test: 2 ** 3 ** 2 = 2 ** 9 = 512
    operators2 := []string{"**", "**"}
    operands2 := []int{2, 3, 2}
    result2 := DoAlgebra(operators2, operands2)
    println("Result2:", result2) // Should print 512
    
    // Test: 10 // 3 = 3
    operators3 := []string{"//"}
    operands3 := []int{10, 3}
    result3 := DoAlgebra(operators3, operands3)
    println("Result3:", result3) // Should print 3
}
