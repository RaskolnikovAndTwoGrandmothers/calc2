package evaluator

import (
	"fmt"
	"math"
	"calc2/internal/parser"
)

// Evaluate вычисляет значение выражения по последовательности токенов
func Evaluate(tokens []parser.Token) (float64, error) {
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}

	stack := make([]float64, 0)

	for _, token := range postfix {
		switch token.Type {
		case parser.Number:
			num, err := parser.ParseNumber(token.Value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		case parser.Operator:
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough operands for operator %s", token.Value)
			}
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			res, err := applyOperator(a, b, token.Value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, res)
		default:
			return 0, fmt.Errorf("unexpected token type in postfix evaluation")
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

func applyOperator(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func infixToPostfix(tokens []parser.Token) ([]parser.Token, error) {
	var output []parser.Token
	var stack []parser.Token

	for _, token := range tokens {
		switch token.Type {
		case parser.Number:
			output = append(output, token)
		case parser.Operator:
			for len(stack) > 0 && isHigherPrecedence(stack[len(stack)-1].Value, token.Value) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case parser.LeftParen:
			stack = append(stack, token)
		case parser.RightParen:
			for len(stack) > 0 && stack[len(stack)-1].Type != parser.LeftParen {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // Pop the left paren
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1].Type == parser.LeftParen {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func isHigherPrecedence(op1, op2 string) bool {
	precedence := map[string]int{
		"^": 4,
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
	}

	return precedence[op1] >= precedence[op2] && op1 != "^"
}
