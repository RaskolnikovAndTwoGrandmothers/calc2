package parser

import (
	"strconv"
	"strings"
)

// Token представляет лексему
type Token struct {
	Type  TokenType
	Value string
}

// TokenType определяет тип лексемы
type TokenType int

const (
	Number TokenType = iota
	Operator
	LeftParen
	RightParen
)

// Parse преобразует входную строку в последовательность токенов
func Parse(input string) ([]Token, error) {
	var tokens []Token
	var buf strings.Builder

	for _, ch := range input {
		switch {
		case isDigit(ch) || ch == '.':
			buf.WriteRune(ch)
		case isOperator(ch):
			if buf.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: buf.String()})
				buf.Reset()
			}
			tokens = append(tokens, Token{Type: Operator, Value: string(ch)})
		case ch == '(':
			if buf.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: buf.String()})
				buf.Reset()
			}
			tokens = append(tokens, Token{Type: LeftParen, Value: string(ch)})
		case ch == ')':
			if buf.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: buf.String()})
				buf.Reset()
			}
			tokens = append(tokens, Token{Type: RightParen, Value: string(ch)})
		case ch == ' ':
			if buf.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: buf.String()})
				buf.Reset()
			}
		default:
			return nil, &ParseError{Message: "invalid character: " + string(ch)}
		}
	}

	if buf.Len() > 0 {
		tokens = append(tokens, Token{Type: Number, Value: buf.String()})
	}

	return tokens, nil
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^'
}

// ParseError представляет ошибку парсинга
type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}

// ParseNumber преобразует строку в число
func ParseNumber(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
