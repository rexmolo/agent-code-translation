package main

import (
	"errors"
	"strconv"
	"strings"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build the expression string similar to Python
	expression := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expression += oprt + strconv.Itoa(operand[i+1])
	}

	// Evaluate the expression
	tokens := tokenize(expression)
	parser := &parser{tokens: tokens, pos: 0}
	result, err := parser.parseExpression()
	if err != nil {
		panic(err)
	}
	return result
}

// Token represents a token in the expression
type token struct {
	TokenType int
	Value     string
}

const (
	TOKEN_NUMBER = iota
	TOKEN_OPERATOR
)

// Tokenize the expression string
func tokenize(expr string) []token {
	var tokens []token
	var current strings.Builder

	for i := 0; i < len(expr); i++ {
		ch := expr[i]

		// Check for double character operators first
		if i+1 < len(expr) && expr[i:i+2] == "**" {
			if current.Len() > 0 {
				tokens = append(tokens, token{TOKEN_NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, token{TOKEN_OPERATOR, "**"})
			i++ // skip next character
			continue
		}

		// Check for // operator
		if i+1 < len(expr) && expr[i:i+2] == "//" {
			if current.Len() > 0 {
				tokens = append(tokens, token{TOKEN_NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, token{TOKEN_OPERATOR, "//"})
			i++ // skip next character
			continue
		}

		// Handle single character operators
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if current.Len() > 0 {
				tokens = append(tokens, token{TOKEN_NUMBER, current.String()})
				current.Reset()
			}
			tokens = append(tokens, token{TOKEN_OPERATOR, string(ch)})
			continue
		}

		// Handle numbers (including multi-digit)
		current.WriteByte(ch)
	}

	// Add the last number
	if current.Len() > 0 {
		tokens = append(tokens, token{TOKEN_NUMBER, current.String()})
	}

	return tokens
}

// Parser structure
type parser struct {
	tokens []token
	pos    int
}

func (p *parser) currentToken() token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return token{TOKEN_OPERATOR, ""}
}

func (p *parser) consume() token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) parseExpression() (int, error) {
	return p.parseAddSub()
}

// parseAddSub handles + and - (lowest precedence)
func (p *parser) parseAddSub() (int, error) {
	left, err := p.parseMulDiv()
	if err != nil {
		return 0, err
	}

	for p.pos < len(p.tokens) {
		tok := p.currentToken()
		if tok.TokenType != TOKEN_OPERATOR || (tok.Value != "+" && tok.Value != "-") {
			break
		}
		p.consume()
		right, err := p.parseMulDiv()
		if err != nil {
			return 0, err
		}
		if tok.Value == "+" {
			left = left + right
		} else {
			left = left - right
		}
	}
	return left, nil
}

// parseMulDiv handles * and // (same precedence, left associative)
func (p *parser) parseMulDiv() (int, error) {
	left, err := p.parsePower()
	if err != nil {
		return 0, err
	}

	for p.pos < len(p.tokens) {
		tok := p.currentToken()
		if tok.TokenType != TOKEN_OPERATOR || (tok.Value != "*" && tok.Value != "//") {
			break
		}
		p.consume()
		right, err := p.parsePower()
		if err != nil {
			return 0, err
		}
		if tok.Value == "*" {
			left = left * right
		} else {
			left = left / right // integer division (floor for positive numbers)
		}
	}
	return left, nil
}

// parsePower handles ** (highest precedence, right associative)
func (p *parser) parsePower() (int, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}

	if p.pos < len(p.tokens) && p.currentToken().Value == "**" {
		p.consume()
		right, err := p.parsePower() // right associative
		if err != nil {
			return 0, err
		}
		left = power(left, right)
	}
	return left, nil
}

// parseUnary handles primary numbers
func (p *parser) parseUnary() (int, error) {
	tok := p.currentToken()
	if tok.TokenType != TOKEN_NUMBER {
		return 0, errors.New("expected number")
	}
	p.consume()
	num, err := strconv.Atoi(tok.Value)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// For testing - simple main function
func main() {
	// Test case from the problem
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	println("Result:", result) // Should be 9
}
