package calculation

import (
	"strconv"
	"strings"
)

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func isValidCharacter(char rune) bool {
	return (char >= '0' && char <= '9') || char == '.' || isOperator(char) || char == '(' || char == ')'
}

func parseNumber(input string) (float64, error) {
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, ErrInvalidExpression
	}
	return num, nil
}

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if len(expression) == 0 {
		return 0, ErrInvalidExpression
	}

	for _, char := range expression {
		if !isValidCharacter(char) {
			return 0, ErrInvalidExpression
		}
	}

	return evaluate(expression)
}

func evaluate(expression string) (float64, error) {
	if strings.Contains(expression, "--") || strings.Contains(expression, "++") {
		return 0, ErrInvalidExpression
	}

	for strings.Contains(expression, "(") {
		start := strings.LastIndex(expression, "(")
		end := strings.Index(expression[start:], ")") + start
		if end <= start {
			return 0, ErrInvalidExpression
		}
		subExpression := expression[start+1 : end]
		result, err := evaluate(subExpression)
		if err != nil {
			return 0, err
		}
		expression = strings.Replace(expression, "("+subExpression+")", strconv.FormatFloat(result, 'f', -1, 64), 1)
	}

	return parseNumber(expression)
}
