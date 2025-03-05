package pkg

import (
    "strconv"
    "strings"
)

func EvaluateExpression(expression string) (float64, error) {
    parts := strings.Fields(expression)
    var result float64
    var operator string

    for i, part := range parts {
        if i == 0 {
            num, err := strconv.ParseFloat(part, 64)
            if err != nil {
                return 0, err
            }
            result = num
            continue
        }

        if isOperator(part) {
            operator = part
            continue
        }

        num, err := strconv.ParseFloat(part, 64)
        if err != nil {
            return 0, err
        }

        switch operator {
        case "+":
            result += num
        case "-":
            result -= num
        case "*":
            result *= num
        case "/":
            result /= num
        }
    }
    return result, nil
}

func isOperator(s string) bool {
    return s == "+" || s == "-" || s == "*" || s == "/"
}
