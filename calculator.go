package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"regexp"
)

func calculator(s string) string {
	match, _ := regexp.MatchString(`^[0-9+\-*/().\s]+$`, s)
	if !match {
		return "Разрешены только цифры и символы + - * /"
	}

	if len(s) > 0 && (s[0] == '+' || s[0] == '*' || s[0] == '/') {
		return "Некорректное выражение. Проверьте ввод."
	}

	expression, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		return "Некорректное выражение. Проверьте ввод."
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return "Некорректное выражение. Проверьте ввод."
	}

	ans, ok := result.(float64)
	if !ok {
		return "Invalid expression"
	}

	if ans == float64(int(ans)) {
		return fmt.Sprintf("%d", int(ans))
	}
	return fmt.Sprintf("%.2f", ans)
}
